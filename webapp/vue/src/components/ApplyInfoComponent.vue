<script setup>
const props = defineProps({
  applyInfo: Object,
  showPrivacyInfo: Boolean,
  commandButtons: Array
})
</script>

<template>
  <div class="card">
    <div class="card-body text-start">
      <p class="mb-0">
        <span class="h1">{{ props.applyInfo.name }}</span>
        <span>{{ props.applyInfo.gender }}</span>
        <span class="small spanRight">{{ props.applyInfo.computerType }}</span>
      </p>
      <p class="small mb-0">学工号:{{ props.applyInfo.cardId }}</p>
      <p class="spanRight h1">{{ props.applyInfo.id }}</p>
      <p class="small mb-0" v-if="props.showPrivacyInfo">邮箱:{{ props.applyInfo.email }}</p>
      <p class="small mb-0" v-if="props.showPrivacyInfo">
        {{ props.applyInfo.contactType }}:{{ props.applyInfo.contact }}
      </p>
      <hr />
      <p>
        {{ props.applyInfo.problem }}
      </p>
    </div>
    <div class="card-footer text-muted text-center">
      <button class="btn btn-secondary fullWidth" v-if="props.applyInfo.isFinish" disabled>
        已完成
      </button>
      <button class="btn btn-secondary fullWidth" v-else-if="props.applyInfo.isAbandoned" disabled>
        已放弃
      </button>
      <div class="d-flex" v-else>
        <button
          class="d-block btn btn-success flex-shrink-1"
          v-for="button in props.commandButtons"
          :onclick="button.onclick"
        >
          {{ button.text }}
        </button>
      </div>
      {{ props.applyInfo.academy }}·{{ props.applyInfo.locationType }}{{ props.applyInfo.location
      }}<br />
      {{ props.applyInfo.createAt }}
    </div>
  </div>
</template>

<style scoped></style>
