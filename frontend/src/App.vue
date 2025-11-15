<template>
  <div id="app" class="app-root">
    <transition name="slide-fade" mode="out-in">
      <router-view />
    </transition>
  </div>
</template>

<script>
export default {
  name: "App"
}
</script>

<style>
/* ====== GLOBAL RESET ====== */
html, body, #app {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  background: white;
  overscroll-behavior-y: contain;
  -webkit-overflow-scrolling: touch;
}

/* Чтобы Vue не ломал позиционирование */
.app-root {
  width: 100%;
  min-height: 100vh;
  overflow-x: hidden;
  position: relative;
}

/* ====== iOS-STYLE PAGE TRANSITION ======
   Плавный вход → выезд справа, fade, scale
   Плавный выход → уход влево с fade и лёгким scale
*/

.slide-fade-enter-active {
  transition: all 0.35s cubic-bezier(0.22, 0.61, 0.36, 1);
}

.slide-fade-leave-active {
  transition: all 0.28s cubic-bezier(0.55, 0.06, 0.68, 0.19);
  position: absolute;
  width: 100%;
  top: 0;
  left: 0;
}

/* ВХОД — новая страница */
.slide-fade-enter-from {
  opacity: 0;
  transform: translateX(40px) scale(0.97);
}

.slide-fade-enter-to {
  opacity: 1;
  transform: translateX(0) scale(1);
}

/* ВЫХОД — старая страница */
.slide-fade-leave-from {
  opacity: 1;
  transform: translateX(0) scale(1);
}

.slide-fade-leave-to {
  opacity: 0;
  transform: translateX(-20px) scale(0.98);
}

</style>