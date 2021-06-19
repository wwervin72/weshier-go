import request from '../request'

export function FileFetch(params) {
	return request({
		baseURL: process.env.VUE_APP_BASE_API + '/file',
		...params
	})
}

/**
 * 登录
 * @param {*} data
 */
export function uploadImg(data) {
	return FileFetch({
		url: '/upload/img',
		method: 'post',
		headers: { 'Content-Type': 'multipart/form-data' },
		data
	})
}

/**
 * 获取缩略图列表
 * @param {*} data
 */
export function fetchThumbnails() {
	return FileFetch({
		url: '/thumbnails'
	})
}

/**
 * 删除缩略图
 * @param {number} id 缩略图 id
 */
export function deleteThumbnail(id) {
	return FileFetch({
		url: '/thumbnails/' + id,
		method: 'delete'
	})
}

export const uploadThumbnailUrl = process.env.VUE_APP_BASE_API + `/file/upload/thumbnail`
export const uploadAnnexUrl = process.env.VUE_APP_BASE_API + `/file/upload/annex`
