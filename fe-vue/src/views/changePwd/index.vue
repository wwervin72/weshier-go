<template>
	<a-form
		ref="form"
		:form="form"
		:label-col="labelCol"
		:wrapper-col="wrapperCol"
	>
		<a-form-item ref="oldPwd" label="旧密码" prop="oldPwd">
			<a-input
				type="password"
				v-decorator="[
					'oldPwd',
					{
						rules: [
							{
								required: true,
								message: '请输入旧密码'
							}
						]
					}
				]"
			/>
		</a-form-item>
		<a-form-item ref="newPwd" label="新密码" prop="newPwd" :min="5" :max="50">
			<a-input
				type="password"
				v-decorator="[
					'newPwd',
					{
						rules: [
							{
								required: true,
								message: '请输入新密码',
								trigger: 'blur'
							}
						]
					}
				]"
			/>
		</a-form-item>
		<a-form-item ref="rePwd" label="重复新密码" prop="rePwd">
			<a-input
				type="password"
				v-decorator="[
					'rePwd',
					{
						rules: [
							{
								required: true,
								message: '请重复输入新密码',
								trigger: 'blur'
							},
							{
								validator: validateRePwd,
								trigger: 'blur'
							}
						]
					}
				]"
			/>
		</a-form-item>
		<a-form-item :wrapperCol="{ push: 4 }">
			<a-button type="primary" @click="onSubmit">提交</a-button>
		</a-form-item>
	</a-form>
</template>

<script>
import { Input, Form, Button } from "ant-design-vue";
import { changePwd } from "@/api/fetch/user.js";
export default {
	name: "WsChangePwd",
	components: {
		AButton: Button,
		AInput: Input,
		AForm: Form,
		AFormItem: Form.Item
	},
	data() {
		return {
			form: this.$form.createForm(this, { name: 'changePwd' }),
			labelCol: { span: 4 },
			wrapperCol: { span: 14 }
		};
	},
	methods: {
		validateRePwd (rule, value, callback) {
			const form = this.form;
			if (value && value !== form.getFieldValue("newPwd")) {
				callback("两次输入新密码不一致");
			} else {
				callback();
			}
		},
		onSubmit() {
			this.form.validateFields(err => {
				if (!err) {
					changePwd(this.form.getFieldsValue())
						.then(res => {
							if (res.code === 200) {
								this.$refs.form.resetFields();
							}
						})
						.catch(err => {
							if (process.env.NODE_ENV === "development") {
								console.log(err);
							}
						});
				}
			})
		}
	}
};
</script>

<style lang="scss" scoped></style>
