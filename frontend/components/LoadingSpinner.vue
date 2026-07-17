<template>
  <div class="flex items-center justify-center" :class="containerClass">
    <div
      :class="[
        'animate-spin rounded-full border-t-2 border-b-2',
        sizeClasses[size],
        colorClasses[color],
      ]"
    ></div>
  </div>
</template>

<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    size?: 'sm' | 'md' | 'lg'
    color?: 'primary' | 'white' | 'gray'
    containerClass?: string
  }>(),
  {
    size: 'md',
    color: 'primary',
    containerClass: 'py-12',
  }
)

const sizeClasses = {
  sm: 'h-6 w-6',
  md: 'h-12 w-12',
  lg: 'h-16 w-16',
}

const colorClasses = {
  primary: 'border-primary',
  white: 'border-white',
  gray: 'border-gray-400',
}
</script>
