package Demo

import (
	"image"
	"os"
	"path/filepath"
	"strings"
	"github.com/qnsoft/web_api/controllers/Token"
	"github.com/qnsoft/web_api/models/Duijie/New"
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"github.com/qnsoft/utils/FileHelper"
	"github.com/qnsoft/utils/PicHelper"
	"github.com/qnsoft/utils/StringHelper"

	"github.com/astaxie/beego"
	"github.com/boombuler/barcode/qr"
	_ "github.com/mattn/go-adodb"
)

/**
*信息实体
 */
type QrCode_Controller struct {
	Token.BaseController
}

func (this *QrCode_Controller) QrCodeDemo() {
	var _rt interface{}
	_is_refresh, _ := this.GetInt("is_refresh")
	_mobile := this.GetString("mobile")
	_model := &New.User{Zhanghao: _mobile}
	results, err := DbHelper.MySqlDb().Get(_model)
	ErrorHelper.CheckErr(err)
	// 文件后缀
	ext := ".jpeg" //path.Ext(filenameWithSuffix)
	if results {

		_Qr_img := "nopic"
		_dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) //获取当前执行文件物理路径
		_Old_img := _dir + "\\upload\\qrcode\\QR_" + _mobile + ext
		if _is_refresh == 0 {
			//判断二维码是否存在
			if FileHelper.CheckNotExist(_Old_img) { //如果不存在直接调取生成
				//获取生成后的二维码地址
				_Qr_img = Create_Qr_pic(_model, ext)
				_rt = map[string]interface{}{"code": 1, "msg": "二维码已生成!", "src": beego.AppConfig.String("server_path::PrefixUrl") + _Qr_img}

			} else { //如果存在直接调取
				_Qr_img = "upload/qrcode/QR_" + _mobile + ext
				_rt = map[string]interface{}{"code": 1, "msg": "二维码获取!", "src": beego.AppConfig.String("server_path::PrefixUrl") + _Qr_img}
			}
		} else {
			//获取生成后的二维码地址
			_Qr_img = Create_Qr_pic(_model, ext)
			_rt = map[string]interface{}{"code": 1, "msg": "二维码已生成!", "src": beego.AppConfig.String("server_path::PrefixUrl") + _Qr_img}
		}
	} else {
		_rt = map[string]interface{}{"code": 0, "msg": "请给接口参数赋正确的值!"}
	}
	this.Data["json"] = _rt
	this.ServeJSON()
}

//生成二维码图片
func Create_Qr_pic(_model *New.User, ext string) string {
	_Pic_Dis := new(PicHelper.Pic_Dispose)
	//获取生成后的二维码地址
	_Qr_img := ""
	qrc := StringHelper.NewQrCode("https://shop.ljz789.com/register/?linkCode="+_model.Zhanghao, 240, 240, qr.Q, qr.Auto)
	path := StringHelper.GetQrCodeFullPath()
	_qr_img, _path, err := qrc.Encode(path)
	ErrorHelper.CheckErr(err)
	ErrorHelper.LogInfo("生成二维码执行", _qr_img)
	bg_pic_path := beego.AppConfig.String("server_path::QrCodebgPic")
	new_pic_path := _path + "66666_" + strings.Replace(_qr_img, ".png", ".jpeg", 1)

	_Bg_Pic_Model := PicHelper.Pic_Model{Path: bg_pic_path}
	_Ft_Pic_Model := PicHelper.Pic_Model{Path: _path + _qr_img, P: image.Pt(255, 658)}
	new_bg_img := _Pic_Dis.Pic_pic_ompose(_Bg_Pic_Model, _Ft_Pic_Model, new_pic_path) //图片与图片合成

	_Ft_Text_Model := PicHelper.Pic_Text{
		Text:  string(_model.RealName), //真实姓名
		Color: [4]uint8{7, 7, 7, 255},
		//	FontFile:  beego.AppConfig.String("server_path::FontPath") + "simhei.ttf",
		Size:      0.36,
		Linewidth: 2,
		Angle:     0.3,
		Space:     "   ",
		Px:        590,
		Py:        786,
	}
	_Bg_Pic_Model1 := PicHelper.Pic_Model{Path: new_bg_img}
	// 图片与文字合成文件
	new_img_qr := beego.AppConfig.String("server_path::RuntimeRootPath") + "QR_" + _model.Zhanghao + ext
	err0 := _Pic_Dis.Pic_text_ompose(_Bg_Pic_Model1, _Ft_Text_Model, new_img_qr) //图片与文字合成
	ErrorHelper.CheckErr(err0)
	_Bg_Pic_Model_a := PicHelper.Pic_Model{Path: new_img_qr}
	_touxiang := beego.AppConfig.String("server_path::RuntimeRootPath") + "LOGO.png"
	if len(_model.Headimgurl) > 5 {
		_touxiang = _model.Headimgurl
	}
	_Ft_Pic_Model_a := PicHelper.Pic_Model{Path: _touxiang, P: image.Pt(146, 330), IsScale: true, Width: 120, Height: 120}
	//new_pic_path_a := "User_" + new_img_qr
	new_bg_img_a := _Pic_Dis.Pic_pic_ompose(_Bg_Pic_Model_a, _Ft_Pic_Model_a, new_img_qr) //图片与头像合成最终效果图
	//执行完之后删除之前的二维码和不含水印文字的合成图
	if FileHelper.CheckNotExist("/" + _path + _qr_img) {
		err := FileHelper.DeleteFile(_path + _qr_img) //删除二维码
		ErrorHelper.CheckErr(err)
	}
	if FileHelper.CheckNotExist("/" + new_pic_path) {
		err := FileHelper.DeleteFile(new_pic_path) //删除合成图
		ErrorHelper.CheckErr(err)
	}
	//判断二维码是否生成成功
	if FileHelper.CheckNotExist("/" + new_bg_img_a) {
		_Qr_img = new_bg_img_a
	} else {
		_Qr_img = "nopic"
	}
	return _Qr_img
}

//生成二维码图片
func Create_Qr_pic_AA(_model *New.User, ext string) string {
	//获取生成后的二维码地址
	_Qr_img := ""
	qrc := StringHelper.NewQrCode("https://shop.ljz789.com/register/?linkCode="+_model.Zhanghao, 218, 218, qr.Q, qr.Auto)
	path := StringHelper.GetQrCodeFullPath()
	_qr_img, _path, err := qrc.Encode(path)
	ErrorHelper.CheckErr(err)
	ErrorHelper.LogInfo("生成二维码执行")
	ErrorHelper.LogInfo(_qr_img)
	//ErrorHelper.LogInfo(_path)
	ErrorHelper.LogInfo("二维码生成出错：", err)
	bg_pic_path := beego.AppConfig.String("server_path::QrCodebgPic")
	new_pic_path := _path + "66666_" + _qr_img
	new_img := FileHelper.Pic_layer(bg_pic_path, _path+_qr_img, new_pic_path)
	markText := FileHelper.MarkText{
		Text:      string(_model.RealName), //真实姓名
		Color:     [4]uint8{8, 8, 8, 122},
		FontFile:  beego.AppConfig.String("server_path::RuntimeRootPath") + "/font/msyh.ttc",
		Size:      0.30,
		Linewidth: 2,
		Angle:     0.3,
		Space:     "   ",
	}
	// 创建水印文件
	new_img_qr := beego.AppConfig.String("server_path::RuntimeRootPath") + "QR_" + _model.Zhanghao + ext
	f, _ := os.Create(new_img_qr)
	// 打水印
	img, _ := FileHelper.MarkingPicture_Text_One(new_img, markText)
	// 将 image 写入 buffur
	buff, _ := FileHelper.ImageToBuffer(img, ext)
	// buffer 写入水印文件
	buff.WriteTo(f)
	defer f.Close()
	//执行完之后删除之前的二维码和不含水印文字的合成图
	if FileHelper.CheckNotExist("/" + _path + _qr_img) {
		err := FileHelper.DeleteFile(_path + _qr_img) //删除二维码
		ErrorHelper.CheckErr(err)
	}
	if FileHelper.CheckNotExist("/" + new_pic_path) {
		err := FileHelper.DeleteFile(new_pic_path) //删除合成图
		ErrorHelper.CheckErr(err)
	}
	//判断二维码是否生成成功
	if FileHelper.CheckNotExist("/" + new_img_qr) {
		_Qr_img = new_img_qr
	} else {
		_Qr_img = "nopic"
	}
	return _Qr_img
}
