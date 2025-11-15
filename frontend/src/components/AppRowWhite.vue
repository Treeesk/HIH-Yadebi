<template>
  <div class="flex items-center py-3 w-full">

    <!-- LEFT -->
    <div class="flex items-center gap-3 flex-shrink">
      <img
          :src="iconUrl"
          class="w-10 h-10 rounded-2xl object-cover flex-shrink-0"
      />
    </div>

    <!-- MIDDLE -->
    <div class="flex flex-col flex-grow px-2 overflow-hidden">
      <h3 class="text-base font-semibold text-gray-900 truncate">
        {{ app.app_name }}
      </h3>

      <p class="text-sm text-gray-500 truncate">
        {{ app.app_category }}
      </p>
    </div>

    <!-- RIGHT (BUTTON) -->
    <div class="flex-shrink-0 pl-2">
      <button
          class="px-4 py-1.5 rounded-full bg-blue-500 text-white text-sm font-semibold
               whitespace-nowrap hover:bg-blue-600 active:scale-95 transition"
      >
        Загрузить
      </button>
    </div>

  </div>
</template>

<script>
export default {
  props: { app: Object },

  computed: {
    iconUrl() {
      return this.app.icon_url ||
          `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${encodeURIComponent(this.app.app_name)}`
    }
  }
}
</script>