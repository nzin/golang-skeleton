<template>
  <el-dialog v-model="createVisible" title="Create Todo" width="60%">
    <table width="100%">
      <tr>
        <td>Title</td>
        <td><el-input v-model="newTitle"></el-input></td>
      </tr>
      <tr>
        <td>Body</td>
        <td><el-input type="textarea" v-model="newBody"></el-input></td>
      </tr>
    </table>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="createTodo">Create</el-button>
      </span>
    </template>
  </el-dialog>
  <el-row>
    <el-col :span="20" :offset="2">
      <div class="container">

        <spinnerApp v-if="!loaded" />

        <div v-if="loaded">
          <span>
            Actions:
          </span>
          <span>
            <el-button type="primary" @click="setCreateVisible">Create a todo</el-button>
          </span>
          <el-divider/>
          <el-row>
            <el-table
              :data="todos"
              :stripe="true"
              :highlight-current-row="false"
              :default-sort="{ prop: 'title', order: 'descending' }"
              v-on:row-click="goToTodo"
              style="width: 100%"
            >
              <el-table-column width="150" prop="title" align="left" label="Todo title" sortable></el-table-column>
              <el-table-column prop="body" align="left" label="Todo body">
                <template v-slot="{row}">
                  {{ row.body.substring(0, 50) }}
                </template>
              </el-table-column>
            </el-table>
          </el-row>
        </div>
      </div>
    </el-col>
  </el-row>
</template>
  
  <script>
  import Axios from "axios";
  
  import constants from "@/constants";
  import SpinnerApp from "@/components/SpinnerApp";
  import helpers from "@/helpers/helpers";
  
  const { handleErr } = helpers;
  
  const { API_URL } = constants;
  
  export default {
    name: "TodosApp",
    components: {
      spinnerApp: SpinnerApp
    },
    data() {
      return {
        loaded: false,
        todos: [],
        createVisible:false,
        newTitle: "",
        newBody: "",
      };
    },
    created() {
      this.getTodos()
    },
    methods: {
      getTodos() {
        Axios.get(`${API_URL}/todos`).then(response => {
          let todos = response.data;
          this.loaded = true;
          this.todos = todos.todos;
        }, handleErr.bind(this));
      },
      goToTodo(row) {
        this.$router.push({ name: "todo", params: { todoid: row.id } });
      },
      setCreateVisible() {
        this.newTitle=""
        this.newBody=""
        this.createVisible=true
      },
      createTodo() {
        Axios.post(`${API_URL}/todos`,{title:this.newTitle,body:this.newBody}).then( () => {
          this.getTodos()
          this.createVisible=false
        }, handleErr.bind(this));

      }
    }
  };
  </script>
    