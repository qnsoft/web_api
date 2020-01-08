﻿(function ($, window) {
    var loadStyles = function (url) {
        var hasSameStyle = false; var links = $('link'); for (var i = 0; i < links.length; i++) { if (links.eq(i).attr('href') == url) { hasSameStyle = true; return } }
        if (!hasSameStyle) { var link = document.createElement("link"); link.type = "text/css"; link.rel = "stylesheet"; link.href = url; document.getElementsByTagName("head")[0].appendChild(link); }
    }
     $.fn.toast = function (options) {
        var $this = $(this); var _this = this; return this.each(function () {
            $(this).css({ position: 'relative' }); var top = ''; var translateInfo = ''; var box = ''; var defaults = { position: "absolute", animateIn: "fadeIn", animateOut: "fadeOut", padding: "10px 20px", background: "rgba(7,17,27,0.66)", borderRadius: "6px", duration: 3000, animateDuration: 500, fontSize: 14, content: "这是一个提示信息", color: "#fff", top: "80%", zIndex: 1000001, isCenter: true, closePrev: true, }
            var opt = $.extend(defaults, options || {}); var t = ''; top = opt.isCenter === true ? '50%' : opt.top; defaults.isLowerIe9 = function () { return (!window.FormData); }
            defaults.createMessage = function () {
                if (opt.closePrev) { $('.cpt-toast').remove(); }
                box = $("<span class='animated " + opt.animateIn + " cpt-toast'></span>").css({ "position": opt.position, "padding": opt.padding, "background": opt.background, "font-size": opt.fontSize, "-webkit-border-radius": opt.borderRadius, "-moz-border-radius": opt.borderRadius, "border-radius": opt.borderRadius, "color": opt.color, "top": top, "z-index": opt.zIndex, "-webkit-transform": 'translate3d(-50%,-50%,0)', "-moz-transform": 'translate3d(-50%,-50%,0)', "transform": 'translate3d(-50%,-50%,0)', '-webkit-animation-duration': opt.animateDuration / 1000 + 's', '-moz-animation-duration': opt.animateDuration / 1000 + 's', 'animation-duration': opt.animateDuration / 1000 + 's', }).html(opt.content).appendTo($this); defaults.colseMessage();
            }
            defaults.colseMessage = function () { var isLowerIe9 = defaults.isLowerIe9(); if (!isLowerIe9) { t = setTimeout(function () { box.removeClass(opt.animateIn).addClass(opt.animateOut).on('webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend', function () { box.remove(); }); }, opt.duration); } else { t = setTimeout(function () { box.remove(); }, opt.duration); } }
            defaults.createMessage();
        })
    };
})(jQuery, window);
var toast = function (content)
{
    var animateIn = 'fadeIn';
    var animateOut = 'fadeOut';
    var duration = '3000';
    var isCenter = true;
    $('body').toast({ position: 'fixed', animateIn: animateIn, animateOut: animateOut, content: content, duration: duration, isCenter: isCenter, });
}