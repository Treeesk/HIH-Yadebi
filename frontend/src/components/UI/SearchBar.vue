<template>
  <div class="mb-6 w-full">
    <div class="w-full flex items-center bg-gray-100 rounded-xl px-4 py-3 shadow-inner">
      <input
          :value="modelValue"
          @input="$emit('update:modelValue', $event.target.value)"
          type="text"
          class="flex-1 bg-transparent text-gray-900 text-base outline-none"
          placeholder="Поиск"
      />
      <span class="text-gray-500 text-xl">🔍</span>
    </div>
  </div>
</template>

<script>
export default {
  props: ["modelValue"]
};
</script>