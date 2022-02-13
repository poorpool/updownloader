<template>
  <el-row gutter="5" justify="center" style="margin-bottom: 5px">
    <el-col :xs="24" :sm="22" :md="20" :lg="18" :xl="16">
      <el-table :data="tableData" border style="width: 100%"
                :row-class-name="tableRowClassName">
        <el-table-column prop="code" label="提取代码">
          <template #default="scope">
            <el-link :href="'/'+scope.row.code" target="_blank" type="primary">{{ scope.row.code }}</el-link>
          </template>
        </el-table-column>
        <el-table-column prop="kind" label="文件类型"/>
        <el-table-column prop="filename" label="文件名"/>
        <el-table-column prop="expireTime" label="过期时间"/>
        <el-table-column prop="actions" label="操作">
          <template #default="scope">
            <el-button type="danger" @click="deleteRecord(scope.row.code)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
import axios from "axios";
import {ElMessage, ElMessageBox} from "element-plus";

export default {
  name: "Admin",
  data() {
    return {
      tableData: []
    }
  },
  methods: {
    // eslint-disable-next-line no-unused-vars
    tableRowClassName(row, rowIndex) {
      return row.row.kind === "文本" ? "text-row" : "file-row";
    },
    queryAll() {
      axios.get("/all").then(({ data }) => {
        if (data.status === 0) {
          this.tableData = [];
          for (const v of data.records) {
            this.tableData.push({
              code: v.Code,
              kind: v.Kind === 1 ? "文本" : "文件",
              filename: v.Kind === 1 ? "-" : v.Filename,
              expireTime: new Date(v.ExpireTime).toLocaleString()
            })
          }
        } else {
          ElMessage.error(data.msg)
        }
      });
    },
    deleteRecord(code) {
      ElMessageBox.confirm('你确定要删除 ' + code + " 对应的文件吗？", '确认删除',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      ).then(() => {
        axios({
          method: 'DELETE',
          // headers: { 'content-type': 'application/x-www-form-urlencoded' },
          // data: { "code": code },
          url: "/record/" + code
        }).then(({ data }) => {
          if (data.status === 0) {
            ElMessage.success("删除成功");
            this.queryAll();
          } else {
            ElMessage.error(data.msg);
          }
        });
      }).catch(() => {
        ElMessage.info("取消删除");
      });
    }
  },
  mounted() {
    this.queryAll()
  }
}
</script>

<style>

.el-table .text-row {
  --el-table-tr-bg-color: var(--el-color-success-lighter);
}
.el-table .file-row {
  --el-table-tr-bg-color: var(--el-color-warning-lighter);
}

</style>