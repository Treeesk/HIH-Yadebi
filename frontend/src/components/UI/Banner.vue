<template>
  <div
      class="w-full h-40 mt-6 rounded-3xl
           bg-gradient-to-br from-indigo-500 to-blue-500
           shadow-lg flex items-center justify-center"
  >
    <p class="text-white text-xl font-semibold px-4 text-center">
      {{ text }}
    </p>
  </div>
</template>

<script>
export default { props: ["text"] }
</script>
