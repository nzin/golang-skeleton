<template>
    <el-row>
        <el-col :span="20" :offset="2">
            <div class="container">
                <spinnerApp v-if="!loaded" />

                <div v-if="loaded">
                    <table width="100%">
                        <tr>
                            <td>Title</td>
                            <td><el-input v-model="todo.title"></el-input></td>
                        </tr>
                        <tr>
                            <td>Body</td>
                            <td><el-input type="textarea" v-model="todo.body"></el-input></td>
                        </tr>
                        <tr>
                            <td></td>
                            <td><span><el-button @click="saveTodo" type="primary">Save</el-button></span><span>&nbsp;</span><span><el-button @click="deleteTodo" type="danger">Delete</el-button></span></td>
                        </tr>
                    </table>
                </div>
            </div>
        </el-col>
    </el-row>
</template>
  
<script>
  import Axios from "axios";
  
  import constants from "@/constants";
  import helpers from "@/helpers/helpers";
  import SpinnerApp from "@/components/SpinnerApp";
  
  const { handleErr } = helpers;
  
  const { API_URL } = constants;
  

  export default {
    name: "TodoApp",
    components: {
        spinnerApp: SpinnerApp,
    },
    data() {
      return {
        loaded: false,
        todo: null,
      };
    },
    computed: {
      todoid() {
        return this.$route.params.todoid;
      },
    },
    methods: {
      fetchTodo() {
        Axios.get(`${API_URL}/todos/${this.todoid}`).then(response => {
          let todo = response.data;
          this.todo = todo;
          this.loaded = true;
        }, handleErr.bind(this));
      },
      saveTodo() {
        Axios.put(`${API_URL}/todos/${this.todoid}`,{title:this.todo.title,body:this.todo.body}).then( () => {
            this.$router.push({ name: "todos"});
        }, handleErr.bind(this));
      },
      deleteTodo() {
        Axios.delete(`${API_URL}/todos/${this.todoid}`).then( () => {
            this.$router.push({ name: "todos"});
        }, handleErr.bind(this));
      }
    },
    mounted() {
      this.fetchTodo();
    }
  };
</script>
  
  