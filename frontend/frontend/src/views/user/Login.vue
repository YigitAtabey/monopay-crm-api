<!-- src/views/user/Login.vue -->
<template>
  <div class="login-page">
    <h1>Giriş Yap</h1>
    <form @submit.prevent="handleLogin" class="login-form">
      <div class="field">
        <label for="email">E-posta</label>
        <input
          id="email"
          type="email"
          v-model="email"
          required
          autocomplete="username"
        />
      </div>
      <div class="field">
        <label for="password">Şifre</label>
        <input
          id="password"
          type="password"
          v-model="password"
          required
          autocomplete="current-password"
        />
      </div>
      <button type="submit" :disabled="loading">
        {{ loading ? 'Giriş Yapılıyor…' : 'Giriş Yap' }}
      </button>
    </form>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router   = useRouter()
const auth     = useAuthStore()

const email    = ref('')
const password = ref('')
const loading  = ref(false)
const error    = ref('')

async function handleLogin() {
  // Önceki hatayı temizle, loading’i aktif et
  error.value   = ''
  loading.value = true

  try {
    // 1) API’den token + profile çek
    await auth.login({ email: email.value, password: password.value })

    // 2) /panel’e git ve tam sayfa yenile
    await router.push('/panel')
    window.location.reload()
  } catch (err) {
    // Hata mesajını al
    error.value =
      err.response?.data?.error ||
      err.response?.data?.message ||
      'Giriş başarısız. Bilgilerini kontrol et.'
  } finally {
    loading.value = false
  }
}

</script>

<style scoped>
.login-page {
  max-width: 400px;
  margin: 3rem auto;
  padding: 2rem;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.08);
}
.login-page h1 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #2c3e50;
}
.login-form .field {
  margin-bottom: 1.25rem;
}
.login-form label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #333;
}
.login-form input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}
.login-form button {
  width: 100%;
  padding: 0.75rem;
  background: #42b983;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, transform 0.1s;
}
.login-form button:disabled {
  background: #99d4af;
  cursor: not-allowed;
}
.login-form button:not(:disabled):hover {
  background: #369f6e;
  transform: translateY(-1px);
}
.error {
  margin-top: 1rem;
  text-align: center;
  color: #e74c3c;
}
</style>
