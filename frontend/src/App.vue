<template>
  <router-view />
</template>

<style>
/* ГЛОБАЛЬНЫЕ СТИЛИ ДЛЯ УДАЛЕНИЯ ОТСТУПОВ */
html, body {
  margin: 0 !important;
  padding: 0 !important;
  width: 100% !important;
  overflow-x: hidden !important;
}

#app {
  margin: 0 !important;
  padding: 0 !important;
  width: 100% !important;
}

/* Для мобильных - УБИРАЕМ ВСЕ ОТСТУПЫ */
@media (max-width: 767px) {
  body {
    -webkit-text-size-adjust: 100%;
    -webkit-tap-highlight-color: transparent;
    margin: 0 !important;
    padding: 0 !important;
  }

  #app {
    margin: 0 !important;
    padding: 0 !important;
    width: 100% !important;
  }

  /* Принудительно убираем отступы у всех контейнеров */
  .container, [class*="container"] {
    max-width: 100% !important;
    padding-left: 0 !important;
    padding-right: 0 !important;
    margin-left: 0 !important;
    margin-right: 0 !important;
  }
}
</style>