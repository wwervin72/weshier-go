<template>
	<a-form-model
		ref="form"
		:model="formData"
		:rules="rules"
		:label-col="labelCol"
		:wrapper-col="wrapperCol"
	>
		<a-form-model-item ref="name" label="标签名" prop="name">
			<a-input v-model="formData.name" />
		</a-form-model-item>
		<a-form-model-item ref="desc" label="描述" prop="desc">
			<a-input v-model="formData.desc" />
		</a-form-model-item>
		<a-form-model-item :wrapperCol="{ push: 4 }">
			<a-button type="primary" @click="onSubmit">提交</a-button>
		</a-form-model-item>
	</a-form-model>
</template>

<script>
import { Input, FormModel, Button, message } from "ant-design-vue";
import { createTag, fetchTagInfo, updateTag } from "@/api/fetch/tag.js";
export default {
	name: "WsTagAddPage",
	components: {
		AButton: Button,
		AInput: Input,
		AFormModel: FormModel,
		AFormModelItem: FormModel.Item
	},
	props: {
		id: {
			type: String,
			default: ""
		}
	},
	computed: {
		isUpdate() {
			return !!this.id;
		}
	},
	created() {
		if (this.isUpdate) {
			this.fetchTagInfo();
		}
	},
	data() {
		return {
			formData: {},
			rules: {},
			labelCol: { span: 4 },
			wrapperCol: { span: 14 }
		};
	},
	methods: {
		fetchTagInfo() {
			fetchTagInfo(this.id).then(res => {
				if (res.code === 200) {
					this.formData = res.data
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		onSubmit() {
			const fn = this.isUpdate ? updateTag : createTag;
			fn(this.formData).then(res => {
				if (res.code === 200) {
					message.success(`标签${this.isUpdate ? '更新' : '添加'}成功`)
					if (this.isUpdate) {
						this.$router.push('/setting/tag')
					} else {
						this.$refs.form.resetFields()
					}
					this.isUpdate || this.$refs.form.resetFields();
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
<style lang="scss" scoped></style>
