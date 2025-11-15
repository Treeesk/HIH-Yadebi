<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
}
</script>

<style>
html, body, #app {
  margin: 0;
  padding: 0;

  /* Разрешаем скролл везде */
  overflow-y: auto !important;

  /* Белый фон — безопасно */
  background: #ffffff !important;

  /* Плавная мобильная прокрутка */
  -webkit-overflow-scrolling: touch;

  /* ❗ Главное исправление — УБРАТЬ высоту 100% */
  height: auto !important;
  min-height: 100%;
}
</style>