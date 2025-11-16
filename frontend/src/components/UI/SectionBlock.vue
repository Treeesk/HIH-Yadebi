<template>
  <section class="mb-10">
    <h2 class="font-bold text-2xl text-gray-900 mb-3">
      {{ title }}
    </h2>
    <slot />
  </section>
</template>

<script>
export default { props: ["title"] }
</script>
