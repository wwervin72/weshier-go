import request from '../request'

export function UserFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/user',
		...params
	})
}

/**
 * 登录
 * @param {*} data
 */
export function login(data) {
	return UserFetch({
		url: '/login',
		method: 'post',
		data
	})
}

/**
 * 登录
 * @param {*} data
 */
export function changePwd(data) {
	return UserFetch({
		url: '/changepwd',
		method: 'put',
		data
	})
}

/**
 * github 登录
 */
export function githubLogin(params) {
	return UserFetch({
		url: '/auth/github/callback',
		params
	})
}

/**
 * 登出
 */
export function logout() {
	return UserFetch({
		url: '/logout'
	})
}

/**
 * 获取用户详情
 */
export function fetchUserInfo() {
	return UserFetch({
		url: '',
	})
}
