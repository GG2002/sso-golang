<template>
  <router-view />
</template>

<script>
import axios from 'axios'

export default {
  created: async function () {
    // localStorage.clear()
    if (!localStorage.getItem("rsaPubKey")) {
      // 获取公钥进行RSA加密
      await axios({
        url: 'http://localhost:8080/getpubkey',
        method: 'post',
      }).then(response => {
        // console.log(response)
        localStorage.setItem("rsaPubKey", response.data.rsaPubKey)
      }).catch(function (error) {
        console.log(error);
      });
    }
    if (window.location == "http://localhost:3000/ok") {
      return
    }
    if ($cookies.isKey("sso_token")) {
      // 检验登录状态
      await axios({
        url: 'http://localhost:8080/logi/logcheck',
        method: 'post',
        withCredentials: true,
      }).then(response => {
        // console.log(response)
        if (response.data.hasOwnProperty("Error")) {
          // 未知错误
          console.log(response.data.Error)
          $cookies.remove("sso_token")
        } else {
          // 验证成功
          if (this.$route.query.hasOwnProperty("redirecturi")) {
            window.location.href = this.$route.query.redirecturi + "?token=" + $cookies.get("sso_token")
          } else {
            window.location.href = "http://localhost:3000/ok"
          }
        }
      }).catch(function (error) {
        console.log(error);
        switch (error.response.status) {
          case 340:
            // Token过期
            $cookies.remove("sso_token")
            break
          case 341:
            // Token错误或用户已下线
            $cookies.remove("sso_token")
            break

        }
      });
    }
  },
  mounted: function () {
    this.$route
  }
}
</script>
