<template>
  <div
      class="w-full rounded-3xl overflow-hidden mt-6 shadow-md bg-gradient-to-br from-indigo-500 to-blue-600"
  >
    <div class="flex flex-col justify-center h-36 px-6 py-4">

      <!-- Заголовок -->
      <p class="text-white text-2xl font-bold leading-tight">
        {{ text }}
      </p>

      <!-- Подзаголовок (опционально) -->
      <p v-if="subtitle" class="text-white/80 text-sm mt-2">
        {{ subtitle }}
      </p>

      <!-- Кнопка (если нужна) -->
      <button
          v-if="button"
          class="mt-4 w-min whitespace-nowrap bg-white text-blue-700 px-4 py-1.5 rounded-xl font-semibold text-sm active:scale-95 transition"
      >
        {{ button }}
      </button>

    </div>
  </div>
</template>

<script>
export default {
  props: {
    text: { type: String, required: true },
    subtitle: { type: String, default: "" },
    button: { type: String, default: "" }
  }
}
</script>