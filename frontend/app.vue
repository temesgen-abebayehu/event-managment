<template>
  <div>
    <NuxtPage />
    <ToastContainer />
  </div>
</template>

<script setup lang="ts">
// Global page transitions
useHead({
  titleTemplate: (titleChunk) => {
    return titleChunk ? `${titleChunk} - EventHub Ethiopia` : 'EventHub Ethiopia'
  },
})
</script>

<style>
/* Global page transitions */
.page-enter-active,
.page-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Smooth scroll behavior */
html {
  scroll-behavior: smooth;
}
</style>
