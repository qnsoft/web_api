package Duijie

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/web_api/models/Duijie/New"
	"github.com/qnsoft/web_api/models/Duijie/Old"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	php2go "github.com/qnsoft/utils/Php2go"
	_ "github.com/qnsoft/utils/WebHelper"

	_ "github.com/mattn/go-adodb"
)

/**
*信息实体
 */
type RegisterAll_Controller struct {
	Token.BaseController
}

/*
导数据
*/
func (this *RegisterAll_Controller) Daoshuju() {
	daoshuju()
	this.ServeJSON()
}

/*
全局用户注册接口
*/
func (this *RegisterAll_Controller) RegisterAll() {
	//_Model_User := new(model.User)
	//fmt.Println(_Model_User)
	this.Data["old"] = map[string]string{"name": Old_Query()}
	this.Data["new"] = map[string]string{"name": New_Query()}
	this.ServeJSON()
}

/*
全局用户从信用卡端登录(如果从信用卡登录在登录成功的同时去商城查询该用户是否存在，如果不存在自动注册！)
*/
func (this *RegisterAll_Controller) LoginXyk() {
	_UserName := this.GetString("UserName")
	_UserPass := this.GetString("UserPass")
	fmt.Println(_UserName + ":" + _UserPass)
	_bl, _user, _user_ext := Old_User_Login(_UserName, _UserPass)
	if _bl { //开始判断信用卡里是否有该用户如果有开始执行下面操作
		//开始post请求新商城
		_new_password := get_lock_url(_UserPass)
		New_User_Add(_user, _user_ext, _new_password)            //注册到商城表
		New_ParentUser_IsExist(strconv.Itoa(_user_ext.Parentid)) //将上级信息注册到商城
		New_User_level_Add(_user, _user_ext)                     //将层级注册到商城数据库
		_login_ok := Login_xyk{}
		_login_ok.StatusCode = 0
		_login_ok.Success = true
		_login_ok.Info = "登录成功!"
		_login_ok.Data = Login_xyk_model{
			Id:         _user.Id,
			Username:   _user.Username,
			Truename:   _user_ext.Truename,
			Telephone:  _user_ext.Telephone,
			Smallimage: _user_ext.Smallimage,
			RongToken:  "",
		}
		_login_ok.Count = 0
		this.Data["json"] = _login_ok
	} else {
		_login_ok := Login_xyk{}
		_login_ok.StatusCode = 0
		_login_ok.Success = true
		_login_ok.Info = "登录失败!"
		_login_ok.Data = Login_xyk_model{}
		_login_ok.Count = 0
		this.Data["json"] = _login_ok
	}
	this.ServeJSON()
}

/*
获取商城密码加密
*/
func get_lock_url(_UserPass string) string {
	url := "https://shop.xhdncppf.com/index.php?store_id=8"
	payload := strings.NewReader("module=app&action=index%20&app=get_lock_url%20%20&password=" + _UserPass)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "shop.xhdncppf.com")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	_login_shop_model := Login_shop_model{}
	json.Unmarshal(body, &_login_shop_model)
	return _login_shop_model.Password
}

/*
信用卡用户登录验证(返回是否登录成功，如果成功返回用户对象不成功返回0)
*/
func Old_User_Login(_UserName, _UserPass string) (bool, *Old.User, *Old.User_ext) {
	_ok := false
	_model := &Old.User{Username: _UserName}
	_model_ext := new(Old.User_ext)
	results, _ := DbHelper.MsSqlDb().Get(_model)
	if results {
		h := md5.New()
		h.Write([]byte(_UserPass + _model.Salt)) // 需要加密的字符串为 _UserPass
		cipherStr := h.Sum(nil)
		str_md5_pwd := hex.EncodeToString(cipherStr)
		if str_md5_pwd == _model.Password {
			_model_extA := &Old.User_ext{Userid: _model.Id}
			results_1, _err := DbHelper.MsSqlDb().Get(_model_extA)
			ErrorHelper.CheckErr(_err)
			if results_1 {
				_model_ext = _model_extA
				_ok = true
			}
		}
	}
	return _ok, _model, _model_ext
}

/*
判断用户是否在商城存在(判断条件主要是手机号)
*/
func Shop_User_Exist(_zhanghao string) bool { //检查新商城中用户账号是否存在不存在才允许注册
	_model := &New.User{Zhanghao: _zhanghao}
	results, _ := DbHelper.MySqlDb().Get(_model)
	return results
}

/*
老用户在信用卡登录的同时到商城验证是否注册如果没注册自动注册，如果已经注册自动修改密码、登录时间、登录ip、登录次数
@_xyk_user 原会员实体信息
@_xyk_user_ext 原会员扩展信息
*/
func New_User_Add(_xyk_user *Old.User, _xyk_user_ext *Old.User_ext, _UserPass string) {
	_is_vip := 0
	_model_vip := &Old.User_layout{Userid: _xyk_user.Id}
	results_vip, _ := DbHelper.MsSqlDb().Get(_model_vip) //根据id判断是否是vip
	if results_vip {
		if _model_vip.Vipflag {
			_is_vip = 1
		} else {
			_is_vip = 0
		}
	}
	_model := &New.User{UserId: strconv.Itoa(_xyk_user.Id)}
	results, _ := DbHelper.MySqlDb().Get(_model)
	//fmt.Println("%v", results)
	if !results {
		_new_model := New.User{
			StoreId:      8,
			UserId:       strconv.Itoa(_xyk_user.Id),                                //id
			UserName:     []byte("我"),                                               //用户昵称
			Zhanghao:     _xyk_user.Username,                                        //账号
			Mima:         _UserPass,                                                 //密码
			Money:        strconv.FormatFloat(_xyk_user_ext.User_money, 'f', 6, 64), //余额
			RegisterData: _xyk_user.Addtime,
			EMail:        _xyk_user_ext.Email,
			RealName:     _xyk_user_ext.Truename,
			Mobile:       _xyk_user.Username,
			Referee:      strconv.Itoa(_xyk_user_ext.Parentid), //推荐人
			IsVip:        _is_vip,                              //是否vip
			IsLock:       0,
		}
		_, err := DbHelper.MySqlDb().Insert(_new_model)
		ErrorHelper.CheckErr(err)
		fmt.Println(_xyk_user.Username + ":已迁移到商城！")
		//fmt.Println("%v", results_1)
	} else { //执行修改密码
		//_update_model := New.User{
		//	Mima:  _UserPass,
		//	IsVip: _is_vip,
		//	}
		count, err := DbHelper.MySqlDb().Table(new(New.User)).Id(_model.Id).Update(map[string]interface{}{"mima": _UserPass, "is_vip": _is_vip})
		ErrorHelper.LogInfo("密码在商城数据库中修改成功！?", strconv.FormatInt(count, 10))
		ErrorHelper.CheckErr(err)
		fmt.Println(_xyk_user.Username + ":商城已经注册过，无需再注册,现在做了密码同步！")
	}
}

/*
判断用户在商城是否存在，不存在自动注册
*/
func New_ParentUser_IsExist(_UserId string) bool {
	//_ok := false
	_model := &New.User{UserId: _UserId}
	results, err := DbHelper.MySqlDb().Get(_model)
	ErrorHelper.CheckErr(err)
	_id, err := strconv.Atoi(_model.UserId)
	ErrorHelper.CheckErr(err)
	_xyk_user := &Old.User{Id: _id}
	_xyk_user_ext := new(Old.User_ext)
	results_old, _ := DbHelper.MsSqlDb().Get(_xyk_user) //获取旧主表数据
	if !results {
		if results_old {
			_model_extA := &Old.User_ext{Userid: _id}
			results_old_1, _err := DbHelper.MsSqlDb().Get(_model_extA)
			ErrorHelper.CheckErr(_err)
			if results_old_1 {
				_xyk_user_ext = _model_extA
				New_User_Add(_xyk_user, _xyk_user_ext, "") //将信用卡用户在线注册到商城表
				//将伞下信息注册到商城
				New_User_level_Add(_xyk_user, _xyk_user_ext) //将用户层级及分润更新到商城分润信息表
			}
		}
	} else {
		_model_extA := &Old.User_ext{Userid: _id}
		results_old_1, _err := DbHelper.MsSqlDb().Get(_model_extA)
		ErrorHelper.CheckErr(_err)
		if results_old_1 {
			_xyk_user_ext = _model_extA
			//将伞下信息注册到商城
			New_User_level_Add(_xyk_user, _xyk_user_ext) //将用户层级及分润更新到商城分润信息表
		}
	}
	return results
}

/*
用户层级添加
_xyk_user：信用卡用户主表
_xyk_user_ext：信用卡用户扩展表
*/
func New_User_level_Add(_xyk_user *Old.User, _xyk_user_ext *Old.User_ext) {

	//Jian_tui(_xyk_user.Id) //间推
	_is_vip := 0
	_model_vip := &Old.User_layout{Userid: _xyk_user.Id}
	results_vip, _ := DbHelper.MsSqlDb().Get(_model_vip) //根据id判断是否是vip
	if results_vip {
		if _model_vip.Vipflag {
			_is_vip = 1
		} else {
			_is_vip = 0
		}
	}
	_new_model := &New.User_level{UserId: strconv.Itoa(_xyk_user.Id)}
	results, _ := DbHelper.MySqlDb().Get(_new_model)
	if !results {
		_insert_model := New.User_level{
			StoreId:     8,
			UserId:      strconv.Itoa(_xyk_user.Id),                                //id
			Referee:     strconv.Itoa(_xyk_user_ext.Parentid),                      //推荐人
			SuperiorAll: _xyk_user_ext.Layout,                                      //隶属上级
			DirectIdAll: Zhi_tui(_xyk_user.Id),                                     //直推人
			CardProfit:  strconv.FormatFloat(_xyk_user_ext.User_money, 'f', 6, 64), //信用卡总返利余额
			IsVip:       _is_vip,                                                   //是否vip
		}
		_, err := DbHelper.MySqlDb().Insert(_insert_model)
		ErrorHelper.CheckErr(err)
		fmt.Println(_xyk_user.Username + ":信用卡分润相关数据迁移成功！")
	} else { //执行余额及层级关系修改
		_ipdate_model := New.User_level{
			IsVip:       _is_vip,
			SuperiorAll: _xyk_user_ext.Layout,
			DirectIdAll: Zhi_tui(_xyk_user.Id),
			CardProfit:  strconv.FormatFloat(_xyk_user_ext.User_money, 'f', 6, 64), //信用卡总返利余额
		}
		_, err := DbHelper.MySqlDb().Where("User_Id=?", strconv.Itoa(_xyk_user.Id)).Update(_ipdate_model)
		ErrorHelper.CheckErr(err)
		//再次从用户表中更新总余额
		_up_user_model := &New.User_level{UserId: strconv.Itoa(_xyk_user.Id)}
		results_up_user, _ := DbHelper.MySqlDb().Get(_up_user_model)
		if results_up_user {
			_update_user_model := New.User{
				//总余额=推人总返利+信用卡总返利+商城消费总返利+商家销售总返利+积分兑换总返利
				Money: _up_user_model.UserProfit + _up_user_model.CardProfit + _up_user_model.ShopProfit + _up_user_model.SellerProfit + _up_user_model.ChangeProfit,
			}
			_, err := DbHelper.MySqlDb().Id(_up_user_model.Id).Update(_update_user_model)
			ErrorHelper.CheckErr(err)
		}
		fmt.Println(_xyk_user.Username + ":信用卡分润相关数据已迁移,现在做了分润统计同步！")
	}
}

/*
根据Userid(递推)伞下直推汇总
UserId用户
*/
func Zhi_tui(UserId int) string {
	_str := ""
	_model := &Old.User_ext{}
	rows, err := DbHelper.MsSqlDb().Where("Parentid=?", UserId).Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(_model)
		ErrorHelper.CheckErr(err)
		_str += strconv.Itoa(_model.Userid) + "|"
	}
	_str = php2go.Rtrim(_str, "|")
	return _str
}

/*
根据Userid(递推)伞下间推汇总
UserId用户
*/
func Jian_tui(UserId int) string {
	_str := ""
	_model := &Old.User_ext{}
	rows, err := DbHelper.MsSqlDb().Where("Parentid=?", UserId).Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(_model)
		//ErrorHelper.CheckErr(err)
		if err == nil {
			_str += Jian_tui(_model.Id) + "|"
		}
	}
	_str = php2go.Rtrim(_str, "|")
	return _str
}

/*
全局用户从商城端登录
*/
func (this *RegisterAll_Controller) LoginShop() {
	_phone := this.GetString("phone")
	_password := this.GetString("password")
	fmt.Println(_phone + ":" + _password)
	this.Data["old"] = map[string]string{"name": Old_Query()}
	this.Data["new"] = map[string]string{"name": New_Query()}
	this.ServeJSON()
}

//-------------------------老数据库操作开始-----------------------------------//

func Old_Query() string {
	_model := &Old.User{Username: "13938202388"}
	results, _ := DbHelper.MsSqlDb().Get(_model)
	//sql := "select * from tb_user"
	//_models := []Old.User{}
	//results := DbHelper.DB().Sql(sql).Find(&_models)

	fmt.Println("%v", results)
	fmt.Println("%v", _model)
	return ""
}

//-------------------------老数据库操作结束-----------------------------------//

//-------------------------新数据库操作结束-----------------------------------//

func New_Query() string {
	_model := &New.User{Mobile: "13938202388"}
	results, _ := DbHelper.MySqlDb().Get(_model)
	//sql := "select * from lkt_user"
	//_models := []New.User{}
	//results := DbHelper.MySqlDb().Sql(sql).Find(&_models)

	fmt.Println("%v", results)
	fmt.Println("%v", _model)
	return ""
}

//-------------------------新数据库操作结束-----------------------------------//

//-------------------------老数登录注册实体结构开始-----------------------------------//
/*
信用卡后台登录返回对象
*/
type Login_xyk struct {
	StatusCode int
	Success    bool
	Info       string
	Data       Login_xyk_model
	Count      int
}

/*
返回内容实体
*/
type Login_xyk_model struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Truename   string `json:"truename"`
	Telephone  string `json:"telephone"`
	Smallimage string `json:"smallimage"`
	RongToken  string `json:"rongToken"`
}

/*
新后台登录成功返回实体
*/
type Login_shop_model struct {
	Code     int    `json:"code"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

//-------------------------老数登录注册实体结构结束-----------------------------------//

func daoshuju() {
	//FunRun("1694", "543", "200", "")
	//测试 更新信用卡订单表
	/*
		_update_model_order := shop.UserCardOrder{
			//UserId:      _UserId,            //用户id
			//Ordernob:    _return_35.OrderNo, //订单编号
			//Rateamountb: _RateAmountA,       //代扣手续费
			ReturnB: "_return_35", //代扣返回信息
			//	AddTime:     time.Now(),         //订单完成时间
		}
		_count, err := DbHelper.MySqlDb().Where("Id=?", strconv.Itoa(1)).Update(_update_model_order)
		ErrorHelper.CheckErr(err)
		fmt.Println(_count)
	*/
	//开始导入绑定银行卡信息
	/*
		_model_old := new(Old.Credit_card_info)
		rows, err := DbHelper.MsSqlDb().SQL("select Id,Memberid,Name,Certno,Bankname,Cardno,Phone,Addtime,Cardtype,Token,Expiretime,Cvv2,Iscashonly from view_tb_credit_card_info where [name]='陈小兵'").Rows(_model_old)
		ErrorHelper.CheckErr(err)
		//ErrorHelper.LogInfo(rows)
		for rows.Next() {
			_ = rows.Scan(_model_old)
			_model_new := new(New.User_bank_card)
			_model_new.StoreId = 8
			_model_new.UserId = _model_old.Memberid
			_model_new.Cardholder = _model_old.Name
			_model_new.IdCard = _model_old.Certno
			_model_new.BankCode = Get_Bank_Code(_model_old.Bankname)
			_model_new.BankName = _model_old.Bankname
			_model_new.BankCardNumber = _model_old.Cardno
			_model_new.Mobile = _model_old.Phone
			_model_new.AddDate = _model_old.Addtime
			_type, _ := strconv.Atoi(_model_old.Cardtype)
			if _type == 0 {
				_model_new.CardType = 1
			} else {
				_model_new.CardType = 2
			}
			_model_new.Token = _model_old.Token
			_model_new.Expiretime = _model_old.Expiretime
			_model_new.Cvv2 = _model_old.Cvv2
			_model_new.Iscashonly = _model_old.Iscashonly
			_, err := DbHelper.MySqlDb().Insert(_model_new)
			ErrorHelper.CheckErr(err)
		}
	*/
	//开始导实名认证信息
	//先删除信息
	/*
		_model_new_delete := New.User_auth{}
		conut, err := DbHelper.MySqlDb().Where("id>0").Delete(_model_new_delete)
		fmt.Println(conut)
		ErrorHelper.CheckErr(err)
		_model_old := new(Old.User_auth)
		rows, err := DbHelper.MsSqlDb().Where("id<>?", 0).Rows(_model_old)
		ErrorHelper.CheckErr(err)
		defer rows.Close()
		for rows.Next() {
			_ = rows.Scan(_model_old)
			_model_new := new(New.User_auth)
			_model_new.StoreId = 8
			_model_new.UserId = _model_old.Userid
			_model_new.Truename = _model_old.Truename
			_model_new.Idcard = _model_old.Idcard
			_model_new.Personimg = _model_old.Personimg
			_model_new.Personimgback = _model_old.Personimgback
			_model_new.Authflag = _model_old.Authflag
			_model_new.Value = _model_old.Value
			_model_new.Ysbflag = _model_old.Ysbflag
			_model_new.Ordertop = _model_old.Ordertop
			_model_new.Addtime = _model_old.Addtime
			_, err := DbHelper.MySqlDb().Insert(_model_new)
			ErrorHelper.CheckErr(err)
		}
	*/
	//开始导银行信息
	/*
		_model_old := new(Old.Channel_bank_kft)
		rows, err := DbHelper.MsSqlDb().Where("id<>?", 0).Rows(_model_old)
		ErrorHelper.CheckErr(err)
		defer rows.Close()
		for rows.Next() {
			_ = rows.Scan(_model_old)
			_, err := DbHelper.MySqlDb().Insert(_model_old)
			ErrorHelper.CheckErr(err)
		}
	*/
}

func Get_Bank_Code(_bank_name string) string {
	_rt := ""
	_model_new := new(New.Channel_bank_kft)
	has, err := DbHelper.MySqlDb().Where("bankName=?", _bank_name).Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		_rt = _model_new.Bankcode
	} else {
		_rt = ""
	}
	return _rt
}

/*
*分润处理
 */
func FunRun(user_id, order_no, amount, mark string) {
	type _rt struct {
		code int
		msg  string
	}
	// _HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	// _BodyData := map[string]interface{}{"user_id": user_id, "order_no": order_no, "type": "2", "amount": amount, "mark": mark}
	// _http_url := "https://shop.xhdncppf.com/index.php?module=app&action=index&store_id=8&app=calc_profit"
	// _req := WebHelper.HttpPost(_http_url, _HeaderData, _BodyData)
	// err := json.Unmarshal([]byte(_req), &_rt{})
	// ErrorHelper.CheckErr(err)
}
