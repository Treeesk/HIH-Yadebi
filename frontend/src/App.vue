<template>
  <router-view />
</template>

<style>
/* ПОЛНЫЙ ФИКС ДЛЯ ВСЕГО ЭКРАНА */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* УБИРАЕМ СКРОЛЛБАРЫ НАВСЕГДА */
html {
  overflow: hidden !important; /* Убираем скролл у html */
  width: 100% !important;
  height: 100% !important;
  background: #f9fafb !important;
}

body {
  margin: 0 !important;
  padding: 0 !important;
  width: 100vw !important; /* 100% ширины viewport */
  height: 100vh !important; /* 100% высоты viewport */
  overflow: auto !important; /* Разрешаем скролл только у body */
  background: #f9fafb !important;

  /* Убираем скроллбары */
  scrollbar-width: none !important; /* Firefox */
  -ms-overflow-style: none !important; /* IE 10+ */
}

body::-webkit-scrollbar {
  display: none !important; /* Chrome, Safari, Edge */
}

#app {
  margin: 0 !important;
  padding: 0 !important;
  width: 100vw !important;
  min-height: 100vh !important;
  background: #f9fafb !important;
}

/* Для мобильных */
@media (max-width: 767px) {
  body {
    -webkit-text-size-adjust: 100%;
    -webkit-tap-highlight-color: transparent;
    width: 100% !important;
    height: 100% !important;
  }
}
</style>