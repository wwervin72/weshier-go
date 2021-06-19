<template>
<div class="thumbnails">
		<a-upload
			name="file"
			list-type="picture-card"
			class="ws-upload"
			:headers="headers"
			:show-upload-list="false"
			:action="uploadAction"
			:before-upload="beforeUpload"
			@change="selectUploadFile"
		>
			<a-icon
				:type="uploadingFile ? 'loading' : 'plus'"
			/>
		</a-upload>
		<div class="thumbnail-item" v-for="(img, index) in thumbnailList" :key="img.value">
			<img :src="baseURL + img.url" :title="img.label"/>
			<div class="btn-area">
				<a-icon type="delete" @click="del(img, index)"/>
			</div>
		</div>
	</div>
</template>

<script>
import { Upload, message, Icon } from 'ant-design-vue'
import { mapGetters } from "vuex";
import {
	uploadThumbnailUrl,
	fetchThumbnails,
	deleteThumbnail
} from "@/api/fetch/file";
export default {
	name: "WsThumbnailMgPage",
	components: {
		AIcon: Icon,
		AUpload: Upload
	},
	data() {
		return {
			uploadAction: uploadThumbnailUrl,
			baseURL: process.env.VUE_APP_PUBLIC_BASE_URL,
			uploadingFile: false,
			thumbnailList: [],
			pagination: {
				current: 0,
				pageSize: 5,
				pageSizeOptions: ['5', '10', '15', '20'],
				total: 0,
			},
		}
	},
	computed: {
		...mapGetters(["token"]),
		headers() {
			return {
				Authorization: `Bearer ${this.token}`
			};
		},
	},
	created() {
		this.fetchThumbnails()
	},
	methods: {
		beforeUpload() {},
		selectUploadFile({ file, fileList }) {
			const res = file.response;
			if (res) {
				// 上传完毕
				if (res.code === 200) {
					this.thumbnailList.push(res.data);
					this.selectThumbnail(res.data);
				} else {
					res.message && message.info(res.message)
					fileList.splice(fileList.length - 1, 1)
				}
			}
		},
		fetchThumbnails() {
			fetchThumbnails().then(res => {
				if (res.code === 200) {
					this.thumbnailList = res.data;
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		del(record, index) {
			console.log(index, record);
			deleteThumbnail(record.value).then(res => {
				if (res.code === 200) {
					this.thumbnailList.splice(index, 1)
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
	}
}
</script>
<style lang="scss" scoped>
.thumbnails {
	padding: 10px;
	display: flex;
	flex-wrap: wrap;
	.thumbnail-item {
		position: relative;
		margin: 10px;
		transition: all ease-in-out .3s;
		transform-origin: center center;
		&:hover {
			transform: scale(1.15);
			.btn-area {
				display: block;
			}
		}
		.btn-area {
			position: absolute;
			top: 0;
			display: none;
			width: 100%;
			padding: 3px 10px;
			background-color: #000;
			color: #fff;
			stroke: #fff;
			fill: #fff;
		}
	}
	img {
		max-width: 300px;
	}
}
.ws-upload {
	width: auto;
}
</style>
