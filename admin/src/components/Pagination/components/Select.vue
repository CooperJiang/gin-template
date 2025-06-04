<template>
  <div class="relative" ref="selectRef">
    <button
      @click="toggleDropdown"
      :disabled="disabled"
      :class="[
        'pagination-select-trigger',
        {
          'pagination-select-trigger--disabled': disabled,
          'pagination-select-trigger--open': isOpen
        }
      ]"
    >
      <span>{{ selectedLabel }}</span>
      <GlobalIcon
        name="chevron-down"
        :size="iconSize"
        :class="[
          'transition-transform duration-200',
          { 'rotate-180': isOpen }
        ]"
      />
    </button>

    <Transition name="dropdown">
      <div
        v-show="isOpen"
        :class="[
          'pagination-select-dropdown',
          dropdownSizeClass
        ]"
      >
        <div
          v-for="option in options"
          :key="option"
          @click="selectOption(option)"
          :class="[
            'pagination-select-option',
            {
              'pagination-select-option--selected': option === modelValue
            }
          ]"
        >
          {{ option }}
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, onUnmounted } from 'vue'

defineOptions({
  name: 'PaginationSelect'
})

interface SelectProps {
  modelValue: number
  options: number[]
  disabled?: boolean
  size?: 'small' | 'medium' | 'large'
}

interface SelectEmits {
  'update:modelValue': [value: number]
}

const props = withDefaults(defineProps<SelectProps>(), {
  disabled: false,
  size: 'medium'
})

const emit = defineEmits<SelectEmits>()

const selectRef = ref<HTMLElement>()
const isOpen = ref(false)

const selectedLabel = computed(() => props.modelValue)

const iconSize = computed(() => {
  switch (props.size) {
    case 'small':
      return 'xs' as const
    case 'large':
      return 'sm' as const
    default:
      return 'xs' as const
  }
})

const dropdownSizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'text-xs'
    case 'large':
      return 'text-base'
    default:
      return 'text-sm'
  }
})

const toggleDropdown = () => {
  if (!props.disabled) {
    isOpen.value = !isOpen.value
  }
}

const selectOption = (option: number) => {
  emit('update:modelValue', option)
  isOpen.value = false
}

const handleClickOutside = (event: Event) => {
  if (selectRef.value && !selectRef.value.contains(event.target as Node)) {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.pagination-select-trigger {
  @apply flex items-center justify-between gap-2 px-3 py-1.5 min-w-16 border border-gray-300 rounded bg-white text-gray-700 hover:border-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors duration-200 cursor-pointer;
}

.pagination-select-trigger--disabled {
  @apply opacity-50 cursor-not-allowed hover:border-gray-300;
}

.pagination-select-trigger--open {
  @apply border-blue-500 ring-2 ring-blue-500;
}

.pagination-select-dropdown {
  @apply absolute top-full left-0 right-0 mt-1 bg-white border border-gray-300 rounded shadow-lg z-50 max-h-40 overflow-y-auto;
}

.pagination-select-option {
  @apply px-3 py-2 hover:bg-gray-50 cursor-pointer transition-colors duration-150;
}

.pagination-select-option--selected {
  @apply bg-blue-50 text-blue-600 font-medium;
}

.dropdown-enter-active,
.dropdown-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-4px) scale(0.95);
}
</style>
