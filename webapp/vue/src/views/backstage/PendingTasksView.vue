<script setup>
import { onBeforeMount, onMounted, reactive, ref } from 'vue'
import router from '@/router'
import ApplyInfoComponent from '@/components/ApplyInfoComponent.vue'
import { getApplyBy } from '@/server/backstageService'

const applys = ref([])

const buttons = ref([])

onBeforeMount(() => {
  getApplyBy('new')
    .then((res) => {
      applys.value = res.data
    })
    .catch((error) => {
      if (error.response.status == 401) {
        alert('登录已过期,请先登录')
        router.replace('/bs/')
      } else if (error.response.status == 400) {
        alert('网页缓存版本可能过老,请ctrl+f5刷新网页或清理缓存')
        router.replace('/')
      }
    })
})
</script>

<template>
  <div class="row w-100 p-3 mx-0">
    <ApplyInfoComponent
      class="mb-3 col-lg-4 col-md-6 col-12"
      v-for="apply in applys"
      :apply-info="apply"
      :command-buttons="buttons"
    />
  </div>
</template>

<style scoped></style>
