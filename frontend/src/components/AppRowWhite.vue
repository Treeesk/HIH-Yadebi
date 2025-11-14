<template>
  <div class="flex items-center justify-between py-3">
    <div class="flex items-center gap-4">
      <img
          :src="iconUrl"
          class="w-14 h-14 rounded-2xl object-cover shadow-sm"
      />
      <div>
        <h3 class="text-base font-semibold text-gray-900">
          {{ app.app_name }}
        </h3>
        <p class="text-sm text-gray-600">{{ app.app_categorie }}</p>
      </div>
    </div>

    <button
        class="px-4 py-2 text-sm rounded-xl font-semibold text-white bg-blue-600 shadow"
    >
      Скачать
    </button>
  </div>
</template>

<script>
export default {
  props: ['app'],
  computed: {
    iconUrl() {
      return `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${encodeURIComponent(this.app.app_name)}`
    }
  }
}
</script>
