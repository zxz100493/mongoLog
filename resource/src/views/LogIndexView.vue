<template>
  <div class="">
    <el-container>
      <el-header>Sslog V1.0.0</el-header>
      <el-container>
        <el-aside width="200px">
          <sideBar />
        </el-aside>
        <el-main>
          <!-- 二级路的出口 -->
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts">
import sideBar from '@/views/sideBar/SideIndex.vue'
export default {
  components: {
    sideBar
  },
  name: 'logIndex',
  methods: {}
}
</script>
<style scoped>
.el-header,
.el-footer {
  background-color: #b3c0d1;
  color: #333;
  text-align: center;
  line-height: 60px;
}
.el-aside {
  background-color: #d3dce6;
  color: #333;
  text-align: center;
  /* line-height: 200px; */
}
/* .el-main {
  background-color: #e9eef3;
  color: #333;
  text-align: center;
  line-height: 160px;
} */
body > .el-container {
  margin-bottom: 40px;
}
.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}
.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>
