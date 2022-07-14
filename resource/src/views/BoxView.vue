<template>
	<h1>产品管理</h1>
	<el-button v-on:click="increment()" type="primary">点击</el-button>
	<label>{{ this.$store.state.count}}</label>
	<el-button @click="showCreate = true" type="primary">新建</el-button>
 
	<el-dialog v-model="showCreate" title="Outer Dialog">
		<template #default>
			<el-form :model="form" label-width="120px">
				<el-form-item label="Activity name">
					<el-input v-model="form.name" />
				</el-form-item>
			</el-form>
		</template>
		<template #footer>
			<div class="dialog-footer">
				<el-button @click="showCreate = false">Cancel</el-button>
				<el-button type="primary" @click="save()">Save</el-button>
			</div>
		</template>
	</el-dialog>
 
	<div style="margin: 60px;">
		<el-table :data="tableData" height="250" style="width: 100%">
			<el-table-column prop="date" label="Date" width="180" />
			<el-table-column prop="name" label="Name" width="180" />
			<el-table-column prop="address" label="Address" />
		</el-table>
	</div>
</template>
<script>
	export default {
		data() {
			return {
				count: 0,
				showCreate: false,
				form: {
					name: ''
				},
				tableData: [{
					date: '2016-05-03',
					name: 'Tom',
					address: 'No. 189, Grove St, Los Angeles',
				}]
			}
		},
		methods: {
			save() {
				this.showCreate = false;
				this.tableData.push({
					date: '2016-05-03',
					name: this.form.name,
					address: 'No. 189, Grove St, Los Angeles',
				});
				this.form.name='';
			},
			increment() {
				this.$store.commit('increment');
			}
		}
	}
</script>
<style>
</style>