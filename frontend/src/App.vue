<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
}
</script>

<style>
html, body, #app {
  margin: 0;
  padding: 0;

  /* Включаем скролл НАВСЕГДА */
  overflow-y: auto !important;

  /* ДЕЛАЕМ ФОН БЕЛЫМ */
  background: #ffffff !important;

  /* Мобильная плавность */
  -webkit-overflow-scrolling: touch;

  height: 100%;
}
</style>