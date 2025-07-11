<!-- src/views/HomeView.vue -->
<template>
  <div class="home">
    <div class="hero-card">
      <h1>Monopay CRM’e Hoşgeldin!</h1>
      <p>
        Kullanıcı ve görev yönetimini kolaylaştıran arayüzümüze göz atmak için aşağıdan bir butona tıkla.
      </p>
      <div class="hero-buttons">
        <router-link
          v-if="!isLoggedIn"
          to="/login"
          class="btn btn-primary"
        >
          Giriş Yap
        </router-link>
        <router-link
          v-if="!isLoggedIn"
          to="/register"
          class="btn btn-secondary"
        >
          Kayıt Ol
        </router-link>
        <router-link
          v-else
          to="/panel"
          class="btn btn-primary"
        >
          Panel
        </router-link>
        <router-link
          v-if="isAdmin"
          to="/admin"
          class="btn btn-danger"
        >
          Admin Paneli
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useAuthStore } from '@/store/auth'

const auth       = useAuthStore()
const isLoggedIn = computed(() => auth.isLoggedIn)
const isAdmin    = computed(() => auth.user?.role === 'admin')
</script>

<style scoped>
.home {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 150px);
}

.hero-card {
  background: #ffffff;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.08);
  max-width: 600px;
  text-align: center;
}

.hero-card h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: #2c3e50;
}

.hero-card p {
  font-size: 1.125rem;
  margin-bottom: 2rem;
  color: #555;
}

.hero-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  justify-content: center;
}

.btn {
  padding: 0.75rem 1.5rem;
  border-radius: 5px;
  text-decoration: none;
  font-weight: 600;
  transition: transform 0.2s, box-shadow 0.2s;
}

.btn-primary {
  background: linear-gradient(90deg, #42b983 0%, #2c3e50 100%);
  color: #fff;
}

.btn-secondary {
  background: #ecf0f1;
  color: #2c3e50;
}

.btn-danger {
  background: #e74c3c;
  color: #fff;
}

.btn:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
</style>
