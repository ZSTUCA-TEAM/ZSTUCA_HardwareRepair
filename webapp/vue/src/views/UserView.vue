<script setup>
import { reactive } from 'vue'
import FormView from '@/components/FormView.vue'
import { postApply } from '@/server/ApplyService.js'

// 表单描述
const tableProtos = [
  {
    name: 'name',
    title: '姓名',
    type: 'text',
    required: true
  },
  {
    name: 'gender',
    title: '性别',
    type: 'select',
    required: true,
    additional: ['先生', '女士']
  },
  {
    name: 'academy',
    title: '学院',
    type: 'text',
    required: true
  },
  {
    name: 'cardId',
    title: '学号(工号)',
    type: 'text',
    required: true
  },
  {
    name: 'email',
    title: '邮箱',
    type: 'email',
    required: true
  },
  {
    name: 'contact',
    title: '即时通讯联系方式',
    type: 'typeSelect-text',
    required: true,
    additional: ['QQ', '微信', '手机']
  },
  {
    name: 'computerType',
    title: '电脑型号',
    type: 'text'
  },
  {
    name: 'problem',
    title: '遇到的问题',
    type: 'text',
    required: true
  },
  {
    name: 'location',
    title: '宿舍(办公室)门牌号',
    type: 'typeSelect-text',
    required: true,
    additional: ['生活一区', '生活二区', '生活三区', '下沙校区']
  }
]

// 表单结果
let apply = reactive({})

//
const submit = () => {
  postApply(apply)
    .then((resCode) => {
      if (resCode == 201) {
        alert(
          '提交成功,已向您的邮箱发送了一份邮件,后续消息将通过邮箱通知您.如果您的邮箱自动拦截了这份邮件,请到垃圾箱中查看并设置正常接收来自我们的邮件.'
        )
      } else if (resCode == 409) {
        alert('您已经提交过了,请勿在10秒内重复提交.')
      } else if (resCode == 500) {
        alert('服务器内部错误,如可以请联系平台管理员.如稍后再试.')
      } else {
        alert('请求发送失败,请检查您的浏览器状态与网络状态.')
      }
    })
    .catch((error) => {
      alert('请求发送失败,请检查您的浏览器状态与网络状态.')
    })
}
</script>

<template>
  <main class="d-flex align-items-center justify-content-center h-100 w-100">
    <div class="d-flex flex-xl-row flex-column" id="main">
      <div id="title" class="text-center col-xl-6 col-12 mt-lg-5 pe-xl-5">
        <p class="h1 mt-5 mb-3"><b>浙理计算机协会</b></p>
        <p class="h1 mb-4"><b>硬件部报修平台</b></p>
        <p class="small">提供免费硬件维修服务.临进考试周后不接收委托.</p>
        <p class="small">
          如委托长时间未被接收,<s>说明维修同学懒癌发作了,</s>可以考虑去找其他维修平台.
        </p>
        <p class="small">后续会推出催接委托功能.</p>
      </div>
      <FormView
        class="col-xl-6 col-12 pt-xl-3 flex-shrink-1 px-3"
        id="formView"
        :protos="tableProtos"
        v-model="apply"
        :onsubmit="submit"
      />
    </div>
  </main>
</template>

<style scoped>
#main {
  height: 95%;
  width: 95%;
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 30px;
}
</style>
