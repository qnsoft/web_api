/**
 * [fix_android_ios 解决沉浸式的问题]
 * @author Hongwei
 * @param {[obj]} color [头部颜色 | darkgray 深灰色 | black 黑色 | green 绿色，如果不是以上这几种颜色，可以自定义颜色 ]
 * @link http://www.kancloud.cn/hongweizhiyuan/apicloud_function/270056
 * @eg fix_android_ios('darkgray')  or fix_android_ios('FF0000')
 */
function fix_android_ios(color){
	switch(color)
	{
		case 'darkgray':
			color = "#303247"
			break;
		case 'black':
			color = "#000000"
			break;
		case 'white':
			color = "#FFFFFF"
			break;
		case 'green':
			color = "#01b980"
			break;
		default:
			color = color
	}
	api.setStatusBarStyle({
		style : 'light',
		color: color
	});
	var header = document.querySelector('header');
	$api.fixIos7Bar(header);
	$api.fixStatusBar(header);
}
	function initRefreshHeader(objFunction){
		api.setCustomRefreshHeaderInfo({
				bgColor : '#ffffff',
				image : {
					pull : ['widget://images/rongCloud/1.png', 'widget://images/rongCloud/2.png', 'widget://images/rongCloud/3.png', 'widget://images/rongCloud/4.png', 'widget://images/rongCloud/5.png','widget://images/rongCloud/6.png','widget://images/rongCloud/7.png', 'widget://images/rongCloud/8.png'],
					load : ['widget://images/rongCloud/8.png', 'widget://images/rongCloud/7.png', 'widget://images/rongCloud/6.png', 'widget://images/rongCloud/5.png', 'widget://images/rongCloud/4.png','widget://images/rongCloud/3.png','widget://images/rongCloud/2.png', 'widget://images/rongCloud/1.png']
				}
			}, function() {
				//objFunction();
				setTimeout(function() {
					api.refreshHeaderLoadDone();
				}, 2000)
			});
	}

function noFindPic(){
	var img=event.srcElement;
	img.src="../../images/nopic.png";
	img.onerror=null; //控制不要一直跳动
} 
//打开新窗口
function openWindowName(name,_url){
	api.openWin({
        name: name,
        url: _url,
        opaque: true,
        reload:true,
        vScrollBarEnabled: false,  
        allowEdit:true,
        animation: {
                type: "movein",                //动画类型（详见动画类型常量）
                subType: "from_right",       //动画子类型（详见动画子类型常量）
                duration: 300                //动画过渡时间，默认300毫秒
      	  }
   	 });
};  
function openWindow(name,pId) {
	// api.closeFrameGroup({name:'product_group'});
	// var aa = parseInt(Math.random()*1000)+'aa';
	api.openWin({
		name:name,
		url:name+'.html',
        pageParam:{id:pId},
        reload:true,
        delay:100,
        bgColor: '#ffffff',
        allowEdit:true,
        animation: {
                type: "movein",                //动画类型（详见动画类型常量）
                subType: "from_right",       //动画子类型（详见动画子类型常量）
                duration: 300                //动画过渡时间，默认300毫秒
      	  }
	});
}
function openWindow(name,url,pId) {
	// api.closeFrameGroup({name:'product_group'});
	// var aa = parseInt(Math.random()*1000)+'aa';
	if(pId=="")pId=0;
	api.openWin({
		name:name,
		url:url,
        pageParam:{id:pId},
        reload:true,
        delay:100,
        bgColor: '#ffffff',
        allowEdit:true,
        animation: {
                type: "movein",                //动画类型（详见动画类型常量）
                subType: "from_right",       //动画子类型（详见动画子类型常量）
                duration: 300                //动画过渡时间，默认300毫秒
      	  }
	});
}
function openWindowObject(name,url,objData) {
	api.openWin({
		name:name,
		url:url,
        pageParam:objData,
        reload:true,
        delay:100,
        bgColor: '#ffffff',
        allowEdit:true,
        animation: {
                type: "movein",                //动画类型（详见动画类型常量）
                subType: "from_right",       //动画子类型（详见动画子类型常量）
                duration: 300                //动画过渡时间，默认300毫秒
      	  }
	});
}

function openWindowAttr(name,url,pId) {
	// api.closeFrameGroup({name:'product_group'});
	// var aa = parseInt(Math.random()*1000)+'aa';
	if(pId=="")pId=0;
	api.openWin({
		name:name,
		url:url,
        pageParam:{attrId:pId},
        reload:true,
        delay:100,
        bgColor: '#ffffff',
        allowEdit:true,
        animation: {
                type: "movein",                //动画类型（详见动画类型常量）
                subType: "from_right",       //动画子类型（详见动画子类型常量）
                duration: 300                //动画过渡时间，默认300毫秒
      	  }
	});
}

//打开新窗口，支持外部URL
function openFrame(name,url) {
	var w = api.screenWidth;
	api.openFrame({
    name: name,
    url: url,
    rect: {
        x: 0,
        y: 0,
        w: 'auto',
        h: 'auto'
    },
    bounces: true,
    bgColor: '#ffffff',
    vScrollBarEnabled: true,
    hScrollBarEnabled: true,
    showProgress:true,
    reload:true,
    allowEdit:true,
    bounces:false
  });
}
function openFrameParam(name,url,pId) {
	if(pId=="")pId=0;
	var w = api.screenWidth;
	api.openFrame({
    name: name,
    url: url,
    rect: {
        x: 0,
        y: 0,
        w: 'auto',
        h: 'auto'
    },
    bounces: true,
    bgColor: '#ffffff',
    pageParam:{id:pId},
    vScrollBarEnabled: true,
    hScrollBarEnabled: true,
    showProgress:true,
    reload:true,
    allowEdit:true,
    animation: {
       type: "movein",                //动画类型（详见动画类型常量）
       subType: "from_right",       //动画子类型（详见动画子类型常量）
       duration: 500                //动画过渡时间，默认300毫秒
    }
  });
}
//头部沉浸效果
function StatusHeaderOnly(){
        $api.fixStatusBar($api.dom('header'));
        api.setStatusBarStyle({
            style: 'light'
        });
}
function StatusHeader(name,url,footerH){
		var header = document.querySelector('header');
        var headerPos = $api.offset(header);	
    	api.openFrame({
	        name: name,
	        url: url,
	        pageParam:api.pageParam,
	        reload:true,
	        rect: {
		        x:0,
		        y:headerPos.h,
		        w:'auto',
		        h:api.frameHeight-footerH
	        }
        });
}
function StatusHeaderParam(name,url,footerH){
        $api.fixStatusBar($api.dom('header'));
        api.setStatusBarStyle({
            style: 'light'
        });
        var header = $api.byId('header');		
        var headerPos = $api.offset(header);	
    	api.openFrame({
	        name: name,
	        url: url+'_con.html',
			reload:true,	
        	rect: {
		        x:0,
		        y:headerPos.h,
		        w:'auto',
		        h:'auto'
	        },
	        pageParam:api.pageParam
        });
}
function initLazyLoad(){
	new auiLazyload({
        errorImage:'../images/person.png'
    });
}
function ConfirmFunc(objMsg,Func){
	api.confirm({
	    title: '提示',
    	msg: objMsg,
    	buttons: ['确定', '取消']
	}, Func);
}
/*消息通知*/
function noticeInfo(objContent){
	api.notification({
	    notify: {
       		content: objContent
	    }
	}, function(ret, err) {
	});
}
/*语音通知*/
function noticeVoice(objContent){
	var speechRecognizer = api.require('speechRecognizer');
	speechRecognizer.read({
   		readStr: objContent,
	    speed: 60,
	    volume: 100,
   		voice: 'vixy',
   		rate: 16000
	}, function(ret, err) {
	});
}
//打开指定的应用程序，没有打开默认的浏览器，指定 URL
function startApp(_appName,_url)
{
    api.openApp({
    	androidPkg: _appName,
    	mimeType: 'text/html',
    	uri: 'http://www.baidu.com'
	},function( ret, err ){
		if(ret)
		{}
		else
		{
			var strFileName=_url.substr(_url.lastIndexOf('/')+1);  
			down_load(_url,strFileName);
		}
 	});
}
//批量绑定 DOT 数据信息
var BindTemplateDot=function(ajaxUrl,tempArry,bindArry,bindSlider)
{
	//alert($api.getStorage($api.CONFIG_HEADER_NAME).appType);
   	api.ajax({
    		url: ajaxUrl,
	    	dataType: 'json',
	    	method: "get",
	    	headers:$api.getStorage($api.CONFIG_HEADER_NAME),
	    	async: false,
	    	cache: true
   	},function(ret, err){
       	if (ret)
       	{
        	//alert(JSON.stringify(ret));
        	for(var i=0;i<tempArry.length;i++)
        	{
	        	dotBind(tempArry[i],ret.Data,bindArry[i]);
        	}
        	if(bindSlider!="")
        	{
	        	initSlide(bindSlider);
        	}
        }
    });
}
/*IOS 获取 位置 信息*/
function GetLocalIOS(){
	api.getLocation(function(ret, err) {
		    if (ret && ret.status) {
		        //获取位置信息成功
		        getLocalIosSucc(ret.longitude, ret.latitude);
		        //initStoreList(ret.longitude, ret.latitude)
		    } else {
		        //alert(JSON.stringify(err));
		    }
	});
}
//绑定  指定 template
//templateId:模板ID <script id="dotId"></script>,
//DataList:传入的 数据
//bindId：需要绑定的 主键
/*
 demo:
var objList=[{id:1,title:"111"},{id:2,title:"2222"}];
dotBind("index-template",objList,"index-template-data")
var arrRes = Enumerable.From(myList).Where("x=>x.Age>18").ToArray();

 * */
/*
 * 
 * */
var dotBind=function(templateId,DataList,bindId){
	var tmpDom=$api.byId(templateId);
    var template =$api.html(tmpDom);
    var _htmlCode=doT.template(template)(DataList);
    $api.html($api.byId(bindId), _htmlCode);
}
var dotBindAppend=function(templateId,DataList,bindId){
	var tmpDom=$api.byId(templateId);
    var template =$api.html(tmpDom);
    var _htmlCode=doT.template(template)(DataList);
    $api.append($api.byId(bindId), _htmlCode);
}
var dotBindSilder=function(templateId,DataList,bindId,bindSlider){
    for(var i=0;i<templateId.length;i++)
    {
		var tmpDom=$api.byId(templateId[i]);
    	var template =$api.html(tmpDom);
    	var _htmlCode=doT.template(template)(DataList);
    	$api.html($api.byId(bindId[i]), _htmlCode);
    }
    if(bindSlider!="")
    {
		initSlide(bindSlider);
    }
}
function loadKeyValue(appKey)
{
	api.loadSecureValue({
	    key:appKey
	},function( ret, err ) {
    	alert(ret.value);
	});
}
//显示加载框 
//param:标题、模态 true/false 
function showProgress(_title,_modal)
{
	api.showProgress({
        title: _title,
        modal: _modal
    });
}
//关闭加载框 
function hidenProgress(){
  	api.hideProgress();
}
//监听返回键
function exitApp(){
		var ci = 0;
    	var mkeyTime = new Date().getTime();
		api.addEventListener({
			name : "keyback"
		}, function(ret, err) {
			//如果当前时间减去标志时间大于2秒，说明是第一次点击返回键，反之为2秒内点了2次，则退出。
			if ((new Date().getTime() - mkeyTime) > 2000) {
				mkeyTime = new Date().getTime();
				api.toast({
					msg : '再按一次退出APP',
					duration : 2000,
					location : 'bottom'
				});
				
			} else {
				api.closeWidget({
					id:api.appId,
					silent:true
				});
			}
		});
}
//后退
 function backOrClose()
 {
	window.location.reload();    //刷新
	window.history.go(-1);
 }

//退出当前应用
function closeApp(_silent)
{
	api.closeWidget({
    	id: api.appId,
    	retData: {
        	name: 'closeWidget'
	    },
    	silent:_silent,
	});

}
//设置 偏好数据
function setPrefs(_key,_value)
{
	api.setPrefs({
		key:_key,
		value:_value
	});
}
//获取偏好数据
function getPrefs(_key)
{
    //同步返回结果：
var userName = api.getPrefs({
    sync: true,
    key: _key
});
return userName;
	/*//异步返回结果
	api.getPrefs({
	    key:_key
    },function(ret,err){
    	//coding...
    	return ret.value;
    });
    */
}
//监听是否有网
//unknown            //未知
//ethernet        //以太网
//wifi            //wifi
//2g                //2G网络
//3g                //3G网络
//4g                //4G网络
//none            //无网络
function add_listener_online()
{
 	var _line=api.connectionType;
 	switch(_line)
 	{
 	 case "none":
 	 	api.alert({
   			title: '提示',
    		msg: '无网络，请检查',
    		buttons: ['确定']
		},function( ret, err ){
    		closeApp(true);
		});
	 	 break;
 	}
}
/*demo
       api.ajax({
	    url:$api.homeUrl,
	    dataType: 'json',
	    method:"post",
	    data: {values: {id:100 }},
	    async:false,
	    cache:true	    
    },function(ret,err){
            alert(JSON.stringify(ret));

    });
 * 
 * */
//跨域获取数据,不带参数
function ajax_url_data(_url,sqlFormat)
{
	api.ajax({
	    url:_url,
	    dataType: 'json',
	    async:false,
	    cache:true	    
    },function(ret,err){
		   _rtn_data= ret;
        if (ret) {     
        var _strTemp="";
			for(var i=0;i<ret.length;i++)
			{
				_strTemp=sqlFormat.format(ret[i]);
				//alert(_strTemp);
				executeSQL(_strTemp);
			}
		} else {
           	_rtn_data=JSON.stringify(err);
        }
    });
}
function ajaxUrl(objModel,ajaxSuccFunc)
{
    $.ajax({
        url: objModel.Url,
        type: objModel.Method,
        dataType: 'json',
        timeout: 5000,
        data: objModel.Data,
        headers: objModel.Headers,
        success: ajaxSuccFunc,
        fail: function (err, status) {
            console.log(err)
        }
    })
}
/*弹出一个选择框，用户选择图片来源*/
function getPicture(){
	api.confirm({
		title:"提示",
		msg:"请选择图片来源",
		buttons:["照相","相册","取消"]
    },function(ret,err){
    	//coding...
    	if(3==ret.buttonIndex){
    		return;
    	}
    	var sourceType="album";
    	if(1==ret.buttonIndex){
    		/*打开相机*/
    		sourceType="camera";
    	}
    	api.getPicture({
        	sourceType:sourceType,
        	encodingType:"jpg",
        	mediaValue:"pic"
        },function(ret,err){
        	//上传图片到服务器
        	if(ret)
        	{
        		var _path=ret.data;
        		uploadImage(_path);
        	}
        });
    });
}

function uploadImage(_path)
{
	api.ajax({
   					url: _uploadPath,
    				method: 'post',
    				timeout: 30,
	    			dataType: 'json',
    				returnAll: false,
			    	data: {
        				values: {name: 'haha'},
        				files: {Filedata: _path}
				    }
				},function( ret, err ){
    				if( ret ){
			    	     alert( JSON.stringify( ret ) );
				    }else{
				         alert( JSON.stringify( err) );
				    }
				});
}

//swip 幻灯JS 特效
  var initSlide=function (obj) {
        var $pointer = $api.byId(obj);
        window.mySlide = Swipe(slide, {
            auto: 3000,
            continuous: true,
            disableScroll: true,
            stopPropagation: true,
            callback: function (index, element) {
                console.log();
                var $actA = $api.dom($pointer, 'a.active');
                $api.removeCls($actA, 'active');
                $api.addCls($api.eq($pointer, index + 1), 'active');
            },
	            transitionEnd: function (index, element) {
            }
        });
    }

/*
 *直接拨打电话
 * */
function btnCall(btnNumber){
    api.call({
   		type: 'tel_prompt',
   		number: btnNumber
	});
}
//下载APK 文件  并实现安装
function down_load(_url,_fileName)
{
	var bool = _url.indexOf("http");
	if(bool<0){
		_url=_serverURL+_url;
	}
	showProgress("正在下载",true);
	var _savePath= "fs://"+_fileName;
	api.download({
		url: _url,
    	savePath: _savePath,
		report: true,
    	cache: true,
    	allowResume: true
		},function( ret, err ){
		    if(ret.state==1){
				api.hideProgress();
				api.installApp({appUri: _savePath});
		    }else{
        		showProgress("已下载:"+ret.percent+"%",true);
		    }
	});
}


//数据库操作,执行sql语句
function executeSQL(_sql)
{
	var db=api.require('db');
	db.openDatabase({
	    name:_DATABASENAME
    },function(ret,err){
    	//coding...
    	if(!ret.status)
    	{
    		toast(err.msg);
    		return;
    	}
    	db.executeSql({
	        name:_DATABASENAME,
	        sql:_sql
        },function(ret,err){
        	//coding...
        	if(!ret.status)
        	{
        		toast(err.msg);
        		return;
        	}
        });
    });
}

//数据库操作,查询sql语句
function selectSQL(_sql)
{
	var db=api.require('db');
	db.openDatabase({
	    name:_DATABASENAME
    },function(ret,err){
    	//coding...
    	if(!ret.status)
    	{
    		toast(err.msg);
    		return;
    	}
    	db.selectSql({
	        name:_DATABASENAME,
	        sql:_sql
        },function(ret,err){
        	//coding...
        	if(!ret.status)
        	{
        		toast(err.msg);
        		return;
        	}        	
        	alert(ret.data);
        });
    });
}

/*图片缓存*/
function fnLoadImage(ele_) {
     var dataUrl = $api.attr(ele_, 'data-url');
     //alert(fnLoadImage)
     if (dataUrl) {
        api.imageCache({
            url : dataUrl
        }, function(ret, err) {
           if (ret) {                        
              ele_.src = ret.url;
              //alert(ret.url);
           $api.attr(ele_, 'data-url', '');
        } else {
        }
      });
   }
}
/**
 * 替换所有匹配exp的字符串为指定字符串
 * @param exp 被替换部分的正则
 * @param newStr 替换成的字符串
 */
String.prototype.replaceAll = function (exp, newStr) {
    return this.replace(new RegExp(exp, "gm"), newStr);
};

/**
 * 原型：字符串格式化
 * @param args 格式化参数值
 */
String.prototype.format = function(args) {
    var result = this;
    if (arguments.length < 1) {
        return result;
    }

    var data = arguments; // 如果模板参数是数组
    if (arguments.length == 1 && typeof (args) == "object") {
        // 如果模板参数是对象
        data = args;
    }
    for ( var key in data) {
        var value = data[key];
        if (undefined != value) {
            result = result.replaceAll("\\{" + key + "\\}", value);
        }
    }
    return result;
}

/**/
/*判断 两个时间 差*/
function diffDateFunc(evalue) {
	var dB = new Date(evalue.replace(/T/g," ").replace(/-/g, "/"));
	if (new Date() > Date.parse(dB)) {return true;} 
	return false;
}
/********************/
function UserLogin(){
	var appConfig=$api.getStorage($api.CONFIG_HEADER_NAME);
	if(appConfig.uid>0)return true;
	return false;		
}