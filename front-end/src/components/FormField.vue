<script setup>
defineProps({
  label: {
    type: String,
    required: true
  },
  labelWidth: {
    type: String,
    default: '120px'
  }
});
</script>

<template>
  <div class="form-field">
    <label class="form-label">{{ label }}</label>
    <div class="form-content">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.form-field {
  @apply flex items-center mb-4;
}

.form-label {
  @apply text-gray-700 font-medium min-w-[var(--label-width)] mr-4;
  --label-width: v-bind(labelWidth);
}

.form-content {
  @apply flex-1;
}
</style> 