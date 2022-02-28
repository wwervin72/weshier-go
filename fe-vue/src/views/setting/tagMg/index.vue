<template>
	<div>
		<a-table :columns="columns" :data-source="list" :pagination="pagination"
			@change="pageChange" rowKey="id">
			<template #operation="text, record, index">
				<a href="javascript:;" @click="edit(index, record)">编辑</a>
				<a href="javascript:;" @click="del(index, record)">删除</a>
			</template>
		</a-table>
		<!-- <a-pagination size="small" show-quick-jumper :total="total" :current="pageNumber" showSizeChanger :page-size.sync="pageSize" @change="pageChange"
			@showSizeChange="pageSizeChange" :pageSizeOptions="pageSizes" /> -->
	</div>
</template>

<script>
import { Table } from 'ant-design-vue'
import { fetchTagPagination } from '@/api/fetch/tag'
// import { fetchTagList } from '../../api/fetch/tag'
export default {
	name: "WsTagMgPage",
	components: {
		ATable: Table,
		// ATableColumn: Table.Column,
	},
	data() {
		return {
			pagination: {
				current: 0,
				pageSize: 10,
				showSizeChanger: true,
				pageSizeOptions: ['5', '10', '15', '20'],
				total: 0,
			},
			list: [],
			columns: [
				{
					title: '标题',
					dataIndex: 'name',
					key: 'name',
				},
				{
					title: '描述',
					dataIndex: 'desc',
					key: 'desc',
				},
				{
					title: '操作',
					dataIndex: 'operation',
					key: 'operation',
					scopedSlots: { customRender: 'operation' }
				},
			]
		}
	},
	created() {
		this.fetchTagPagination()
	},
	methods: {
		fetchTagPagination() {
			fetchTagPagination({
				pageSize: this.pagination.pageSize,
				pageNumber: this.pagination.current,
			}).then(res => {
				if (res.code === 200) {
					const { data } = res.data
					this.total = data.total
					this.list = data.list
				}
			}).catch(err => {
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		},
		pageChange () {

		},
		pageSizeChange (pagination) {
			const { pageSize, current } = pagination
			this.pagination.pageSize = pageSize
			this.pagination.current = current
			this.fetchTagPagination()
		},
		edit(index, record) {
			this.$router.push("/setting/tag/" + record.id)
		},
		del(index, record) {
			console.log(index, record);
		},
	}
}
</script>
<style lang="scss" scoped>

</style>
