<script setup>
import { ref } from 'vue'
import FormView from '@/components/FormView.vue'

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

let res = ref({})
</script>

<template>
  <main class="d-flex align-items-center justify-content-center h-100 w-100">
    <div class="row" id="main">
      <div id="title" class="text-center col-xl-6 col-12 mt-lg-5 pe-xl-5">
        <p class="h1 mt-5 mb-3">浙理计算机协会</p>
        <p class="h1 mb-4">硬件部报修平台</p>
        <p class="small">提供免费硬件维修服务.临进考试周后不接收委托.</p>
        <p class="small">
          如委托长时间未被接收,<s>说明维修同学懒癌发作了,</s>可以考虑去找其他维修平台.
        </p>
        <p class="small">后续会推出催接委托功能.</p>
      </div>
      <FormView class="col-xl-6 col-12 pt-xl-3" id="formView" :protos="tableProtos" v-model="res" />
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

#formView {
  flex: 1;
}
</style>
