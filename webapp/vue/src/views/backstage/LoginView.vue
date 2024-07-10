<script setup>
import FormComponent from '@/components/FormComponent.vue'
import router from '@/router'
import { postLogin } from '@/server/backstageService'
import { reactive } from 'vue'

const tableProtos = [
  {
    name: 'username',
    title: '用户名',
    type: 'text',
    require: true
  },
  {
    name: 'password',
    title: '密码',
    type: 'password',
    require: true
  }
]

let userInfo = reactive({})

const submit = () => {
  postLogin(userInfo)
    .then((resCode) => {
      router.replace('pendingTasks')
    })
    .catch((error) => {
      if (error.response.stauts == 500) {
        alert('服务器内部错误,如可以请联系平台管理员.或稍后再试.')
      } else if (error.response.stauts == 401) {
        alert('用户名或密码错误,请检查后重试.')
      } else {
        alert('请求发送失败,请检查您的浏览器状态与网络状态.')
      }
    })
}
</script>

<template>
  <div
    class="d-flex h-100 w-100 flex-xl-row flex-column justify-content-center align-content-center flex-wrap"
  >
    <div class="col-xl-6 col-12">
      <p class="h1 text-center mt-5">浙理计协硬件部报修后台</p>
      <hr />
      <p class="h5 text-center">如没有账号或忘记密码,请联系硬件部部长.</p>
    </div>
    <FormComponent
      class="col-xl-6 col-12 pt-xl-3 flex-shrink-1 px-3 overflow-auto overflow-x-hidden"
      :protos="tableProtos"
      v-model="userInfo"
      button-text="登录"
      :onsubmit="submit"
    />
  </div>
</template>

<style scoped></style>
