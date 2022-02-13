<template>
  <el-container>
    <el-header style="padding: 0;">
      <Navibar></Navibar>
    </el-header>
    <el-main>
      <router-view></router-view>
    </el-main>
    <el-footer>
      UPDOWNLOADER Copyright © 2022-2022 poorpool
    </el-footer>
  </el-container>
</template>

<script>
import Navibar from "@/components/Navibar";

export default {
  name: "App",
  components: {
    Navibar
  }
}
</script>

<style>
#app {
  font-family: Noto Sans CJK SC, Avenir, Helvetica, Arial, sans-serif;
  font-size: 16px;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.el-container {
  min-height: 100vh;
}
.el-footer {
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
  border-top: 1px solid var(--el-border-color-base);
}
</style>