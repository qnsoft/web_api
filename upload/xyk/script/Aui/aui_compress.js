/**
 * aui -- apicloud图片压缩处理  流浪男  QQ：343757327  
 */
(function() {
	/**
	* 压缩图片处理函数  image  options
	* success  压缩成功的回调   fail 失败  start  开始压缩  over 压缩结束
	*/
	function auiCompress(image, options) {
		this.image = image;
		this.options = {
			start: null,
			over: null,
			success: null,
			fail: null,
			quality: 0.7
		};

		for (var p in options) {
			this.options[p] = options[p];
		}
		if (this.options.quality > 1) this.options.quality = 1;

		this.ret = {
			origin: null,
			base64: null,
			base64Len: null
		};
		this.init();
	}

	auiCompress.prototype = {
		constructor: auiCompress,

		/*初始化*/
		init: function() {
			var that = this;
			that.create(that.image);
		},

		/**
		* 生成base64
		* @param image
		* @param callback
		*/
		create: function(image) {
			var that = this,
			img = new Image(),
			ret = that.ret,
			blob = (typeof image === 'string') ? image : URL.createObjectURL(image);

			img.crossOrigin = "*";
			img.onerror = function() {
				var error = new Error('图片加载失败');
				// 读取图片失败
				if (typeof that.options.fail === 'function') {
					that.options.fail(error);
				}
				// 压缩结束回调
				if (typeof that.options.over === 'function') {
					that.options.over();
				}
			};

			img.onload = function() {
				// 获得图片缩放尺寸
				var resize = that.resize(this);
				// 初始化canvas
				var canvas = document.createElement('canvas');
				var ctx = canvas.getContext('2d');
				canvas.width = resize.w;
				canvas.height = resize.h;
				
				ctx.fillStyle = '#fffff';
				ctx.fillRect(0, 0, resize.w, resize.h);

				ret.origin = image;
				ctx.drawImage(img, 0, 0, resize.w, resize.h);
				ret.base64 = canvas.toDataURL('image/jpeg', that.options.quality);
				_retCallback(ret);

				/*回调处理*/
				function _retCallback(ret) {
					// 释放内存
					canvas = null;
					img = null;
					URL.revokeObjectURL(blob);

					// 加入base64Len，方便后台校验是否传输完整
					ret.base64Len = ret.base64.length;

					// 压缩成功回调
					if (typeof that.options.success === 'function') {
						that.options.success(ret);
					}

					// 压缩结束回调
					if (typeof that.options.over === 'function') {
						that.options.over();
					}
				}
			};

			// 压缩开始前回调
			if (typeof this.options.start === 'function') {
				this.options.start();
			}

			img.src = blob;
		},

		/**
		* 获得图片的缩放尺寸
		* @param img
		* @returns {{w: (Number), h: (Number)}}
		*/
		resize: function(img) {
			var w = this.options.width,
			h = this.options.height,
			scale = img.width / img.height,
			ret = {
				w: img.width,
				h: img.height
			};

			if (w & h) {
				ret.w = w;
				ret.h = h;
			} else if (w) {
				ret.w = w;
				ret.h = Math.ceil(w / scale);
			} else if (h) {
				ret.w = Math.ceil(h * scale);
				ret.h = h;
			}

			// 超过这个值base64无法生成，在IOS上
			if (ret.w >= 3264 || ret.h >= 2448) {
				ret.w *= 0.8;
				ret.h *= 0.8;
			}

			return ret;
		}
	};
	window.auiCompress = function(image, options) {
		return new auiCompress(image, options);
	};
})();
