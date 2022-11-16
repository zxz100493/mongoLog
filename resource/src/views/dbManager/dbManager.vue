<template>
  <div>
    <div class="top-button">
      <el-button type="primary" round @click="createDb">+ Add Db</el-button>
    </div>
    <div>
      <el-table :data="form.tableData" border style="width: 100%">
        <el-table-column prop="db" label="dbName" width="180" />
        <el-table-column prop="clsNum" label="collectionsNum" width="180" />
        <el-table-column prop="size" label="diskUsage" />
        <el-table-column fixed="right" label="Operations" width="120">
          <template #default="scope">
            <el-button
            link
            type="danger"
            :icon="Delete"
            size="small"
            :disabled="scope.row.disabled"
            @click.prevent="deleteRow(scope.$index)"
            >
            Remove
            </el-button>
            <!-- <el-button link type="danger" size="small">Edit</el-button> -->
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div>
      <el-dialog v-model="form.dialogFormVisible" title="Create Db">
        <el-form :model="form">
          <el-form-item label="Db Name" :label-width="form.formLabelWidth">
            <el-input v-model="form.createDb.name" autocomplete="off" />
          </el-form-item>
          <!-- <el-form-item label="Zones" :label-width="form.formLabelWidth">
            <el-select v-model="form.createDb.name" placeholder="Please select a zone">
              <el-option label="Zone No.1" value="shanghai" />
              <el-option label="Zone No.2" value="beijing" />
            </el-select>
          </el-form-item> -->
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="cancelCreateDb">Cancel</el-button>
            <el-button type="primary" @click="createDbOk">
              Confirm
            </el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, getCurrentInstance, onMounted, ref } from 'vue'
import {
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
// do not use same name with ref
const form = reactive({
  dialogFormVisible: false,
  tableData: [{
    db: '2016-05-03',
    clsNum: 'Tom',
    size: 'No. 189, Grove St, Los Angeles',
    disabled: false
  }],
  createDb: {
    name: ''
  },
  formLabelWidth: '30'
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
  getDbList()
})
const getDbList = () => {
  proxy.axios.get('api/log/db/list', { card: 111 })
    .then((e:any) => {
      var data = e.data.data
      form.tableData = data
    })
}
const deleteRow = (index: number) => {
  form.tableData.splice(index, 1)
  console.log('click')
}
const createDb = () => {
  console.log('created db')
  form.dialogFormVisible = true
}
const cancelCreateDb = () => {
  form.dialogFormVisible = false
  form.createDb.name = ''
}
const errorMsg = () => {
  ElMessage.error('Oops, this is a error message.')
}
const createDbOk = () => {
  const params = {
    name: form.createDb.name
  }
  proxy.axios.post('api/log/db/create', params)
    .then((e:any) => {
      var data = e.data.data
      if (data.status) {
        errorMsg()
        return
      }

      getDbList()
    })
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
  .top-button{
    display: flex;
    margin-bottom: 10px;
  }
</style>
