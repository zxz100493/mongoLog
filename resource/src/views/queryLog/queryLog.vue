<template>
  <div>
    <div class="top-button">
      <!-- <el-button type="primary" round @click="createDb">+ Add Db</el-button> -->
      <el-row class="demo-autocomplete">
        <el-col>
          <el-autocomplete
            v-model="state1"
            :fetch-suggestions="querySearch"
            clearable
            class="inline-input w-150"
            placeholder="select unique remark"
            @select="handleSelect"
          />
        </el-col>
      </el-row>
      <div class="time-range">
        <el-date-picker
          v-model="value2"
          type="datetimerange"
          :shortcuts="shortcuts"
          range-separator="To"
          start-placeholder="Start date"
          end-placeholder="End date"
          value-format="YYYY-MM-DD HH:mm:ss"
        />
      </div>
      <div>
        <el-input v-model="input" placeholder="Please input keywords" />
      </div>
      <div>
        <el-button type="primary" round @click="queryLog">Query</el-button>
      </div>
    </div>
    <div>
      <el-table :data="form.tableData" border style="width: 100%">
        <el-table-column prop="datetime" label="datetime" width="180" />
        <el-table-column prop="context" label="context" width="180" />
        <el-table-column prop="host" label="host" />
        <el-table-column prop="type" label="type" />
        <el-table-column prop="level" label="level" />
        <el-table-column fixed="right" label="Operations" width="120">
          <template #default="scope">
            <el-button
            link
            type="primary"
            :icon="Search"
            size="small"
            @click.prevent="viewJson(scope.row.context)"
            >
            View Detail
            </el-button>
            <!-- <el-button link type="danger" size="small">Edit</el-button> -->
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div>
        <el-pagination background layout="sizes, prev, pager, next" :total="1000" />
    </div>
    <div style="text-align:left">
      <el-dialog v-model="form.dialogFormVisible" title="View Detail">
        <json-viewer :value="form.jsonData"
        :expand-depth=5
        copyable
        boxed
        >
        </json-viewer>
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
import 'vue-json-viewer/style.css'
// do not use same name with ref
const input = ref('')
const value1 = ref('')
const value2 = ref('')
const value3 = ref('')
const defaultTime = new Date(2000, 1, 1, 12, 0, 0)

const shortcuts = [
  {
    text: 'Today',
    value: new Date()
  },
  {
    text: 'Yesterday',
    value: () => {
      const date = new Date()
      date.setTime(date.getTime() - 3600 * 1000 * 24)
      return date
    }
  },
  {
    text: 'A week ago',
    value: () => {
      const date = new Date()
      date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
      return date
    }
  }
]
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
  formLabelWidth: '30',
  selectList: [],
  showJson: false,
  jsonData: '',
  selectedCls: ''
})

interface RestaurantItem {
  value: string
  link: string
}

const state1 = ref('')
const state2 = ref('')

const restaurants = ref<RestaurantItem[]>([])
const querySearch = (queryString: string, cb: any) => {
  const results = queryString
    ? restaurants.value.filter(createFilter(queryString))
    : restaurants.value
  // call callback function to return suggestions
  cb(results)
}
const createFilter = (queryString: string) => {
  return (restaurant: RestaurantItem) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    )
  }
}

const handleSelect = (item: RestaurantItem) => {
  console.log(item)
}

const queryLog = () => {
  const uniqueMark = state1.value
  const timeStart = value2.value[0]
  const timeEnd = value2.value[1]
  const keywords = input.value
  console.log(keywords)
  if (!uniqueMark && !timeStart && !timeEnd && !keywords) {
    ElMessage.error('please check your query condition!!!')
    return
  }
  const params = {
    uniqueMark: uniqueMark,
    timeStart: timeStart,
    timeEnd: timeEnd,
    keywords: keywords
  }
  proxy.axios.get('api/log/db/query', { params: params })
    .then((e:any) => {
      var data = e.data.data
      form.tableData = data
    })
}

const viewJson = (data: string) => {
  form.dialogFormVisible = true
  form.jsonData = data
}

const tableData = reactive([
  {
    datetime: '2022-11-04 17:07:14',
    host: 'Tom',
    type: 'Tom',
    project: 'Los Angeles'
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
  proxy.axios.get('api/log/db/cls/names', { card: 111 })
    .then((e:any) => {
      var data = e.data.data
      form.selectList = data
      const arr = []
      for (let i = 0; i < data.length; i++) {
        arr[i] = {
          value: data[i],
          link: ''
        }
      }
      restaurants.value = arr
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

<style scoped>
.demo-autocomplete{
  /* width: 4000px; */
}
.time-range{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.time-range-from{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.time-range-to{
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-left: 10px;
}
.demo-datetime-picker {
  display: flex;
  width: 100%;
  padding: 0;
  flex-wrap: wrap;
}
.demo-datetime-picker .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}
.demo-datetime-picker .block:last-child {
  border-right: none;
}
.demo-datetime-picker .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 12px;
  margin-bottom: 20px;
  margin-right: 10px;
}
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
    justify-content:flex-start;
    align-items:center;
    margin-bottom: 10px;
  }
  .demo-range{
    margin-right: 10px;
  }
  .search-form{
    display:flex
  }
</style>
