<template>
	<div class="ws_editor">
		<mavon-editor
			class="editor"
			codeStyle="atom-one-dark"
			ref="editor"
			v-model="article.content"
			@save="save"
			@imgAdd="imgAdd"
		/>
		<div class="editor_menu">
			<a-form-model
				labelAlign="left"
				ref="form"
				:model="article"
				:rules="rules"
				:label-col="labelCol"
				:wrapper-col="wrapperCol"
			>
				<a-form-model-item label="标题" prop="title">
					<a-input
						v-model="article.title"
						placeholder="请输入文章标题"
					/>
				</a-form-model-item>
				<a-form-model-item label="标签" prop="tags">
					<a-select
						v-model="article.tags"
						mode="multiple"
						placeholder="请选择文章标签"
					>
						<a-select-option
							v-for="item in tagList"
							:key="item.value"
							>{{ item.label }}</a-select-option
						>
					</a-select>
				</a-form-model-item>
				<a-form-model-item label="分类" prop="categoryId">
					<a-select
						v-model="article.categoryId"
						placeholder="请选择文章分类"
					>
						<a-select-option
							v-for="item in cateList"
							:key="item.value"
							>{{ item.label }}</a-select-option
						>
					</a-select>
				</a-form-model-item>
				<a-form-model-item label="密码" prop="password">
					<a-input
						v-model="article.password"
						placeholder="文章查看密码"
					/>
				</a-form-model-item>
				<a-form-model-item label="是否允许评论" prop="allowComment">
					<a-switch v-model="article.allowComment">
						<a-icon slot="checkedChildren" type="check" />
						<a-icon slot="unCheckedChildren" type="close" />
					</a-switch>
				</a-form-model-item>
				<a-form-model-item label="是否置顶" prop="topping">
					<a-switch v-model="article.topping">
						<a-icon slot="checkedChildren" type="check" />
						<a-icon slot="unCheckedChildren" type="close" />
					</a-switch>
				</a-form-model-item>
				<a-form-model-item
					class="thumbnail"
					label="缩略图"
					prop="thumbnail"
				>
					<img
						width="100%"
						height="100%"
						:src="thumbnail"
						title="缩略图"
					/>
					<div class="ws-thumbnail-btns">
						<a-icon
							class="chose-image"
							type="file-image"
							@click="openChoseThumbnail"
						/>
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
							<img v-if="imageUrl" :src="imageUrl" alt="avatar" />
							<div v-else>
								<a-icon
									:type="uploadingFile ? 'loading' : 'plus'"
								/>
							</div>
						</a-upload>
					</div>
				</a-form-model-item>
				<a-form-model-item label="附件" prop="annexs">
					<a-upload
						name="file"
						:headers="headers"
						:action="uploadAnnexAction"
						:default-file-list="defaultAnnexList"
						@change="selectUploadAnnex"
					>
						<a-button><a-icon type="upload" />上传附件</a-button>
					</a-upload>
				</a-form-model-item>
			</a-form-model>
		</div>
		<a-modal
			title="选择缩略图"
			:visible="showSelectThumbnail"
			@cancel="cancelSelectThumbnail"
		>
			<img
				class="thumbnail-item"
				v-for="img in thumbnailList"
				:key="img.value"
				@dblclick="selectThumbnail(img)"
				:src="baseURL + img.url"
				:title="img.label"
			/>
		</a-modal>
	</div>
</template>

<script>
import { mapGetters } from "vuex";
import {
	FormModel,
	Input,
	Select,
	Switch,
	Icon,
	message,
	Upload,
	Modal,
	Button
} from "ant-design-vue";
import {
	uploadImg,
	uploadThumbnailUrl,
	fetchThumbnails,
	uploadAnnexUrl
} from "../../api/fetch/file";
import { fetchUserCategoryList } from "../../api/fetch/category";
import { fetchUserTagList } from "../../api/fetch/tag";
import {
	fetchArticleInfo,
	createArticle,
	updateArticle
} from "../../api/fetch/article";
const maxAbstractLen = 400
export default {
	name: "WsEditorPage",
	components: {
		AFormModel: FormModel,
		AFormModelItem: FormModel.Item,
		AInput: Input,
		AButton: Button,
		ASelect: Select,
		AUpload: Upload,
		AModal: Modal,
		ASelectOption: Select.Option,
		ASwitch: Switch,
		AIcon: Icon
	},
	props: {
		articleId: {
			type: String,
			default: ""
		}
	},
	data() {
		return {
			baseURL: process.env.VUE_APP_PUBLIC_BASE_URL,
			uploadAction: uploadThumbnailUrl,
			uploadAnnexAction: uploadAnnexUrl,
			showSelectThumbnail: false,
			value: "",
			tagList: [],
			cateList: [],
			thumbnailList: [],
			labelCol: {
				offset: 1,
				span: 22
			},
			wrapperCol: {
				offset: 1,
				span: 22
			},
			rules: {
				title: [
					{
						required: true,
						message: "请输入文章名称",
						trigger: "blur"
					}
				]
			},
			thumbnailId: null,
			imageUrl: "",
			uploadingFile: false,
			article: {},
			defaultAnnexList: [
				// {
				// 	id: 1,
				// 	name: 'xxx.png',
				// 	status: 'done',
				// 	url: 'http://www.baidu.com/xxx.png',
				// },
			]
		};
	},
	created() {
		if (this.articleId) {
			this.fetchArticleInfo(this.articleId);
		}
		this.fetchThumbnails();
		this.fetchUserCategoryList();
		this.fetchUserTagList();
	},
	computed: {
		...mapGetters(["token"]),
		isUpdate() {
			return !!this.articleId;
		},
		headers() {
			return {
				Authorization: `Bearer ${this.token}`
			};
		},
		thumbnail() {
			return (
				this.baseURL +
				(this.article.thumbnail ||
					`${process.env.VUE_APP_PUBLIC_PREFIXER}/asset/image/thumbnail.jpg`)
			);
		}
	},
	methods: {
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
		selectThumbnail(img) {
			this.$set(this.article, 'thumbnail', img.url)
			this.thumbnailId = img.value
			this.cancelSelectThumbnail();
		},
		openChoseThumbnail() {
			this.showSelectThumbnail = true;
		},
		cancelSelectThumbnail() {
			this.showSelectThumbnail = false;
		},
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
		selectUploadAnnex({ file }) {
			const res = file.response;
			if (res) {
				// fileList.splice(fileList.length - 1, 1)
				// // 上传完毕
				// if (res.code === 200) {
				// 	this.defaultAnnexList.push(this.transformFileListData(res.data));
				// } else {
				// 	res.message && message.info(res.message)
				// }
			}
		},
		transformFileListData(data) {
			return {
				...data,
				uid: data.id.toString(),
				status: 'done',
				response: 'success',
				url: this.baseURL + data.url
			}
		},
		fetchUserCategoryList() {
			fetchUserCategoryList().then(res => {
				if (res.code === 200) {
					this.cateList = res.data;
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		fetchArticleInfo(id) {
			fetchArticleInfo(id).then(res => {
				if (res.code === 200) {
					res.data.tags = res.data.tags || []
					this.thumbnailId = res.data.thumbnailEntity.id;
					this.article = res.data;
					this.article.allowComment = !!this.article.allowComment
					this.article.topping = !!this.article.topping
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		fetchUserTagList() {
			fetchUserTagList().then(res => {
				if (res.code === 200) {
					this.tagList = res.data;
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		save(val, render) {
			this.$refs.form.validate(res => {
				if (res) {
					const fn = this.isUpdate ? updateArticle : createArticle;
					let {
						content,
						title,
						tags,
						abstract,
						categoryId,
						password,
						allowComment,
						topping,
						id
					} = this.article;
					const body = {
						content,
						title,
						tags,
						abstract,
						categoryId,
						password,
						allowComment: allowComment ? 1 : 0,
						topping: topping ? 1 : 0,
					};
					if (!abstract) {
						const domParse = new DOMParser()
						const contentText = domParse.parseFromString(render.replace(/\s+/, ""), "text/html")
							.documentElement.textContent.replace("\n", "")
						const abstractLen = Math.ceil(Math.max(Math.random() * maxAbstractLen, 200))
						abstract = contentText.slice(0, abstractLen)
						body.abstract = abstract
					}
					if (this.thumbnailId) {
						body.thumbnail = this.thumbnailId.toString()
					}
					if (this.isUpdate) {
						body.id = id
					}
					fn(body).then(res => {
						if (res.code == 200) {
							message.success(
								`文章${this.isUpdate ? "更新" : "创建"}成功`
							);
							this.$router.push(`/a/${id || res.data.id}`)
						}
					}).catch(err => {
						if (process.env.NODE_ENV === 'development') {
							console.log(err);
						}
					})
				}
			});
		},
		imgAdd(pos, $file) {
			var formdata = new FormData();
			formdata.append("file", $file);
			uploadImg(formdata).then(res => {
				if (res.code === 200) {
					this.$refs.editor.$img2Url(
						pos,
						process.env.VUE_APP_PUBLIC_BASE_URL + res.data.url
					);
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		}
	}
};
</script>
<style lang="scss" scoped>
.ws_editor {
	display: flex;
	height: 100%;
	overflow-y: auto;
}
.editor {
	flex: 1;
}
.editor_menu {
	flex: 0 0 400px;
	height: 100%;
	overflow-y: auto;
}
.ws-thumbnail-btns {
	width: 100px;
	display: flex;
	justify-content: space-between;
	position: absolute;
	align-items: center;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);
	.chose-image {
		line-height: 45px;
		font-size: 25px;
		background-color: #fafafa;
		border: 1px solid #d9d9d9;
		border-radius: 4px;
		&:hover {
			border-color: #1890ff;
			border-style: dashed;
		}
	}
	.chose-image,
	.ws-upload {
		display: none;
		width: 40px;
		height: 40px;
	}
	/deep/ {
		.ant-upload {
			width: 100%;
			height: 100%;
		}
	}
}
.thumbnail-item {
	width: 100px;
	height: 100px;
	border-radius: 3px;
}
.thumbnail {
	position: relative;
	/deep/ {
		.ant-form-item-children {
			display: block;
			width: 100%;
			height: 100%;
			min-height: 250px;
		}
	}
	&.ant-form-item {
		margin-bottom: 0;
	}
	&:hover {
		.ws-thumbnail-btns {
			.chose-image,
			.ws-upload {
				display: block;
			}
		}
	}
}
/deep/ {
	.ant-form-item-label {
		line-height: 32px;
	}
	.ant-form-item {
		margin-bottom: 19px;
	}
	.ant-form-item-with-help {
		margin-bottom: 0;
	}
	.ant-modal-mask,
	.ant-modal-wrap {
		z-index: 10000;
	}
}
</style>
