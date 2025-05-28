import { ref, watch, type Ref } from 'vue'

export function useLocalStorage<T>(key: string, defaultValue: T): [Ref<T>, (newValue: T) => void, () => void] {
  // Try to get initial value from localStorage
  const storedValue = localStorage.getItem(key)
  let initialValue = defaultValue

  if (storedValue !== null) {
    try {
      initialValue = JSON.parse(storedValue)
    } catch {
      // If parsing fails, use default value
      initialValue = defaultValue
    }
  }

  const value = ref<T>(initialValue) as Ref<T>

  // Watch for changes and update localStorage
  watch(value, (newValue) => {
    localStorage.setItem(key, JSON.stringify(newValue))
  }, { deep: true })

  const setValue = (newValue: T) => {
    value.value = newValue
  }

  const removeValue = () => {
    localStorage.removeItem(key)
    value.value = defaultValue
  }

  return [value, setValue, removeValue]
}
