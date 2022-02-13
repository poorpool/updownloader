<template>
  <el-container>
    <el-header style="padding: 0;">
      <Navibar></Navibar>
    </el-header>
    <el-main>
      <router-view></router-view>
    </el-main>
    <el-footer>
      UPDOWNLOADER Copyright © 2022-2022 poorpool | 文件、文本定期清理 | 只可用于教育、非政治用途
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