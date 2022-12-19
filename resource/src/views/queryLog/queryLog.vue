<template>
  <div>
    <div class="top-button">
      <!-- <el-button type="primary" round @click="createDb">+ Add Db</el-button> -->
      <el-row class="demo-autocomplete">
    <el-col :span="12">
      <el-autocomplete
        v-model="state1"
        :fetch-suggestions="querySearch"
        clearable
        class="inline-input w-50"
        placeholder="Please Input"
        @select="handleSelect"
      />
    </el-col>
    <el-col :span="12">
      <el-autocomplete
        v-model="state2"
        :fetch-suggestions="querySearch"
        :trigger-on-focus="false"
        clearable
        class="inline-input w-50"
        placeholder="Please Input"
        @select="handleSelect"
      />
    </el-col>
  </el-row>
    </div>
    <div>
      <el-table :data="form.tableData" border style="width: 100%">
        <el-table-column prop="datetime" label="datetime" width="180" />
        <el-table-column prop="context" label="context" width="180" />
        <el-table-column prop="size" label="diskUsage" />
        <el-table-column prop="size" label="diskUsage" />
        <el-table-column prop="size" label="diskUsage" />
        <el-table-column fixed="right" label="Operations" width="120">
          <template #default="scope">
            <el-button
            link
            type="danger"
            :icon="Delete"
            size="small"
            :disabled="scope.row.disabled"
            @click.prevent="deleteRow(scope.$index,scope.row.db)"
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
            <el-input v-model="form.createDb.dbName" autocomplete="off" />
          </el-form-item>
          <el-form-item label="cls Name" :label-width="form.formLabelWidth">
            <el-input v-model="form.createDb.clsName" autocomplete="off" />
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
    dbName: '',
    clsName: ''
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
const deleteRow = (index: number, name: string) => {
  deleteDb(name, index)
}
const createDb = () => {
  console.log('created db')
  form.dialogFormVisible = true
}
const cancelCreateDb = () => {
  form.dialogFormVisible = false
  form.createDb.dbName = ''
  form.createDb.clsName = ''
}
const errorMsg = (msg: string) => {
  ElMessage.error(msg)
}
const createDbOk = () => {
  const params = {
    dbName: form.createDb.dbName,
    clsName: form.createDb.clsName
  }
  proxy.axios.post('api/log/db/create', params)
    .then((e:any) => {
      var data = e.data
      if (data.status) {
        errorMsg(data.data)
        return
      }
      getDbList()
      cancelCreateDb()
    })
}

const deleteDb = (name: string, index: number) => {
  const params = {
    dbName: name
  }
  proxy.axios.post('api/log/db/delete', params)
    .then((e:any) => {
      var data = e.data
      if (data.status) {
        errorMsg(data.data)
        return
      }
      form.tableData.splice(index, 1)
      getDbList()
      return true
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
  .search-form{
    display:flex
  }
</style>
