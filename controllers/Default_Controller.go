package controllers

import (
	"fmt"

	"qnsoft/web_api/controllers/Token"
	"qnsoft/web_api/models/oauth"
)

/**
*信息实体
 */
type Default_Controller struct {
	Token.BaseController
}

func (this *Default_Controller) Get() {
	_Model_User := new(oauth.OauthUser)
	fmt.Println(_Model_User)
	this.Data["json"] = map[string]string{"name": "api服务!"}
	this.ServeJSON()
}

func (this *Default_Controller) TestToken() {
	if this.Check_Token() {
		this.Data["json"] = map[string]string{"name": "成功通过token验证!"}
		this.ServeJSON()
	}
}
