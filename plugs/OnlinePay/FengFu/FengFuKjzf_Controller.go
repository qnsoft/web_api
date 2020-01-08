package FengFu

import (
	"fmt"
	"qnsoft/online_pay/utils/pay/fengfukjzf"
	"qnsoft/online_pay/utils/pay/paytool"
	"qnsoft/qn_card/controllers/Token"
	"qnsoft/qn_card/models/shop"
	"qnsoft/web_api/utils/DbHelper"
	"qnsoft/web_api/utils/ErrorHelper"
	"strconv"
)

/**
*快捷支付——丰付控制器实体
 */
type FengFuKjzf_Controller struct {
	Token.BaseController
}

var (
	Api_notifyUrl_kjzf string = "https://api1.xhdncppf.com" //异步回调地址 http://testapi.ljz789.com
	_SubMerchId_kjzf   string = ""                          //由101接口获取
)

/*
3.3. 支付卡开通发短信(快捷支付必走接口)------------------------------
*/
func (this *FengFuKjzf_Controller) Card_Treaty() {
	var _json_model interface{}
	_UserId := this.GetString("UserId") //用户id
	fmt.Println(_UserId)
	_HolderName := this.GetString("HolderName")               //持卡人真实姓名
	_BankCardNo := this.GetString("BankCardNo")               //银行卡号
	_MobileNo := this.GetString("MobileNo")                   //预留手机号码
	_CertificateNo := this.GetString("CertificateNo")         //证件号
	_CustCardValidDate := this.GetString("CustCardValidDate") //客户信用卡有效期
	_CustCardCvv2 := this.GetString("CustCardCvv2")           //客户信用卡的cvv2
	//3.6. 支付卡开通发短信接口--------------
	cardModel := paytool.CardModel{
		AccountNumber: _BankCardNo,        //卡号
		Phone:         _MobileNo,          //手机号
		HolderName:    _HolderName,        //姓名
		IdCard:        _CertificateNo,     //身份证号
		Cvv2:          _CustCardCvv2,      //Cvv2
		Expired:       _CustCardValidDate, //有效期
	}
	payTool := fengfukjzf.FengFuPay{}
	_channelType := "ffapi" //快捷支付通道channelType：ffkjapi
	rtModel := payTool.CardSms(cardModel, _channelType)
	fmt.Println(rtModel)
	if rtModel.Code == "00" { //如果Code=0 获取开通短信成功
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "验证码发送成功！"}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "快捷支付协议申请失败！"}
	}
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
*3.4支付卡开通发短信(快捷支付必走接口)
 */
func (this *FengFuKjzf_Controller) Card_Treaty_Confirm() {
	var _json_model interface{}
	_UserId := this.GetString("UserId") //用户id
	fmt.Println(_UserId)
	_EndDate := this.GetString("EndDate")             //协议失效日期
	_BankCardType := this.GetString("BankCardType")   //银行卡类型 1、借记卡 2、信用卡
	_MobileNo := this.GetString("MobileNo")           //预留手机号码
	_CertificateNo := this.GetString("CertificateNo") //证件号
	/*----捎带传参完毕，以下是确认必传参数-----*/
	_merOrderNumber := this.GetString("OrderNo")              //商户订单号
	_smsCode := this.GetString("AuthCode")                    //手机动态校验码
	_HolderName := this.GetString("HolderName")               //持卡人真实姓名
	_BankCardNo := this.GetString("BankCardNo")               //银行卡号
	_CustCardValidDate := this.GetString("CustCardValidDate") //客户信用卡有效期
	_CustCardCvv2 := this.GetString("CustCardCvv2")           //客户信用卡的cvv2
	_TreatyType := this.GetString("TreatyType")               //协议类型：11：借记卡扣款 12：信用卡扣款
	cardModel := paytool.CardModel{
		AccountNumber: _BankCardNo,        //卡号
		Phone:         _MobileNo,          //手机号
		HolderName:    _HolderName,        //姓名
		IdCard:        _CertificateNo,     //身份证号
		Cvv2:          _CustCardCvv2,      //Cvv2
		Expired:       _CustCardValidDate, //有效期
	}
	_channelType := "ffapi" //快捷支付通道channelType：ffkjapi
	payTool := fengfukjzf.FengFuPay{}
	rtModel := payTool.CardSmsConfirm(cardModel, _channelType, _merOrderNumber, _smsCode)
	//fmt.Println(rtModel)
	if rtModel.Code == "00" { //如果Code=0 获取开通短信成功
		_CardType, err := strconv.Atoi(_BankCardType)
		ErrorHelper.CheckErr(err)
		_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankCardNo}
		results, _ := DbHelper.MySqlDb().Get(_model)
		if results {
			_update_model := shop.UserBankCard{
				StoreId:        8,
				UserId:         _UserId,                        //用户id
				Treatyid:       "",                             //协议号
				Treatytype:     _TreatyType,                    //协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：余额+借记卡扣款 15： 余额+信用卡扣款
				Treatyenddate:  _EndDate,                       //协议失效日期
				Cardholder:     _HolderName,                    //持卡人真实姓名
				BankCode:       Get_Bank_Code(_model.BankName), //银行行别代码
				CardType:       _CardType,                      //银行卡类型 1、借记卡 2、信用卡
				BankCardNumber: _BankCardNo,                    //银行卡号
				Mobile:         _MobileNo,                      //预留手机号
				IdCard:         _CertificateNo,                 //身份证号
				Expiretime:     _CustCardValidDate,             //客户信用卡有效期
				Cvv2:           _CustCardCvv2,                  //cvv2
				Delflag:        1,                              //是否签约标志
				FfKjzf:         rtModel.Content.MerOrderNumber, //丰付快捷支付-商户签约订单号
			}
			_, err := DbHelper.MySqlDb().Id(_model.Id).Update(_update_model)
			ErrorHelper.CheckErr(err)
		} else {
			_new_model := shop.UserBankCard{
				StoreId:       8,
				UserId:        _UserId,     //用户id
				Treatyid:      "",          //协议号
				Treatytype:    _TreatyType, //协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：余额+借记卡扣款 15： 余额+信用卡扣款
				Treatyenddate: _EndDate,    //协议失效日期
				Cardholder:    _HolderName, //持卡人真实姓名
				//	BankCode:       Get_Bank_Code(_model.BankName), //银行行别代码
				CardType:       _CardType,                      //银行卡类型 1、借记卡 2、信用卡
				BankCardNumber: _BankCardNo,                    //银行卡号
				Mobile:         _MobileNo,                      //预留手机号
				IdCard:         _CertificateNo,                 //身份证号
				Expiretime:     _CustCardValidDate,             //客户信用卡有效期
				Cvv2:           _CustCardCvv2,                  //cvv2
				Delflag:        1,                              //是否签约标志
				FfKjzf:         rtModel.Content.MerOrderNumber, //丰付快捷支付-商户签约订单号
			}
			_, err := DbHelper.MySqlDb().Insert(_new_model)
			ErrorHelper.CheckErr(err)
		}
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "快捷支付协议申请成功！"}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "快捷支付协议申请失败！"}
	}
	fmt.Printf("%+v", _json_model)
	ErrorHelper.LogInfo("【"+_UserId+"】签约成功%+v", _json_model)
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
快捷支付
*/
func (this *FengFuKjzf_Controller) Quick_Pay() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")                               //用户id
	_BankCardNo := this.GetString("BankCardNo")                       //用户出金卡号
	_MerchantBankAccountNo := this.GetString("MerchantBankAccountNo") //用户入金卡号
	_Amount, _ := this.GetFloat("Amount")                             //支付金额
	//根据userid和出金卡号取快捷支付签约信息
	_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankCardNo}
	results, _ := DbHelper.MySqlDb().Get(_model)
	if results { //判断该卡存在并且签约
		//信用卡出金-----------------------------------------
		cardModelA := paytool.CardModel{
			AccountNumber: _BankCardNo,       //出金卡号
			Phone:         _model.Mobile,     //手机号
			HolderName:    _model.Cardholder, //姓名
			IdCard:        _model.IdCard,     //身份证号
			Cvv2:          _model.Cvv2,       //Cvv2
			Expired:       _model.Expiretime, //有效期
		}
		//储蓄卡入金-----------------------------------------
		cardModelB := paytool.CardModel{
			AccountNumber: _MerchantBankAccountNo, //入金卡号
			Phone:         _model.Mobile,          //手机号
		}
		fmt.Printf("%v", cardModelA)
		_channelType := "ffapi" //通道标示小额channelType：ffqrdh大额channelType：ffkj
		_fee := 0.58            //费率
		_ewfee := 1.0
		_city := "郑州市"
		payTool := fengfukjzf.FengFuPay{}
		rtModel := payTool.CardPay(cardModelA, cardModelB, _channelType, _city, _Amount, _fee, _ewfee)
		if rtModel.Code == "00" { //如果Code=0 下单成功写入数据库
			//开始将代扣执行结果以json方式返回
			_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "快捷支付交易提交成功！"}
		} else {
			_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "快捷支付交易提交失败！"}
		}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "出金未签约不能交易！"}
	}
	ErrorHelper.LogInfo("【"+_UserId+"】丰富快捷支付：从"+_BankCardNo+"到"+_MerchantBankAccountNo+"===结果==%+v", _json_model)
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
*判断用户当前出金卡是否签订代扣协议
 */
func (this *FengFuKjzf_Controller) Card_Treaty_IsOk() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")         //用户id
	_BankCardNo := this.GetString("BankCardNo") //用户出金卡号
	//根据userid和出金卡号取代扣签约信息
	_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankCardNo}
	results, _ := DbHelper.MySqlDb().Get(_model)
	if results {
		if len(_model.FfKjzf) > 0 {
			_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": _model, "info": "已签订代扣协议！"}
		} else {
			_json_model = map[string]interface{}{"code": 2, "msg": "fail", "data": _model, "info": "还未签订代扣协议！"}
		}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": _model}
	}
	this.Data["json"] = _json_model
	this.ServeJSON()
}
