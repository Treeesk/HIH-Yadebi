<template>
  <div class="w-full bg-white shadow-sm flex items-center justify-between px-4 py-3 sticky top-0 z-50">

    <!-- BACK -->
    <button @click="$router.back()" class="text-gray-700 text-2xl">
      ←
    </button>

    <!-- SEARCH -->
    <div class="flex items-center bg-gray-100 rounded-full px-4 py-2 w-7/12 shadow-sm">
      <input
          v-model="search"
          type="text"
          placeholder="Поиск"
          class="bg-transparent outline-none text-gray-800 flex-1"
      />
      <span class="text-gray-500 text-xl">🔍</span>
    </div>

    <!-- PROFILE -->
    <button class="text-gray-700 text-2xl">
      <span class="material-icons">person_outline</span>
    </button>

  </div>
</template>

<script>
export default {
  name: "TopBar",
  data() {
    return {
      search: ""
    }
  }
}
</script>

<style scoped>
.material-icons {
  font-size: 26px;
}
</style>