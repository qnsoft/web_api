/*
 * APICloud JavaScript Library
 * Copyright (c) 2014 apicloud.com
 */
(function(window){
    var u = {};
    var isAndroid = (/android/gi).test(navigator.appVersion);
    var uzStorage = function(){
        var ls = window.localStorage;
        if(isAndroid){
           ls = os.localStorage();
        }
        return ls;
    };
    function parseArguments(url, data, fnSuc, dataType) {
        if (typeof(data) == 'function') {
            dataType = fnSuc;
            fnSuc = data;
            data = undefined;
        }
        if (typeof(fnSuc) != 'function') {
            dataType = fnSuc;
            fnSuc = undefined;
        }
        return {
            url: url,
            data: data,
            fnSuc: fnSuc,
            dataType: dataType
        };
    }
    u.trim = function(str){
        if(String.prototype.trim){
            return str == null ? "" : String.prototype.trim.call(str);
        }else{
            return str.replace(/(^\s*)|(\s*$)/g, "");
        }
    };
    u.trimAll = function(str){
        return str.replace(/\s*/g,'');
    };
    u.isElement = function(obj){
        return !!(obj && obj.nodeType == 1);
    };
    u.isArray = function(obj){
        if(Array.isArray){
            return Array.isArray(obj);
        }else{
            return obj instanceof Array;
        }
    };
    u.isEmptyObject = function(obj){
        if(JSON.stringify(obj) === '{}'){
            return true;
        }
        return false;
    };
    u.addEvt = function(el, name, fn, useCapture){
        if(!u.isElement(el)){
            console.warn('$api.addEvt Function need el param, el param must be DOM Element');
            return;
        }
        useCapture = useCapture || false;
        if(el.addEventListener) {
            el.addEventListener(name, fn, useCapture);
        }
    };
    u.rmEvt = function(el, name, fn, useCapture){
        if(!u.isElement(el)){
            console.warn('$api.rmEvt Function need el param, el param must be DOM Element');
            return;
        }
        useCapture = useCapture || false;
        if (el.removeEventListener) {
            el.removeEventListener(name, fn, useCapture);
        }
    };
    u.one = function(el, name, fn, useCapture){
        if(!u.isElement(el)){
            console.warn('$api.one Function need el param, el param must be DOM Element');
            return;
        }
        useCapture = useCapture || false;
        var that = this;
        var cb = function(){
            fn && fn();
            that.rmEvt(el, name, cb, useCapture);
        };
        that.addEvt(el, name, cb, useCapture);
    };
    u.dom = function(el, selector){
        if(arguments.length === 1 && typeof arguments[0] == 'string'){
            if(document.querySelector){
                return document.querySelector(arguments[0]);
            }
        }else if(arguments.length === 2){
            if(el.querySelector){
                return el.querySelector(selector);
            }
        }
    };
    u.domAll = function(el, selector){
        if(arguments.length === 1 && typeof arguments[0] == 'string'){
            if(document.querySelectorAll){
                return document.querySelectorAll(arguments[0]);
            }
        }else if(arguments.length === 2){
            if(el.querySelectorAll){
                return el.querySelectorAll(selector);
            }
        }
    };
    u.byId = function(id){
        return document.getElementById(id);
    };
    u.first = function(el, selector){
        if(arguments.length === 1){
            if(!u.isElement(el)){
                console.warn('$api.first Function need el param, el param must be DOM Element');
                return;
            }
            return el.children[0];
        }
        if(arguments.length === 2){
            return this.dom(el, selector+':first-child');
        }
    };
    u.last = function(el, selector){
        if(arguments.length === 1){
            if(!u.isElement(el)){
                console.warn('$api.last Function need el param, el param must be DOM Element');
                return;
            }
            var children = el.children;
            return children[children.length - 1];
        }
        if(arguments.length === 2){
            return this.dom(el, selector+':last-child');
        }
    };
    u.eq = function(el, index){
        return this.dom(el, ':nth-child('+ index +')');
    };
    u.not = function(el, selector){
        return this.domAll(el, ':not('+ selector +')');
    };
    u.prev = function(el){
        if(!u.isElement(el)){
            console.warn('$api.prev Function need el param, el param must be DOM Element');
            return;
        }
        var node = el.previousSibling;
        if(node.nodeType && node.nodeType === 3){
            node = node.previousSibling;
            return node;
        }
    };
    u.next = function(el){
        if(!u.isElement(el)){
            console.warn('$api.next Function need el param, el param must be DOM Element');
            return;
        }
        var node = el.nextSibling;
        if(node.nodeType && node.nodeType === 3){
            node = node.nextSibling;
            return node;
        }
    };
    u.closest = function(el, selector){
        if(!u.isElement(el)){
            console.warn('$api.closest Function need el param, el param must be DOM Element');
            return;
        }
        var doms, targetDom;
        var isSame = function(doms, el){
            var i = 0, len = doms.length;
            for(i; i<len; i++){
                if(doms[i].isEqualNode(el)){
                    return doms[i];
                }
            }
            return false;
        };
        var traversal = function(el, selector){
            doms = u.domAll(el.parentNode, selector);
            targetDom = isSame(doms, el);
            while(!targetDom){
                el = el.parentNode;
                if(el != null && el.nodeType == el.DOCUMENT_NODE){
                    return false;
                }
                traversal(el, selector);
            }

            return targetDom;
        };

        return traversal(el, selector);
    };
    u.contains = function(parent,el){
        var mark = false;
        if(el === parent){
            mark = true;
            return mark;
        }else{
            do{
                el = el.parentNode;
                if(el === parent){
                    mark = true;
                    return mark;
                }
            }while(el === document.body || el === document.documentElement);

            return mark;
        }
        
    };
    u.remove = function(el){
        if(el && el.parentNode){
            el.parentNode.removeChild(el);
        }
    };
    u.attr = function(el, name, value){
        if(!u.isElement(el)){
            console.warn('$api.attr Function need el param, el param must be DOM Element');
            return;
        }
        if(arguments.length == 2){
            return el.getAttribute(name);
        }else if(arguments.length == 3){
            el.setAttribute(name, value);
            return el;
        }
    };
    u.removeAttr = function(el, name){
        if(!u.isElement(el)){
            console.warn('$api.removeAttr Function need el param, el param must be DOM Element');
            return;
        }
        if(arguments.length === 2){
            el.removeAttribute(name);
        }
    };
    u.hasCls = function(el, cls){
        if(!u.isElement(el)){
            console.warn('$api.hasCls Function need el param, el param must be DOM Element');
            return;
        }
        if(el.className.indexOf(cls) > -1){
            return true;
        }else{
            return false;
        }
    };
    u.addCls = function(el, cls){
        if(!u.isElement(el)){
            console.warn('$api.addCls Function need el param, el param must be DOM Element');
            return;
        }
        if('classList' in el){
            el.classList.add(cls);
        }else{
            var preCls = el.className;
            var newCls = preCls +' '+ cls;
            el.className = newCls;
        }
        return el;
    };
    u.removeCls = function(el, cls){
        if(!u.isElement(el)){
            console.warn('$api.removeCls Function need el param, el param must be DOM Element');
            return;
        }
        if('classList' in el){
            el.classList.remove(cls);
        }else{
            var preCls = el.className;
            var newCls = preCls.replace(cls, '');
            el.className = newCls;
        }
        return el;
    };
    u.toggleCls = function(el, cls){
        if(!u.isElement(el)){
            console.warn('$api.toggleCls Function need el param, el param must be DOM Element');
            return;
        }
       if('classList' in el){
            el.classList.toggle(cls);
        }else{
            if(u.hasCls(el, cls)){
                u.removeCls(el, cls);
            }else{
                u.addCls(el, cls);
            }
        }
        return el;
    };
    u.val = function(el, val){
        if(!u.isElement(el)){
            console.warn('$api.val Function need el param, el param must be DOM Element');
            return;
        }
        if(arguments.length === 1){
            switch(el.tagName){
                case 'SELECT':
                    var value = el.options[el.selectedIndex].value;
                    return value;
                    break;
                case 'INPUT':
                    return el.value;
                    break;
                case 'TEXTAREA':
                    return el.value;
                    break;
            }
        }
        if(arguments.length === 2){
            switch(el.tagName){
                case 'SELECT':
                	var kk = el.options;  
					for (var i=0; i<kk.length; i++) {  
						if (kk[i].value==val) {  
							kk[i].selected=true;
						} 
					} 
                    return el;
                    break;
                case 'INPUT':
                    el.value = val;
                    return el;
                    break;
                case 'TEXTAREA':
                    el.value = val;
                    return el;
                    break;
            }
        }
        
    };
    u.prepend = function(el, html){
        if(!u.isElement(el)){
            console.warn('$api.prepend Function need el param, el param must be DOM Element');
            return;
        }
        el.insertAdjacentHTML('afterbegin', html);
        return el;
    };
    u.append = function(el, html){
        if(!u.isElement(el)){
            console.warn('$api.append Function need el param, el param must be DOM Element');
            return;
        }
        el.insertAdjacentHTML('beforeend', html);
        return el;
    };
    u.before = function(el, html){
        if(!u.isElement(el)){
            console.warn('$api.before Function need el param, el param must be DOM Element');
            return;
        }
        el.insertAdjacentHTML('beforebegin', html);
        return el;
    };
    u.after = function(el, html){
        if(!u.isElement(el)){
            console.warn('$api.after Function need el param, el param must be DOM Element');
            return;
        }
        el.insertAdjacentHTML('afterend', html);
        return el;
    };
    u.html = function(el, html){
        if(!u.isElement(el)){
            console.warn('$api.html Function need el param, el param must be DOM Element');
            return;
        }
        if(arguments.length === 1){
            return el.innerHTML;
        }else if(arguments.length === 2){
            el.innerHTML = html;
            return el;
        }
    };
    u.text = function(el, txt){
        if(!u.isElement(el)){
            console.warn('$api.text Function need el param, el param must be DOM Element');
            return;
        }
        if(arguments.length === 1){
            return el.textContent;
        }else if(arguments.length === 2){
            el.textContent = txt;
            return el;
        }
    };
    u.offset = function(el){
        if(!u.isElement(el)){
            console.warn('$api.offset Function need el param, el param must be DOM Element');
            return;
        }
        var sl = Math.max(document.documentElement.scrollLeft, document.body.scrollLeft);
        var st = Math.max(document.documentElement.scrollTop, document.body.scrollTop);

        var rect = el.getBoundingClientRect();
        return {
            l: rect.left + sl,
            t: rect.top + st,
            w: el.offsetWidth,
            h: el.offsetHeight
        };
    };
    u.css = function(el, css){
        if(!u.isElement(el)){
            console.warn('$api.css Function need el param, el param must be DOM Element');
            return;
        }
        if(typeof css == 'string' && css.indexOf(':') > 0){
            el.style && (el.style.cssText += ';' + css);
        }
    };
    u.cssVal = function(el, prop){
        if(!u.isElement(el)){
            console.warn('$api.cssVal Function need el param, el param must be DOM Element');
            return;
        }
        if(arguments.length === 2){
            var computedStyle = window.getComputedStyle(el, null);
            return computedStyle.getPropertyValue(prop);
        }
    };
    u.jsonToStr = function(json){
        if(typeof json === 'object'){
            return JSON && JSON.stringify(json);
        }
    };
    u.strToJson = function(str){
        if(typeof str === 'string'){
            return JSON && JSON.parse(str);
        }
    };
    u.setStorage = function(key, value){
        if(arguments.length === 2){
            var v = value;
            if(typeof v == 'object'){
                v = JSON.stringify(v);
                v = 'obj-'+ v;
            }else{
                v = 'str-'+ v;
            }
            var ls = uzStorage();
            if(ls){
                ls.setItem(key, v);
            }
        }
    };
    u.getStorage = function(key){
        var ls = uzStorage();
        if(ls){
            var v = ls.getItem(key);
            if(!v){return;}
            if(v.indexOf('obj-') === 0){
                v = v.slice(4);
                return JSON.parse(v);
            }else if(v.indexOf('str-') === 0){
                return v.slice(4);
            }
        }
    };
    u.rmStorage = function(key){
        var ls = uzStorage();
        if(ls && key){
            ls.removeItem(key);
        }
    };
    u.clearStorage = function(){
        var ls = uzStorage();
        if(ls){
            ls.clear();
        }
    };

   
    /*by king*/
    u.fixIos7Bar = function(el){
	    return u.fixStatusBar(el);
        if(!u.isElement(el)){
            console.warn('$api.fixIos7Bar Function need el param, el param must be DOM Element');
            return;
        }
        var strDM = api.systemType;
        if (strDM == 'ios') {
            var strSV = api.systemVersion;
            var numSV = parseInt(strSV,10);
            var fullScreen = api.fullScreen;
            var iOS7StatusBarAppearance = api.iOS7StatusBarAppearance;
            if (numSV >= 7 && !fullScreen && iOS7StatusBarAppearance) {
                el.style.paddingTop = '20px';
            }
        }
    };
    u.fixStatusBar = function(el){
	    if(!u.isElement(el)){
	        console.warn('$api.fixStatusBar Function need el param, el param must be DOM Element');
	        return 0;
	    }
	    el.style.paddingTop = api.safeArea.top+'px';
	    return el.offsetHeight;
    
        if(!u.isElement(el)){
            console.warn('$api.fixStatusBar Function need el param, el param must be DOM Element');
            return;
        }
        var sysType = api.systemType;
        if(sysType == 'ios'){
            u.fixIos7Bar(el);
        }else if(sysType == 'android'){
            var ver = api.systemVersion;
            ver = parseFloat(ver);
            if(ver >= 4.4){
                el.style.paddingTop = '25px';
            }
        }
    };
    u.fixTabBar = function(el){
    	if(!u.isElement(el)){
        	console.warn('$api.fixTabBar Function need el param, el param must be DOM Element');
	        return 0;
    	}
    	el.style.paddingBottom = api.safeArea.bottom + 'px';
	    return el.offsetHeight;
	}
    u.toast = function(title, text, time){
        var opts = {};
        var show = function(opts, time){
            api.showProgress(opts);
            setTimeout(function(){
                api.hideProgress();
            },time);
        };
        if(arguments.length === 1){
            var time = time || 500;
            if(typeof title === 'number'){
                time = title;
            }else{
                opts.title = title+'';
            }
            show(opts, time);
        }else if(arguments.length === 2){
            var time = time || 500;
            var text = text;
            if(typeof text === "number"){
                var tmp = text;
                time = tmp;
                text = null;
            }
            if(title){
                opts.title = title;
            }
            if(text){
                opts.text = text;
            }
            show(opts, time);
        }
        if(title){
            opts.title = title;
        }
        if(text){
            opts.text = text;
        }
        time = time || 500;
        show(opts, time);
    };
    u.post = function(/*url,data,fnSuc,dataType*/){
        var argsToJson = parseArguments.apply(null, arguments);
        var json = {};
        var fnSuc = argsToJson.fnSuc;
        argsToJson.url && (json.url = argsToJson.url);
        argsToJson.data && (json.data = argsToJson.data);
        if(argsToJson.dataType){
            var type = argsToJson.dataType.toLowerCase();
            if (type == 'text'||type == 'json') {
                json.dataType = type;
            }
        }else{
            json.dataType = 'json';
        }
        json.method = 'post';
        api.ajax(json,
            function(ret,err){
                if (ret) {
                    fnSuc && fnSuc(ret);
                }
            }
        );
    };
    u.get = function(/*url,fnSuc,dataType*/){
        var argsToJson = parseArguments.apply(null, arguments);
        var json = {};
        var fnSuc = argsToJson.fnSuc;
        argsToJson.url && (json.url = argsToJson.url);
        //argsToJson.data && (json.data = argsToJson.data);
        if(argsToJson.dataType){
            var type = argsToJson.dataType.toLowerCase();
            if (type == 'text'||type == 'json') {
                json.dataType = type;
            }
        }else{
            json.dataType = 'text';
        }
        json.method = 'get';
        api.ajax(json,
            function(ret,err){
                if (ret) {
                    fnSuc && fnSuc(ret);
                }
            }
        );
    };

/*end*/

/*========自定义内容========*/
u.imgServerUrl ="http://api.xhdncppf.com";  //正式
u.serverUrl ="";  //正式
u.webUrl ="http://credit.xhdncppf.com";  //正式

u.StoreId=2;
u.pageSize=20;
u.CONFIG_USER_INFO="user_info";
u.CONFIG_USER_ID="userId";
//$api.getStorage($api.CONFIG_HEADER_NAME)
//u.CONFIG_HEADER_INFO="appType="+u.systemType+"&appVersion="+u.systemVersion+"&appOS="+u.appId+"&sid="+u.deviceId+"&uid={0}";
u.CONFIG_HEADER_NAME="app_info";
u.CONFIG_HEADER_JSON={"appType":"","appVersion":"","appOS":"","sid":"","uid":0,"userName":"","encryptPwd":""};
u.CONFIG_LOCATION_NAME="locationInfo";
//u.CONFIG_LOCATION_JSON={"Lng":"113.666902","Lat":"34.71436","Province":"河南省","City":"郑州市","AreaInfo":"","Address":"","SearchCity":"郑州市"};
u.CONFIG_LOCATION_JSON={"Lng":"","Lat":"","Province":"河南省","City":"郑州市","AreaInfo":"","Address":"","SearchCity":"郑州市"};
u.CONFIG_STOREINFO="storeFlag";
u.CONFIG_LOCATION_INFO="location";
u.CONFIG_LOCATION_INFO_JSON={"Lng":"","Lat":""};
u.CONFIG_STOREINFO_JSON={storeFlag:false,IsSuccess:false,foodFlag:false};
u.CONFIG_RONGCLOUD_TOKEN="rongToken";

/*获取  APP  Flag标志*/
u.Flag_MenuList=u.serverUrl+"/ApiTest/GetMenuFlag";
u.Flag_FriendList=u.serverUrl+"/ApiTest/GetFriendFlag";

//获取 广告信息列表
u.Index_AdvertUrl=u.serverUrl+"/category/GetAdvertList?categoryid={0}&topNumber={1}";
//获取 产品 特性 列表信息 参数：特征(0代表精品,1代表热销,2代表新品)
u.Index_ProductTrait=u.serverUrl+"/api/catalog/getstoreproductlist?count={0}&storeId={1}&storeCid={2}&trait={3}";

//获取 产品类别列表 信息
u.Category_List=u.serverUrl+"/category/list";
u.Category_All=u.serverUrl+"/category/getall";

/*Article*/
u.ArticleList=u.serverUrl+"/Article/Index";


/*产品 相关接口*/
//获取 产品列表 信息
//参数：分类 ID、排序列[列名]、排序方向、当前页
u.Win_ProductList=u.serverUrl+"/product/List";//?id={0}&title={1}&lowPrice={2}&maxPrice={3}&categoryId={4}&orderby={5}&ordername={6}&priceId={7}&product_type={8}&pageSize={9}&lng={10}&lat={11}&city={12}
//var objData={id:0,title:"",lowPrice:-1,maxPrice:-1,categoryId:-1,orderby:"desc",ordername:"price_auction",priceId:-1,pageSize:10,Lng:"",Lat:"",City:""};
u.Win_ProductListJson=u.serverUrl+"/product/List_Param";
//获取 产品 信息
u.Win_ProductDetails=u.serverUrl+"/api/catalog/product?pid={0}";
//获取产品详细描述
u.Win_ProductDetailsInfo=u.serverUrl+"/product/detail";
u.Win_ProductDetailsBase=u.serverUrl+"/product/BaseDetail";

//商品评价
u.Product_Comment=u.serverUrl+"/product/CommentList?id={0}";

u.SearchList=u.serverUrl+"/product/SearchList?topNumber={0}";
u.SearchAdd=u.serverUrl+"/product/SearchInsert?title={0}";
u.Message=u.serverUrl+"/account/message";
u.NearStore=u.serverUrl+"/user/GetNearStoreList"; //附近

/*购物车 相关接口*/
//批量添加商品
u.Cart_AddList=u.serverUrl+"/api/cart/AddProductList?cartJson={0}";
u.Cart_List_Json={"Pid":0,"SaleCount":0};
//添加商品到购物车
u.Cart_Add=u.serverUrl+"/cart/Insert";//?productId={0}&number={1}";
//更新购物车
u.Cart_Update=u.serverUrl+"/cart/Update";//?productId={0}&number={1}";
//删除购物车
u.Cart_Delete=u.serverUrl+"/cart/Delete";//?productId={0}";
//清空购物车
u.Cart_Clear=u.serverUrl+"/cart/clear";
//获取购物车列表信息
u.Cart_List=u.serverUrl+"/cart/list";

//提交订单
u.Order_SubmitOrder=u.serverUrl+"/order/ConfirmOrder";//?payName={0}&payCreditCount={1}&couponIdList={2}&fullCut={3}";
//获取 确认订单 信息
u.Order_ConfirmOrder=u.serverUrl+"/order/ConfirmOrder";//?user_address={0}&payment_id={1}";
// 更新 订单状态 
u.Order_UpdateStatus=u.serverUrl+"/order/UpdateSuccess"
//提交订单，返回订单号、总额
u.Order_SubmitOrder=u.serverUrl+"/api/order/SubmitOrder?payName={0}&saId={1}";
//订单汇总
u.Order_StatusSum=u.serverUrl+"/user/OrderStatusList";

u.Order_Comment=u.serverUrl+"/flow/Comment"; //data:orderNumber,content

u.Order_Payment_WxApp=u.serverUrl+"/wxpayapp/wxpay"
u.Order_PayWxAppSucc=u.webUrl+"/Alipay/AppWxPayment"


//积分 
u.User_GetUserPoint=u.serverUrl+"/user/GetUserPoint"; //data:orderNumber,content
u.User_PointList=u.serverUrl+"/user/GetUserPointList"; //data:orderNumber,content
u.User_PointOrder=u.serverUrl+"/user/UserPointOrder"; //data:orderNumber,content
u.User_PointRecharge=u.serverUrl+"/user/UserPointRecharge"; //data:orderNumber,content
u.User_PointExchange=u.serverUrl+"/user/Point_Exchange"; //data:orderNumber,content

//会员登陆 注册 
u.User_Login=u.serverUrl+"/account/login?username={0}&userPass={1}";
u.User_Register=u.serverUrl+"/account/register";//?userName={0}&userPass={1}&regCode={2}&telePhone={3}&smsCaptcha={4}";
u.User_Info=u.serverUrl+"/user/GetModel";//?Lng={0}&Lat={1}
u.User_InfoName=u.serverUrl+"/user/GetUserInfo";//?Lng={0}&Lat={1}
u.User_BaseInfo=u.serverUrl+"/user/GetBaseInfo";
u.User_UpdateInfo=u.serverUrl+"/user/UpdateInfo";
u.User_Image=u.serverUrl+"/user/UpdateImage";
u.User_GetImage=u.serverUrl+"/user/GetUserImage";
u.User_Forget=u.serverUrl+"/account/ForgetPass";
u.User_UpdateLoginInfo=u.serverUrl+"/user/UpdateLoginInfo";

/*实名认证*/
u.User_Auth_Create=u.serverUrl+"/user/UserAuthCreate";
u.User_Auth_Image=u.serverUrl+"/user/UserAuthImageCreate";
u.User_Auth_Info=u.serverUrl+"/user/UserAuthInfo";
u.User_Auth_Check=u.serverUrl+"/user/UserAuthCheck";
u.User_Auth_Token=u.serverUrl+"/user/GetCreditToken";

//获取 会员 收货地址列表
u.User_AddressList=u.serverUrl+"/user/Select_User_Address";
u.User_AddressAdd=u.serverUrl+"/user/Add_User_Address";
u.User_AddressDel=u.serverUrl+"/user/Delete_User_Address?id={0}";
u.User_AddressGet=u.serverUrl+"/user/Get_User_Address?id={0}";
u.User_CheckPass=u.serverUrl+"/user/Check_User_Pass?userPass={0}"; //Post
u.User_UpdatePass=u.serverUrl+"/user/Update_Pass"; //Post
u.User_Draw_Pass=u.serverUrl+"/user/Update_Draw_Pass"; //提现 密码 Post

u.User_Order_List=u.serverUrl+"/order/OrderList?status={0}&orderNumber={1}&pageIndex={2}&pageSize={3}";
u.User_Transfer=u.serverUrl+"/user/Transfer";
u.User_Browser=u.serverUrl+"/user/Get_User_Browser";//会员浏览记录

/*视频列表*/
u.Audio_GetModel=u.serverUrl+"/Article/GetAudioModel";//获取 音频信息
u.Attach_Detail=u.serverUrl+"/Article/AttrchDetail";//获取附件

/*招聘列表*/
u.ZhaoPin_List=u.serverUrl+"/zhaopin/list";//会员浏览记录
u.ZhaoPin_WebList=u.serverUrl+"/zhaopin/ListWeb";//会员浏览记录
u.ZhaoPin_Create=u.serverUrl+"/zhaopin/create";//会员浏览记录
u.ZhaoPin_Update=u.serverUrl+"/zhaopin/update";//会员浏览记录
u.ZhaoPin_Delete=u.serverUrl+"/zhaopin/delete";//会员浏览记录
u.ZhaoPin_Model=u.serverUrl+"/zhaopin/GetModel";//会员浏览记录

/*在线充值*/
    u.User_Recharge = u.serverUrl + "/order/OrderRecharge";//会员充值
    u.User_RechargeSucc = u.serverUrl + "/order/OrderRehageSucc";//会员充值
    u.UserMoneyAll = u.serverUrl + "/user/GetUserMoneyAll";//会员余额
    u.UserCheckShare = u.serverUrl + "/user/CheckUserShare";//会员余额
    u.UserMoney = u.serverUrl + "/user/GetUserMoney";//会员余额
    u.UserBankList = u.serverUrl + "/user/GetUserBankList";//会员银行列表
    u.UserWithDraw = u.serverUrl + "/user/UserWithDraw";//会员提现
    u.UserWithDrawList = u.serverUrl + "/user/GetUserDrawList";//会员提现列表
    u.UserShare = u.serverUrl + "/user/GetUserShare";//会员提现
    u.UserCoupon = u.serverUrl + "/user/GetUserCoupon";//会员优惠券
/*通知*/
u.UserNotice=u.serverUrl+"/user/GetUserNotice";//会员消息通知
u.UserNoticeCreate=u.serverUrl+"/user/UserNoticeCreate";//消息通知
u.UserNoticeUpdate=u.serverUrl+"/user/UserNoticeUpdate";//消息通知
u.UserNoticeList=u.serverUrl+"/Notice/GetNoticeList";//消息通知
u.UserNoticeNear=u.serverUrl+"/Notice/GetNearList";//消息通知

u.UserNoticeCheck=u.serverUrl+"/Notice/Check";//消息通知
u.UserNoticeDelete=u.serverUrl+"/Notice/Delete";//消息通知
u.SystemNoticeList=u.serverUrl+"/Notice/GetSystemNoticeList";//消息通知

/*信用卡接口*/
u.CreditList=u.serverUrl+"/Credit/GetUserCardList";
u.CreditUnBind=u.serverUrl+"/Credit/H5_UnBind_Card";
u.CreditUpdate=u.serverUrl+"/Credit/CreditUpdate";
u.CreditBill=u.serverUrl+"/Credit/CreateBill";
u.CreditBillConfirm=u.serverUrl+"/Credit/CreateBillConfirm";
u.CreditBillList=u.serverUrl+"/Credit/GetBillList";
u.CreditBillCancel=u.serverUrl+"/Credit/CancelBill";
u.CreditDebitCard=u.serverUrl+"/Credit/CreditDebitCard";
u.CreditAccInfo=u.serverUrl+"/Credit/GetBillDay";
u.CreditAccCreate=u.serverUrl+"/Credit/CreatBillDay";
u.CreditDebitSucc=u.serverUrl+"/Credit/GetCreditSucc";
    u.CreditBindApi = u.serverUrl + "/credit/CreditBind";
u.CreditBindChannel = u.serverUrl + "/credit/CreditBindByChannelNew";

u.CreditUnBindH5=u.webUrl+"/Credit/H5_Bind_Card";
u.CreditPayFirst=u.webUrl+"/Credit/H5_OnlinePay_First";

/*朋友列表*/
u.friendList=u.serverUrl+"/Friend/GetFirendList";
u.friendInvite=u.serverUrl+"/Friend/Invite";
u.friendInviteList=u.serverUrl+"/Friend/GetInviteList";
u.friendAccept=u.serverUrl+"/Friend/FriendAccept";
u.friendNearList=u.serverUrl+"/Friend/NearFirendList";
u.friendSearch=u.serverUrl+"/Friend/Search";

/*融云相关接口*/
u.rong_Token=u.serverUrl+"/Rongcloud/GetToken";
//根据 会员主键  获取token
u.rong_UserToken=u.serverUrl+"/Rongcloud/GetUserToken";


/*VIP 中心*/
u.Vip_Money=u.serverUrl+"/user/vipmoney";
u.Vip_Order=u.serverUrl+"/user/VipOrder?paymentid={0}";
u.Offline=u.serverUrl+"/user/Select_Offline";
u.AgentList=u.serverUrl+"/user/AgentList";

/*云片网 短信验证*/
u.SmsSend=u.serverUrl+"/account/sendsms?telephone={0}";
u.SendSMSForget=u.serverUrl+"/account/sendsmsForget?telephone={0}";
u.SmsValidForm=u.serverUrl+"/api/home/smsCheck?telephone={0}";

/*会员中心*/
u.center_baseInfo=u.serverUrl+"/api/ucenter/baseInfo";
u.center_updatePass=u.serverUrl+"/api/ucenter/UpdatePassword?oldPasswrod={0}&password={1}&confirmPwd={2}";
u.center_orderList=u.serverUrl+"/api/ucenter/OrderList?page={0}&startAddTime={1}&endAddTime{2}&orderState={3}";
u.center_orderInfo=u.serverUrl+"/api/ucenter/OrderInfo?oid={0}";
u.center_cancelOrder=u.serverUrl+"/api/ucenter/CancelOrder?oid={0}&cancelReason={1}";
u.center_ReceiveOrder=u.serverUrl+"/api/ucenter/ReceiveOrder?oid={0}";
u.center_Favorite=u.serverUrl+"/api/ucenter/AjaxFavoriteProductList?page={0}";
u.center_FavoriteAdd=u.serverUrl+"/api/ucenter/AddProductToFavorite?pid={0}";
u.center_FavoriteDel=u.serverUrl+"/api/ucenter/DelFavoriteProduct?pid={0}";

u.center_StoreFavorite=u.serverUrl+"/api/ucenter/AjaxFavoriteStoreList?page={0}";
u.center_StoreFavoriteAdd=u.serverUrl+"/api/ucenter/AddStoreToFavorite?storeId={0}";
u.center_StoreFavoriteDel=u.serverUrl+"/api/ucenter/DelFavoriteStore?storeId={0}";

u.center_AddressList=u.serverUrl+"/api/ucenter/ShipAddressList";
u.center_AddressAdd=u.serverUrl+"/api/ucenter/AddShipAddress?regionId={0}&alias{1}&consignee={2}&mobile={3}&phone={4}&email={5}&zipcode={6}&address={7}&isDefault={8}";
u.center_AddressEdit=u.serverUrl+"/api/ucenter/EditShipAddress?regionId={0}&alias{1}&consignee={2}&mobile={3}&phone={4}&email={5}&zipcode={6}&address={7}&isDefault={8}&saId={9}";
u.center_AddressDel=u.serverUrl+"/api/ucenter/DelShipAddress?saId={0}";
u.center_AddressDefault=u.serverUrl+"/api/ucenter/SetDefaultShipAddress?saId={0}";

u.center_CouponList=u.serverUrl+"/api/ucenter/CouponList?type={0}";


/*店铺 信息*/
u.store_Apply=u.serverUrl+"/user/Apply_Store";
u.store_GetStoreInfo=u.serverUrl+"/user/GetStoreInfo";
u.store_GetStoreStatus=u.serverUrl+"/user/GetStoreStatus";
u.store_ProductList=u.serverUrl+"/product/Store_Product_List";
u.store_ProductAdd=u.serverUrl+"/product/Create";
u.store_ProductDel=u.serverUrl+"/user/Store_Delete";
u.store_ProductInfo=u.serverUrl+"/user/Store_GetProduct";
u.store_Browser=u.serverUrl+"/user/Create_User_Browser";
u.store_ProImgDel=u.serverUrl+"/product/Delete_Image";

/*银行卡信息*/
u.BankList=u.serverUrl+"/user/bankList";
u.BankModel=u.serverUrl+"/user/GetBankModel";
u.BankCreate=u.serverUrl+"/user/BankCreate";
u.BankDelete=u.serverUrl+"/user/BankDelete";

/*========车险========*/
u.CarSafeList=u.serverUrl+"/carSafe/Index"; //条件 筛选
u.Province=u.serverUrl+"/carSafe/GetProvince";
u.CarSafeOrder=u.serverUrl+"/carSafe/CarSafeOrder";
u.CarSafeOrderUpdate=u.serverUrl+"/carSafe/UpdateCarImg";
u.CarSafeOrderUpdate=u.serverUrl+"/carSafe/UpdateCarImg";
u.CarSafeOrderSucc=u.serverUrl+"/carSafe/CarSafeOrderSucc";
u.CarSafeOrderList=u.serverUrl+"/carSafe/OrderList";
u.GetCarOrderDetail=u.serverUrl+"/carSafe/GetCarOrderDetail";
u.CarSafeModel=u.serverUrl+"/carSafe/CarSafeModel";
u.GetCarOrder=u.serverUrl+"/carSafe/GetCarOrder"; //获取 车险订单信息
u.CarSafePay=u.serverUrl+"/carSafe/CarSafePay";

/*========课科========*/
u.BookOrder=u.serverUrl+"/book/CreateOrder";
u.BookOrderSucc=u.serverUrl+"/book/BookOrderSucc";
u.BookOrderCheck=u.serverUrl+"/book/Check";
u.BookDetail=u.serverUrl+"/product/BookDetail";
u.BookOrderList=u.serverUrl+"/book/OrderList";

/*=========直播===========*/
u.LiveUserList=u.serverUrl+"/live/index";
u.LiveRoomModel=u.serverUrl+"/live/GetLiveRoomInfo";
u.LiveUserModel=u.serverUrl+"/live/GetLiveUserModel";
u.LiveViewModel=u.serverUrl+"/live/GetLiveViewModel";
u.LiveUpdateFlag=u.serverUrl+"/live/UpdateFlag";


/*========在线留言========*/
u.MessageCreate=u.serverUrl+"/Message/Create"; //条件 筛选

/*========收入统计========*/
u.ReportDay=u.serverUrl+"/Report/Day"; //条件 筛选
u.ReportMonth=u.serverUrl+"/Report/Month"; //条件 筛选
u.ReportPocket=u.serverUrl+"/Report/PocketIndex"; //条件 筛选
u.ReportClassRate=u.serverUrl+"/Report/BookRateList"; //条件 筛选
u.ReportConsoleRate=u.serverUrl+"/Report/ConsoleRateList"; //条件 筛选


/*=======咨询列表========*/
u.ConsultList=u.serverUrl+"/product/ConsultList"; //获取 咨询列表
u.ConsultDetail=u.serverUrl+"/product/ConsultDetail"; //获取 咨询列表
u.ConsultCheck=u.serverUrl+"/Consult/Check"; //
u.ConsultOrder=u.serverUrl+"/Consult/CreateOrder"; //咨询订单
u.ConsultSucc = u.serverUrl + "/Consult/OrderSucc"; //支付成功
u.GetCardPhone = u.serverUrl + "/credit/GetCardPhone"; //获取 签约 手机号

u.ChannelRate = u.webUrl + "/Credit/GetChannelByCardId";//获取通道费率
u.ChannelKftSign = u.serverUrl + "/Credit/ChannelKftSign";//快付通发送短信
u.ChannelKftSignConfirm = u.serverUrl + "/Credit/ChannelKftSignConfirm";//快付通确认短信，绑定卡
u.ChannelKftConfirmOrder = u.serverUrl + "/credit/ChannelKftConfirmOrder"; //签约成功更新订单
u.GetChannelList = u.serverUrl + "/credit/GetChannelList"; //签约成功更新订单

u.ChannelJFTSign = u.serverUrl + "/Credit/ChannelJftSign";//佳付通发送短信
u.ChannelJFTSignConfirm = u.serverUrl + "/Credit/ChannelJftSignConfirm";//佳付通发送短信

/*========自定义内容========*/
//获取 广告缩略图地址
u.GetAdvertImage=function(objUrl){
	var imgUrl= u.serverUrl+objUrl;	//var imgUrl= "http://106.3.45.164:9901"+objUrl;
	return imgUrl;
}
u.GetUserImage=function(objUrl){
	var imgUrl= u.serverUrl+objUrl;
	return imgUrl;
}
u.GetBookFile=function(objUrl){
	var filePath=u.webUrl+objUrl;
	return filePath;
}
//获取 商品缩略图地址
u.GetProductImage=function(storeId,width,height,objUrl){
	return u.serverUrl+"/upload/store/"+storeId+"/product/show/thumb"+width+"_"+height+"/"+objUrl;
}
/*
 * 获取文件名，带 扩展名
 * */
u.getFileName=function(o){
	var pos=o.lastIndexOf("/");
   	return o.substring(pos+1);
}
/*
 * 获取文件名，带 扩展名
 * */
u.getFileExtion=function(o){
	var pos=o.lastIndexOf(".");
   	return o.substring(pos).toLowerCase();
}

/*
  格式化日期 时间戳
  参数：\/Date(1467273890000)\/
 */
u.formatDate=function(dt)   {       
        var   year=dt.getYear()+1900;       
        var   month=dt.getMonth()+1;       
        var   date=dt.getDate();       
        var   hour=dt.getHours();       
        var   minute=dt.getMinutes();       
        var   second=dt.getSeconds();       
        return year+"-"+month+"-"+date;       
}
u.formatJsonDateTime=function(objDate){
  var objData=new Date(objDate).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '');
 	return objData;       
}
u.formatJsonDate=function(objDate){
  	var objData=u.Substring(objDate,10);
 	return objData;       
}
u.formatMobile=function(tel){
	var objTel=tel.substr(0, 3) + '****' + tel.substr(7);  
	return objTel;
}
u.formatDateStamp=function(Dtime){
     var NewDtime = new Date(parseInt(Dtime.slice(6, 19)));
     return u.formatDate(NewDtime);
}
u.Uinx2Time=function(objNumber){
	var timestamp3 = 1403058804;
	var newDate = new Date();
	newDate.setTime(objNumber * 1000);
	return newDate.toLocaleString();
}
// 生成guid,主要用于生成随机文件名
u.NewGuid=function() {
	function S4() {
      return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
     }
     return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}
u.Regex_Phone=function(mobile)
{
   if(mobile.length==0 || mobile.length!=11) 
      return false;         
   var myreg = /^(((13[0-9]{1})|(15[0-9]{1})|(17[0-9]{1})|(18[0-9]{1}))+\d{8})$/; 
   if(!myreg.test(mobile)) 
	return false;
  return true;
} 
u.Regex_Money=function(objNumber){
   var myreg =/(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)/; 
   if(!myreg.test(objNumber)) 
	return false;
  return true;
}
u.stringOrEmpty=function(objString){
	if(typeof(objString)=="undefined" || objString==undefined || objString==""){
		return true;
	} 
	return false;
}
u.NullToEmpty=function(objString){
	if(u.stringOrEmpty(objString))return "";
	return objString;
}
u.Substring=function funcstring(objString,objLen){
	if(objString=="" || objString==undefined || objString=="undefined")
		return "";
	if(objString.length<=objLen)
		return objString;
	return objString.substr(0,objLen);
}
u.isContains=function(str, substr) {
	var _index=str.indexOf(substr);
    return  _index>= 0;
}
u.IsNumber=function(objNumber){
	var Letters = "1234567890";
	var len=Letters.length;
	var i;
	var c;
	for(var i = 0; i<len; i++ )   {   
		//Letters.length() ->>>>取字符长度
		c = objNumber.charAt(i);
		if(Letters.indexOf(c) ==-1)   { 
			//在"Letters"中找不到"c" 见下面的此函数的返回值
			return false;
　　　   　}
　　　}
	return true; 
}
u.isNumberBy100=function(ssn) {
 var re = /^[0-9]*[0-9]$/i;       //校验是否为数字
 if(re.test(ssn) && ssn%100===0) {
  return true;
 }else {
  return false;
 }
}

u.endWith=function(s){
  if(s==null||s==""||this.length==0||s.length>this.length)
     return false;
  if(this.substring(this.length-s.length)==s)
     return true;
  else
     return false;
  return true;
 }  
 
 u.startWith=function(objContent,s){
  if(s==null||s==""||s.length==0||s.length>objContent.length)
   return false;
  if(objContent.substr(0,s.length)==s)
     return true;
  else
     return false;
  return true;
 }
 
 u.bindInputVal=function(objId,objValue){
 	u.val(u.byId(objId),objValue);
 }
 u.getInputVal=function(objId){
 	return u.val(u.byId(objId));
 }
 
window.$api = u;

})(window);


