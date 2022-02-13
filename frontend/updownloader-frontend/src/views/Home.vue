<template>
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
            <el-button class="button" type="text" v-if="cardKind===1" @click="copyToClip(cardText)">复制到剪贴板</el-button>
          </div>
        </template>
        <div class="show-text" style="white-space: pre-wrap;" v-if="cardKind===1" >{{cardText}}</div>
        <el-link :href="cardText" target="_blank" type="primary" v-if="cardKind===2" >{{ cardText }}</el-link>
        <!--            <pre>{{cardText}}</pre>-->
      </el-card>
    </el-col>
  </el-row>
</template>

<script>

import axios from "axios";
import {ElMessage} from "element-plus";

export default {
  name: 'Home',
  components: {
  },
  data() {
    return {
      codeInput: "",
      ifShowCard: false,
      cardTitle: "",
      cardText: "",
      cardKind: 1,
    }
  },
  methods: {
    query() {
      axios.get("/record/"+this.codeInput).then(({ data }) => {
        if (data.status === 0 && data.kind === 1) {
          this.cardTitle = this.codeInput;
          this.cardText = data.msg;
          this.cardKind = 1;
          this.ifShowCard = true;
          ElMessage.success("获取成功")
        } else if (data.status === 0 && data.kind === 2) {
          this.cardTitle = this.codeInput;
          this.cardText = data.msg;
          this.cardKind = 2;
          this.ifShowCard = true;
          ElMessage.success("获取成功")
        } else {
          ElMessage.error(data.msg)
        }
      });
    },

    copyToClip(arg) {
      this.$copyText(arg).then(() => {
        ElMessage.success("已复制到剪贴板")
      }, function (e) {
        ElMessage.error(e)
      });
    },
  },
  mounted() {
    let code = this.$route.params.code;
    if (code !== undefined && code !== "") {
      this.codeInput = code;
      this.query();
    }
  },
}
</script>

<style>

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.show-text {
  font-family: SF Mono, Monaco, JetBrains Mono, Fira Code, Consolas, Courier New, monospace, Noto Sans CJK SC, Avenir, Helvetica, Arial, sans-serif;
}
</style>
