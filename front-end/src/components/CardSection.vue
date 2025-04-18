<script setup>
defineProps({
  title: {
    type: String,
    required: true
  }
});
</script>

<template>
  <div class="card-section">
    <h3 class="card-title">{{ title }}</h3>
    <div class="card-content">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.card-section {
  @apply bg-white rounded-lg shadow-sm p-6 mb-6;
}

.card-title {
  @apply text-xl font-bold text-gray-800 mb-4 pb-3 border-b border-gray-200;
}

.card-content {
  @apply space-y-4;
}
</style> 