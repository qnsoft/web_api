package plugs

import (
	"github.com/qnsoft/web_api/plugs/OnlinePay/FengFu"

	"github.com/astaxie/beego"
)

/*
工易付相关路由
*/
func Plugs_onlinepay_FengFu() {
	//支付卡开通发短信
	beego.Router("/api/pay/FengFu/Card_Treaty", &FengFu.FengFu_Controller{}, "post:Card_Treaty")
	//支付卡开通发短信
	beego.Router("/api/pay/FengFu/Card_Treaty_Confirm", &FengFu.FengFu_Controller{}, "post:Card_Treaty_Confirm")
	//手动还款(代扣-代还)
	beego.Router("/api/pay/FengFu/Quick_Pay", &FengFu.FengFu_Controller{}, "post:Quick_Pay")
	//判断用户是否签约
	beego.Router("/api/pay/FengFu/Card_Treaty_IsOk", &FengFu.FengFu_Controller{}, "post,get:Card_Treaty_IsOk")

	//用户卡余额查询
	beego.Router("/api/pay/FengFu/YeCx", &FengFu.FengFu_Controller{}, "post,get:YeCx")
	//用户卡余额手提
	beego.Router("/api/pay/FengFu/RuJin", &FengFu.FengFu_Controller{}, "post,get:RuJin")
	//--------------以下是快捷支付----------------//
	//支付卡快捷支付开通发短信
	beego.Router("/api/pay/FengFukjzf/Card_Treaty", &FengFu.FengFuKjzf_Controller{}, "post,get:Card_Treaty")
	//支付卡快捷支付开通发短信确认
	beego.Router("/api/pay/FengFukjzf/Card_Treaty_Confirm", &FengFu.FengFuKjzf_Controller{}, "post:Card_Treaty_Confirm")
	//快捷支付
	beego.Router("/api/pay/FengFukjzf/Quick_Pay", &FengFu.FengFuKjzf_Controller{}, "post:Quick_Pay")
	//判断快捷支付用户是否签约
	beego.Router("/api/pay/FengFukjzf/Card_Treaty_IsOk", &FengFu.FengFuKjzf_Controller{}, "post,get:Card_Treaty_IsOk")
}
