package FengFu

import (
	"fmt"
	"github.com/qnsoft/online_pay/utils/pay/fengfu"
	"github.com/qnsoft/online_pay/utils/pay/paytool"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/web_api/models/shop"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"strconv"
	"sync"
	"time"
)

/**
*支付——丰付控制器实体
 */
type FengFu_Controller struct {
	Token.BaseController
}

var (
	Api_notifyUrl string = "https://api1.xhdncppf.com" //异步回调地址 http://testapi.ljz789.com
	_SubMerchId   string = ""                          //由101接口获取
)

/*
3.6. 支付卡开通发短信(大额必走接口)------------------------------
*/
func (this *FengFu_Controller) Card_Treaty() {
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
	payTool := fengfu.FengFuPay{}
	_channelType := "fk" //签约小额为空 //通道标示小额channelType：ffqrdh大额channelType：fk
	rtModel := payTool.CardSign(cardModel, _channelType)
	fmt.Println(rtModel)
	if rtModel.Code == "00" { //如果Code=0 获取开通短信成功
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "验证码发送成功！"}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "代扣协议申请失败！"}
	}
	this.Data["json"] = _json_model
	this.ServeJSON()

	//信用卡出金订单查询
	//_merOrderNumber := "2019110510301919"
	//rtModel := payTool.CardOutOrder(_merOrderNumber)
	//信用卡入金
	// payTool := fengfu.FengFuPay{}
	// cardModelB := paytool.CardModel{
	// 	AccountNumber: "6225258803477017",   //卡号
	// 	Phone:         "13938202388",        //手机号
	// 	HolderName:    "宋晓辉",                //姓名
	// 	IdCard:        "410185198209190514", //身份证号
	// 	Cvv2:          "307",                //Cvv2
	// 	Expired:       "1222",               //有效期
	// }
	// fmt.Printf("%v", cardModelB)
	// _channelType := "ffqrdh" //通道标示小额channelType：ffqr大额channelType：ffkj
	// _money := 99.70
	// _extraFee := 0.5
	// rtModel := payTool.CardIn(cardModelB, _channelType, _money, _extraFee)
	//信用卡入金订单查询
	// _merOrderNumber := "201911051635117832"
	// rtModel := payTool.CardInOrder(_merOrderNumber)
	//信用卡余额查询
	/* _channelType := "ffqr"
	_accountNumber := "6225258803477017"
	payTool := fengfu.FengFuPay{}
	rtModel := payTool.CardBanlance(_channelType, _accountNumber) */
	//fmt.Printf("%v", rtModel)
}

/*
*支付卡开通发短信(大额必走接口)
 */
func (this *FengFu_Controller) Card_Treaty_Confirm() {
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
	_channelType := "fk" //签约小额为空 //通道标示小额channelType：ffqrdh大额channelType：fk
	payTool := fengfu.FengFuPay{}
	rtModel := payTool.CardSignConfirm(cardModel, _channelType, _merOrderNumber, _smsCode)
	fmt.Println(rtModel)
	if rtModel.Code == "00" { //如果Code=0 获取开通短信成功
		_CardType, err := strconv.Atoi(_BankCardType)
		ErrorHelper.CheckErr(err)
		_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankCardNo}
		results, _ := DbHelper.MySqlDb().Get(_model)
		if results {
			_update_model := shop.UserBankCard{
				StoreId:          8,
				UserId:           _UserId,                        //用户id
				Treatyid:         "",                             //协议号
				Treatytype:       _TreatyType,                    //协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：余额+借记卡扣款 15： 余额+信用卡扣款
				Treatyenddate:    _EndDate,                       //协议失效日期
				Cardholder:       _HolderName,                    //持卡人真实姓名
				BankCode:         Get_Bank_Code(_model.BankName), //银行行别代码
				CardType:         _CardType,                      //银行卡类型 1、借记卡 2、信用卡
				BankCardNumber:   _BankCardNo,                    //银行卡号
				Mobile:           _MobileNo,                      //预留手机号
				IdCard:           _CertificateNo,                 //身份证号
				Expiretime:       _CustCardValidDate,             //客户信用卡有效期
				Cvv2:             _CustCardCvv2,                  //cvv2
				Delflag:          1,                              //是否签约标志
				FfMerordernumber: rtModel.Content.MerOrderNumber, //丰付-商户签约订单号
				FfIssign:         "f",                            //丰付-商户签约标识
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
				CardType:         _CardType,                      //银行卡类型 1、借记卡 2、信用卡
				BankCardNumber:   _BankCardNo,                    //银行卡号
				Mobile:           _MobileNo,                      //预留手机号
				IdCard:           _CertificateNo,                 //身份证号
				Expiretime:       _CustCardValidDate,             //客户信用卡有效期
				Cvv2:             _CustCardCvv2,                  //cvv2
				Delflag:          1,                              //是否签约标志
				FfMerordernumber: rtModel.Content.MerOrderNumber, //丰付-商户签约订单号
				FfIssign:         "f",                            //丰付-商户签约标识
			}
			_, err := DbHelper.MySqlDb().Insert(_new_model)
			ErrorHelper.CheckErr(err)

		}
		_json_model = map[string]interface{}{"code": 1, "msg": "success", "data": rtModel, "info": "代扣协议申请成功！"}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "data": rtModel, "info": "代扣协议申请失败！"}
	}
	fmt.Printf("%+v", _json_model)
	ErrorHelper.LogInfo("【"+_UserId+"】签约成功%+v", _json_model)
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
手动还款(代扣-代还)
*/
func (this *FengFu_Controller) Quick_Pay() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")                               //用户id
	_BankCardNo := this.GetString("BankCardNo")                       //用户出金卡号
	_MerchantBankAccountNo := this.GetString("MerchantBankAccountNo") //用户入金卡号
	_Amount, _ := this.GetFloat("Amount")                             //代扣金额 手动输入 从提交获取
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
		_channelType := "ffqrdh" //通道标示小额channelType：ffqrdh大额channelType：ffkj
		_fee := 0.36             //小额费率
		_city := "郑州市"
		//if _Amount > 1000 {
		_channelType = "fk"
		_fee = 0.58 //大额费率
		//	}
		//_money := 100.0
		payTool := fengfu.FengFuPay{}
		rtModel := payTool.CardOut(cardModelA, _channelType, _city, _Amount, _fee)
		_str_rt := fmt.Sprintf("%+v", rtModel)
		if rtModel.Code == "00" { //如果Code=0 下单成功写入数据库
			//开始将代扣信息写入到信用卡订单表[lkt_user_card_order]
			_new_model_order := shop.UserCardOrder{
				UserId:                _UserId,                        //用户id
				Ordernoa:              rtModel.Content.MerOrderNumber, //订单编号
				Amount:                int(_Amount * 100),             //订单金额
				Rateamounta:           int(_fee * 100),                //代扣手续费
				Bankcardno:            _BankCardNo,                    //出金卡号
				Merchantbankaccountno: _MerchantBankAccountNo,         //入金卡号
				ReturnA:               _str_rt,                        //代扣返回信息
				AddTime:               time.Now(),                     //订单创建时间
			}
			_count, err := DbHelper.MySqlDb().Insert(_new_model_order)
			ErrorHelper.CheckErr(err)
			if _count > 0 && err == nil {
				_AmountB := _Amount - _Amount*(_fee/100)
				var wg = &sync.WaitGroup{}
				wg.Add(1)
				go func() { //开始异步执行
					this.Quick_Pay_Ok(cardModelA, _UserId, rtModel.Content.MerOrderNumber, _BankCardNo, _MerchantBankAccountNo, _AmountB)
					wg.Done()
				}()
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
	ErrorHelper.LogInfo("【"+_UserId+"】手动代付执行：从"+_BankCardNo+"到"+_MerchantBankAccountNo+"===结果==%+v", _json_model)
	this.Data["json"] = _json_model
	this.ServeJSON()
}

/*
手动还款(代付)
@_UserId 用户id
@_OrderNoA 订单出金编号
@_BankCardNo 出金卡号
@_MerchantBankAccountNo 入金卡号
@_AmountB 出金金额
*/
func (this *FengFu_Controller) Quick_Pay_Ok(cardModelA paytool.CardModel, _UserId, _OrderNoA, _BankCardNo, _MerchantBankAccountNo string, _AmountB float64) {
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
		_extraFee := 1.0 //到账手续费
		payTool := fengfu.FengFuPay{}
		//入金操作之前必须保证出金订单已经完成
	L_on_job:
		chujinOrder := payTool.CardOutOrder(cardModelA, _OrderNoA)
		if chujinOrder.Content.OrderStatus == "00" {
			rtModel := payTool.CardIn(cardModelB, _channelType, _AmountB, _extraFee) //入金操作
			_str_rt := fmt.Sprintf("%+v", rtModel)
			_update_model_order := shop.UserCardOrder{
				UserId:      _UserId,                        //用户id
				Ordernob:    rtModel.Content.MerOrderNumber, //订单编号
				Rateamountb: int(_extraFee * 100),           //代扣手续费
				ReturnB:     _str_rt,                        //代扣返回信息
			}
			_count, err := DbHelper.MySqlDb().Where("OrderNoA=?", _OrderNoA).Update(_update_model_order)
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
			time.Sleep(time.Second * 60)
			goto L_on_job
		}
	} else {
		_json_model = map[string]interface{}{"code": 0, "msg": "fail", "info": "入金未签约不能交易！"}
	}
	fmt.Printf("%+v", _json_model)
	ErrorHelper.LogInfo("【"+_UserId+"】手动代付执行：从"+_BankCardNo+"到"+_MerchantBankAccountNo+"===结果==%+v", _json_model)
}

/*
*判断用户当前出金卡是否签订代扣协议
 */
func (this *FengFu_Controller) Card_Treaty_IsOk() {
	var _json_model interface{}
	_UserId := this.GetString("UserId")         //用户id
	_BankCardNo := this.GetString("BankCardNo") //用户出金卡号
	//根据userid和出金卡号取代扣签约信息
	_model := &shop.UserBankCard{UserId: _UserId, BankCardNumber: _BankCardNo}
	results, _ := DbHelper.MySqlDb().Get(_model)
	if results {
		if len(_model.FfMerordernumber) > 2 && len(_model.FfIssign) > 0 {
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

/*
余额查询
*/
func (this *FengFu_Controller) YeCx() {
	var jsonModel interface{}
	channelType := this.GetString("channelType", "ffqrdh") //通道标识
	accountNumber := this.GetString("accountNumber")       //卡号
	payTool := fengfu.FengFuPay{}
	jsonModel = payTool.CardBanlance(channelType, accountNumber)
	this.Data["json"] = jsonModel
	this.ServeJSON()
}

/*
入金操作
*/
func (this *FengFu_Controller) RuJin() {
	var jsonModel interface{}
	channelType := this.GetString("channelType", "ffqrdh") //通道标识
	accountNumber := this.GetString("accountNumber")       //卡号
	amountB, _ := this.GetFloat("amount")                  //金额
	//根据用户id和银行卡号取用户和卡的相关信息
	_model := &shop.UserBankCard{BankCardNumber: accountNumber} //这里是入金卡 臻方便费率应该从出金卡扣除
	results, _ := DbHelper.MySqlDb().Get(_model)
	if results {
		//信用卡入金-----------------------------------------
		cardModelB := paytool.CardModel{
			AccountNumber: accountNumber,     //入金卡号
			Phone:         _model.Mobile,     //手机号
			HolderName:    _model.Cardholder, //姓名
			IdCard:        _model.IdCard,     //身份证号
			Cvv2:          _model.Cvv2,       //Cvv2
			Expired:       _model.Expiretime, //有效期
		}
		fmt.Printf("%v", cardModelB)

		//if amountB > 1000 {
		channelType = "fk"
		// } else {
		// 	channelType = "ffqrdh"
		// }
		_extraFee := 1.0 //到账手续费
		payTool := fengfu.FengFuPay{}
		jsonModel = payTool.CardIn(cardModelB, channelType, amountB, _extraFee)
	} else {
		jsonModel = map[string]interface{}{"code": 0, "msg": "fail", "info": "入金未签约不能交易！"}
	}
	this.Data["json"] = jsonModel
	this.ServeJSON()
}

/*
*根据银行代码获取当前银行别代码
 */
func Get_Bank_Code(_bank_name string) string {
	_rt := ""
	_model_new := &shop.ChannelBankKft{Bankname: _bank_name}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		_rt = _model_new.Bankcode
	} else {
		_rt = ""
	}
	return _rt
}

/*
*根据银行卡号查询银行名称
 */
func Get_Bank_Name(_Bank_No string) string {
	_rt := ""
	_model_new := &shop.UserBankCard{BankCardNumber: _Bank_No}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		_rt = _model_new.BankName
	}
	return _rt
}

/*
*根据银行卡号获取银行行别代码
 */
func Get_Bank_CodeA(_Bank_No string) string {
	_rt := ""
	_model_new := &shop.UserBankCard{BankCardNumber: _Bank_No}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		_rt = _model_new.BankCode
	}
	return _rt
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
