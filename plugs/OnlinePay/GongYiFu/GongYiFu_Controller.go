package gongyifuF

import (
	"fmt"
	"strconv"
	"time"
	GongYiFu "github.com/qnsoft/online_pay/utils/pay/gongyifu"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/web_api/models/shop"
	date "github.com/qnsoft/utils/DateHelper"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/StringHelper"
)

/**
*支付——工易付控制器实体
 */
type GongYiFu_Controller struct {
	Token.BaseController
}

var (
	Api_notifyUrl string = "https://api1.xhdncppf.com" //异步回调地址 http://testapi.ljz789.com
	_SubMerchId   string = ""                          //由101接口获取
)

/*
工易付-S101.商户信息提交
*/
func (this *GongYiFu_Controller) Gyf_101() {
	//_user_id := this.GetString("UserId", "")
	//fmt.Println(_user_id)
	//首次使用该接口
	_model_101 := GongYiFu.Model_101{
		MerchName: this.GetString("MerchName", "河南联九舟网络科技有限公司"), //商户名称
		// Name:         "李金领",                //商户姓名
		// Phone:        "15638905677",        //预留电话
		// IdNo:         "41012219710510305X", //身份证号
		// CardId:       "6226223001466633",   //绑定结算卡号 民生银行
		Name:         this.GetString("Name"),   //法人姓名
		Phone:        this.GetString("Phone"),  //法人电话
		IdNo:         this.GetString("IdNo"),   //身份证号
		CardId:       this.GetString("CardId"), //交易卡号
		MerchAddress: "郑州市郑汴路环球大厦B座2303",       //商户地址
		FeeRate:      "45",                     //费率0.68%
		ExternFee:    "100",                    //手续费单位分
		Remark:       "",
	}
	_model_rt_101 := GongYiFu.Gyf_S101(_model_101)
	_SubMerchId = _model_rt_101.Data.SubMerchId
	//将_user_id  _SubMerchId 关系 存入
	ErrorHelper.LogInfo("工易付-S101.商户信息提交返回：", _model_rt_101)
	this.Data["json"] = _model_rt_101
	this.ServeJSON()
}

/*
工易付-S102.商户信息修改
*/
func (this *GongYiFu_Controller) Gyf_102() {
	_model_102 := GongYiFu.Model_102{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId),    //商户编号
		Phone:      this.GetString("Phone", "15638905677"),       //手机号
		CardId:     this.GetString("CardId", "6226223001466633"), //结算卡号
	}
	_model_rt_102 := GongYiFu.Gyf_S102(_model_102)
	ErrorHelper.LogInfo("工易付-S102.商户信息修改返回：", _model_rt_102)
	this.Data["json"] = _model_rt_102
	this.ServeJSON()
}

/*
工易付-S103.商户信息费率
*/
func (this *GongYiFu_Controller) Gyf_103() {
	_model_103 := GongYiFu.Model_103{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId), //商户编号
		FeeRate:    this.GetString("FeeRate", "32"),           //交易费率
		ExternFee:  this.GetString("ExternFee", "50"),         //附加费，单位为分
	}
	_model_rt_103 := GongYiFu.Gyf_S103(_model_103)
	ErrorHelper.LogInfo("工易付-S103.商户信息费率返回：", _model_rt_103)
	this.Data["json"] = _model_rt_103
	this.ServeJSON()
}

/*
工易付-S201.银联页面绑卡
*/
func (this *GongYiFu_Controller) Gyf_201() {
	_OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	ErrorHelper.LogInfo("绑卡订单号：", _OrderId)
	_model_201 := GongYiFu.Model_201{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId),                 //商户编号
		Name:       this.GetString("Name"),                                    //法人姓名
		Phone:      this.GetString("Phone"),                                   //法人电话
		IdNo:       this.GetString("IdNo"),                                    //身份证号
		CardId:     this.GetString("CardId"),                                  //交易卡号
		NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S401",              //异步通知地址
		FrontUrl:   Api_notifyUrl + "/api/pay/gongyifu/Gyf_201_Ok",            //页面通知地址
		OrderId:    _OrderId,                                                  //请求流水号
		DeviceId:   this.GetString("DeviceId", StringHelper.GetRandomNum(15)), //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
		IpAddress:  this.GetString("IpAddress", this.Ctx.Request.RemoteAddr),  //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
	}
	_model_rt_201 := GongYiFu.Gyf_S201(_model_201)
	ErrorHelper.LogInfo("工易付-S201.银联页面绑卡返回：", _model_rt_201)
	if len(_model_rt_201.Data.Html) > 0 {
		this.Data["qianyue_json"] = _model_rt_201
		this.TplName = "gongyifu/qianyue.html"
	} else {
		this.Data["json"] = _model_rt_201
		this.ServeJSON()
	}
}

/*
签约成功通知页面
*/
func (this *GongYiFu_Controller) Gyf_201_Ok() {
	this.Data["qianyue_ok"] = nil
	this.TplName = "gongyifu/qianyue_ok.html"
}

/*
工易付-S202.绑卡查询
*/
func (this *GongYiFu_Controller) Gyf_202() {
	_model_202 := GongYiFu.Model_202{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId),    //商户编号
		CardId:     this.GetString("CardId", "6225768311530149"), //交易卡号
	}
	_model_rt_202 := GongYiFu.Gyf_S202(_model_202)
	ErrorHelper.LogInfo("工易付-S202.绑卡查询返回：", _model_rt_202)
	this.Data["json"] = _model_rt_202
	this.ServeJSON()
}

/*
工易付-S301.订单查询
*/
func (this *GongYiFu_Controller) Gyf_301() {
	_model_301 := GongYiFu.Model_301{
		OrderId:    this.GetString("OrderId", "1910091448252077"),
		SubMerchId: this.GetString("SubMerchId", _SubMerchId), //商户编号
	}
	_model_rt_301 := GongYiFu.Gyf_S301(_model_301)
	ErrorHelper.LogInfo("工易付-S301.订单查询返回：", _model_rt_301)
	this.Data["json"] = _model_rt_301
	this.ServeJSON()
}

/*
工易付-S302.快捷支付
*/
func (this *GongYiFu_Controller) Gyf_302() {
	_OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	_model_302 := GongYiFu.Model_302{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId),                 //商户编号                                          //子商户号
		OrderId:    _OrderId,                                                  //订单号
		Name:       this.GetString("Name"),                                    //法人姓名
		Phone:      this.GetString("Phone"),                                   //法人电话
		IdNo:       this.GetString("IdNo"),                                    //身份证号
		CardId:     this.GetString("CardId"),                                  //交易卡号
		NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S402?type=302",     //异步通知地址
		Amount:     this.GetString("Amount", "500"),                           //交易金额 单位（分）
		GoodsName:  "代扣订单-" + _OrderId,                                        //订单名称
		CardType:   this.GetString("CardType", "2"),                           //卡类型 01 借记卡 02 贷记卡
		Cvv:        this.GetString("Cvv", "999"),                              //安全码
		ExpDate:    this.GetString("ExpDate", "1120"),                         //有效期
		DeviceId:   this.GetString("DeviceId", StringHelper.GetRandomNum(15)), //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
		IpAddress:  this.GetString("IpAddress", this.Ctx.Request.RemoteAddr),  //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
	}
	_model_rt_302 := GongYiFu.Gyf_S302(_model_302)
	ErrorHelper.LogInfo("工易付-S302.快捷支付返回", _model_rt_302)
	this.Data["json"] = _model_rt_302
	this.ServeJSON()
}

/*
工易付-S302.落地云闪付支付接口
*/
func (this *GongYiFu_Controller) Gyf_302_A() {
	_OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	ErrorHelper.LogInfo("云闪付支付交易订单号：", _OrderId)
	_model_302_A := GongYiFu.Model_302_A{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId),                 //子商户号
		OrderId:    _OrderId,                                                  //订单号
		Name:       this.GetString("Name"),                                    //法人姓名
		Phone:      this.GetString("Phone"),                                   //法人电话
		IdNo:       this.GetString("IdNo"),                                    //身份证号
		CardId:     this.GetString("CardId"),                                  //交易卡号
		NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S402?type=302A",    //异步通知地址
		Amount:     this.GetString("Amount", "500"),                           //交易金额 单位（分）
		GoodsName:  "代扣订单-" + _OrderId,                                        //订单名称
		CardType:   this.GetString("CardType", "2"),                           //卡类型 01 借记卡 02 贷记卡
		Cvv:        this.GetString("Cvv", "999"),                              //安全码
		ExpDate:    this.GetString("ExpDate", "1120"),                         //有效期
		DeviceId:   this.GetString("DeviceId", StringHelper.GetRandomNum(15)), //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
		IpAddress:  this.GetString("IpAddress", this.Ctx.Request.RemoteAddr),  //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
	}
	_model_rt_302_A := GongYiFu.Gyf_S302_A(_model_302_A)
	ErrorHelper.LogInfo("工易付-S302.落地云闪付支付接口返回", _model_rt_302_A)
	this.Data["json"] = _model_rt_302_A
	this.ServeJSON()
}

/*
工易付-S302.子商户结算接口
*/
func (this *GongYiFu_Controller) Gyf_302_B() {
	_OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	ErrorHelper.LogInfo("子商户结算订单号：", _OrderId)
	_model_302_B := GongYiFu.Model_302_B{
		SubMerchId: this.GetString("SubMerchId", _SubMerchId),              //子商户号
		OrderId:    _OrderId,                                               //订单号
		Name:       this.GetString("Name"),                                 //法人姓名
		Phone:      this.GetString("Phone"),                                //法人电话
		IdNo:       this.GetString("IdNo"),                                 //身份证号
		CardId:     this.GetString("CardId"),                               //交易卡号
		NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S402?type=302B", //异步通知地址
		Amount:     this.GetString("Amount", "50"),                         //交易金额 单位（分）
		Remark:     "",                                                     //备注
	}
	_model_rt_302_B := GongYiFu.Gyf_S302_B(_model_302_B)
	ErrorHelper.LogInfo("工易付-S302.子商户结算接口返回", _model_rt_302_B)
	this.Data["json"] = _model_rt_302_B
	this.ServeJSON()
}

/*
工易付-S302.子商户余额查询
*/
func (this *GongYiFu_Controller) Gyf_302_C() {
	_OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	ErrorHelper.LogInfo("子商户结算订单号：", _OrderId)
	_model_302_C := GongYiFu.Model_302_C{
		SubMerchId: this.GetString("SubMerchId"), //子商户号
	}
	_model_rt_302_C := GongYiFu.Gyf_S302_C(_model_302_C)
	ErrorHelper.LogInfo("工易付-S302.子商户余额查询返回", _model_rt_302_C)
	this.Data["json"] = _model_rt_302_C
	this.ServeJSON()
}

//////----------------异步回调处理-------------------------///////////////
/*
绑卡异步通知
*/
func (this *GongYiFu_Controller) Gyf_S401() {
	_OrderId := this.GetString("orderId")
	_CardId := this.GetString("cardId")
	_Token := this.GetString("token")
	_RespCode := this.GetString("respCode")
	_RespDesc := this.GetString("respDesc")
	_OrderTime := this.GetString("orderTime")
	_Sign := this.GetString("sign")
	_json := map[string]interface{}{"orderId": _OrderId, "cardId": _CardId, "token": _Token, "respCode": _RespCode, "respDesc": _RespDesc, "orderTime": _OrderTime, "sign": _Sign}
	ErrorHelper.LogInfo("异步：==>>工易付绑卡通知：", _json)
	_model := shop.UserBankCard{BankCardNumber: _CardId}
	results, _ := DbHelper.MySqlDb().Get(&_model)
	if results {
		_model.GyfStauts = _Sign
		_model.GyfToken = _Token
		//卡签约结果保存在数据库
		results_update, _ := DbHelper.MySqlDb().Id(_model.Id).Update(&_model)
		fmt.Println(results_update)
	} else {
		_json = map[string]interface{}{"code": 0, "msg": "fail", "info": _CardId + "-该卡还没有绑定！"}
	}
	ErrorHelper.LogInfo("异步：==>>工易付", _CardId, "-卡绑定结果：", _json)
	this.Data["json"] = _json
	this.ServeJSON()
}

/*
S402.交易异步通知
*/
func (this *GongYiFu_Controller) Gyf_S402() {
	//以下是本平台参数
	_type := this.GetString("type")
	//以下是工易付参数
	_status := this.GetString("status")
	_OrderId := this.GetString("orderId")
	_respCode := this.GetString("respCode")
	_json := map[string]interface{}{"orderId": _OrderId, "status": _status, "respCode": _respCode, "respDesc": this.GetString("respDesc"), "orderTime": this.GetString("orderTime"), "sign": this.GetString("sign")}
	if _type == "302A" {
		_OrderId_df := "DH" + date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4) //开始准备代付订单号
		switch _status {
		case "02":
			//ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]：", "执行代扣成功！02", _json)

			//先处理智还订单执行
			_model_job_order := shop.UserCardJobOrder{ //根据自动订单号获取订单继续执行--------------自动动订单
				Ordernoa: _OrderId, //代付订单编号
			}
			_order_job_ok, err := DbHelper.MySqlDb().Get(&_model_job_order) //获取代扣订单
			ErrorHelper.CheckErr(err)
			if _order_job_ok { //如果代扣订单存在 继续往下执行代付
				_model_user_auth := shop.UserAuth{
					UserId: _model_job_order.UserId,
				}
				_user_auth_ok, err := DbHelper.MySqlDb().Get(&_model_user_auth) //获取商户工易付签约信息
				ErrorHelper.CheckErr(err)
				if _user_auth_ok {
					_str_rt_302_A := fmt.Sprintf("%+v", _json)
					_model_job_order.ReturnA = "工易付智还代扣-" + _str_rt_302_A //修改当前订单状态内容
					_model_job_order.Ordernob = _OrderId_df
					results_job_order_update, _ := DbHelper.MySqlDb().Id(_model_job_order.Id).Update(&_model_job_order)
					//fmt.Println(results_update)
					if results_job_order_update > 0 {
						ErrorHelper.LogInfo("执行智能代扣成功后-执行智能代付提交：", _model_user_auth.Submerchid, _model_job_order.Bankcardno, _model_job_order.Merchantbankaccountno, strconv.Itoa(_model_job_order.Amount/100), _OrderId_df)
						Gyf_DaiFu(_model_user_auth.Submerchid, _model_job_order.Bankcardno, _model_job_order.Merchantbankaccountno, strconv.Itoa(_model_job_order.Amount/100), _OrderId, _OrderId_df) //开始执行代付
					}
				} else {
					ErrorHelper.LogInfo("异步：==>>工易付[" + _type + "]智能代扣执行失败，商户为签约！：" + _OrderId)
				}
			} else {
				_model_order := shop.UserCardOrder{ //根据订单号获取订单继续执行--------------手动订单
					Ordernoa: _OrderId, //代付订单编号
				}
				_order_ok, err := DbHelper.MySqlDb().Get(&_model_order) //获取代扣订单
				ErrorHelper.CheckErr(err)
				if _order_ok { //如果代扣订单存在 继续往下执行代付
					_model_user_auth := shop.UserAuth{
						UserId: _model_order.UserId,
					}
					_user_auth_ok, err := DbHelper.MySqlDb().Get(&_model_user_auth) //获取商户工易付签约信息
					ErrorHelper.CheckErr(err)
					if _user_auth_ok {
						_str_rt_302_A := fmt.Sprintf("%+v", _json)
						_model_order.ReturnA = "工易付收款代扣-" + _str_rt_302_A //修改当前订单状态内容
						_model_order.Ordernob = _OrderId_df
						results_order_update, _ := DbHelper.MySqlDb().Id(_model_order.Id).Update(&_model_order)
						//fmt.Println(results_update)
						if results_order_update > 0 {
							ErrorHelper.LogInfo("执行代扣成功后手动代付执行提交：", _model_user_auth.Submerchid, _model_order.Bankcardno, _model_order.Merchantbankaccountno, strconv.Itoa(_model_order.Amount/100), _OrderId_df)
							Gyf_DaiFu(_model_user_auth.Submerchid, _model_order.Bankcardno, _model_order.Merchantbankaccountno, strconv.Itoa(_model_order.Amount/100), _OrderId, _OrderId_df) //开始执行代付
						}
					} else {
						ErrorHelper.LogInfo("异步：==>>工易付[" + _type + "]手动代扣执行失败，302A-02原因是没查到订单：" + _OrderId)
					}
				}
			}
		case "03":
			//ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]：", "执行代扣失败！03", _json)
			//修改数据库返回订单记录状态
			//开始执行自动订单
			_model_job_order := shop.UserCardJobOrder{ //根据订单号获取订单继续执行---智还订单
				Ordernoa: _OrderId, //代付订单编号
			}
			_order_job_ok, err := DbHelper.MySqlDb().Get(&_model_job_order) //获取代扣订单
			ErrorHelper.CheckErr(err)
			if _order_job_ok {
				_str_rt_302_A := fmt.Sprintf("%+v", _json)
				_model_job_order.ReturnA = "工易付智还代扣-" + _str_rt_302_A //修改当前订单状态内容
				_model_job_order.IsFinish = 1
				results_order_job_update, _ := DbHelper.MySqlDb().Id(_model_job_order.Id).Update(&_model_job_order)
				if results_order_job_update > 0 {
					ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]", "执行代扣失败！", _json)
				}
			} else {
				_model_order := shop.UserCardOrder{ //根据订单号获取订单继续执行---手动订单
					Ordernoa: _OrderId, //代付订单编号
				}
				_order_ok, err := DbHelper.MySqlDb().Get(&_model_order) //获取代扣订单
				ErrorHelper.CheckErr(err)
				if _order_ok {
					_str_rt_302_A := fmt.Sprintf("%+v", _json)
					_model_order.ReturnA = "工易付收款代扣-" + _str_rt_302_A //修改当前订单状态内容
					results_order_update, _ := DbHelper.MySqlDb().Id(_model_order.Id).Update(&_model_order)
					if results_order_update > 0 {
						ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]", "执行代扣失败！", _json)
					}
				}
			}
		default:
			ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]：", "执行代扣交易处理中！", _json)
		}

	} else if _type == "302B" {
		switch _status {
		case "02":
			ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]", "执行代付成功！！！！！", _OrderId)
			//修改数据库返回订单记录状态
			// _model_job_order := shop.UserCardJobOrder{ //根据订单号获取订单继续执行---智还订单
			// 	Ordernob: _rb, //代付订单编号
			// }
			// _order_job_ok, err := DbHelper.MySqlDb().Get(&_model_job_order) //获取代扣订单
			// ErrorHelper.CheckErr(err)
			// if _order_job_ok {
			// 	_str_rt_302_B := fmt.Sprintf("%+v", _json)
			// 	_model_job_order.ReturnB = "工易付智还代付-" + _str_rt_302_B
			// 	_model_job_order.IsFinish = 1
			// 	results_job_order_update, _ := DbHelper.MySqlDb().Id(_model_job_order.Id).Update(&_model_job_order)
			// 	if results_job_order_update > 0 {
			// 		ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]", "执行智能代付成功！", _json)
			// 	}
			// } else {
			// 	_model_order := shop.UserCardOrder{ //根据订单号获取订单继续执行
			// 		Ordernob: _rb, //代付订单编号
			// 	}
			// 	_order_ok, err := DbHelper.MySqlDb().Get(&_model_order) //获取代扣订单
			// 	ErrorHelper.CheckErr(err)
			// 	if _order_ok {
			// 		_str_rt_302_B := fmt.Sprintf("%+v", _json)
			// 		_model_order.ReturnB = "工易付收款代付-" + _str_rt_302_B
			// 		results_order_update, _ := DbHelper.MySqlDb().Id(_model_order.Id).Update(&_model_order)
			// 		if results_order_update > 0 {
			// 			ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]", "执行代付成功！", _json)
			// 		}
			// 	}
			// }

		case "03":
			//ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]", "执行代付失败！", _json)
		default:
			ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]：", "执行代付交易处理中！", _json)
		}
	} else {
		ErrorHelper.LogInfo("异步：==>>工易付["+_type+"]：", "其它交易！", _json)
	}
	this.Data["json"] = _json
	this.ServeJSON()
}

/*
S501.地区编码查询
*/
func (this *GongYiFu_Controller) Gyf_S501() {
	str_501 := GongYiFu.Gyf_S501()
	//fmt.Print("%+v", str_501)
	this.Data["json"] = str_501
	this.ServeJSON()
}

////-------------开始与数据库对接-----------------////////////
/*
检测商户是否注册
*/
func (this *GongYiFu_Controller) Gyf_IsZhuCe() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")
	_Name := this.GetString("Name")                         //法人姓名
	_Phone := this.GetString("Phone")                       //法人电话
	_IdNo := this.GetString("IdNo")                         //身份证号
	_CardId := this.GetString("CardId", "6225768311530149") //交易卡号
	fmt.Println(_UserId)
	_model := shop.UserAuth{UserId: _UserId}
	results, _ := DbHelper.MySqlDb().Get(&_model)
	if results && len(_model.Submerchid) > 3 {
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": _model, "info": "商户已注册！"}
	} else {
		//开始请求工易付接口
		_model_101 := GongYiFu.Model_101{
			MerchName:    this.GetString("MerchName", "河南联九舟网络科技有限公司"), //商户名称
			Name:         _Name,                                        //法人姓名
			Phone:        _Phone,                                       //法人电话
			IdNo:         _IdNo,                                        //身份证号
			CardId:       _CardId,                                      //交易卡号
			MerchAddress: "郑州市郑汴路环球大厦C座1007",                           //商户地址
			FeeRate:      "36",                                         //费率0.36%
			ExternFee:    "100",                                        //手续费单位分
			Remark:       _Name + "的注册信息",
		}
		_model_rt_101 := GongYiFu.Gyf_S101(_model_101)
		_model.Submerchid = _model_rt_101.Data.SubMerchId
		if len(_model.Submerchid) > 3 {
			//商户注册结果保存在数据库
			results_update, _ := DbHelper.MySqlDb().Id(_model.Id).Update(&_model)
			//fmt.Println(results_update)
			if results_update > 0 {
				// _model_103 := GongYiFu.Model_103{
				// 	SubMerchId: _model.Submerchid, //商户编号
				// 	FeeRate:    "36",              //交易费率
				// 	ExternFee:  "100",             //附加费，单位为分
				// }
				// _model_rt_103 := GongYiFu.Gyf_S103(_model_103)
				// ErrorHelper.LogInfo("商户"+_Name+"["+_SubMerchId+"]签约费率信息费率返回：", _model_rt_103)
				_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": _model, "info": "商户注册完成！"}
			} else {
				_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "商户注册失败！！"}
			}
		} else {
			_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "商户注册失败！"}
		}
	}
	ErrorHelper.LogInfo(_UserId, "-商户注册结果：", _json_model)
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
检测用户卡是否签约
*/
func (this *GongYiFu_Controller) Gyf_IsQianYue() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")
	_CardId := this.GetString("CardId") //交易卡号
	_model := shop.UserBankCard{UserId: _UserId, BankCardNumber: _CardId}
	results, _ := DbHelper.MySqlDb().Get(&_model)
	if results && len(_model.GyfStauts) > 3 && len(_model.GyfToken) > 3 {
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "info": "该卡已签约"}

	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "该卡未签约"}
	}
	//ErrorHelper.LogInfo(_UserId, "-签约请求验证", _json_model)
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
检测用户卡是否签约HTML
*/
func (this *GongYiFu_Controller) Gyf_IsQianYueHtml() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")
	_CardId := this.GetString("CardId") //交易卡号
	_Name := this.GetString("Name")     //法人姓名
	_Phone := this.GetString("Phone")   //法人电话
	_IdNo := this.GetString("IdNo")     //身份证号
	_OrderId := date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4)
	_model := shop.UserBankCard{UserId: _UserId, BankCardNumber: _CardId}
	results, _ := DbHelper.MySqlDb().Get(&_model)
	if results && len(_model.GyfStauts) > 3 && len(_model.GyfToken) > 3 {
		_json_model = map[string]interface{}{"Code": 000, "Data": map[string]string{"Html": "<html><head></head><body><a href=\"http://www.xhdncppf.com/xyk/index.html\">该卡已签约！</a></body></html>"}}
		this.Data["qianyue_json"] = _json_model
		this.TplName = "gongyifu/qianyue.html"
	} else {
		//ErrorHelper.LogInfo("绑卡订单号：", _OrderId)
		_model_201 := GongYiFu.Model_201{
			SubMerchId: this.GetString("SubMerchId", _SubMerchId),                 //商户编号
			Name:       _Name,                                                     //法人姓名
			Phone:      _Phone,                                                    //法人电话
			IdNo:       _IdNo,                                                     //身份证号
			CardId:     _CardId,                                                   //交易卡号
			NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S401",              //异步通知地址
			FrontUrl:   Api_notifyUrl + "/api/pay/gongyifu/Gyf_201_Ok",            //页面通知地址
			OrderId:    _OrderId,                                                  //请求流水号
			DeviceId:   this.GetString("DeviceId", StringHelper.GetRandomNum(15)), //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
			IpAddress:  this.GetString("IpAddress", this.Ctx.Request.RemoteAddr),  //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
		}
		_model_rt_201 := GongYiFu.Gyf_S201(_model_201)
		//ErrorHelper.LogInfo("工易付-S201.银联页面绑卡返回：", _model_rt_201)
		if len(_model_rt_201.Data.Html) > 0 {
			this.Data["qianyue_json"] = _model_rt_201
			this.TplName = "gongyifu/qianyue.html"
		} else { //如果没有返回银联绑卡html则清空SubMerchId
			// _model := shop.UserAuth{UserId: _UserId}
			// _model.Submerchid = ""
			// results_update, _ := DbHelper.MySqlDb().Id(_model.Id).Update(&_model)
			// if results_update > 0 {
			// }
			this.Data["qianyue_json"] = _json_model
			this.TplName = "gongyifu/qianyue.html"
		}
	}
}

/*
根据id查询用户商户号
*/
func (this *GongYiFu_Controller) Gyf_Submerchid() {
	var _json_model interface{}
	_UserId := this.GetString("UserId") //用户id
	_model := shop.UserAuth{UserId: _UserId}
	_results, _ := DbHelper.MySqlDb().Get(_model)
	if _results {
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "Submerchid": _model.Submerchid, "info": "获取商户id成功！"}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取商户id失败！"}
	}
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
工易付云闪付代扣(在臻方便调取的条件为Amount<=1000,SubMerchId不为空,CardId必须签约))
*/
func (this *GongYiFu_Controller) Gyf_DaiKou() {
	_CardId := this.GetString("CardId") //代扣卡号
	_Amount := this.GetString("Amount") //交易金额 这里接收的是元
	_Amount_float64, _ := strconv.ParseFloat(_Amount, 64)
	_SubMerchId_New := this.GetString("SubMerchId")                                               //子商户号
	_CardId_To := this.GetString("CardId_To")                                                     //代付卡号
	_OrderId := "DK" + date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4) //代扣订单号
	//_OrderId_df := "DH" + date.FormatDate(time.Now(), "yyMMddHHmmss") + StringHelper.GetRandomNum(4) //代付订单号
	//ErrorHelper.LogInfo("云闪付支付交易订单号：", _OrderId)
	_model := shop.UserBankCard{BankCardNumber: _CardId}
	results, _ := DbHelper.MySqlDb().Get(&_model)
	if results && len(_model.GyfStauts) > 3 && len(_model.GyfToken) > 3 {
		_model_302_A := GongYiFu.Model_302_A{
			SubMerchId: _SubMerchId_New,                                               //子商户号
			OrderId:    _OrderId,                                                      //订单号
			Name:       _model.Cardholder,                                             //法人姓名
			Phone:      _model.Mobile,                                                 //法人电话
			IdNo:       _model.IdCard,                                                 //身份证号
			CardId:     _CardId,                                                       //交易卡号
			NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S402?type=302A",        //异步通知地址
			Amount:     strconv.FormatFloat(float64(_Amount_float64*100), 'f', 0, 64), //交易金额 单位（分）
			GoodsName:  "代扣-" + _OrderId,                                              //订单名称
			CardType:   this.GetString("CardType", "2"),                               //卡类型 01 借记卡 02 贷记卡
			Cvv:        this.GetString("Cvv", _model.Cvv2),                            //安全码
			ExpDate:    this.GetString("ExpDate", _model.Expiretime),                  //有效期
			DeviceId:   this.GetString("DeviceId", StringHelper.GetRandomNum(15)),     //设备ID 安卓:IMEI，iOS:IDFV，PC:硬盘序列号（若不填大额交易限额会被银联风控）
			IpAddress:  this.GetString("IpAddress", this.Ctx.Request.RemoteAddr),      //请求IP地址 公网IP地址（若不填大额交易限额会被风控）（付款客户端IP）
			CityCode:   "4910",                                                        //城市编码 写死到河南省
		}
		_model_rt_302_A := GongYiFu.Gyf_S302_A(_model_302_A)
		//_model_rt_302_A := ""
		//	ErrorHelper.LogInfo("工易付卡号：", _CardId+"-代扣执行提交：", _model_302_A, "返回：", _model_rt_302_A)
		//开始将代扣订单信息写入到信用卡订单表[lkt_user_card_order]
		_str_rt_302_A := fmt.Sprintf("%+v", _model_rt_302_A)
		_, _shouxufei := Gyf_Fl(_model.UserId) //获取当前卡费率
		_new_model_order := shop.UserCardOrder{
			UserId:                _model.UserId,              //用户id
			Ordernoa:              _OrderId,                   //代扣订单编号
			Ordernob:              "000000",                   //代付订单编号
			Amount:                int(_Amount_float64 * 100), //订单金额
			Rateamounta:           int(_shouxufei * 100),      //代扣手续费
			Bankcardno:            _CardId,                    //出金卡号
			Merchantbankaccountno: _CardId_To,                 //入金卡号
			ReturnA:               "工易付收款代扣-" + _str_rt_302_A, //代扣返回信息
			AddTime:               time.Now(),                 //订单创建时间
		}
		_, err := DbHelper.MySqlDb().Insert(_new_model_order)
		ErrorHelper.CheckErr(err)
		this.Data["json"] = _model_rt_302_A
		this.ServeJSON()
	} else {
		ErrorHelper.LogInfo("工易付卡号：", _CardId+"-代扣执行结果：因未签约执行失败！")
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": "fail", "info": "该卡未签约或未绑定！"}
		this.ServeJSON()
	}
}

/*
工易付代付
*/
func Gyf_DaiFu(_SubMerchId, _CardId, _CardId_To, _Amount, _OrderId, _OrderId_df string) {
	_model_bank_card := shop.UserBankCard{BankCardNumber: _CardId_To}
	results_bank_card, _ := DbHelper.MySqlDb().Get(&_model_bank_card)
	if results_bank_card {
		_feilv, _shouxufei := Gyf_Fl(_model_bank_card.UserId) //获取当前卡费率
		_Amount_float64, _ := strconv.ParseFloat(_Amount, 64)
		_Amount_new := float64(_Amount_float64*100) - float64(_Amount_float64*100*_feilv) - _shouxufei //计算出代付金额 代付金额=总金额-费用-手续费
		_Amount_to := strconv.FormatFloat(_Amount_new, 'f', 0, 64)                                     //转为字符串
		_model_302_B := GongYiFu.Model_302_B{
			SubMerchId: _SubMerchId,                                                              //子商户号
			OrderId:    _OrderId_df,                                                              //订单号
			Name:       _model_bank_card.Cardholder,                                              //法人姓名
			Phone:      _model_bank_card.Mobile,                                                  //法人电话
			IdNo:       _model_bank_card.IdCard,                                                  //身份证号
			CardId:     _CardId_To,                                                               //交易卡号
			NotifyUrl:  Api_notifyUrl + "/api/pay/gongyifu/Gyf_S402?type=302B&rb=" + _OrderId_df, //异步通知地址
			Amount:     _Amount_to,                                                               //交易金额 单位（分）
			Remark:     "代付" + _OrderId_df,                                                       //备注
		}
		_model_rt_302_B := GongYiFu.Gyf_S302_B(_model_302_B)
		//fmt.Println(_model_rt_302_B)
		//修改订单状态
		//开始执行智还自动订单
		//修改数据库返回订单记录状态
		_model_job_order := shop.UserCardJobOrder{ //根据订单号获取订单继续执行---智还订单
			Ordernoa: _OrderId, //代付订单编号
		}
		_order_job_ok, err := DbHelper.MySqlDb().Get(&_model_job_order) //获取代扣订单
		ErrorHelper.CheckErr(err)
		if _order_job_ok {
			_str_rt_302_B := fmt.Sprintf("%+v", _model_rt_302_B)
			_model_job_order.Ordernob = _OrderId_df
			//_model_job_order.ReturnB = "工易付智还代付-" + _str_rt_302_B //修改当前订单状态内容
			_model_job_order.ReturnB = "工易付智还代付-提交成功！"
			_model_job_order.IsFinish = 1
			results_job_order_update, _ := DbHelper.MySqlDb().Id(_model_job_order.Id).Update(&_model_job_order)
			if results_job_order_update > 0 {
				ErrorHelper.LogInfo("同步：==>>工易付["+_CardId+"]", "智能代付提交！", _str_rt_302_B)
			}
		} else {
			_model_order := shop.UserCardOrder{ //根据订单号获取订单继续执行
				Ordernoa: _OrderId, //代付订单编号
			}
			_order_ok, err := DbHelper.MySqlDb().Get(&_model_order) //获取代扣订单
			ErrorHelper.CheckErr(err)
			if _order_ok {
				_str_rt_302_B := fmt.Sprintf("%+v", _model_rt_302_B)
				_model_order.Ordernob = _OrderId_df
				//_model_order.ReturnB = "工易付收款代付-" + _str_rt_302_B //修改当前订单状态内容
				_model_order.ReturnB = "工易付收款代付-提交成功！"
				results_order_update, _ := DbHelper.MySqlDb().Id(_model_order.Id).Update(&_model_order)
				if results_order_update > 0 {
					ErrorHelper.LogInfo("同步：==>>工易付["+_CardId+"]", "手动代付提交！", _str_rt_302_B)
				}
			}
		}

		//ErrorHelper.LogInfo("工易付代付提交：", _model_302_B, "返回：", _model_rt_302_B)
	}
}

/*
工易付费率查询,暂时写死
*/
func Gyf_Fl(_UserId string) (float64, float64) {
	_feilv := 0.0036     //费率0.0036
	_shouxufei := 100.00 //手续费
	if _UserId == "1694" {
		_feilv = 0.0032    //最低0.0032
		_shouxufei = 50.00 //最新手续费0.5
	}
	return _feilv, _shouxufei
}
