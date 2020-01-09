package bll

import (
	"github.com/qnsoft/web_api/models/shop"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
)

/*
*获取单条信息信息
 */
func User_Auth_Get_One(_model *shop.UserAuth) (bool, *shop.UserAuth) {
	_model_new := &shop.UserAuth{UserId: _model.UserId}
	results, err := DbHelper.MySqlDb().Get(_model)
	ErrorHelper.CheckErr(err)
	return results, _model_new
}
