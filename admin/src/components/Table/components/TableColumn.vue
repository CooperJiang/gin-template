<template>
  <div style="display: none;"></div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'GlobalTableColumn',
  props: {
    prop: String,
    label: String,
    width: [Number, String],
    minWidth: Number,
    maxWidth: Number,
    align: {
      type: String as () => 'left' | 'center' | 'right',
      default: 'left'
    },
    sortable: {
      type: Boolean,
      default: false
    },
    filterable: {
      type: Boolean,
      default: false
    },
    fixed: String as () => 'left' | 'right',
    ellipsis: {
      type: Boolean,
      default: false
    },
    type: String as () => 'selection' | 'index' | 'expand',
    showOverflowTooltip: {
      type: Boolean,
      default: false
    }
  }
})
</script>
