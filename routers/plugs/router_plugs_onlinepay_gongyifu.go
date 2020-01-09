package plugs

import (
	gongyifuF "github.com/qnsoft/web_api/plugs/OnlinePay/GongYiFu"

	"github.com/astaxie/beego"
)

/*
工易付相关路由
*/
func Plugs_onlinepay_gongyifu() {
	//S101.商户信息提交
	beego.Router("/api/pay/gongyifu/Gyf_101", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_101")
	//S102.商户信息修改
	beego.Router("/api/pay/gongyifu/Gyf_102", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_102")
	//S103.商户信息费率
	beego.Router("/api/pay/gongyifu/Gyf_103", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_103")
	//S201.银联页面绑卡
	beego.Router("/api/pay/gongyifu/Gyf_201", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_201")
	//S201.银联页面绑卡成功通知页面
	beego.Router("/api/pay/gongyifu/Gyf_201_Ok", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_201_Ok")
	//S202.绑卡查询
	beego.Router("/api/pay/gongyifu/Gyf_202", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_202")
	//S301.订单查询
	beego.Router("/api/pay/gongyifu/Gyf_301", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_301")
	//S302.快捷支付
	beego.Router("/api/pay/gongyifu/Gyf_302", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_302")
	//S302.落地云闪付支付
	beego.Router("/api/pay/gongyifu/Gyf_302_A", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_302_A")
	//S302.子商户结算接口
	beego.Router("/api/pay/gongyifu/Gyf_302_B", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_302_B")
	//S302.子商户余额查询
	beego.Router("/api/pay/gongyifu/Gyf_302_C", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_302_C")

	//S401.绑卡异步通知
	beego.Router("/api/pay/gongyifu/Gyf_S401", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_S401")
	//S402.交易异步通知
	beego.Router("/api/pay/gongyifu/Gyf_S402", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_S402")
	//S501.地区编码查询
	beego.Router("/api/pay/gongyifu/Gyf_S501", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_S501")

	//---------------以下与数据库对接----------------------//
	//判断用户是否注册 如果没有注册直接注册并保存到数据库
	beego.Router("/api/pay/gongyifu/Gyf_IsZhuCe", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_IsZhuCe")
	//判断用户是否签约
	beego.Router("/api/pay/gongyifu/Gyf_IsQianYue", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_IsQianYue")
	//获取工易付签约商户号
	beego.Router("/api/pay/gongyifu/Gyf_Submerchid", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_Submerchid")
	//判断用户是否签约 如果没签约跳到签约界面签约
	beego.Router("/api/pay/gongyifu/Gyf_IsQianYueHtml", &gongyifuF.GongYiFu_Controller{}, "post,get:Gyf_IsQianYueHtml")
	//执行代扣操作，落地云闪付与数据库对接
	beego.Router("/api/pay/gongyifu/Gyf_DaiKou", &gongyifuF.GongYiFu_Controller{}, "post:Gyf_DaiKou")

}
