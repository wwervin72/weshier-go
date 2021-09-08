import store from "@/store/index";

export function goLogin() {
	if (!this.$router) return
	const redirectUrl = window.encodeURIComponent(this.$route.fullPath)
	this.$router.push(`/login?redirectUrl=${redirectUrl}`)
}

export function loginJump() {
	if (!this.$router) return
	const query = this.$route.query
	if (query.redirectUrl) {
		const redirectUrl = window.decodeURIComponent(query.redirectUrl)
		this.$router.push(redirectUrl)
	} else {
		this.$router.push('/')
	}
}

export function isHeart(data = []) {
	const user = store.getters.user
	if (!user) return
	const {id} = user
	data.forEach(el => {
		judgeHearted(el, id)
	})
}

export function judgeHearted(item, uid) {
	item.hearted = item.heartUserID && item.heartUserID.includes(uid)
	if (Array.isArray(item.comments)) {
		item.comments.forEach(ele => {
			judgeHearted(ele, uid)
		})
	}
}

/**
 * 切换浏览器 tab 页监听事件
 */
 export function switchBrowserTabs() {
	let selfTitle = document.title;
	document.addEventListener("visibilitychange", function () {
		if (document.visibilityState == "hidden") {
			document.title = "糟糕！出BUG了，快看";
		} else {
			document.title = selfTitle;
		}
	});
}

/**
 * 粘贴事件，往粘贴板增加自定义内容
 */
 export function copySiteInfo() {
	document.body.addEventListener("copy", function (evt) {
		var clipboardData = evt.clipboardData || window.clipboardData;
		var selection = window.getSelection().toString();
		if (clipboardData && selection) {
			evt.preventDefault();
			var siteInfo = [
				"作者：ervin",
				"来自：微识",
				"链接：" + window.location.href,
				"",
				selection,
			];
			clipboardData.setData("text/html", siteInfo.join("<br>")),
				clipboardData.setData("text/plain", siteInfo.join("\n"));
		}
	});
}
