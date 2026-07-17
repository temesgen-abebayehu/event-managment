<template>
  <div ref="container" class="relative overflow-hidden">
    <transition name="fade">
      <div
        v-if="!loaded"
        class="absolute inset-0 bg-gray-200 animate-pulse flex items-center justify-center"
      >
        <svg
          v-if="showPlaceholder"
          class="w-12 h-12 text-gray-400"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
          />
        </svg>
      </div>
    </transition>
    <transition name="fade">
      <img
        v-if="shouldLoad"
        :src="src"
        :alt="alt"
        :class="className"
        @load="onLoad"
        @error="onError"
      />
    </transition>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  src: string
  alt: string
  className?: string
  showPlaceholder?: boolean
}>()

const container = ref<HTMLElement | null>(null)
const shouldLoad = ref(false)
const loaded = ref(false)

const onLoad = () => {
  loaded.value = true
}

const onError = () => {
  loaded.value = true
}

onMounted(() => {
  if (!container.value) return

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          shouldLoad.value = true
          observer.disconnect()
        }
      })
    },
    { rootMargin: '50px' }
  )

  observer.observe(container.value)

  onUnmounted(() => {
    observer.disconnect()
  })
})
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
