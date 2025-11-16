<template>
  <div
      class="w-full bg-gray-100 rounded-full px-5 py-3 flex items-center gap-3 active:opacity-70 transition shadow-sm"
      @click="$router.push('/search')"
  >
    <span class="text-gray-500 text-xl">🔍</span>
    <span class="text-gray-500 text-base">Поиск приложений</span>
  </div>
</template>

<script>
export default {
  name: "SearchBar"
}
</script>

<style scoped>
/* Дополнительная мягкая тень */
</style>