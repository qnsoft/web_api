package bll

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/qnsoft/web_api/models/shop"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"

	"github.com/ajg/form"
)

/*
*获取登录信息
 */
func User_Get_IsLogin(_model shop.User) (bool, *shop.User) {
	_rt := false
	_model_new := &shop.User{Zhanghao: _model.Zhanghao}
	results, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if results {
		if Get_unlock_url(_model_new.Mima) == _model.Mima {
			_rt = true
		}
	}
	return _rt, _model_new

}

/*
*获取单条信息信息
 */
func User_Get_One(_model *shop.User) (bool, *shop.User) {
	_model_new := &shop.User{UserId: _model.UserId}
	results, err := DbHelper.MySqlDb().Get(_model)
	ErrorHelper.CheckErr(err)
	return results, _model_new
}

/*
获取商城密码加密
*/
func Get_lock_url(_UserPass string) string {
	var _rt Login_shop_model
	formData := map[string]string{"password": _UserPass}
	_req := HttpPostCall("https://shop.xhdncppf.com/index.php?store_id=8&module=app&action=index&app=get_lock_url", formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt.Password
}

/*
获取商城密码解密
*/
func Get_unlock_url(_UserPass string) string {
	var _rt ULogin_shop_model
	formData := map[string]string{"password": _UserPass}
	_req := HttpPostCall("https://shop.xhdncppf.com/index.php?store_id=8&module=app&action=index&app=get_unlock_url", formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt.U_Password
}

/*
新后台登录成功返回实体加密
*/
type Login_shop_model struct {
	Code     int    `json:"code"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

/*
新后台登录成功返回实体解密
*/
type ULogin_shop_model struct {
	Code       int    `json:"code"`
	U_Password string `json:"u_password"`
	Message    string `json:"message"`
}

func HttpPostCall(url string, formData map[string]string) string {
	var _rt = ""
	ErrorHelper.LogInfo("这是支付提交的URL：", url)
	ErrorHelper.LogInfo("这是支付提交的参数对象：", formData)

	_formData, err := form.EncodeToString(formData)
	ErrorHelper.CheckErr(err)
	if len(url) > 0 && len(formData) > 0 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		payload := strings.NewReader(_formData)
		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; text/html; charset=utf-8")
		req.Header.Add("Accept", "text/html")
		res, err := client.Do(req)
		//res, err := http.DefaultClient.Do(req)
		ErrorHelper.CheckErr(err)
		defer res.Body.Close()
		if res != nil {
			body, _ := ioutil.ReadAll(res.Body)
			_rt = string(body)
		}
	}
	return _rt
}
