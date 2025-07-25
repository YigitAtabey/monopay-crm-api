// frontend/.eslintrc.js
module.exports = {
    root: true,
    env: {
      node: true,
      browser: true
    },
    extends: [
      'plugin:vue/vue3-essential',
      'eslint:recommended'
    ],
    rules: {
      // multi-word component kuralını kapat
      'vue/multi-word-component-names': 'off'
    }
  }
  