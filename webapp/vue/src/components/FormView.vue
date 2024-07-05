<script setup>
import { onBeforeMount, onMounted, onUpdated } from 'vue'

const props = defineProps({
  // 表单描述对象
  protos: {
    type: Array,
    require: true
  },
  // 提交按钮文本
  buttonText: {
    type: String
  },
  // 提交事件
  onsubmit: {
    type: Function
  }
})

// 获取当前表单项实际类型,位于最后一个-之后
const getTrueType = (type) => {
  const trueType = type.match(/(?<=-).+$/)
  return trueType ? trueType[0] : type
}

const model = defineModel()

// 初始化列表选择项的初始选项为数组的第一项
const initSelect = () => {
  props.protos.forEach((proto) => {
    if (getTrueType(proto.type) == 'select') {
      model.value[proto.name] = proto.additional[0]
    } else if (proto.type == 'typeSelect-text') {
      model.value[proto.name + 'Type'] = proto.additional[0]
    }
  })
}

// 在页面加载前调用initSelect
onBeforeMount(() => {
  initSelect()
})
</script>

<template>
  <form class="h-100 overflow-scroll overflow-x-hidden" @submit.prevent="props.onsubmit">
    <div class="my-3" v-for="proto in props.protos">
      <!-- 表单单项标签 -->
      <label class="form-label" :for="proto.name">{{ proto.title }}</label>
      <div class="input-group">
        <!-- 带类型选择项的表单项 -->
        <div class="col-5 col-xl-3" v-if="proto.type == 'typeSelect-text'">
          <select class="form-select" v-model="model[proto.name + 'Type']">
            <option v-for="addi in proto.additional">{{ addi }}</option>
          </select>
        </div>

        <!-- 文本输入 -->
        <input
          class="form-control"
          v-if="(getTrueType(proto.type) == 'text') | (getTrueType(proto.type) == 'email')"
          :required="proto.required"
          :type="getTrueType(proto.type)"
          :id="proto.name"
          v-model="model[proto.name]"
        />
        <!-- 列表选择 -->
        <select
          class="form-select"
          v-else-if="getTrueType(proto.type) == 'select'"
          :required="proto.require"
          :id="proto.name"
          v-model="model[proto.name]"
        >
          <option v-for="addi in proto.additional">{{ addi }}</option>
        </select>
      </div>
    </div>
    <!-- 提交按钮 -->
    <div class="text-center mt-5 mb-3">
      <button type="submit" class="btn btn-primary text-nowrap w-100">
        {{ props.buttonText ? props.buttonText : '提交' }}
      </button>
    </div>
  </form>
</template>

<style scoped></style>
