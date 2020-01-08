package xyk

import (
	"fmt"
	"strconv"
	"time"
	"qnsoft/web_api/controllers/OnlinePay/Kuaifutong"
	"qnsoft/web_api/controllers/Token"
	"qnsoft/web_api/models/shop"
	"qnsoft/web_api/utils/DateHelper"
	"qnsoft/web_api/utils/DbHelper"
	"qnsoft/web_api/utils/ErrorHelper"
	"qnsoft/web_api/utils/StringHelper"

	_ "github.com/mattn/go-adodb"
)

/**
*信用卡控制器
 */
type Xyk_Controller struct {
	Token.BaseController
}

/*
公众号卡手动订单消费明细
*/
func (this *Xyk_Controller) Bind_User_Order_List() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId") //用户id
	_model := new(shop.UserCardOrder)
	rows, err := DbHelper.MySqlDb().Where(" user_id=? and isjob=0 ", _UserId).Desc("add_time").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	_list := make([]*shop.UserCardJobOrder, 0)
	for rows.Next() {
		_ = rows.Scan(_model)
		_model_new := new(shop.UserCardJobOrder)
		_model_new.Id = _model.Id
		//	_model_new.ParentOrderId = _model.ParentOrderId
		_model_new.UserId = _model.UserId
		_model_new.Ordernoa = _model.Ordernoa
		_model_new.Ordernob = _model.Ordernob
		_model_new.Amount = _model.Amount
		_model_new.Rateamounta = _model.Rateamounta
		_model_new.Rateamountb = _model.Rateamountb
		_model_new.Bankcardno = _model.Bankcardno
		_model_new.Merchantbankaccountno = _model.Merchantbankaccountno
		_model_new.ReturnA = _model.ReturnA
		_model_new.ReturnB = _model.ReturnB
		_model_new.UserParentaId = _model.UserParentaId
		_model_new.UserParentbId = _model.UserParentbId
		_model_new.UserParentaProfit = _model.UserParentaProfit
		_model_new.UserParentbProfit = _model.UserParentbProfit
		_model_new.AddTime = _model.AddTime
		_list = append(_list, _model_new)
	}
	_rt_json = map[string]interface{}{"code": 1, "msg": "success", "list": _list, "info": "手动明细列表获取成功！"}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
公众号卡自动订单消费明细
*/
func (this *Xyk_Controller) Bind_User_Order_JobList() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId") //用户id
	_model := new(shop.UserCardJobOrder)
	rows, err := DbHelper.MySqlDb().Where(" user_id=? and status=1 and is_finish=1", _UserId).Desc("implement_time").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	_list := make([]*shop.UserCardJobOrder, 0)
	for rows.Next() {
		_ = rows.Scan(_model)
		_model_new := new(shop.UserCardJobOrder)
		_model_new.Id = _model.Id
		_model_new.ParentOrderId = _model.ParentOrderId
		_model_new.UserId = _model.UserId
		_model_new.Ordernoa = _model.Ordernoa
		_model_new.Ordernob = _model.Ordernob
		_model_new.Amount = _model.Amount
		_model_new.Rateamounta = _model.Rateamounta
		_model_new.Rateamountb = _model.Rateamountb
		_model_new.Bankcardno = _model.Bankcardno
		_model_new.Merchantbankaccountno = _model.Merchantbankaccountno
		_model_new.ReturnA = _model.ReturnA
		_model_new.ReturnB = _model.ReturnB
		_model_new.UserParentaId = _model.UserParentaId
		_model_new.UserParentbId = _model.UserParentbId
		_model_new.UserParentaProfit = _model.UserParentaProfit
		_model_new.UserParentbProfit = _model.UserParentbProfit
		_model_new.Treatyid = _model.Treatyid
		_model_new.AddTime = _model.AddTime
		_model_new.TaskName = _model.TaskName
		_model_new.CronSpec = _model.CronSpec
		_model_new.ImplementTime = _model.ImplementTime
		_model_new.Status = _model.Status
		_model_new.IsFinish = _model.IsFinish
		_list = append(_list, _model_new)
	}
	_rt_json = map[string]interface{}{"code": 1, "msg": "success", "list": _list, "info": "自动明细列表获取成功！"}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
公众号卡计划任务自动代还订单生成列表
*/
func (this *Xyk_Controller) Bind_Card_Order_JobList() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId")         //用户id
	_BankcardNo := this.GetString("BankcardNo") //用户id
	_model := new(shop.UserCardJobOrder)
	rows, err := DbHelper.MySqlDb().Where(" user_id=? and BankCardNo=? and status=1 and is_finish=0", _UserId, _BankcardNo).Asc("implement_time").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	_list := make([]*shop.UserCardJobOrder, 0)
	for rows.Next() {
		_ = rows.Scan(_model)
		_model_new := new(shop.UserCardJobOrder)
		_model_new.Id = _model.Id
		_model_new.ParentOrderId = _model.ParentOrderId
		_model_new.UserId = _model.UserId
		_model_new.Ordernoa = _model.Ordernoa
		_model_new.Ordernob = _model.Ordernob
		_model_new.Amount = _model.Amount
		_model_new.Rateamounta = _model.Rateamounta
		_model_new.Rateamountb = _model.Rateamountb
		_model_new.Bankcardno = _model.Bankcardno
		_model_new.Merchantbankaccountno = _model.Merchantbankaccountno
		_model_new.ReturnA = _model.ReturnA
		_model_new.ReturnB = _model.ReturnB
		_model_new.UserParentaId = _model.UserParentaId
		_model_new.UserParentbId = _model.UserParentbId
		_model_new.UserParentaProfit = _model.UserParentaProfit
		_model_new.UserParentbProfit = _model.UserParentbProfit
		_model_new.Treatyid = _model.Treatyid
		_model_new.AddTime = _model.AddTime
		_model_new.TaskName = _model.TaskName
		_model_new.CronSpec = _model.CronSpec
		_model_new.ImplementTime = _model.ImplementTime
		_model_new.Status = _model.Status
		_model_new.IsFinish = _model.IsFinish
		_list = append(_list, _model_new)
	}
	_rt_json = map[string]interface{}{"code": 1, "msg": "success", "list": _list, "info": "任务列表获取成功！"}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
*取消
 */
func (this *Xyk_Controller) Bind_Card_Order_JobList_Cancel() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId")         //用户id
	_BankcardNo := this.GetString("BankcardNo") //用户id
	_model := shop.UserCardJobOrder{UserId: _UserId, Bankcardno: _BankcardNo}
	_count, err := DbHelper.MySqlDb().Where(" is_finish=0 ").Delete(&_model)
	ErrorHelper.CheckErr(err)
	fmt.Println(_count)
	_rt_json = map[string]interface{}{"code": 1, "msg": "success", "info": "取消成功！"}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
*解绑
 */
func (this *Xyk_Controller) Bind_Card_List_Cancel() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId")         //用户id
	_BankcardNo := this.GetString("BankcardNo") //用户id
	_modelA := shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankcardNo}
	_ok, err := DbHelper.MySqlDb().Get(&_modelA)
	ErrorHelper.CheckErr(err)
	if _ok {
		_kuaifutong := Kuaifutong.KuaiPayHelper{}
		//开始查询代扣协议
		_model_39 := Kuaifutong.Model_39{
			OrderNo:  "",
			TreatyNo: _modelA.Treatyid,
		}
		_rerurn_39 := _kuaifutong.Gbp_query_treaty_info(_model_39)
		//_model := shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankcardNo}
		_count, err := DbHelper.MySqlDb().Where(" id>0 ").Delete(&_modelA)
		ErrorHelper.CheckErr(err)
		fmt.Println(_count)
		if _count > 0 {
			_model_310 := Kuaifutong.Model_310{
				OrderNo:  _rerurn_39.OrderNo, //订单编号
				TreatyNo: _modelA.Treatyid,   //协议号
			}
			_rerurn_310 := _kuaifutong.Gbp_cancel_treaty_info(_model_310)
			fmt.Println(_rerurn_310)
			//	if _rerurn_310.OrderNo != "" {
			_rt_json = map[string]interface{}{"code": 1, "msg": "success", "info": "解绑成功！"}
			//	} else {
			//		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "order": _rerurn_310, "info": "解绑失败！快付通代扣协议有误！"}
			//	}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "解绑失败-数据库操作有误！"}
		}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
*提现操作(必传三要素卡号，金额，手续费)
 */
func (this *Xyk_Controller) Bind_Card_TiXian() {
	var _rt_json interface{}
	_CustBankAccountNo := this.GetString("CustBankAccountNo") //往客户的哪张卡上付钱
	_Amount, err := this.GetFloat("Amount")                   //金额（元转分）
	ErrorHelper.CheckErr(err)
	_RateAmount, err := this.GetFloat("RateAmount") //手续费（元转分）
	ErrorHelper.CheckErr(err)
	_CustBankNo := Get_Bank_Model(_CustBankAccountNo).BankCode //客户银行账户行别
	_CustName := Get_Bank_Model(_CustBankAccountNo).Cardholder //客户姓名
	_CustID := Get_Bank_Model(_CustBankAccountNo).IdCard       //客户证件号码
	str_order := date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(6)
	_model_35 := Kuaifutong.Model_35{
		OrderNo:   str_order, //订单编号 用于标识商户发起的一笔交易,在批量交易中,此编号可写在批量请求文件中,用于标识批量请求中的每一笔交易
		TradeName: "签约",      //交易名称 由商户填写,简要概括此次交易的内容.用于在查询交易记录时,提醒用户此次交易具体做了什么事情
		//MerchantBankAccountNo:    "",                                            //商户银行账号 可空 商户用于付款的银行账户，资金到账T+0模式时必填。
		//MerchantBindPhoneNo:      "",                                            //商户开户时绑定的手机号（可空）对于有些银行账户被扣款时，需要提供此绑定手机号才能进行交易；此手机号和短信通知客户的手机号可以一致，也可以不一致
		Amount:                   strconv.FormatFloat(float64(_Amount*100), 'f', 0, 64),     //交易金额 此次交易的具体金额,单位:分,不支持小数点
		CustBankNo:               _CustBankNo,                                               //客户银行账户行别 客户银行账户所属的银行的编号,具体见第5.3.1章节
		CustBankAccountIssuerNo:  "",                                                        //客户开户行网点号 可空 指支付系统里的行号，具体到某个支行（网点）号；
		CustBankAccountNo:        _CustBankAccountNo,                                        //客户银行账户号 本次交易中,往客户的哪张卡上付钱
		CustName:                 _CustName,                                                 //客户姓名 收钱的客户的真实姓名
		CustBankAcctType:         "",                                                        //客户银行账户类型 可空 指客户的银行账户是个人账户还是企业账户
		CustAccountCreditOrDebit: "",                                                        //客户账户借记贷记类型 可空 若是信用卡，则以下两个参数“信用卡有效期”和“信用卡cvv2”； 1借记 2贷记 4 未知
		CustCardValidDate:        "",                                                        //客户信用卡有效期 可空 信用卡的正下方的四位数，前两位是月份，后两位是年份；
		CustCardCvv2:             "",                                                        //客户信用卡的cvv2 可空 信用卡的背面的三位数
		CustID:                   _CustID,                                                   //客户证件号码
		CustPhone:                "",                                                        //客户手机号 如果商户购买的产品中勾选了短信通知功能，则当商户填写了手机号时,快付通会在交易成功后通过短信通知客户
		Messages:                 "",                                                        //发送客户短信内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
		CustEmail:                "",                                                        //客户邮箱地址 可空 如果商户购买的产品中勾选了邮件通知功能，则当商户填写了邮箱地址时,快付通会在交易成功后通过邮件通知客户
		EmailMessages:            "",                                                        //发送客户邮件内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
		Remark:                   "",                                                        //备注 可空 商户可额外填写备注信息,此信息会传给银行,会在银行的账单信息中显示(具体如何显示取决于银行方,快付通不保证银行肯定能显示)
		CustProtocolNo:           "",                                                        //客户协议编号 可空 扣款人在快付通备案的协议号。
		ExtendParams:             "",                                                        //扩展参数 可空 用于商户的特定业务信息传递，只有商户与快付通约定了传递此参数且约定了参数含义，此参数才有效。参数格式：参数名 1^参数值 1|参数名 2^参数值 2|……多条数据用“|”间隔注意: 不能包含特殊字符（如：#、%、&、+）、敏感词汇, 如果必须使用特殊字符,则需要自行做URL Encoding
		RateAmount:               strconv.FormatFloat(float64(_RateAmount*100), 'f', 0, 64), //商户手续费 可空 本次交易需要扣除的手续费。单位:分,不支持小数点 如不填，则手续费默认0元；
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	_return_35 := _kuaifutong.Gbp_same_id_credit_card_pay(_model_35)
	fmt.Println(_return_35)
	if _return_35.Status == "1" {
		_rt_json = map[string]interface{}{"code": 1, "msg": "success", "info": "提现成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "success", "info": _return_35}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
*提现操作(必传三要素卡号，金额，手续费)(小额)
 */
func (this *Xyk_Controller) Bind_Card_TiXian_Small() {
	var _rt_json interface{}
	_CustBankAccountNo := this.GetString("CustBankAccountNo") //往客户的哪张卡上付钱
	_Amount, err := this.GetFloat("Amount")                   //金额（元转分）
	ErrorHelper.CheckErr(err)
	_RateAmount, err := this.GetFloat("RateAmount") //手续费（元转分）
	ErrorHelper.CheckErr(err)
	_CustBankNo := Get_Bank_Model(_CustBankAccountNo).BankCode //客户银行账户行别
	_CustName := Get_Bank_Model(_CustBankAccountNo).Cardholder //客户姓名
	_CustID := Get_Bank_Model(_CustBankAccountNo).IdCard       //客户证件号码
	str_order := date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(6)
	_model_35 := Kuaifutong.Model_35{
		OrderNo:   str_order, //订单编号 用于标识商户发起的一笔交易,在批量交易中,此编号可写在批量请求文件中,用于标识批量请求中的每一笔交易
		TradeName: "签约",      //交易名称 由商户填写,简要概括此次交易的内容.用于在查询交易记录时,提醒用户此次交易具体做了什么事情
		//MerchantBankAccountNo:    "",                                            //商户银行账号 可空 商户用于付款的银行账户，资金到账T+0模式时必填。
		//MerchantBindPhoneNo:      "",                                            //商户开户时绑定的手机号（可空）对于有些银行账户被扣款时，需要提供此绑定手机号才能进行交易；此手机号和短信通知客户的手机号可以一致，也可以不一致
		Amount:                   strconv.FormatFloat(float64(_Amount*100), 'f', 0, 64),     //交易金额 此次交易的具体金额,单位:分,不支持小数点
		CustBankNo:               _CustBankNo,                                               //客户银行账户行别 客户银行账户所属的银行的编号,具体见第5.3.1章节
		CustBankAccountIssuerNo:  "",                                                        //客户开户行网点号 可空 指支付系统里的行号，具体到某个支行（网点）号；
		CustBankAccountNo:        _CustBankAccountNo,                                        //客户银行账户号 本次交易中,往客户的哪张卡上付钱
		CustName:                 _CustName,                                                 //客户姓名 收钱的客户的真实姓名
		CustBankAcctType:         "",                                                        //客户银行账户类型 可空 指客户的银行账户是个人账户还是企业账户
		CustAccountCreditOrDebit: "",                                                        //客户账户借记贷记类型 可空 若是信用卡，则以下两个参数“信用卡有效期”和“信用卡cvv2”； 1借记 2贷记 4 未知
		CustCardValidDate:        "",                                                        //客户信用卡有效期 可空 信用卡的正下方的四位数，前两位是月份，后两位是年份；
		CustCardCvv2:             "",                                                        //客户信用卡的cvv2 可空 信用卡的背面的三位数
		CustID:                   _CustID,                                                   //客户证件号码
		CustPhone:                "",                                                        //客户手机号 如果商户购买的产品中勾选了短信通知功能，则当商户填写了手机号时,快付通会在交易成功后通过短信通知客户
		Messages:                 "",                                                        //发送客户短信内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
		CustEmail:                "",                                                        //客户邮箱地址 可空 如果商户购买的产品中勾选了邮件通知功能，则当商户填写了邮箱地址时,快付通会在交易成功后通过邮件通知客户
		EmailMessages:            "",                                                        //发送客户邮件内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
		Remark:                   "",                                                        //备注 可空 商户可额外填写备注信息,此信息会传给银行,会在银行的账单信息中显示(具体如何显示取决于银行方,快付通不保证银行肯定能显示)
		CustProtocolNo:           "",                                                        //客户协议编号 可空 扣款人在快付通备案的协议号。
		ExtendParams:             "",                                                        //扩展参数 可空 用于商户的特定业务信息传递，只有商户与快付通约定了传递此参数且约定了参数含义，此参数才有效。参数格式：参数名 1^参数值 1|参数名 2^参数值 2|……多条数据用“|”间隔注意: 不能包含特殊字符（如：#、%、&、+）、敏感词汇, 如果必须使用特殊字符,则需要自行做URL Encoding
		RateAmount:               strconv.FormatFloat(float64(_RateAmount*100), 'f', 0, 64), //商户手续费 可空 本次交易需要扣除的手续费。单位:分,不支持小数点 如不填，则手续费默认0元；
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	_return_35 := _kuaifutong.Gbp_same_id_credit_card_pay_small(_model_35)
	fmt.Println(_return_35)
	if _return_35.Status == "1" {
		_rt_json = map[string]interface{}{"code": 1, "msg": "success", "info": "提现成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "success", "info": _return_35}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
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
