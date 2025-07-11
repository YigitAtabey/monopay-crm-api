<!-- src/components/NavBar.vue -->
<template>
  <nav class="navbar">
    <!-- Marka Logo / Başlık -->
    <div class="navbar-brand">
      MonopayCRM
    </div>

    <!-- Menü Öğeleri -->
    <ul class="navbar-menu">
      <li><router-link to="/">Ana Sayfa</router-link></li>
      <li v-if="!isLoggedIn"><router-link to="/login">Giriş</router-link></li>
      <li v-if="!isLoggedIn"><router-link to="/register">Kayıt</router-link></li>

      <li v-if="isLoggedIn"><router-link to="/panel">Panel</router-link></li>
      <li v-if="isLoggedIn"><router-link to="/tasks">Görev Listesi</router-link></li>
      <li v-if="isLoggedIn"><router-link to="/tasks/manage">Görev Yönetimi</router-link></li>
      <li v-if="isLoggedIn && isAdmin"><router-link to="/admin">Admin Paneli</router-link></li>
      <li v-if="isLoggedIn"><router-link to="/logout">Çıkış</router-link></li>
    </ul>
  </nav>
</template>

<script setup>
import { computed } from 'vue'
import { useAuthStore } from '@/store/auth'

const auth       = useAuthStore()
const isLoggedIn = computed(() => auth.isLoggedIn)
const isAdmin    = computed(() => auth.user?.role === 'admin')
</script>

<style scoped>
.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.75rem 2rem;
  background: linear-gradient(90deg, #42b983 0%, #2c3e50 100%);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  margin-bottom: 1.5rem;
}

.navbar-brand {
  color: #ffffff;
  font-size: 1.5rem;
  font-weight: 700;
  letter-spacing: 1px;
  text-shadow: 1px 1px 2px rgba(0,0,0,0.3);
}

.navbar-menu {
  display: flex;
  gap: 1rem;
  list-style: none;
  margin: 0;
  padding: 0;
}

.navbar-menu li {
  margin: 0;
}

.navbar-menu a {
  display: block;
  padding: 0.5rem 1rem;
  color: #ecf0f1;
  font-weight: 500;
  border-radius: 4px;
  transition: background-color 0.3s, transform 0.2s;
}

.navbar-menu a:hover {
  background-color: rgba(236, 240, 241, 0.2);
  transform: translateY(-2px);
}

.navbar-menu a.router-link-active {
  background-color: rgba(236, 240, 241, 0.3);
  box-shadow: 0 2px 6px rgba(0,0,0,0.1);
}
</style>
