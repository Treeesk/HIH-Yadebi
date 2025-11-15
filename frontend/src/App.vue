<template>
  <div id="app">
    <transition name="calm" mode="out-in">
      <router-view />
    </transition>
  </div>
</template>

<style>
/* ГЛОБАЛЬНЫЕ СТИЛИ ДЛЯ УДАЛЕНИЯ ОТСТУПОВ */
html, body {
  margin: 0 !important;
  padding: 0 !important;
  width: 100% !important;
  overflow-x: hidden !important;
  background: white !important;
}

#app {
  margin: 0 !important;
  padding: 0 !important;
  width: 100% !important;
  min-height: 100vh;
}

/* Для мобильных - УБИРАЕМ ВСЕ ОТСТУПЫ */
@media (max-width: 767px) {
  body {
    -webkit-text-size-adjust: 100%;
    -webkit-tap-highlight-color: transparent;
  }

  .container,
  [class*="container"] {
    max-width: 100% !important;
    padding-left: 0 !important;
    padding-right: 0 !important;
    margin: 0 !important;
  }
}

/* =========================================================
   CALM TRANSITION — мягкая, премиальная, без резких эффектов
   ========================================================= */

/* ВХОД/ВЫХОД — плавность */
.calm-enter-active,
.calm-leave-active {
  transition: opacity 0.30s ease, transform 0.30s ease;
}

/* ВХОД — очень лёгкое смещение */
.calm-enter-from {
  opacity: 0;
  transform: translateY(6px);
}

.calm-enter-to {
  opacity: 1;
  transform: translateY(0);
}

/* ВЫХОД — лёгкое плавное исчезновение */
.calm-leave-from {
  opacity: 1;
  transform: translateY(0);
}

.calm-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>