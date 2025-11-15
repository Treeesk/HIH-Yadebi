<template>
  <div class="flex items-center justify-between py-3">

    <!-- LEFT -->
    <div class="flex items-center gap-4">

      <!-- ICON -->
      <img
          :src="iconUrl"
          class="w-14 h-14 rounded-2xl object-cover shadow-sm"
      />

      <!-- TEXT -->
      <div class="flex flex-col leading-tight">
        <h3 class="text-lg font-semibold text-gray-900">
          {{ app.app_name }}
        </h3>

        <!-- GENRE instead of stars -->
        <p class="text-sm text-gray-600 mt-1">
          {{ app.app_category }}
        </p>
      </div>

    </div>

    <!-- DOWNLOAD BUTTON -->
    <button
        class="px-4 py-2 bg-blue-600 text-white text-sm font-semibold rounded-xl shadow-sm
             hover:bg-blue-700 active:scale-95 transition"
    >
      Скачать
    </button>

  </div>
</template>

<script>
export default {
  props: { app: Object },

  computed: {
    iconUrl() {
      return `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${encodeURIComponent(this.app.app_name)}`
    }
  }
}
</script>
