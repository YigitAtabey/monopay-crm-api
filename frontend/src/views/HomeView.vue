<template>
  <div class="home">
    <div :class="['hero-card', { 'admin-theme': isAdmin }]">
      <h1>
        <span v-if="isAdmin" class="admin-crown">👑</span>
        Monopay CRM’e Hoşgeldin!
      </h1>
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
          class="btn btn-admin"
        >
          <span class="admin-crown-btn">👑</span> Admin Paneli
        </router-link>
      </div>
      <div v-if="isAdmin" class="admin-welcome">
        <hr />
        <h2>Yönetici Modu Aktif!</h2>
        <p>
          Yönetimsel işlemler ve kullanıcı kontrolü için <b>Admin Paneli</b>’ni kullanabilirsin.<br>
          <span class="admin-hint">Ekstra yetkilerle sistemin tamamını yönetebilirsin.</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useAuthStore } from '@/store/auth'

const auth       = useAuthStore()
const isLoggedIn = computed(() => auth.isLoggedIn)
// isAdmin artık role stringine göre!
const isAdmin    = computed(() => auth.user?.role === 'admin')
</script>

<style scoped>
.home {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 150px);
  background: #23272f;
}

.hero-card {
  background: #fff;
  padding: 2.2rem 2rem 2rem 2rem;
  border-radius: 14px;
  box-shadow: 0 8px 32px 0 rgba(60,60,90,0.12), 0 2px 4px rgba(50,50,100,0.09);
  max-width: 600px;
  text-align: center;
  transition: background 0.25s, color 0.25s;
}

.hero-card.admin-theme {
  background: linear-gradient(120deg, #42b983 80%, #23272f 100%);
  color: #fff;
  box-shadow: 0 12px 40px 0 rgba(66,185,131,0.19);
  border: 2.5px solid #42b983;
  position: relative;
  overflow: hidden;
}

.hero-card h1 {
  font-size: 2.5rem;
  margin-bottom: 1.1rem;
  color: #2c3e50;
  font-weight: 800;
  letter-spacing: 1px;
  text-shadow: 0 2px 10px rgba(66,185,131,0.04);
}

.hero-card.admin-theme h1 {
  color: #fff;
  text-shadow: 0 2px 12px rgba(66,185,131,0.13);
}

.admin-crown {
  font-size: 2.1rem;
  margin-right: 0.4rem;
  vertical-align: -0.25em;
  filter: drop-shadow(0 2px 8px rgba(66,185,131,0.18));
  animation: admin-crown-pop 0.7s cubic-bezier(.21,.68,.41,1.12);
}
@keyframes admin-crown-pop {
  0%   { transform: scale(0.7) translateY(-25px); opacity: 0.1; }
  70%  { transform: scale(1.15) translateY(3px); opacity: 1; }
  100% { transform: scale(1) translateY(0); opacity: 1; }
}

.hero-card p {
  font-size: 1.125rem;
  margin-bottom: 2rem;
  color: #555;
}
.hero-card.admin-theme p {
  color: #eafff2;
}

.hero-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  justify-content: center;
  margin-bottom: 1rem;
}

.btn {
  padding: 0.75rem 1.5rem;
  border-radius: 5px;
  text-decoration: none;
  font-weight: 600;
  transition: transform 0.2s, box-shadow 0.2s, background 0.18s;
  font-size: 1.08em;
  border: none;
  outline: none;
  display: inline-flex;
  align-items: center;
  gap: 0.3em;
}

.btn-primary {
  background: linear-gradient(90deg, #42b983 0%, #2c3e50 100%);
  color: #fff;
}

.btn-secondary {
  background: #ecf0f1;
  color: #2c3e50;
}

.btn-danger,
.btn-admin {
  background: linear-gradient(90deg, #ec4646 60%, #23272f 100%);
  color: #fff;
  font-weight: 700;
  border: 2px solid #ec4646;
}
.btn-admin .admin-crown-btn {
  font-size: 1.1em;
  margin-right: 0.1em;
  filter: drop-shadow(0 2px 8px rgba(66,185,131,0.17));
}

.btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 4px 12px rgba(66, 185, 131, 0.17);
}

.admin-welcome {
  background: rgba(44, 62, 80, 0.19);
  margin-top: 2.2rem;
  border-radius: 13px;
  padding: 1.2rem 1.1rem 1.5rem 1.1rem;
  box-shadow: 0 1px 5px rgba(66,185,131,0.07);
}
.admin-welcome h2 {
  margin: 0.5rem 0 0.7rem 0;
  color: #fff;
  font-size: 1.45em;
  letter-spacing: 0.5px;
  font-weight: 800;
  text-shadow: 0 2px 8px rgba(66,185,131,0.09);
}
.admin-hint {
  color: #d2fff1;
  opacity: 0.9;
  font-size: 1.04em;
  margin-top: 0.5rem;
  display: block;
}
</style>