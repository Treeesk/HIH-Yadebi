<template>
  <div class="flex items-center justify-between py-3">

    <!-- LEFT PART -->
    <div class="flex items-center gap-4">

      <!-- ICON -->
      <img
          :src="iconUrl"
          class="w-14 h-14 rounded-2xl object-cover shadow-sm"
      />

      <!-- TEXT -->
      <div class="flex flex-col">

        <div class="flex items-center gap-2">
          <h3 class="text-base font-semibold text-gray-900 leading-tight">
            {{ app.app_name }}
          </h3>

          <!-- BADGE -->
          <span
              v-if="app.badge"
              class="px-2 py-0.5 text-xs rounded-md font-medium text-white"
              :class="badgeClass(app.badge)"
          >
            {{ app.badge }}
          </span>
        </div>

        <p class="text-sm text-gray-600">{{ app.app_categorie }}</p>

        <div class="flex items-center gap-1 mt-1">
          <span class="text-yellow-500 text-sm">★</span>
          <span class="text-sm text-gray-700">{{ app.rating }}</span>
        </div>

      </div>
    </div>

    <!-- DOWNLOAD BUTTON -->
    <button
        class="px-4 py-2 text-sm rounded-xl font-semibold text-white bg-blue-600 shadow hover:bg-blue-700 active:scale-[0.97]"
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
      return `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${encodeURIComponent(this.app.app_name)}`;
    }
  },

  methods: {
    badgeClass(type) {
      return {
        "bg-green-600": type === "ТОП",
        "bg-purple-600": type === "ХИТ",
        "bg-blue-600": type === "НОВОЕ"
      }
    }
  }
};
</script>
