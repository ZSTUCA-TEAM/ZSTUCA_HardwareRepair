<script setup>
// 强制按日刷新请求
document.body.style += ';background-image: url(https://api.kdcc.cn?' + new Date().getDate() + ');'
</script>

<template>
  <RouterView />
</template>

<style>
html,
body {
  height: 100%;
}

body {
  background: url(https://api.kdcc.cn) no-repeat center center;
}
</style>
