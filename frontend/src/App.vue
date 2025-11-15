<!--&lt;!&ndash; src/App.vue &ndash;&gt;-->
<!--<template>-->
<!--  <div id="app">-->
<!--    <router-view/>-->
<!--  </div>-->
<!--</template>-->

<!--<script>-->
<!--export default {-->
<!--  name: 'App',-->
<!--  mounted() {-->
<!--    this.checkAndroidWebView();-->
<!--  },-->
<!--  methods: {-->
<!--    checkAndroidWebView() {-->
<!--      // Проверяем, открыто ли в WebView Android-->
<!--      const isAndroidWebView = navigator.userAgent.includes('Android') &&-->
<!--          !navigator.userAgent.includes('Chrome');-->

<!--      // Показываем OpenBoard только в WebView Android-->
<!--      if (isAndroidWebView && this.$route.name !== 'OpenBoard') {-->
<!--        this.$router.push('/openboard');-->
<!--      }-->
<!--    }-->
<!--  },-->
<!--  watch: {-->
<!--    '$route'(to) {-->
<!--      // При смене маршрута также проверяем-->
<!--      this.checkAndroidWebView();-->
<!--    }-->
<!--  }-->
<!--}-->
<!--</script>-->
<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  // ВРЕМЕННО: убрали всю логику проверки
  // mounted() {
  //   this.checkAndroidWebView();
  // },
  // methods: {
  //   checkAndroidWebView() {
  //     // логика проверки
  //   }
  // },
  // watch: {
  //   '$route'(to) {
  //     this.checkAndroidWebView();
  //   }
  // }
}
</script>