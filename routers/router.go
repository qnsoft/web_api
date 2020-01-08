package routers

import (
	"qnsoft/web_api/controllers"
	"qnsoft/web_api/controllers/Demo"
	"qnsoft/web_api/controllers/Duijie"
	"qnsoft/web_api/controllers/ImageUpload"
	"qnsoft/web_api/controllers/OnlinePay"
	"qnsoft/web_api/controllers/Sms"
	"qnsoft/web_api/controllers/Token"
	api "qnsoft/web_api/controllers/api/login"
	"qnsoft/web_api/controllers/api/xyk"
	"qnsoft/web_api/routers/plugs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//加入跨域权限
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		AllowCredentials: true}))
	beego.SetStaticPath("/upload", "upload") //重定向静态路径
	beego.SetStaticPath("/views", "views")   //重定向静态路径
	//-------------------------基础接口开始-----------------------------------//
	//根目录
	beego.Router("/", &controllers.Default_Controller{}, "get,post:Get")
	//先获取access_token
	beego.Router("/access_token", &Token.AccesstokenController{}, "post:Access_Token")
	//token测试
	beego.Router("/testtoken", &controllers.Default_Controller{}, "get,post:TestToken")
	//上传图片
	beego.Router("/img/upload_pic", &ImageUpload.Image_Uplaod_Controller{}, "post:Uplaod_Pic")
	//获取图片
	beego.Router("/img/info_pic", &ImageUpload.Image_Uplaod_Controller{}, "get:Info_Pic")
	//发送短信验证码
	beego.Router("/sms/MsgSend", &Sms.Sms_Controller{}, "post:MsgSend")
	//-------------------------基础接口结束-----------------------------------//
	//-------------------------对接接口开始-----------------------------------//
	beego.Router("/duijie/RegisterAll", &Duijie.RegisterAll_Controller{}, "post:RegisterAll")
	//老用户信用卡登录
	beego.Router("/duijie/LoginXyk", &Duijie.RegisterAll_Controller{}, "post:LoginXyk")
	//-------------------------对接接口结束-----------------------------------//

	//-------------------------DEMO接口开始-----------------------------------//
	//二维码推广图片
	beego.Router("/demo/QrCodeDemo", &Demo.QrCode_Controller{}, "post:QrCodeDemo")
	//-------------------------信用卡接口开始-----------------------------------//
	//信用卡测试
	beego.Router("/OnlinePay/kft/XYK_Demo1", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_Demo1")
	//信用卡 3.5单笔付款接口
	beego.Router("/OnlinePay/kft/XYK_35", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_35")
	//信用卡 3.5单笔付款接口(小额)
	beego.Router("/OnlinePay/kft/XYK_35_Small", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_35_Small")
	//信用卡 3.6快捷协议代扣协议申请接口
	beego.Router("/OnlinePay/kft/XYK_36", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_36")
	//信用卡 3.7快捷协议代扣协议确定接口
	beego.Router("/OnlinePay/kft/XYK_37", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_37")
	//信用卡 3.8快捷协议代扣接口
	beego.Router("/OnlinePay/kft/XYK_38", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_38")
	//信用卡 3.9快捷协议代扣协议查询接口
	beego.Router("/OnlinePay/kft/XYK_39", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_39")
	//信用卡 3.10快捷协议代扣协议解除接口
	beego.Router("/OnlinePay/kft/XYK_310", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_310")
	//信用卡 3.11用户资金查询接口
	beego.Router("/OnlinePay/kft/XYK_311", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_311")
	//信用卡 3.11用户资金查询接口(小额)
	beego.Router("/OnlinePay/kft/XYK_311_Small", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_311_Small")
	//信用卡 3.12交易查询接口
	beego.Router("/OnlinePay/kft/XYK_312", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_312")
	//信用卡 3.13查询账户余额接口
	beego.Router("/OnlinePay/kft/XYK_313", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_313")
	//信用卡 3.14交易类对账文件获取接口
	beego.Router("/OnlinePay/kft/XYK_314", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_314")
	//信用卡 3.15银行卡三要素验证接口
	beego.Router("/OnlinePay/kft/XYK_315", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_315")
	//信用卡 3.16验证类对账文件获取接口
	beego.Router("/OnlinePay/kft/XYK_316", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_316")
	//信用卡 3.17对账结算单笔付款接口
	beego.Router("/OnlinePay/kft/XYK_317", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_317")
	//信用卡 对账单笔付款协议申请接口
	beego.Router("/OnlinePay/kft/XYK_318", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_318")
	//信用卡 单笔付款协议查询接口
	beego.Router("/OnlinePay/kft/XYK_319", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_319")
	//信用卡 结算查询账户余额接口
	beego.Router("/OnlinePay/kft/XYK_322", &OnlinePay.OnlinePay_KFT_Controller{}, "post:XYK_322")
	/*--------线上正式交易接口-------------*/
	//信用卡用户登录
	beego.Router("/api/login/UserLogin", &api.Login_Controller{}, "post:UserLogin")
	//信用卡用户修改密码
	beego.Router("/api/login/UserChangePwd", &api.Login_Controller{}, "post:UserChangePwd")
	//信用卡用户找回密码
	beego.Router("/api/login/UserFindPwd", &api.Login_Controller{}, "post:UserFindPwd")
	//信用卡用户获取余额(大额)
	beego.Router("/api/login/Get_User_XykYe", &api.Login_Controller{}, "post:Get_User_XykYe")
	//信用卡用户获取余额(小额)
	beego.Router("/api/login/Get_User_XykYe_Small", &api.Login_Controller{}, "post:Get_User_XykYe_Small")
	//信用卡用户根据卡获取还款任务计划列表
	beego.Router("/api/xyk/Bind_Card_Order_JobList", &xyk.Xyk_Controller{}, "post:Bind_Card_Order_JobList")
	//信用卡用户根据UserID获取自动消费明细列表
	beego.Router("/api/xyk/Bind_User_Order_JobList", &xyk.Xyk_Controller{}, "post:Bind_User_Order_JobList")
	//信用卡用户根据UserID获取手动消费明细列表
	beego.Router("/api/xyk/Bind_User_Order_List", &xyk.Xyk_Controller{}, "post:Bind_User_Order_List")
	//信用卡 智能代还任务订单取消
	beego.Router("/api/xyk/Bind_Card_Order_JobList_Cancel", &xyk.Xyk_Controller{}, "post:Bind_Card_Order_JobList_Cancel")
	//信用卡 信用卡解绑
	beego.Router("/api/xyk/Bind_Card_List_Cancel", &xyk.Xyk_Controller{}, "post:Bind_Card_List_Cancel")
	//信用卡 提现(大额)
	// beego.Router("/api/xyk/Bind_Card_TiXian", &xyk.Xyk_Controller{}, "post:Bind_Card_TiXian")
	// //信用卡 提现(小额额)
	// beego.Router("/api/xyk/Bind_Card_TiXian_Small", &xyk.Xyk_Controller{}, "post:Bind_Card_TiXian_Small")
	//信用卡 代扣接口
	beego.Router("/OnlinePay/kft/Quick_Pay", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Quick_Pay")
	//信用卡 代还接口
	beego.Router("/OnlinePay/kft/Quick_Pay_Ok", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Quick_Pay_Ok")
	//信用卡 代还成功返回订单接口
	beego.Router("/OnlinePay/kft/Card_Order_GetOne", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Order_GetOne")
	//信用卡 智能代还任务订单生成
	beego.Router("/OnlinePay/kft/Get_Create_ZnDhList", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Get_Create_ZnDhList")
	//信用卡 智能代还任务订单确认
	beego.Router("/OnlinePay/kft/Get_Create_ZnDhList_Confirm", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Get_Create_ZnDhList_Confirm")
	//信用卡 代扣异步回传接口
	beego.Router("/OnlinePay/kft/Notify_kft", &OnlinePay.OnlinePay_KFT_Controller{}, "get,post:Notify_kft")
	//信用卡 公众号银行列表
	beego.Router("/OnlinePay/kft/Bind_Bank_List", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Bind_Bank_List")
	//信用卡 公众号用户绑卡
	beego.Router("/OnlinePay/kft/Bind_Card_Add", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Bind_Card_Add")
	//信用卡 公众号用户卡列表
	beego.Router("/OnlinePay/kft/Bind_Card_List", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Bind_Card_List")
	//信用卡 公众号用户卡解绑
	beego.Router("/OnlinePay/kft/Bind_Card_Relieve", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Bind_Card_Relieve")
	//信用卡 公众号获取用户单卡信息
	beego.Router("/OnlinePay/kft/Bind_Card_GetOne", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Bind_Card_GetOne")
	//信用卡 根据用户id和卡号检测用户当前出金卡是否签约
	beego.Router("/OnlinePay/kft/Card_Treaty_IsOk", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Treaty_IsOk")
	//信用卡 根据用户id和卡号检测用户当前出金卡是否签约小额
	beego.Router("/OnlinePay/kft/Card_Treaty_IsOk_Samll", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Treaty_IsOk_Samll")
	//信用卡 信用卡协议签约
	beego.Router("/OnlinePay/kft/Card_Treaty", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Treaty")
	//信用卡 信用卡协议签约(小额)
	beego.Router("/OnlinePay/kft/Card_Treaty_Small", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Treaty_Small")
	//信用卡 信用卡协议签约确认
	beego.Router("/OnlinePay/kft/Card_Treaty_Confirm", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Treaty_Confirm")
	//信用卡 信用卡协议签约确认(小额)
	beego.Router("/OnlinePay/kft/Card_Treaty_Confirm_Small", &OnlinePay.OnlinePay_KFT_Controller{}, "post:Card_Treaty_Confirm_Small")
	//信用卡 检测用户是否实名认证接口
	beego.Router("/OnlinePay/kft/User_Is_Auth", &OnlinePay.OnlinePay_KFT_Controller{}, "post:User_Is_Auth")
	//信用卡 用户实名认证接口
	beego.Router("/OnlinePay/kft/User_Auth_Add", &OnlinePay.OnlinePay_KFT_Controller{}, "post:User_Auth_Add")
	//信用卡 用户实名认证后台审核接口
	beego.Router("/OnlinePay/kft/User_Auth_Allow", &OnlinePay.OnlinePay_KFT_Controller{}, "post:User_Auth_Allow")
	//-------------------------导数据接口开始-----------------------------------//
	//导数据
	beego.Router("/demo/Daoshuju", &Duijie.RegisterAll_Controller{}, "post:Daoshuju")
	//-------------------------DEMO接口结束-----------------------------------//
	//与淘宝对接接口
	Mallother()
	//-------------------------信用卡支付相关接口--------------------------------//
	//工易付对接接口,新接口升级老接口暂停
	//plugs.Plugs_onlinepay_gongyifu()
	//丰付接口对接
	plugs.Plugs_onlinepay_FengFu()
}
