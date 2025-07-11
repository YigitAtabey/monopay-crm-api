<template>
    <div class="register-page">
      <h1>Kayıt Ol</h1>
      <form @submit.prevent="handleRegister">
        <div class="field">
          <label for="name">İsim:</label>
          <input id="name" v-model="name" type="text" required />
        </div>
        <div class="field">
          <label for="email">E-posta:</label>
          <input id="email" v-model="email" type="email" required />
        </div>
        <div class="field">
          <label for="password">Şifre:</label>
          <input
            id="password"
            v-model="password"
            type="password"
            minlength="6"
            required
          />
        </div>
        <button
          type="submit"
          :disabled="loading || !name || !email || password.length < 6"
        >
          <span v-if="loading">Bekleyin…</span>
          <span v-else>Kayıt Ol</span>
        </button>
      </form>
      <p v-if="error" class="error">{{ error }}</p>
    </div>
  </template>
  
  <script setup>
  import { ref }       from 'vue'
  import axios          from 'axios'
  import { useRouter }  from 'vue-router'
  
  const router   = useRouter()
  const name     = ref('')
  const email    = ref('')
  const password = ref('')
  const error    = ref('')
  const loading  = ref(false)
  
  const handleRegister = async () => {
    error.value   = ''
    loading.value = true
  
    try {
      // proxy sayesinde /register → backend’e gider
      await axios.post('/register', {
        name:     name.value,
        email:    email.value,
        password: password.value
      })
      // Kayıt başarılı → login sayfasına yönlendir
      router.push('/login')
    } catch (err) {
      console.error('Register error detay:', err.response || err)
      error.value =
        err.response?.data?.message ||
        'Kayıt sırasında bir hata oluştu.'
    } finally {
      loading.value = false
    }
  }
  </script>
  
  <style scoped>
  .register-page {
    max-width: 400px;
    margin: 2rem auto;
    padding: 1.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  .field {
    margin-bottom: 1rem;
  }
  label {
    display: block;
    margin-bottom: 0.25rem;
  }
  input {
    width: 100%;
    padding: 0.5rem;
    box-sizing: border-box;
  }
  button {
    width: 100%;
    padding: 0.75rem;
  }
  .error {
    color: red;
    margin-top: 1rem;
    text-align: center;
  }
  </style>
  