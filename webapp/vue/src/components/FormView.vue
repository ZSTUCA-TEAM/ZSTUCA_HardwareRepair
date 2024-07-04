<script setup>
const props = defineProps({
  protos: {
    type: Array,
    require: true
  },
  buttonText: {
    type: String
  },
  onsubmit: {
    type: Function
  }
})

const getTrueType = (type) => {
  const trueType = type.match(/(?<=-).+$/)
  return trueType ? trueType[0] : type
}

const model = defineModel({
  res: {
    require: true
  }
})
</script>

<template>
  <form class="h-100 overflow-scroll overflow-x-hidden" onsubmit="return false">
    <div class="my-3" v-for="proto in props.protos">
      <label class="form-label" :for="proto.name">{{ proto.title }}</label>
      <div class="input-group">
        <div class="col-5 col-xl-3" v-if="proto.type == 'typeSelect-text'">
          <select class="form-select" :id="proto.name + 'Type'">
            <option v-for="addi in proto.additional">{{ addi }}</option>
          </select>
        </div>
        <input
          class="form-control"
          v-if="(getTrueType(proto.type) == 'text') | (getTrueType(proto.type) == 'email')"
          :required="proto.required"
          :type="getTrueType(proto.type)"
          :id="proto.name"
        />
        <select
          class="form-select"
          v-else-if="getTrueType(proto.type) == 'select'"
          :required="proto.require"
          :id="proto.name"
        >
          <option v-for="addi in proto.additional">{{ addi }}</option>
        </select>
      </div>
    </div>
    <div class="text-center mt-5 mb-3">
      <button type="submit" class="btn btn-primary text-nowrap w-100" @click="props.onsubmit">
        {{ props.buttonText ? props.buttonText : '提交' }}
      </button>
    </div>
  </form>
</template>

<style scoped></style>
