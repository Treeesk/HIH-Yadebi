<template>
  <div id="app" class="w-full min-h-screen bg-white overflow-x-hidden">
    <router-view />
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
  width: 100%;
  height: 100%;
  background: white;
  overscroll-behavior-y: contain;
}
</style>