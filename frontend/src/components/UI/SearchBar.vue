<template>
  <div class="mb-6">
    <div class="flex items-center bg-gray-100 rounded-xl px-4 py-2 shadow-inner">
      <input
          :value="modelValue"
          @input="$emit('update:modelValue', $event.target.value)"
          type="text"
          class="bg-transparent flex-1 text-gray-900 text-base outline-none"
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
