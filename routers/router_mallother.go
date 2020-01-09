package routers

import (
	"github.com/qnsoft/mall_online/controllers/mall_other/taobao"

	"github.com/astaxie/beego"
)

func Mallother() {
	//阿里商品列表
	beego.Router("/api/mallother/taobao/goods", &taobao.Mall_other_taobao_Controller{}, "post:Goods")
	//阿里妈妈推广券
	beego.Router("/api/mallother/taobao/Taobao_YouhuiQuan", &taobao.Mall_other_taobao_Controller{}, "post:Taobao_YouhuiQuan")
	//阿里妈妈推广券新接口
	beego.Router("/api/mallother/taobao/Taobao_YouhuiQuan1", &taobao.Mall_other_taobao_Controller{}, "post:Taobao_YouhuiQuan1")
	//阿里妈妈推广券新接口搜索
	beego.Router("/api/mallother/taobao/Taobao_Search", &taobao.Mall_other_taobao_Controller{}, "post:Taobao_Search")
	//维易淘宝客API--获取授权地址url,该url地址提供给app用户点击授权
	beego.Router("/api/mallother/taobao/Get_auth_url?:user_id", &taobao.Wytklm_Controller{}, "get:Get_auth_url")
	//维易淘宝客API--高佣转链接口
	beego.Router("/api/mallother/taobao/Get_directhc?:num_iid/:relation_id", &taobao.Wytklm_Controller{}, "get:Get_directhc")
	//接收三方接口回传
	beego.Router("/api/mallother/taobao/Callback_other", &taobao.Mall_other_taobao_Controller{}, "post,get:Callback_other")
	//工具接口
	//beego.Router("/api/mallother/taobao/Get_xianjia", &taobao.Mall_other_taobao_Controller{}, "post:Get_xianjia")
}
