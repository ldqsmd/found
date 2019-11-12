jQuery(document).ready(function() {
	var murl = new objURL();
	initMemberMenu();

	$("[data-toggle='tooltip']").tooltip();

	$("[data-whetherread='0']").each(function(index, el) {
		$(this).animate({
			opacity: '1'
		}, 'slow');
		$(this).css('background-color', '#fefefe');
	});

	//刷新当前页面
	$(".btn-refresh").click(function() {
		location.reload();
	});

	var ma_count = 0;
	//筛选订单
	$(".btn-order-screen").click(function() {
		$t = ma_count > 0 ? 0 : 500;
		if ($(this).hasClass('on')) {
			$(".more-action").slideUp($t);
			$(this).removeClass('on').html('<i class="uicon-toggle-off"></i> 更多筛选');
		} else {
			$(this).addClass('on').html('<i class="uicon-toggle-on"></i> 关闭筛选');
			$(".more-action").slideDown($t);
		}
	});

	//初始化筛选
	$('.more-action ul').each(function(index, el) {
		var name = $(this).data('name');
		$v = murl.get(name)
		if (typeof($v) != "undefined") {
			$selected_li = $(this).find("li[data-value='" + $v + "']");
			$selected_li.addClass('selected').siblings().removeClass('selected');
			button_value = "<span>" + $(this).attr('title') + '</span>：' + $selected_li.text();
			hidden_name = $(this).data('name');
			data_value = $selected_li.data("value");
			condition = "tag-" + $(this).attr('class');
			$("." + condition).remove();
			html_serach_tag = '<div class="tag-more-order  ' + condition + '"><button class="btn-remove-action"  title="单击移除条件">' + button_value + '</button><input type="hidden" name="' + hidden_name + '" value="' + data_value + '" /> </div>';
			$(html_serach_tag).hide().appendTo('.screen-list').fadeIn("fast");
			ma_count++;
		}
	});

	//筛选条件大于0时,显示筛选列表
	if (ma_count > 0) {
		$(".btn-order-screen").trigger('click');
		$(".btn-order-total").removeClass('hidden');
	}

	//选择条件
	$(".more-action ul li").click(function() {
		if ($(this).hasClass('title')) {
			return;
		}
		$(this).addClass('selected').siblings().removeClass('selected');
		$ul = $(this).parents("ul");
		//条件外部容器class
		condition = "tag-" + $ul.attr('class');
		//条件按钮值
		button_value = "<span>" + $ul.attr('title') + '</span>：' + $(this).text();
		data_value = $(this).data("value");
		hidden_name = $ul.data('name');
		//移除当前存在name
		$("." + condition).remove();
		if ($(this).text() != "不限") {
			html_serach_tag = '<div class="tag-more-order  ' + condition + '"><button class="btn-remove-action"  title="单击移除条件">' + button_value + '</button><input type="hidden" name="' + hidden_name + '" value="' + data_value + '" /> </div>';
			$(html_serach_tag).hide().appendTo('.screen-list').fadeIn("fast");
		}
	});

	//移除条件
	$(".screen-list").on('click', '.btn-remove-action', function() {
		$name = $(this).next().attr('name');
		$(".more-action ul[data-name='" + $name + "'] li").removeClass('selected').eq(1).addClass('selected');
		$(this).parent().slideUp().remove();
	});

	//清除所有条件
	$(".btn-sc-clear").click(function() {
		$(".more-action ul").each(function(index, el) {
			$(this).find('li').removeClass('selected').eq(1).addClass('selected');
		});
		$(".screen-list .btn-remove-action").each(function() {
			$(this).parent().slideUp().remove();
		});
	});

	//查询更多筛选
	$(".btn-search-more-action").click(function() {
		$(this).html("正在执行查询");
	});


	//领取优惠券
	$(".btn-receive-coupon").click(function() {
		_self = $(this);
		ctid = $(this).attr('rel');
		var ajax_data = {
			action: 'receive_coupons_by_ctid',
			ctid: ctid,
			random: new Date().getTime()
		};
		_self.addClass('loading').attr('disabled', 'disabled').text('正在领取');
		jQuery.post("/wp-admin/admin-ajax.php", ajax_data,
			function(data) {
				if (data.status == "n") {
					banner_alert(data.msg);
					_self.removeClass('loading').removeAttr('disabled').html('领取优惠券');
				} else {
					//已领数量
					$span_received = $(".gift-" + ctid).find('.coupon-received');
					//全部数量
					$span_num = $(".gift-" + ctid).find('.coupon-ct_mumbleofcoupon');

					$received = parseInt($span_received.text()) + 1;
					$num = parseInt($span_num.text());
					console.log($received);
					if ($received == $num) {
						_self.attr('disabled', 'disabled');
						_self.removeClass('loading').html('<i class="uicon-nothing"></i>  已经领完');
					} else {
						_self.removeClass('loading').removeAttr('disabled').html('领取优惠券');
					}
					$(".gift-" + ctid).find('.coupon-received').text($received);

					banner_alert(data.msg, "success");
				}
			}
		);
	});

	//发表订单留言
	$(".btn-send-comment").click(function() {
		var __self = jQuery(this);
		var messageContent = $(".comment-content").val();

		if (messageContent == "") {
			banner_alert("评论不能为空.");
			return false;
		}

		__self.attr('disabled', true);

		ajax_data = $(".order-comment-form").serializeArray();

		jQuery.post('/wp-admin/admin-ajax.php', ajax_data,
			function(data) {
				if (data.status == 1) {
					console.log(data);
					$("#comment .empty").remove();
					$comment_rel = [
						'<div class="comment-list-item">',
						'<div class="comment-avatar"><a href="javascript:;" class="comment-author-avatar">',
						data.user_avatar,
						'</a></div>',
						'<div class="comment-c-content">',
						'<div class="comment-meta">',
						'<span class="comment-author-name">',
						data.wp_b_fromUserName,
						'</span><time class="comment-author-time pull-right">刚刚',
						'</time> </div> <div class="comment-content">',
						data.wp_b_Content,
						'</div></div></div>'
					].join("");
					$($comment_rel).hide().prependTo($('.comment-list')).slideDown();;
					$(".comment-content").val('').focus();

				} else {
					banner_alert("留言提交失败...");
				}
				__self.html('发 送').removeClass('loading').removeAttr('disabled');
			}
		);
	});

	//更新资料
	$(".btn-update-profile").click(function() {
		_self = $(this);
		//判断昵称是否填写为空，为空用用户名代替
		if ($("input[name='usermeta[nickname]']").val() == "") {
			$("input[name='usermeta[nickname]']").val($(".username").text());
		}
		ajax_data = $(".profile-form").serializeArray();
		$action = {
			name: "action",
			value: "ui_user_update_profile"
		};
		ajax_data.unshift($action); //增加执行方法。
		_self.html(' 正在更新资料 ').addClass('loading').attr('disabled', 'disabled');
		jQuery.post('/wp-admin/admin-ajax.php', ajax_data, function(data) {
			if (data.status == 1) {
				alert("设置已成功修改！");
			}
			_self.html(' 保存设置 ').removeClass('loading').removeAttr('disabled');
		});
	});


});



//初始化主菜单
function initMemberMenu() {
	var index = 1;
	$('.member-menu .nav li a').each(function() {
		if ($($(this))[0].href == String(window.location).split("?")[0]) {
			$(this).parent().addClass('active');
			index = 0;
		}
	});
}



//更换用户头像
function tx($uid) {
	ajax_data = "action=ui_user_update_avatar&uid=" + $uid + "&avatar=" + myFrame.window.tx + "&ran=" + Math.random();
	jQuery.post('/wp-admin/admin-ajax.php', ajax_data,
		function(data) {
			if (data.status == 1) {
				$(".avatar,.ui-account").find('img').attr('src', myFrame.window.tx);
			}
		}
	);
}



//全选取消按钮函数
function checkAll(chkobj) {
	if ($.trim($(chkobj).text()) == "全选") {
		$(chkobj).children("span").text("取消");
		$(".checkall input:enabled").prop("checked", true);
	} else {
		$(chkobj).children("span").text("全选");
		$(".checkall input:enabled").prop("checked", false);
	}
}
