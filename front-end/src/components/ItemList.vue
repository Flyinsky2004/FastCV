<script setup>
defineProps({
  items: {
    type: Array,
    required: true
  },
  onEdit: {
    type: Function,
    required: true
  },
  onDelete: {
    type: Function,
    required: true
  }
});
</script>

<template>
  <div class="item-list">
    <div v-for="(item, index) in items" 
         :key="index" 
         class="item-card">
      <slot name="item" :item="item"></slot>
      <div class="item-actions">
        <button @click="onEdit(index)" class="action-btn edit">
          <span class="text-sm">✎</span>
        </button>
        <button @click="onDelete(index)" class="action-btn delete">
          <span class="text-xl">×</span>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.item-list {
  @apply space-y-3;
}

.item-card {
  @apply bg-white rounded-lg p-4 shadow-sm relative hover:shadow-md transition-shadow duration-200;
}

.item-actions {
  @apply absolute top-2 right-2 flex space-x-2;
}

.action-btn {
  @apply p-1 rounded hover:bg-gray-100 transition-colors duration-200;
}

.action-btn.edit {
  @apply text-blue-500 hover:text-blue-700;
}

.action-btn.delete {
  @apply text-red-500 hover:text-red-700;
}
</style> 