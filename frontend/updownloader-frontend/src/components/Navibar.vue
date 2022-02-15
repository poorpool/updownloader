<template>

  <el-dialog
      v-model="dialogAboutVisible"
      title="关于 UPDOWNLOADER"
      width="80%"
  >
    <p>
      随时随地，在PC、手机、云服务器上上传和下载文本或者文件。
    </p>
    <p>
      项目地址 <a href="https://github.com/poorpool/updownloader" target="_blank">https://github.com/poorpool/updownloader</a>
    </p>
    <p>
      上传的文件将于3个小时后自动删除。本项目意在解决简单的个人同步需求，不对其中文件内容负责。本项目只得用于教育、非政治用途，不得上传违法文件。
    </p>
    <pre class="show-text">
upload: curl -X POST http://BACKEND_ADDRESS/updown/file -F "file=@YOUR_FILE_PATH"
query:  curl http://BACKEND_ADDRESS/updown/record/YOUR_CODE
    </pre>
    </el-dialog>

  <el-dialog
      v-model="dialogCodeVisible"
      title="确认你的提取代码"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="false"
      width="80%"
  >
    <p>
      你的提取代码是 <span class="show-text" style="color: #409EFF" @click="copyToClip(showCode)">{{showCode}}</span>
    </p>
    <p>
      或者使用 <span class="show-text" style="color: #409EFF" @click="copyToClip(frontendHost+'/'+showCode)">{{frontendHost}}/{{showCode}}</span> 直接提取内容
    </p>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="dialogCodeVisible = false">确认</el-button>
      </span>
    </template>
  </el-dialog>

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
        :action="backendBaseUrl+'/file'"
        :limit="1"
        :on-exceed="handleExceed"
        :auto-upload="false"
        :before-upload="beforeUpload"
        :on-success="uploadSuccess"
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
        <el-button type="primary" @click="this.$refs.upload.submit()">提交</el-button>
      </span>
    </template>
  </el-dialog>

  <el-menu mode="horizontal" >
    <el-menu-item index="1" @click="this.$router.push('/')">UPDOWNLOADER</el-menu-item>
    <el-menu-item index="2" @click="dialogTextVisible=true">上传文字</el-menu-item>
    <el-menu-item index="3" @click="dialogFileVisible=true">上传文件</el-menu-item>
    <el-menu-item index="4" @click="dialogAboutVisible=true">关于</el-menu-item>
<!--    <el-menu-item index="4" @click="this.$router.push('/admin')">后台</el-menu-item>-->
  </el-menu>
</template>

<script>
  import axios from "axios";
  import {ElMessage} from "element-plus";

  export default {
    name: "Navibar",
    data() {
      return {
        dialogTextVisible: false,
        dialogFileVisible: false,
        dialogCodeVisible: false,
        dialogAboutVisible: false,
        showCode: "",
        textarea: "",
        backendBaseUrl: axios.defaults.baseURL,
        frontendHost: window.location.host + "/updown"
      }
    },
    methods: {
      handleExceed(files) {
        this.$refs.upload.clearFiles()
        this.$refs.upload.handleStart(files[0])
      },
      submitText() {
        axios({
          method: 'POST',
          // headers: { 'content-type': 'application/x-www-form-urlencoded' },
          data: { "text": this.textarea },
          url: "/text"
        }).then(({ data }) => {
          if (data.status === 0) {
            this.dialogTextVisible = false;
            this.showTiqudaima(data.msg);
          } else {
            ElMessage.error(data.msg);
          }
        });
      },
      beforeUpload(file) {
        const isLt2M = file.size / 1024 / 1024 < 500
        if (!isLt2M) {
          ElMessage.error('文件大小不能超过 500MB!')
        }
        return isLt2M;
      },
      uploadSuccess(response) {
        if (response.status === 0) {
          this.dialogFileVisible = false;
          this.showTiqudaima(response.msg);
        } else {
          ElMessage.error(response.msg);
        }
      },
      copyToClip(arg) {
        this.$copyText(arg).then(() => {
          ElMessage.success("已复制到剪贴板")
        }, function (e) {
          ElMessage.error(e)
        });
      },
      showTiqudaima(code) {
        this.showCode = code;
        this.dialogCodeVisible = true;
      }
    }
  }
</script>

<style>
.show-text {
  font-family: SF Mono, Monaco, JetBrains Mono, Fira Code, Consolas, Courier New, monospace, Noto Sans CJK SC, Avenir, Helvetica, Arial, sans-serif;
}
</style>