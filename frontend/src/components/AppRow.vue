<template>
  <div class="flex items-center gap-4">

    <!-- ICON -->
    <img
        :src="iconUrl"
        class="w-14 h-14 rounded-2xl object-cover"
        alt=""
    />

    <!-- TEXT -->
    <div class="flex flex-col flex-1">
      <!-- 🔥 НАЗВАНИЕ -->
      <h3 class="text-base font-semibold text-gray-900 leading-tight">
        {{ app.app_name }}
      </h3>

      <!-- 🔥 КАТЕГОРИЯ -->
      <p class="text-sm text-gray-600">
        {{ app.app_categorie }}
      </p>

      <!-- 🔥 РЕЙТИНГ -->
      <div class="flex items-center gap-1 mt-1">
        <span class="text-yellow-500 text-sm">★</span>
        <span class="text-sm text-gray-700">{{ app.rating }}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  props: {
    app: { type: Object, required: true }
  },

  computed: {
    iconUrl() {
      return `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${encodeURIComponent(this.app.app_name)}`;
    }
  }
};
</script>
