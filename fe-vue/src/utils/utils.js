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
