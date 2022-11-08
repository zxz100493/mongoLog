<template>
  <div class="tag gap-4 my-4">
    <el-tag
      key="Host"
      type=""
      class="mx-1"
      effect="dark"
      round
      size="large"
    >
    Host
    </el-tag>
  </div>
  <el-row :gutter="12">
    <el-col :span="8">
      <el-card shadow="hover"> Platform:{{form.platform}}  </el-card>
    </el-col>
    <el-col :span="8">
      <el-tooltip
        class="box-item"
        effect="light"
        content="(used / total)"
        placement="top"
      >
      <el-card shadow="hover"> Memory Info:{{form.memUsed}} / {{form.mem}} </el-card>
    </el-tooltip>
    </el-col>
    <el-col :span="8">
      <el-card shadow="hover"> uptime:{{form.hostUptime}}h </el-card>
    </el-col>
  </el-row>
  <el-divider />

  <div class="tag gap-4 my-4">
    <el-tag
      key="MongoDB"
      type=""
      class="mx-1"
      effect="dark"
      round
      size="large"
    >
    MongoDB
    </el-tag>
  </div>
  <el-row :gutter="12">
    <el-col :span="8">
      <el-card shadow="hover"> Version:{{form.version}}  </el-card>
    </el-col>
    <el-col :span="8">
      <el-tooltip
        class="box-item"
        effect="light"
        content="(physical / virtual)"
        placement="top"
      >
      <el-card shadow="hover"> Memory Usage:{{form.memory}} Mb / {{form.virtualMemory}} Mb </el-card>
    </el-tooltip>
    </el-col>
    <el-col :span="8">
      <el-card shadow="hover"> uptime:{{form.uptime}}h </el-card>
    </el-col>
  </el-row>
  <el-divider />
  <div class="dbSelection">
    <el-select v-model="value" filterable placeholder="select db">
      <el-option
        v-for="item in form.options"
        :key="item"
        :label="item"
        :value="item"
      />
    </el-select>
  </div>
  <div class="card-flex">
    <div>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>database info</span>
            <el-button class="button" text>Operation button</el-button>
          </div>
        </template>
        <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
      </el-card>
    </div>
    <div>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>database statistics</span>
            <el-button class="button" text>Operation button</el-button>
          </div>
        </template>
        <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
      </el-card>
    </div>
  </div>
  <el-divider />
  <h2>A Simple and Slight Log System For You!</h2>
</template>

<script lang="ts" setup>
import { reactive, getCurrentInstance, onMounted, ref } from 'vue'

// do not use same name with ref
const form = reactive({
  version: '',
  memory: '',
  virtualMemory: '',
  uptime: '',
  cpuNum: '',
  platform: '',
  hostUptime: '',
  mem: '',
  memUsed: '',
  options: []
})
const value = ref('')

const onSubmit = () => {
  console.log('submit!')
}

const { proxy }:any = getCurrentInstance()

onMounted(() => {
  proxy.axios.get('api/log/sys', { card: 111 })
    .then((e:any) => {
      var data = e.data.data
      form.version = data.version
      form.uptime = (data.uptime / 60 / 60).toFixed(2)
      form.hostUptime = (data.hostUptime / 60 / 60).toFixed(2)
      form.memory = data.memory
      form.virtualMemory = data.virtualMemory
      form.cpuNum = data.cpuNum
      form.platform = data.platform
      form.mem = data.mem
      form.memUsed = data.memUsed
    })
  proxy.axios.get('api/log/names', { card: 111 })
    .then((e:any) => {
      form.options = e.data.data
    })
})

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
