package jobs

import (
	"bytes"
	"fmt"
	"html/template"
	"os/exec"
	"github.com/qnsoft/online_pay/utils/pay/fengfu"
	"github.com/qnsoft/online_pay/utils/pay/paytool"
	"github.com/qnsoft/web_api/models/shop"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego"
)

var mailTpl *template.Template

func init() {
	mailTpl, _ = template.New("mail_tpl").Parse(`
	你好 {{.username}}，<br/>

<p>以下是任务执行结果：</p>

<p>
任务ID：{{.task_id}}<br/>
任务名称：{{.task_name}}<br/>       
执行时间：{{.start_time}}<br />
执行耗时：{{.process_time}}秒<br />
执行状态：{{.status}}
</p>
<p>-------------以下是任务执行输出-------------</p>
<p>{{.output}}</p>
<p>
--------------------------------------------<br />
本邮件由系统自动发出，请勿回复<br />
如果要取消邮件通知，请登录到系统进行设置<br />
</p>
`)

}

type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *shop.UserCardJobOrder                            // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(task *shop.UserCardJobOrder) (*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(task.Id, task.TaskName, "task.Command")
	job.task = task
	//	job.Concurrent = task.Concurrent == 1
	return job, nil
}

func NewCommandJob(id int, name string, command string) *Job {
	job := &Job{
		id:   id,
		name: name,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		cmd.Start()
		err, isTimeout := runCmdWithTimeout(cmd, timeout)

		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}

func (j *Job) Status() int {
	return j.status
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetId() int {
	return j.id
}

func (j *Job) GetLogId() int64 {
	return j.logId
}

func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		beego.Warn(fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.id))
		return
	}

	defer func() {
		if err := recover(); err != nil {
			beego.Error(err, "\n", string(debug.Stack()))
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	//开始执行计划任务 //根据任务id,从数据库表中取出要执行的任务
	beego.Debug(fmt.Sprintf("开始执行信用卡代还: %d", strconv.Itoa(j.id)+"["+j.name+"]"))

	j.status++
	defer func() {
		j.status--
	}()
	ZN_DF(j.id)
	//fmt.Println("信用卡代还执行完毕!")
}

/*
*手动代还代付
 */
func ZN_DF(_id int) {
	_UserId := ""
	var _json_model interface{}
	//根据订单id获取当前要执行的订单
	_job_order_model := &shop.UserCardJobOrder{Id: _id}
	results, err := DbHelper.MySqlDb().Get(_job_order_model)
	ErrorHelper.CheckErr(err)
	_model_card := shop.UserBankCard{BankCardNumber: _job_order_model.Bankcardno}
	results_card, _ := DbHelper.MySqlDb().Get(&_model_card)
	if results {
		_channelType := "ffqrdh" //通道标示小额channelType：ffqrdh大额channelType：ffkj
		_fee := 0.36
		if results_card && len(_model_card.FfMerordernumber) > 3 && len(_model_card.FfIssign) > 0 && _job_order_model.Amount > 100000 {
			_channelType = "fk"
			_fee = 0.58
		}
		_UserId = _job_order_model.UserId                                               //用户id
		_BankCardNo := _job_order_model.Bankcardno                                      //用户出金卡号
		_MerchantBankAccountNo := _job_order_model.Merchantbankaccountno                //用户入金卡号
		_Amount, _ := strconv.ParseFloat(strconv.Itoa(_job_order_model.Amount/100), 10) //代扣金额
		//根据userid和出金卡号取代扣签约信息
		_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankCardNo}
		results, _ := DbHelper.MySqlDb().Get(_model)
		if results {
			//信用卡出金-----------------------------------------
			cardModelA := paytool.CardModel{
				AccountNumber: _BankCardNo,       //出金卡号
				Phone:         _model.Mobile,     //手机号
				HolderName:    _model.Cardholder, //姓名
				IdCard:        _model.IdCard,     //身份证号
				Cvv2:          _model.Cvv2,       //Cvv2
				Expired:       _model.Expiretime, //有效期
			}
			fmt.Printf("%v", cardModelA)

			_city := "郑州市"
			//if _Amount > 1000 {
			_channelType = "fk"
			_fee = 0.58
			//}
			//_money := 100.0
			payTool := fengfu.FengFuPay{}
			rtModel := payTool.CardOut(cardModelA, _channelType, _city, _Amount, _fee) //出金操作
			_str_rt := fmt.Sprintf("%+v", rtModel)
			if rtModel.Code == "00" { //如果Code=0 下单成功写入数据库
				//开始将代扣信息写入到信用卡订单表[lkt_user_card_order]
				_update_model_job_order := shop.UserCardJobOrder{
					UserId:                _model_card.UserId,                     //用户id
					Ordernoa:              rtModel.Content.MerOrderNumber,         //代扣订单编号
					Ordernob:              "000000",                               //代付订单编号
					Amount:                _job_order_model.Amount,                //订单金额
					Rateamounta:           int(_Amount * (_fee / 100)),            //代扣手续费
					Bankcardno:            _job_order_model.Bankcardno,            //出金卡号
					Merchantbankaccountno: _job_order_model.Merchantbankaccountno, //入金卡号
					ReturnA:               "丰付智能代扣-" + _str_rt,                    //代扣返回信息
					//AddTime:               time.Now(),                             //订单创建时间
				}
				_count, err := DbHelper.MySqlDb().Id(_job_order_model.Id).Update(&_update_model_job_order) //修改智能代还订单
				ErrorHelper.CheckErr(err)
				if _count > 0 && err == nil {
					_AmountB := _Amount - _Amount*(_fee/100)
					var wg = &sync.WaitGroup{}
					wg.Add(1)
					go func() { //开始异步执行
						Quick_Pay_Ok(cardModelA, _UserId, rtModel.Content.MerOrderNumber, _BankCardNo, _MerchantBankAccountNo, _AmountB)
						wg.Done()
					}()
					//	time.Sleep(time.Second * 2)
					wg.Wait()
					//开始将代扣执行结果以json方式返回
					_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "代扣交易提交成功！"}
				} else {
					_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "代扣订单写入数据库失败！"}
				}
			} else {
				_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "代扣下单失败！"}
			}
		} else {
			_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "出金未签约不能交易！"}
		}
	}
	ErrorHelper.LogInfo("【"+_UserId+"】智能订单执行：从"+_job_order_model.Bankcardno+_job_order_model.Bankcardno+"到"+_job_order_model.Merchantbankaccountno+"===结果==%+v", _json_model)
}

/*
自动还款(代付)
@_UserId 用户id
@_OrderNoA 订单出金编号
@_BankCardNo 出金卡号
@_MerchantBankAccountNo 入金卡号
@_AmountB 出金金额
*/
func Quick_Pay_Ok(cardModelA paytool.CardModel, _UserId, _OrderNoA, _BankCardNo, _MerchantBankAccountNo string, _AmountB float64) {
	_json_model := make(map[string]interface{}, 0)
	//根据用户id和银行卡号取用户和卡的相关信息
	_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _MerchantBankAccountNo} //这里是入金卡 臻方便费率应该从出金卡扣除
	results, _ := DbHelper.MySqlDb().Get(_model)
	if results {
		//信用卡入金-----------------------------------------
		cardModelB := paytool.CardModel{
			AccountNumber: _MerchantBankAccountNo, //入金卡号
			Phone:         _model.Mobile,          //手机号
			HolderName:    _model.Cardholder,      //姓名
			IdCard:        _model.IdCard,          //身份证号
			Cvv2:          _model.Cvv2,            //Cvv2
			Expired:       _model.Expiretime,      //有效期
		}
		fmt.Printf("%v", cardModelB)
		_channelType := "ffqrdh" //通道标示小额channelType：ffqrdh大额channelType：ffkj
		//if _AmountB > 1000 {
		_channelType = "fk"
		//}
		_extraFee := 1.0
		payTool := fengfu.FengFuPay{}
		//入金操作之前必须保证出金订单已经完成
	L_on_job:
		chujinOrder := payTool.CardOutOrder(cardModelA, _OrderNoA) //出金操作
		if chujinOrder.Content.OrderStatus == "00" {
			rtModel := payTool.CardIn(cardModelB, _channelType, _AmountB, _extraFee) //入金操作
			_str_rt := fmt.Sprintf("%+v", rtModel)
			_update_model_order := shop.UserCardJobOrder{
				UserId:      _UserId,                        //用户id
				Ordernob:    rtModel.Content.MerOrderNumber, //订单编号
				Rateamountb: int(_extraFee * 100),           //代扣手续费
				ReturnB:     "丰付智能代还-" + _str_rt,            //代扣返回信息
				IsFinish:    1,                              //订单完成状态
			}
			_count, err := DbHelper.MySqlDb().Where("OrderNoA=?", _OrderNoA).Update(_update_model_order) //修改智能代还订单
			ErrorHelper.CheckErr(err)
			//fmt.Println(_count)
			if _count > 0 && err == nil {
				if rtModel.Code == "00" {
					_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "代付交易提交成功！"}
				} else {
					_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "代付交易失败！"}
				}
			}
		} else {
			//time.Sleep(time.Second * 60)
			goto L_on_job
		}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "入金未签约不能交易！"}
	}
	ErrorHelper.LogInfo("【"+_UserId+"】智能订单【"+_OrderNoA+"】代付执行：从"+_BankCardNo+"到"+_MerchantBankAccountNo+"===结果==%+v", _json_model)
}

/*
*根据银行卡号获取获取绑卡信息
 */
func Get_Bank_Model(_Bank_No string) *shop.UserBankCard {
	_rt := &shop.UserBankCard{}
	_model_new := &shop.UserBankCard{BankCardNumber: _Bank_No}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		_rt = _model_new
	}
	return _rt
}

/*
*信用卡分润处理
 */
func FunRun(user_id, order_no, amount, mark string) {
	/* 	type _rt struct {
	   		code int
	   		msg  string
	   	}
	   	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	   	_BodyData := map[string]interface{}{"user_id": user_id, "order_no": order_no, "type": "2", "amount": amount, "mark": mark}
	   	_http_url := "https://shop.xhdncppf.com/index.php?module=app&action=index&store_id=8&app=calc_profit"
	   	_req := WebHelper.HttpPost(_http_url, _HeaderData, _BodyData)
	   	err := json.Unmarshal([]byte(_req), &_rt{})
	   	ErrorHelper.CheckErr(err) */
}
