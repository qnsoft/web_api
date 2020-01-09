package api

import (
	"github.com/qnsoft/web_api/bll/user"
	"github.com/qnsoft/web_api/controllers/OnlinePay/Kuaifutong"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/web_api/models/shop"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"

	_ "github.com/mattn/go-adodb"
)

/**
*信息实体
 */
type Login_Controller struct {
	Token.BaseController
}

/*
全局用户从信用卡端登录
*/
func (this *Login_Controller) UserLogin() {
	var _rt_json interface{}
	_UserName := this.GetString("UserName")
	_UserPass := this.GetString("UserPass")

	_model := shop.User{Zhanghao: _UserName, Mima: _UserPass}
	_results, _bll_user := bll.User_Get_IsLogin(_model)
	if _results {
		_str_IdCard, _SubMerchId := "", ""
		_model_auth := shop.UserAuth{UserId: _bll_user.UserId}
		results_auth, _ := DbHelper.MySqlDb().Get(&_model_auth)
		if results_auth {
			_str_IdCard = _model_auth.Idcard     //身份证号
			_SubMerchId = _model_auth.Submerchid //商户签约号
		}
		_rt_json = map[string]interface{}{"StatusCode": 0, "Success": true, "Data": map[string]interface{}{
			"id":         _bll_user.UserId,
			"username":   _bll_user.Zhanghao,
			"truename":   _bll_user.RealName,
			"telephone":  _bll_user.Mobile,
			"smallimage": _bll_user.Headimgurl,
			"IdCard":     _str_IdCard,
			"SubMerchId": _SubMerchId,
			"rongToken":  "",
		}, "info": "登录成功！"}
	} else {
		_rt_json = map[string]interface{}{"StatusCode": -1, "Fail": false, "Info": "登录失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
全局用户从信用卡端修改密码
*/
func (this *Login_Controller) UserChangePwd() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId")     //用户id
	_UserPass := this.GetString("UserPass") //用户密码

	_model := shop.User{UserId: _UserId}
	_results, _ := bll.User_Get_One(&_model)
	if _results {
		_update_model := shop.User{
			Mima: bll.Get_lock_url(_UserPass),
		}
		_count, err := DbHelper.MySqlDb().Id(_model.Id).Update(&_update_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			_rt_json = map[string]interface{}{"code": 1, "msg": "success", "info": "密码修改成功！"}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "密码修改失败！"}
		}

	} else {
		_rt_json = map[string]interface{}{"StatusCode": -1, "fail": false, "Info": "密码修改失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
全局用户从信用卡端修改密码
*/
func (this *Login_Controller) UserFindPwd() {
	var _rt_json interface{}
	_Zhanghao := this.GetString("Mobile")   //用户id
	_UserPass := this.GetString("UserPass") //用户密码

	_model := shop.User{Zhanghao: _Zhanghao}
	_results, _ := bll.User_Get_One(&_model)
	if _results {
		_update_model := shop.User{
			Mima: bll.Get_lock_url(_UserPass),
		}
		_count, err := DbHelper.MySqlDb().Id(_model.Id).Update(&_update_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			_rt_json = map[string]interface{}{"code": 1, "msg": "success", "info": "找回成功！"}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "密码找回失败！"}
		}

	} else {
		_rt_json = map[string]interface{}{"StatusCode": -1, "fail": false, "Info": "密码找回失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
*获取用户余额
 */
func (this *Login_Controller) Get_User_XykYe() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId") //用户id
	_model := shop.UserAuth{UserId: _UserId}
	_results, _ := bll.User_Auth_Get_One(&_model)
	if _results {
		_model_311 := Kuaifutong.Model_311{
			ReqNo:   "",
			CustID:  _model.Idcard,
			PageNum: "1",
		}
		_kuaifutong := Kuaifutong.KuaiPayHelper{}
		_order := _kuaifutong.Gbp_same_id_credit_card_not_pay_balance(_model_311)
		_rt_json = map[string]interface{}{"code": 1, "msg": "success", "order": _order, "info": "获取大额余额成功!"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
*获取用户余额(小额)
 */
func (this *Login_Controller) Get_User_XykYe_Small() {
	var _rt_json interface{}
	_UserId := this.GetString("UserId") //用户id
	_model := shop.UserAuth{UserId: _UserId}
	_results, _ := bll.User_Auth_Get_One(&_model)
	if _results {
		_model_311 := Kuaifutong.Model_311{
			ReqNo:   "",
			CustID:  _model.Idcard,
			PageNum: "1",
		}
		_kuaifutong := Kuaifutong.KuaiPayHelper{}
		_order := _kuaifutong.Gbp_same_id_credit_card_not_pay_balance_small(_model_311)
		_rt_json = map[string]interface{}{"code": 1, "msg": "success", "order": _order, "info": "获取小额余额成功!"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}
