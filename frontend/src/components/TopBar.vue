<template>
  <div class="w-full flex items-center justify-between py-3 px-4 bg-white">

    <!-- КНОПКА НАЗАД -->
    <button
        v-if="showBack"
        @click="$router.back()"
        class="text-gray-700 text-2xl active:scale-90 transition"
    >
      ←
    </button>

    <!-- ПОИСК (нажимаем → переходим на SearchScreen) -->
    <div
        @click="$router.push('/search')"
        class="flex items-center bg-gray-100 rounded-full px-4 py-2 w-full mx-3 active:scale-95 transition cursor-pointer"
    >
      <span class="text-gray-400 mr-3">🔍</span>
      <span class="text-gray-500 flex-1">Поиск приложений</span>
    </div>

    <!-- ПРОФИЛЬ -->
    <button
        @click="$router.push('/profile')"
        class="text-gray-700 text-2xl active:scale-90 transition"
    >
      👤
    </button>

  </div>
</template>

<script>
export default {
  props: {
    showBack: { type: Boolean, default: false }
  }
}
</script>