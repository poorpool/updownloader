<template>
  <el-dialog v-model="dialogTextVisible" title="提交文字" width="80%">
    <el-input
        v-model="textarea"
        maxlength="50000"
        placeholder="请输入文字"
        show-word-limit
        type="textarea"
        rows="15"
    />
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="submitText">提交</el-button>
      </span>
    </template>
  </el-dialog>
  <el-dialog v-model="dialogFileVisible" title="提交文件" width="80%">
    <el-upload
        ref="upload"
        class="upload-demo"
        action="https://jsonplaceholder.typicode.com/posts/"
        :limit="1"
        :on-exceed="handleExceed"
        :auto-upload="false"
    >
      <el-button type="primary" plain>选择文件</el-button>
      <template #tip>
        <div class="el-upload__tip" style="color: red">
          仅限一个文件
        </div>
      </template>
    </el-upload>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary">提交</el-button>
      </span>
    </template>
  </el-dialog>
  <el-container>
    <el-header style="padding: 0;">
      <el-menu mode="horizontal" >
        <el-menu-item index="1">UPDOWNLOADER</el-menu-item>
        <el-menu-item index="2" @click="dialogTextVisible=true">上传文字</el-menu-item>
        <el-menu-item index="3" @click="dialogFileVisible=true">上传文件</el-menu-item>
        <el-menu-item index="4">后台</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>


      <el-row gutter="5" justify="center" style="margin-bottom: 5px">
        <el-col :xs="18" :sm="16" :md="12" :lg="12" :xl="12">
          <el-input v-model="codeInput" placeholder="输入获取代码"
                    size="large" maxlength="6"
                    @keyup.enter="query"
          />
        </el-col>
        <el-col :xs="6" :sm="4" :md="2" :lg="2" :xl="2">
          <el-button type="primary" size="large" style="display: flex" @click="query">查询</el-button>
        </el-col>
      </el-row>

      <el-row justify="center">
        <el-col :xs="24" :sm="24" :md="20" :lg="18" :xl="16">
          <el-card class="box-card" v-if="ifShowCard">
            <template #header>
              <div class="card-header">
                <span>{{cardTitle}}</span>
                <!--            <el-button class="button" type="text">Operation button</el-button>-->
              </div>
            </template>
            <div style="white-space: pre-wrap;">{{cardText}}</div>

          </el-card>
        </el-col>
      </el-row>


    </el-main>
  </el-container>
</template>

<script>

import axios from "axios";
import {ElMessage, ElMessageBox} from "element-plus";
import qs from "qs";

export default {
  name: 'App',
  components: {
  },
  data() {
    return {
      codeInput: "",
      ifShowCard: false,
      cardTitle: "",
      textarea: "",
      dialogTextVisible: false,
      dialogFileVisible: false,
      cardText: "",
      cardKind: 1,
    }
  },
  methods: {
    query() {
      axios.get("http://localhost:10370/updown/record/"+this.codeInput).then(({ data }) => {
        if (data.status === 0 && data.kind === 1) {
          this.cardTitle = this.codeInput;
          this.cardText = data.msg;
          this.cardKind = 1;
          this.ifShowCard = true;
          ElMessage.success("获取成功")
        } else if (data.status === 0 && data.kind === 2) {
          ElMessage.error(data.msg)
        } else {
          ElMessage.error(data.msg)
        }
      });
    },
    handleExceed(files) {
      this.$refs.upload.clearFiles()
      this.$refs.upload.handleStart(files[0])
    },
    submitText() {

      const data = { "text": this.textarea };
      const options = {
        method: 'POST',
        headers: { 'content-type': 'application/x-www-form-urlencoded' },
        data: qs.stringify(data),
        url: "http://localhost:10370/updown/text"
      };
      axios(options).then(({ data }) => {
        if (data.status === 0) {
          ElMessageBox.alert('提取代码是: ' + data.msg, '确认你的提取代码', {
            confirmButtonText: '记下来了',
          })
          this.dialogTextVisible = false;
        } else {
          ElMessage.error(data.msg);
        }
      });
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
