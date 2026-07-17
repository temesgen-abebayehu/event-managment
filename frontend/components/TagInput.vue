<template>
  <div>
    <label class="block text-sm font-medium text-gray-700 mb-2">
      Tags (Optional) <span class="text-gray-500 text-xs">Press Enter to add</span>
    </label>
    
    <!-- Tag Display -->
    <div v-if="selectedTags.length > 0" class="flex flex-wrap gap-2 mb-2">
      <span
        v-for="(tag, index) in selectedTags"
        :key="index"
        class="inline-flex items-center gap-1 bg-blue-100 text-blue-800 px-3 py-1 rounded-full text-sm"
      >
        {{ tag }}
        <button
          type="button"
          @click="removeTag(index)"
          class="hover:text-blue-900"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </span>
    </div>
    
    <!-- Tag Input -->
    <input
      v-model="newTag"
      type="text"
      @keydown.enter.prevent="addTag"
      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary"
      placeholder="Type a tag and press Enter"
    />
    <p class="text-xs text-gray-500 mt-1">
      Examples: concert, workshop, outdoor, free-food, networking
    </p>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  modelValue: string[]
}>()

const emit = defineEmits<{
  'update:modelValue': [tags: string[]]
}>()

const selectedTags = ref<string[]>([...props.modelValue])
const newTag = ref('')

const addTag = () => {
  const tag = newTag.value.trim().toLowerCase()
  
  if (tag && !selectedTags.value.includes(tag)) {
    selectedTags.value.push(tag)
    emit('update:modelValue', selectedTags.value)
    newTag.value = ''
  }
}

const removeTag = (index: number) => {
  selectedTags.value.splice(index, 1)
  emit('update:modelValue', selectedTags.value)
}

watch(() => props.modelValue, (newVal) => {
  selectedTags.value = [...newVal]
})
</script>
