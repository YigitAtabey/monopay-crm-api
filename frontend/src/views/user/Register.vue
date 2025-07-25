<template>
  <div class="register-page">
    <h1>Kayıt Ol</h1>
    <form @submit.prevent="handleRegister" class="register-form">
      <div class="field">
        <label for="name">İsim</label>
        <input
          id="name"
          v-model="name"
          type="text"
          required
          autocomplete="name"
        />
      </div>
      <div class="field">
        <label for="email">E-posta</label>
        <input
          id="email"
          v-model="email"
          type="email"
          required
          autocomplete="email"
        />
      </div>
      <div class="field">
        <label for="password">Şifre</label>
        <input
          id="password"
          v-model="password"
          type="password"
          minlength="6"
          required
          autocomplete="new-password"
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
import axios         from 'axios'
import { useRouter } from 'vue-router'

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
      name:     name.value.trim(),
      email:    email.value.trim(),
      password: password.value
    })
    // Kayıt başarılı → login sayfasına yönlendir
    await router.push('/login')
  } catch (err) {
    // Uygun hata mesajı göster
    error.value =
      err.response?.data?.error ||
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
  margin: 3rem auto;
  padding: 2rem;
  background: #23272f;
  border-radius: 12px;
  box-shadow: 0 4px 18px rgba(60,60,90,0.13);
  color: #f7f7f7;
}
.register-page h1 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #42b983;
  font-weight: 800;
  letter-spacing: 1px;
  text-shadow: 0 2px 10px rgba(66,185,131,0.08);
}
.register-form .field {
  margin-bottom: 1.15rem;
}
.register-form label {
  display: block;
  margin-bottom: 0.4rem;
  font-weight: 500;
  color: #d3efdf;
}
.register-form input {
  width: 100%;
  padding: 0.5rem;
  border: 1.5px solid #42b983;
  border-radius: 6px;
  background: #2c313c;
  color: #fafafa;
  box-sizing: border-box;
  transition: border 0.18s;
}
.register-form input:focus {
  border: 1.5px solid #43e39b;
  outline: none;
  background: #23272f;
  color: #fff;
}
button {
  width: 100%;
  padding: 0.8rem;
  background: linear-gradient(90deg, #42b983 60%, #2c3e50 100%);
  color: #fff;
  border: none;
  border-radius: 6px;
  font-weight: 700;
  font-size: 1.08em;
  cursor: pointer;
  transition: background 0.2s, transform 0.1s;
  box-shadow: 0 1px 4px rgba(66,185,131,0.13);
}
button:disabled {
  background: #3a3a3a;
  cursor: not-allowed;
  opacity: 0.7;
}
button:not(:disabled):hover {
  background: #369f6e;
  transform: translateY(-1px) scale(1.01);
}
.error {
  color: #ec4646;
  margin-top: 1rem;
  text-align: center;
  font-weight: bold;
  letter-spacing: 0.5px;
}
</style>