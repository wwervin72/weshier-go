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

