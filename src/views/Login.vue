<template>
  <v-main>
    <v-container fluid fill-height>
      <div class="welcome mx-auto text-h4 text-lg-h1 text-sm-h2 font-weight-bold">バンピンぐへようこうそ</div>
      <v-row cols="12" align="center" justify="center">
        <v-col :sm="6" :lg="2" :md="4" class="mx-auto">
          <form>
            <v-text-field
              v-model="username"
              :counter="10"
              :rules="usernameRules"
              label="Username"
              required
              outlined
              dense
              color="pink lighten-3"
            ></v-text-field>

            <v-text-field
              v-model="password"
              :rules="passwordRules"
              label="Password"
              required
              outlined
              dense
              :counter="16"
              color="purple lighten-3"
            ></v-text-field>

            <v-btn class="mr-4" @click="doLogin">Login</v-btn>
            <v-btn @click="doClear">Clear</v-btn>
          </form>
        </v-col>
      </v-row>
    </v-container>
    <v-snackbar
      v-model="snackbar"
      :timeout="2000"
    >
      {{ text }}

      <template v-slot:action="{ attrs }">
        <v-btn
          color="pink lighten-1"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-main>
</template>

<script>
import { reactive, toRefs } from "@vue/composition-api";
import axios from 'axios'
axios.defaults.baseURL = "/api/v1/"
export default {
  setup(props, { root }) {
    
    const state = reactive({
      username: "",
      password: "",
      usernameRules: [
        value => !!value || "Required.",
        value => /^[a-z0-9A-Z]{4,10}$/.test(value) || "特殊符号不可输入"
      ],
      passwordRules: [
        value => !!value || "Required.",
        value => /^[a-z0-9A-Z]{6,16}$/.test(value) || "特殊符号不可输入"
      ],
      snackbar:false,
      text:"",
    });
    
    const doLogin = async () => {
     let {data:res} = await axios.post('/login',JSON.stringify({
            username:state.username,
            password:state.password
        }));
      if(res.code !== 1000){
         state.snackbar = true;
         state.text = "用户名或密码错误"
         return 
      }
      state.snackbar = true;
      state.text = "登录成功";
      window.localStorage.setItem('token',res.data);
      root.$router.push('home');
      return
    }
    const doClear = () =>{
        state.username = "";
        state.password = "";
    }
    return {
      ...toRefs(state),
      doLogin,
      doClear,
    };
  }
};
</script>

<style scoped>
.welcome {
  background: linear-gradient(
    to right,
    rgba(249, 152, 226, 1),
    rgba(42, 121, 215, 0.849645390070922)
  );
  -webkit-background-clip: text;
  color: transparent;
}
</style>