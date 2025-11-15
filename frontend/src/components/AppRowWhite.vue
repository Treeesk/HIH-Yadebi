<template>
  <div class="flex items-center py-3 w-full">

    <!-- LEFT: FIXED ICON -->
    <div
        class="w-14 h-14 rounded-2xl bg-gray-100 flex items-center justify-center flex-shrink-0"
    >
      <span class="text-3xl">
        {{ emoji }}
      </span>
    </div>

    <!-- MIDDLE -->
    <div class="flex flex-col flex-grow px-3 overflow-hidden">
      <h3 class="text-base font-semibold text-gray-900 truncate">
        {{ app.app_name }}
      </h3>

      <p class="text-sm text-gray-500 truncate">
        {{ app.app_category }}
      </p>
    </div>

    <!-- RIGHT BUTTON -->
    <div class="flex-shrink-0">
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
    emoji() {
      // можно заменить на любую логику
      return this.app.emoji || "🤖"
    }
  }
}
</script>