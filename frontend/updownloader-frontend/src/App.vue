<template>

  <el-container>
    <el-header style="padding: 0;">
      <el-menu mode="horizontal" >
        <el-menu-item index="1">UPDOWNLOADER</el-menu-item>
        <el-menu-item index="2">上传文字</el-menu-item>
        <el-menu-item index="3">上传文件</el-menu-item>
        <el-menu-item index="4">后台</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>


      <el-row gutter="5" justify="center" style="margin-bottom: 5px">
        <el-col :xs="18" :sm="16" :md="12" :lg="12" :xl="12">
          <el-input v-model="codeInput" placeholder="输入六位代码"
                    size="large" maxlength="6" minlength="6"
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
            <div v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</div>
          </el-card>
        </el-col>
      </el-row>


    </el-main>
  </el-container>
</template>

<script>

export default {
  name: 'App',
  components: {
  },
  data() {
    return {
      codeInput: "",
      ifShowCard: false,
      cardTitle: ""
    }
  },
  methods: {
    query() {
      this.cardTitle = this.codeInput;
      this.ifShowCard = true;
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
