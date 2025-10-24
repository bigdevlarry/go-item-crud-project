<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
  <div class="text-center">
    <h1 class="text-4xl font-bold text-green-600 mb-4">{{ msg }}</h1>
    <p class="text-lg text-gray-600">
      You've successfully created a project with
      <a href="https://vuejs.org/" target="_blank" rel="noopener" class="text-blue-600 hover:text-blue-800 underline">Vue 3</a>. What's next?
    </p>
  </div>
</template>

<style scoped>
</style>
