<template>
    <v-sheet class="bg-deep pa-12">
        <v-card class="mx-auto px-6 py-8" max-width="344" variant="outlined">
            <v-form v-model="form" @submit.prevent="onSubmit">
                <v-text-field v-model="username" label="用户名" :readonly="loading" :rules="[required]"
                    :error-messages="unameExisted ? '用户名已存在' : ''" @input="checkUsername" clearable></v-text-field>
                <v-text-field v-model="password" label="密码" :append-inner-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                    :type="show ? 'text' : 'password'" @click:append-inner="show = !show" :readonly="loading"
                    :rules="[required, repeatPwdEqual, pwdMin]" clearable></v-text-field>
                <v-text-field v-model="repeatPassword" label="确认密码" :append-inner-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                    :type="show ? 'text' : 'password'" @click:append-inner="show = !show" :readonly="loading"
                    :rules="[required, repeatPwdEqual]" clearable></v-text-field>
                <br>
                <v-btn :disabled="!form || loading" :loading="loading" block size="large" type="submit"
                    variant="outlined">
                    注册
                </v-btn>
            </v-form>
        </v-card>
    </v-sheet>
    <v-snackbar v-model="snackbar" multi-line>
        {{ errorText }}
        <template v-slot:actions>
            <v-btn color="red" variant="text" @click="snackbar = false">
                Close
            </v-btn>
        </template>
    </v-snackbar>
</template>
  
<script>
import axios from 'axios'
import EncryptData from '../utils';

export default {
    data: () => ({
        form: false,
        username: sessionStorage.getItem("tmp_username"),
        password: null,
        repeatPassword: null,
        loading: false,
        unameExisted: false,
        show: false,
        snackbar: false,
    }),
    created: async () => {
        sessionStorage.removeItem("tmp_username")
    },
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

            axios({
                url: 'http://localhost:8080/sign/signup',
                method: 'post',
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
                let url = '/login'
                if (this.$route.query.hasOwnProperty("redirecturi")) {
                    url += "?redirecturi=" + this.$route.query.redirecturi
                }
                this.$router.push(url)
            }).catch((error) => {
                console.log(error, error.response.status);
                if (error.response.status == 337) {
                    this.unameExisted = true
                }
                this.loading = false;
            });

        },
        checkUsername() {
            let userData = JSON.stringify({
                username: this.username,
            })
            let encrpytedData = EncryptData(userData)
            // console.log(userData, encrpytedData)

            axios({
                url: 'http://localhost:8080/sign/usernameexisted',
                method: 'post',
                data: encrpytedData,
            }).then(response => {
                // console.log(response)
                this.unameExisted = false;
            }).catch((error) => {
                if (error.response.status == 337) {
                    this.unameExisted = true
                }
            });
        },
        required(v) {
            return !!v || '必填项'
        },
        repeatPwdEqual(v) {
            return this.password == v || '重复密码需与密码一致'
        },
        pwdMin(v) {
            return v.length >= 6 || '最少需要6个字符'
        },
    },
}
</script>
  