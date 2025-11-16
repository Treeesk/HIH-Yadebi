<template>
  <div class="w-full rounded-3xl overflow-hidden shadow-lg bg-white">

    <!-- BIG IMAGE -->
    <img
        :src="image"
        class="w-full h-44 object-cover"
    />

    <!-- TEXT AREA -->
    <div class="p-4">
      <h3 class="text-xl font-semibold text-gray-900">{{ title }}</h3>
      <p class="text-gray-600 text-sm mt-1">{{ subtitle }}</p>

      <button
          class="mt-3 px-4 py-2 bg-blue-600 text-white rounded-xl shadow
               active:scale-95 transition font-medium"
      >
        Открыть
      </button>
    </div>

  </div>
</template>

<script>
export default {
  props: ["image", "title", "subtitle"]
}
</script>
