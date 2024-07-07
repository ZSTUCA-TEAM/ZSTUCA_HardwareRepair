<script setup></script>

<template>
  <div class="d-flex flex-column h-100">
    <!-- 顶部导航栏 -->
    <nav class="navbar navbar-expand-md bg-dark navbar-dark">
      <!-- 标题 -->
      <a class="navbar-brand ms-3" href="javascript:void(0)">浙理计协硬件部报修后台</a>

      <!-- 屏幕尺寸不足时的折叠导航栏 -->
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#collapsibleNavbar"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <!-- 导航栏内容 -->
      <div class="collapse navbar-collapse" id="collapsibleNavbar">
        <ul class="navbar-nav">
          <li class="nav-item">
            <RouterLink class="nav-link" to="pendingTasks/">待接取的任务</RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link" to="receiptedTasks/">已接取的任务</RouterLink>
          </li>
          <li class="nav-item d-none">
            <RouterLink class="nav-link" to="allTasks/">所有委托信息</RouterLink>
          </li>
          <li class="nav-item d-none">
            <RouterLink class="nav-link" to="adminRegister/">硬件部成员注册</RouterLink>
          </li>
        </ul>
      </div>
    </nav>

    <main class="flex-shrink-1 h-100">
      <RouterView class="container-fluid row"></RouterView>
    </main>
  </div>
</template>

<style scoped>
main {
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 0px 0px 30px 30px;
}
</style>
