<template>
  <v-sheet class="bg-deep pa-12">
    <v-card class="mx-auto px-6 py-8" max-width="344" variant="outlined">
      <v-form v-model="form" @submit.prevent="onSubmit">
        <v-text-field v-model="username" label="用户名" :readonly="loading" :rules="[required]"
          :error="!pwdIsCorrect || !userExisted" :error-messages="userExisted ? '' : '用户不存在'"
          @input="userExisted = pwdIsCorrect = true" clearable></v-text-field>
        <v-text-field v-model="password" label="密码" :readonly="loading" :rules="[required]"
          :error-messages="pwdIsCorrect ? '' : '密码或用户名错误'" @input="pwdIsCorrect = true" type="password"
          clearable></v-text-field>
        <br>
        <v-btn :disabled="!form || loading" :loading="loading" block size="large" type="submit"
          variant="outlined">
          登录
        </v-btn>
        <br>
        <v-btn :disabled="loading" block size="large" @click="GotoSignup" variant="outlined">
          注册
        </v-btn>
      </v-form>
    </v-card>
  </v-sheet>
</template>

<script>
import axios from 'axios'
import EncryptData from '../utils';

export default {
  data: () => ({
    form: false,
    username: null,
    password: null,
    loading: false,
    userExisted: true,
    pwdIsCorrect: true,
  }),
  methods: {
    onSubmit() {
      if (!this.form) return

      this.loading = true
      let userData = JSON.stringify({
        username: this.username,
        password: this.password
      })
      let encrpytedData = EncryptData(userData)
      // console.log(userData, encrpytedData)
      // ?redirecturi=http://localhost:3000/login
      let redirectUri = this.$route.query.redirecturi

      axios({
        url: 'http://localhost:8080/logi/login',
        method: 'post',
        headers: {
          "Redirect": redirectUri
        },
        data: encrpytedData,
      }).then(response => {
        this.loading = false;
        console.log(response)
        if (response.data != null) {
          if (response.data.hasOwnProperty("Error")) {
            this.snackbar = true
            this.errorText = response.data.Error
            return
          }
        }
      }).catch((error) => {
        console.log(error, error.response.status);
        switch (error.response.status) {
          case 302:
            // 登陆成功，重定向回原网页
            // console.log(error.response.headers)
            if (error.response.headers.hasOwnProperty("redirect")) {
              window.location.href = error.response.headers.redirect + "?token=" + error.response.data.token
            } else {
              this.$router.push("/ok")
            }
            $cookies.config("24h")
            $cookies.set("sso_token", error.response.data.token)
            break
          case 338:
            // 用户名不存在
            this.userExisted = false
            break
          case 339:
            // 密码错误
            this.pwdIsCorrect = false
            break
        }
        this.loading = false;
      });

    },
    GotoSignup() {
      if (this.username != "") {
        sessionStorage.setItem("tmp_username", this.username)
      }
      let url='/signup'
      if (this.$route.query.hasOwnProperty("redirecturi")) {
        url+="?redirecturi="+this.$route.query.redirecturi
      }
      this.$router.push(url)
    },
    required(v) {
      return !!v || '必填项'
    },
  },
}
</script>
