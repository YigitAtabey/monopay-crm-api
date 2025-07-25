// .eslintrc.js — proje kökünde
module.exports = {
    root: true,                // bu dosyanın kök config olduğunu belirtir
    env: {
      node: true,              // Node.js global’lerini aktif eder
      browser: true            // tarayıcı global’lerini aktif eder
    },
    extends: [
      'plugin:vue/vue3-essential', // Vue 3 için temel kurallar
      'eslint:recommended'          // ESLint’in önerilen kuralları
    ],
    rules: {
      // Sadece bu kuralı kapatıyoruz:
      'vue/multi-word-component-names': 'off'
    }
  }
  