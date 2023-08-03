<template>
    <v-sheet class="bg-deep pa-12">
        <v-card class="mx-auto " max-width="344" variant="outlined">
            <v-card-item>
                <div>
                    <div class="text-overline mb-1 text-center">
                        当前状态
                    </div>
                    <div class="text-h6 mb-1 text-center">
                        {{ logStatus[logStatusStr].statusText }}
                    </div>
                    <div class="text-caption text-center">{{ logStatus[logStatusStr].tips }}</div>
                </div>
            </v-card-item>
            <v-card-actions>
                <v-btn variant="outlined" block @click="SwitchLogStatus">
                    {{ logStatus[logStatusStr].btnText }}
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-sheet>
</template>

<script>
import axios from 'axios'

export default {
    data: () => ({
        "logStatus": {
            "logIn": {
                "statusText": "登录成功",
                "tips": "无重定向地址时，你会看见此界面",
                "btnText": "注销登录",
            },
            "logOut": {
                "statusText": "未登录",
                "tips": "请前往登录页面进行登录",
                "btnText": "登录",
            },
            "unknown": {
                "statusText": "获取状态中",
                "tips": "✐.ɴɪᴄᴇ ᴅᴀʏ 〰",
                "btnText": "(◍•ᴗ•◍)",
            }
        },
        "logStatusStr": "unknown"
    }),
    created: function () {

        if ($cookies.isKey("sso_token")) {
            // 检验登录状态
            axios({
                url: 'http://hustmaths.top/sso/logi/logcheck',
                method: 'post',
                withCredentials: true,
            }).then(response => {
                // console.log(response)
                if (response.data.hasOwnProperty("Error")) {
                    // 未知错误
                    console.log(response.data.Error)
                    this.logStatusStr = "logOut"
                } else {
                    this.logStatusStr = "logIn"
                }
            }).catch((error) => {
                console.log(error);
                this.logStatusStr = "logOut"
            });
        } else {
            this.logStatusStr = "logOut"
        }
    },
    methods: {
        SwitchLogStatus() {
            switch (this.logStatusStr) {
                case "logIn":
                    axios({
                        url: 'http://hustmaths.top/sso/logi/logout',
                        method: 'post',
                        withCredentials: true,
                    }).then(response => {
                        console.log(response)
                        this.logStatusStr = "logOut"
                        this.$router.push("ok")
                    }).catch(function (error) {
                        console.log(error);
                    });
                    break
                case "logOut":
                    let url = '/login'
                    if (this.$route.query.hasOwnProperty("redirecturi")) {
                        url += "?redirecturi=" + this.$route.query.redirecturi
                    }
                    this.$router.push(url)
                    break
                case "unknown":
                    break
            }
        }
    }
}
</script>