import { globalIgnores } from 'eslint/config'
import { defineConfigWithVueTs, vueTsConfigs } from '@vue/eslint-config-typescript'
import pluginVue from 'eslint-plugin-vue'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

// To allow more languages other than `ts` in `.vue` files, uncomment the following lines:
// import { configureVueProject } from '@vue/eslint-config-typescript'
// configureVueProject({ scriptLangs: ['ts', 'tsx'] })
// More info at https://github.com/vuejs/eslint-config-typescript/#advanced-setup

export default defineConfigWithVueTs(
  {
    name: 'app/files-to-lint',
    files: ['**/*.{ts,mts,tsx,vue}'],
  },

  globalIgnores(['**/dist/**', '**/dist-ssr/**', '**/coverage/**']),

  pluginVue.configs['flat/essential'],
  vueTsConfigs.recommended,
  skipFormatting,

  // 自定义规则
  {
    name: 'app/custom-rules',
    rules: {
      // 允许页面组件使用单个单词命名（如 index.vue）
      'vue/multi-word-component-names': ['error', {
        ignores: ['index', 'Dashboard', 'Login', 'Register', 'Profile', 'Debug']
      }],

      // 完全关闭 any 类型检查（在某些情况下 any 类型是合理的）
      '@typescript-eslint/no-explicit-any': 'off',

      // 放宽未使用变量的限制，允许以下划线开头的变量
      '@typescript-eslint/no-unused-vars': ['error', {
        argsIgnorePattern: '^_',
        varsIgnorePattern: '^_',
        ignoreRestSiblings: true
      }]
    }
  }
)
