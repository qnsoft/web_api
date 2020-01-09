package XYK

import (
	"fmt"
	"strconv"
	"time"
	"github.com/qnsoft/web_api/controllers/OnlinePay/Kuaifutong"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/utils/DateHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/StringHelper"

	_ "github.com/mattn/go-adodb"
)

/**
*信息实体
 */
type XYK_Controller struct {
	Token.BaseController
}

/*
测试接口1
*/
func (this *XYK_Controller) XYK_Demo1() {
	//_Model_User := new(model.User)
	//fmt.Println(_Model_User)

	//this.Data["old"] = map[string]string{"name": Old_Query()}
	//this.Data["new"] = map[string]string{"name": New_Query()}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_pay(10)
	//	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.5单笔付款接口
*/
func (this *XYK_Controller) XYK_35() {
	_Amount, err := this.GetFloat("Amount") //金额（元转分）
	ErrorHelper.CheckErr(err)
	_CustBankNo := this.GetString("CustBankNo")               //客户银行账户行别
	_CustBankAccountNo := this.GetString("CustBankAccountNo") //往客户的哪张卡上付钱
	_CustName := this.GetString("CustName")                   //客户姓名
	_CustID := this.GetString("CustID")                       //客户证件号码
	_RateAmount, err := this.GetFloat("RateAmount")           //手续费（元转分）
	ErrorHelper.CheckErr(err)
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
	sss := _kuaifutong.Gbp_same_id_credit_card_pay(_model_35)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.6快捷协议代扣协议申请接口
*/
func (this *XYK_Controller) XYK_36() {
	_TreatyType := this.GetString("TreatyType")               //协议类型：11：借记卡扣款 12：信用卡扣款
	_Note := this.GetString("Note")                           //说明 参数可空
	_EndDate := this.GetString("EndDate")                     //协议失效日期
	_HolderName := this.GetString("HolderName")               //持卡人真实姓名
	_BankType := this.GetString("BankType")                   //银行行别
	_BankCardType := this.GetString("BankCardType")           //银行卡类型 1、借记卡 2、信用卡
	_BankCardNo := this.GetString("BankCardNo")               //银行卡号
	_MobileNo := this.GetString("MobileNo")                   //预留手机号码
	_CertificateNo := this.GetString("CertificateNo")         //证件号
	_CustCardValidDate := this.GetString("CustCardValidDate") //客户信用卡有效期 如果协议类型为12时不可为空
	_CustCardCvv2 := this.GetString("CustCardCvv2")           //客户信用卡的cvv2 如果协议类型为12时不可为空
	_model_36 := Kuaifutong.Model_36{
		TreatyType:        _TreatyType,
		Note:              _Note,                                   //说明
		StartDate:         date.FormatDate(time.Now(), "yyyyMMdd"), //协议生效日 //根据当前系统自动生成！
		EndDate:           _EndDate,
		HolderName:        _HolderName,
		BankType:          _BankType,
		BankCardType:      _BankCardType,
		BankCardNo:        _BankCardNo,
		MobileNo:          _MobileNo,
		CertificateType:   "0", //证件类型：0表示身份证
		CertificateNo:     _CertificateNo,
		CustCardValidDate: _CustCardValidDate,
		CustCardCvv2:      _CustCardCvv2,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_treaty_collect_apply(_model_36)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.7快捷协议代扣协议确定接口
商户平台通过此接口确认开通代扣协议，进行四要素鉴权获取进行快捷协议代扣的协议号
*/
func (this *XYK_Controller) XYK_37() {
	_SmsSeq := this.GetString("SmsSeq")                       //短信流水号
	_AuthCode := this.GetString("AuthCode")                   //手机动态校验码
	_HolderName := this.GetString("HolderName")               //持卡人真实姓名
	_BankCardNo := this.GetString("BankCardNo")               //银行卡号
	_CustCardValidDate := this.GetString("CustCardValidDate") //客户信用卡有效期
	_CustCardCvv2 := this.GetString("CustCardCvv2")           //客户信用卡的cvv2
	_model_37 := Kuaifutong.Model_37{
		SmsSeq:            _SmsSeq,
		AuthCode:          _AuthCode,
		HolderName:        _HolderName,
		BankCardNo:        _BankCardNo,
		CustCardValidDate: _CustCardValidDate,
		CustCardCvv2:      _CustCardCvv2,
	}
	//协议类型
	TreatyType := "1"
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_confirm_treaty_collect_apply(_model_37, TreatyType)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.8快捷协议代扣接口
此接口用于商户平台协议代扣。此接口需要用户先签定协议。
*/
func (this *XYK_Controller) XYK_38() {
	_OrderNo := this.GetString("OrderNo")                             //订单编号
	_TreatyNo := this.GetString("TreatyNo")                           //协议号
	_TradeTime := this.GetString("TradeTime")                         //商户方交易时间
	_Amount := this.GetString("Amount")                               //代扣金额
	_CustAccountId := this.GetString("CustAccountId")                 //账户ID
	_HolderName := this.GetString("HolderName")                       //持卡人真实姓名
	_BankType := this.GetString("BankType")                           //银行行别
	_BankCardNo := this.GetString("BankCardNo")                       //银行卡号
	_ExtendParams := this.GetString("ExtendParams")                   //扩展字段
	_MerchantBankAccountNo := this.GetString("MerchantBankAccountNo") //商户银行账户
	_RateAmount := this.GetString("RateAmount")                       //商户手续费
	_CustCardValidDate := this.GetString("CustCardValidDate")         //客户信用卡有效期
	_CustCardCvv2 := this.GetString("CustCardCvv2")                   //客户信用卡的cvv2
	_NotifyUrl := this.GetString("NotifyUrl")                         //商户后台通知URL
	_CityCode := this.GetString("CityCode", "4910")                   //城市编码
	_SourceIP := this.GetString("SourceIP")                           //公网IP地址
	_DeviceID := this.GetString("DeviceID")                           //设备标识

	_model_38 := Kuaifutong.Model_38{
		OrderNo:               _OrderNo,
		TreatyNo:              _TreatyNo,
		TradeTime:             _TradeTime,
		Amount:                _Amount,
		CustAccountId:         _CustAccountId,
		HolderName:            _HolderName,
		BankType:              _BankType,
		BankCardNo:            _BankCardNo,
		ExtendParams:          _ExtendParams,
		MerchantBankAccountNo: _MerchantBankAccountNo,
		RateAmount:            _RateAmount,
		CustCardValidDate:     _CustCardValidDate,
		CustCardCvv2:          _CustCardCvv2,
		NotifyUrl:             _NotifyUrl,
		CityCode:              _CityCode,
		SourceIP:              _SourceIP,
		DeviceID:              _DeviceID,
	}
	//协议类型
	TreatyType := "12"
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_credit_card_treaty_collect(_model_38, TreatyType)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
*3.9快捷协议代扣协议查询接口
*此接口用于商户平台通过此接口查询协议信息。
 */
func (this *XYK_Controller) XYK_39() {
	//订单编号
	_OrderNo := this.GetString("OrderNo")
	//协议号
	_TreatyNo := this.GetString("TreatyNo")

	_model_39 := Kuaifutong.Model_39{
		OrderNo:  _OrderNo,
		TreatyNo: _TreatyNo,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_query_treaty_info(_model_39)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
*3.10快捷协议代扣协议解除接口
*此接口用于商户平台通过此接口解除快捷协议收款协议信息。
 */
func (this *XYK_Controller) XYK_310() {
	//订单编号
	_OrderNo := this.GetString("OrderNo")
	//协议号
	_TreatyNo := this.GetString("TreatyNo")

	_model_310 := Kuaifutong.Model_310{
		OrderNo:  _OrderNo,
		TreatyNo: _TreatyNo,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_cancel_treaty_info(_model_310)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.11用户资金查询接口(对应我的余额时custID必传)
*/
func (this *XYK_Controller) XYK_311() {
	_ReqNo := this.GetString("reqNo")
	_CustID := this.GetString("custID")
	_PageNum := this.GetString("pageNum")
	_model_311 := Kuaifutong.Model_311{
		ReqNo:   _ReqNo,
		CustID:  _CustID,
		PageNum: _PageNum,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_credit_card_not_pay_balance(_model_311)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.12交易查询接口
* 3.12交易查询接口
*用于查询指定的一笔或多笔交易的结果,例如购买支付交易状态
*/
func (this *XYK_Controller) XYK_312() {
	_StartDate := this.GetString("StartDate")
	_EndDate := this.GetString("EndDate")
	_OrderNo := this.GetString("OrderNo")
	_TradeType := this.GetString("TradeType")
	_Status := this.GetString("Status")
	_model_312 := Kuaifutong.Model_312{
		StartDate: _StartDate,
		EndDate:   _EndDate,
		OrderNo:   _OrderNo,
		TradeType: _TradeType,
		Status:    _Status,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_credit_card_trade_record_query(_model_312)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
信用卡3.13账户余额查询接口
*/
func (this *XYK_Controller) XYK_313() {

	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Query_available_balance()
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
*3.14交易类对账文件获取接口
此功能用于给商户提供对账数据。
*/
func (this *XYK_Controller) XYK_314() {
	_CheckDate := this.GetString("CheckDate")
	_model_314 := Kuaifutong.Model_314{
		CheckDate: _CheckDate,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_credit_card_recon_result_query(_model_314)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
*3.15银行卡三要素验证接口
*此接口用于校验指定的银行卡和用户身份信息是否匹配及正确，
 */
func (this *XYK_Controller) XYK_315() {
	_CustBankNo := this.GetString("CustBankNo")
	_CustName := this.GetString("CustName")
	_CustBankAccountNo := this.GetString("CustBankAccountNo")
	_CustAccountCreditOrDebit := this.GetString("CustAccountCreditOrDebit")
	_CustCertificationType := this.GetString("CustCertificationType")
	_CustID := this.GetString("CustID")
	_model_315 := Kuaifutong.Model_315{
		CustBankNo:               _CustBankNo,
		CustName:                 _CustName,
		CustBankAccountNo:        _CustBankAccountNo,
		CustAccountCreditOrDebit: _CustAccountCreditOrDebit,
		CustCertificationType:    _CustCertificationType,
		CustID:                   _CustID,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	orderStarWith := ""
	sss := _kuaifutong.Gbp_threeMessage_verification(_model_315, orderStarWith)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}

/*
3.16验证类对账文件获取接口
此功能用于给商户提供对账数据。
*/
func (this *XYK_Controller) XYK_316() {
	_CheckDate := this.GetString("CheckDate")
	_model_316 := Kuaifutong.Model_316{
		CheckDate: _CheckDate,
	}
	_kuaifutong := Kuaifutong.KuaiPayHelper{}
	sss := _kuaifutong.Gbp_same_id_credit_card_verify_result_query(_model_316)
	fmt.Println(sss)
	ErrorHelper.LogInfo(sss)
	this.Data["json"] = sss
	this.ServeJSON()
}
