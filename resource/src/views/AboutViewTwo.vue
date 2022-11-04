<template>
  <el-row :gutter="12">
    <el-col :span="8">
      <el-card shadow="always"> Version  </el-card>
    </el-col>
    <el-col :span="8">
      <el-card shadow="hover"> Memory Usage </el-card>
    </el-col>
    <el-col :span="8">
      <el-card shadow="never"> Dbs </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { reactive, getCurrentInstance, onMounted, ref } from 'vue'

// do not use same name with ref
const form = reactive({
  name: '',
  region: '',
  date1: '',
  date2: '',
  delivery: false,
  type: [],
  resource: '',
  desc: ''
})

const onSubmit = () => {
  console.log('submit!')
}

const { proxy }:any = getCurrentInstance()

onMounted(() => {
  proxy.axios.get('api/log/test3', { card: 111 })
    .then((e:any) => {
      form.desc = e.data.data
    })
})
</script>
