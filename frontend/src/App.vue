<template>
  <router-view />
</template>

<style>
/* Сбрасываем всё, что может ломать layout */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* Корневой html — фиксируем горизонтальный скролл */
html {
  width: 100% !important;
  height: 100% !important;
  overflow-x: hidden !important;  /* <-- главное */
  overflow-y: auto !important;
  background: #f9fafb !important;
}

/* Body — только вертикальный скролл, никакого 100vw */
body {
  margin: 0 !important;
  padding: 0 !important;

  width: 100% !important;      /* без 100vw */
  max-width: 100% !important;  /* запрещаем переполнение */

  height: 100% !important;
  overflow-x: hidden !important;
  overflow-y: auto !important;

  background: #f9fafb !important;

  /* скрываем скроллбар */
  scrollbar-width: none !important;      /* Firefox */
  -ms-overflow-style: none !important;   /* IE/Edge */
}

body::-webkit-scrollbar {
  display: none !important; /* Chrome/Safari/Opera */
}

/* #app — тоже ТОЛЬКО % */
#app {
  margin: 0 !important;
  padding: 0 !important;

  width: 100% !important;
  max-width: 100% !important;

  min-height: 100vh !important;
  overflow-x: hidden !important;

  background: #f9fafb !important;
}

/* Мобильная адаптация */
@media (max-width: 767px) {
  body {
    width: 100% !important;
    height: 100% !important;
    -webkit-text-size-adjust: 100%;
    -webkit-tap-highlight-color: transparent;
  }
}
</style>