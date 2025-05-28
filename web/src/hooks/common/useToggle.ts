import { ref, type Ref } from 'vue'

export function useToggle(initialValue = false): [Ref<boolean>, () => void] {
  const value = ref(initialValue)

  const toggle = () => {
    value.value = !value.value
  }

  return [value, toggle]
}
