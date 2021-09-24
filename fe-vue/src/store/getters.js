export default {
	token: state => state.user.token,
	user: state => state.user.user,
	role: state => state.user.role,
	tocHtml: state => state.article.tocHtml,
}
