import { login, githubLogin, fetchUserInfo, logout } from '@/api/fetch/user'
import { setToken, getToken, removeToken, roleCookieName } from '@/utils/cookie'
import { message } from 'ant-design-vue'

const state = {
	user: null,
	token: getToken() || null,
	role: getToken(roleCookieName) || null,
};

const mutations = {
	updateToken(state, token) {
		state.token = token
	},
	updateUser(state, user) {
		state.user = user
	},
	updateRole(state, role) {
		state.role = role
	},
};

const actions = {
	login({ dispatch }, data) {
		return login(data).then(res => {
			if (res.code === 200) {
				dispatch('handleLogin', res.data)
				return res.data
			} else {
				message.error(res.message || '登录失败')
				return Promise.reject()
			}
		})
	},
	githubLogin({ dispatch }, data) {
		return githubLogin(data).then(res => {
			if (res.code === 200) {
				dispatch('handleLogin', res.data)
				return res.data
			} else {
				message.error(res.message || '登录失败')
				return Promise.reject()
			}
		})
	},
	logout({dispatch}) {
		return logout().then(res => {
			if (res.code === 200) {
				dispatch('handleLogout')
				return Promise.resolve()
			} else {
				message.error(res.message || '登出失败')
				return Promise.reject()
			}
		})
	},
	fetchUserInfo({ dispatch }) {
		return fetchUserInfo().then(res => {
			if (res.code === 200) {
				dispatch('handleLogin', res.data)
				return res.data
			}
		}).catch(err => {
			// dispatch('handleLogout')
			if (process.env.NODE_ENV === 'development') {
				console.log(err)
			}
		})
	},
	handleLogin({commit}, data) {
		commit('updateUser', data)
		commit('updateToken', data.token)
		commit('updateRole', data.role)
		setToken(data.token)
		setToken(data.role, roleCookieName)
		return data
	},
	handleLogout({commit}) {
		commit('updateUser', null)
		commit('updateToken', null)
		commit('updateRole', null)
		removeToken()
		removeToken(roleCookieName)
		return
	}
}

export default {
	state,
	mutations,
	actions
};
