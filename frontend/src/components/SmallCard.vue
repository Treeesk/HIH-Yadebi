<template>
  <div class="w-full bg-white rounded-2xl flex items-center gap-4 px-3 py-3">

    <img
        :src="iconUrl"
        class="w-12 h-12 rounded-xl object-cover"
    />

    <div class="flex flex-col leading-tight">
      <h3 class="text-base font-semibold text-gray-900">
        {{ app.app_name }}
      </h3>

      <p class="text-sm text-gray-500">
        {{ app.app_category }}
      </p>
    </div>

  </div>
</template>

<script>
export default {
  props: { app: Object },

  computed: {
    iconUrl() {
      // ЕДИНЫЙ СТИЛЬ ИКОНКИ (bottts-neutral)
      return `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${encodeURIComponent(
          this.app.app_name
      )}&radius=10&size=96`;
    }
  }
}
</script>

<style scoped>
/* clean */
</style>