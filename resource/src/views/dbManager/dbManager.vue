<template>
  <el-table :data="form.tableData" style="width: 100%">
    <el-table-column prop="db" label="DbName" width="180" />
    <el-table-column prop="clsNum" label="collectionsNum" width="180" />
    <el-table-column prop="size" label="diskUsage" />
    <el-table-column prop="address" label="operation" />
  </el-table>
</template>

<script lang="ts" setup>
import { reactive, getCurrentInstance, onMounted, ref } from 'vue'

// do not use same name with ref
const form = reactive({
  tableData: [{
    db: '2016-05-03',
    clsNum: 'Tom',
    size: 'No. 189, Grove St, Los Angeles'
  }]
})

const tableData = reactive([
  {
    db: '2016-05-03',
    clsNum: 'Tom',
    size: 'No. 189, Grove St, Los Angeles'
  }
])

const value = ref('')

const onSubmit = () => {
  console.log('submit!')
}

const { proxy }:any = getCurrentInstance()

onMounted(() => {
  proxy.axios.get('api/log/db/list', { card: 111 })
    .then((e:any) => {
      var data = e.data.data
      form.tableData = data
    })
})
function getCollections () {
  const params = {
    params: {
      name: form.dbName
    }
  }
}

function getClsDetail (name) {
  const params = {
    params: {
      name: name
    }
  }
}
</script>

<style>
  .tag{
    display: flex;
    align-items: flex-start;
    margin: 10px 10px 10px;
  }
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .box-card {
    width: 480px;
  }
  .card-flex{
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .dbSelection{
    margin-bottom: 10px;
    display: flex;
    align-items: left;
  }
</style>
