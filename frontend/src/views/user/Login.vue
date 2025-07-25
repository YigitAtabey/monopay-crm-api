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
  error.value   = ''
  loading.value = true

  try {
    // Giriş isteği (token ve profil çek)
    const user = await auth.login({ email: email.value, password: password.value })

    // Blocked kullanıcının girişi frontend'de engellenir
    if (user && user.role === 'blocked') {
      await auth.logout()
      error.value = 'Hesabınız engellenmiş. Lütfen yöneticiyle iletişime geçin.'
      loading.value = false
      return
    }

    // Başarılı giriş: panele yönlendir
    await router.push('/panel')
    window.location.reload()
  } catch (err) {
    // Hata mesajı
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
  background: #23272f;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(60,60,90,0.13);
}
.login-page h1 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #42b983;
  letter-spacing: 1px;
  font-weight: 800;
  text-shadow: 0 2px 10px rgba(66,185,131,0.08);
}
.login-form .field {
  margin-bottom: 1.25rem;
}
.login-form label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #d3efdf;
}
.login-form input {
  width: 100%;
  padding: 0.5rem;
  border: 1.5px solid #42b983;
  border-radius: 6px;
  background: #2c313c;
  color: #fafafa;
  box-sizing: border-box;
  transition: border 0.18s;
}
.login-form input:focus {
  border: 1.5px solid #43e39b;
  outline: none;
  background: #23272f;
  color: #fff;
}
.login-form button {
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
.login-form button:disabled {
  background: #3a3a3a;
  cursor: not-allowed;
  opacity: 0.7;
}
.login-form button:not(:disabled):hover {
  background: #369f6e;
  transform: translateY(-1px) scale(1.01);
}
.error {
  margin-top: 1rem;
  text-align: center;
  color: #ec4646;
  font-weight: bold;
  letter-spacing: 0.5px;
}
</style>