<template>
  <div class="admin-panel">
    <h2>Admin Paneli</h2>

    <!-- Kullanıcı Listesi ve Yönetim -->
    <section class="user-management">
      <h3>Kullanıcıları Yönet</h3>
      <p v-if="userError" class="error">{{ userError }}</p>
      <ul>
        <li v-for="u in users" :key="u.email">
          <div class="user-info">
            <span class="user-name">{{ u.name }}</span>
            <span class="user-email">({{ u.email }})</span>
            <span v-if="u.role === 'admin'" class="admin-badge">[Yönetici]</span>
            <span v-if="u.role === 'blocked'" class="blocked-badge">[Engelli]</span>
          </div>
          <div class="actions">
            <button
              class="admin-btn"
              @click="makeAdmin(u.email)"
              :disabled="u.role === 'admin'"
              :title="u.role === 'admin' ? 'Zaten yönetici' : 'Yönetici yap'"
            >
              👑 Yönetici Yap
            </button>
            <button
              class="block-btn"
              @click="blockByEmail(u.email)"
              :disabled="u.role === 'blocked'"
              :title="u.role === 'blocked' ? 'Zaten engelli' : 'Engelle'"
            >
              🚫 Engelle
            </button>
            <button
              class="unblock-btn"
              @click="unblockByEmail(u.email)"
              :disabled="u.role !== 'blocked'"
              :title="u.role !== 'blocked' ? 'Engelli değil' : 'Engeli kaldır'"
            >
              🔓 Engeli Kaldır
            </button>
          </div>
        </li>
      </ul>
    </section>

    <!-- E-posta ile Hızlı İşlemler -->
    <section class="quick-actions">
      <h3>Toplu İşlem</h3>
      <div class="quick-form">
        <input v-model="emailAction" placeholder="E-posta gir" />
        <button class="admin-btn" @click="makeAdmin(emailAction)">👑 Yönetici Yap</button>
        <button class="block-btn" @click="blockByEmail(emailAction)">🚫 Engelle</button>
        <button class="unblock-btn" @click="unblockByEmail(emailAction)">🔓 Engeli Kaldır</button>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const users       = ref([])
const userError   = ref('')
const emailAction = ref('')

// Kullanıcıları çek
const fetchUsers = async () => {
  try {
    const res = await axios.get('/users')
    users.value = res.data
    userError.value = ''
  } catch {
    userError.value = 'Kullanıcılar alınırken hata oluştu.'
  }
}

// Yönetici yap (email ile)
const makeAdmin = async (email) => {
  if (!email || !email.trim()) {
    userError.value = 'Geçerli bir e-posta giriniz.'
    return
  }
  try {
    await axios.put('/admin/make', { email })
    await fetchUsers()
    userError.value = ''
  } catch (err) {
    userError.value = err?.response?.data?.message || 'Yönetici yapma başarısız.'
  }
}

// Kullanıcıyı engelle (email ile)
const blockByEmail = async (email) => {
  if (!email || !email.trim()) {
    userError.value = 'Geçerli bir e-posta giriniz.'
    return
  }
  try {
    await axios.put('/admin/block', { email })
    await fetchUsers()
    userError.value = ''
  } catch (err) {
    userError.value = err?.response?.data?.message || 'Engelleme başarısız.'
  }
}

// Engel kaldır (email ile)
const unblockByEmail = async (email) => {
  if (!email || !email.trim()) {
    userError.value = 'Geçerli bir e-posta giriniz.'
    return
  }
  try {
    await axios.put('/admin/unblock', { email })
    await fetchUsers()
    userError.value = ''
  } catch (err) {
    userError.value = err?.response?.data?.message || 'Engel kaldırma başarısız.'
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.admin-panel {
  max-width: 700px;
  margin: 2rem auto;
  padding: 2rem 1.5rem 1.5rem 1.5rem;
  background: #23272f;
  border-radius: 18px;
  box-shadow: 0 8px 32px 0 rgba(60,60,90,0.10), 0 2px 4px rgba(50,50,100,0.07);
  color: #f7f7f7;
}

h2 {
  text-align: center;
  margin-bottom: 2rem;
  letter-spacing: 1px;
  font-weight: 800;
  color: #42b983;
  text-shadow: 0 2px 10px rgba(66,185,131,0.05);
}

section {
  margin-bottom: 2.5rem;
  background: rgba(40,44,52,0.93);
  border-radius: 12px;
  padding: 1.5rem 1.2rem;
  box-shadow: 0 1px 4px rgba(66,185,131,0.03);
}

.user-management ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.user-management li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: rgba(55,60,75,0.23);
  border-radius: 6px;
  padding: 0.8rem 1rem;
  margin-bottom: 0.9rem;
  box-shadow: 0 1px 2px rgba(66,185,131,0.03);
  transition: background 0.2s;
}

.user-management li:hover {
  background: rgba(66,185,131,0.12);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  font-size: 1.09em;
}

.user-name {
  font-weight: 600;
  color: #fafafa;
}
.user-email {
  color: #b9b9b9;
  font-size: 0.98em;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

.admin-btn, .block-btn, .unblock-btn {
  font-size: 0.98em;
  border: none;
  border-radius: 5px;
  padding: 0.38rem 0.95rem;
  cursor: pointer;
  font-weight: 600;
  transition: background 0.18s, color 0.18s;
}

.admin-btn {
  background: linear-gradient(90deg, #42b983 60%, #2c3e50 100%);
  color: #fff;
  box-shadow: 0 1px 4px rgba(66,185,131,0.13);
}

.admin-btn:disabled {
  background: #3a3a3a;
  color: #aaa;
  cursor: not-allowed;
  opacity: 0.6;
}

.block-btn {
  background: linear-gradient(90deg, #ec4646 50%, #2c3e50 100%);
  color: #fff;
  box-shadow: 0 1px 4px rgba(236,70,70,0.10);
}

.block-btn:disabled {
  background: #3a3a3a;
  color: #aaa;
  cursor: not-allowed;
  opacity: 0.6;
}

.unblock-btn {
  background: linear-gradient(90deg, #42b983 50%, #1a8c7c 100%);
  color: #fff;
  box-shadow: 0 1px 4px rgba(66,185,131,0.10);
}

.unblock-btn:disabled {
  background: #3a3a3a;
  color: #aaa;
  cursor: not-allowed;
  opacity: 0.6;
}

.admin-badge {
  color: #42b983;
  font-weight: 700;
  margin-left: 0.7rem;
  letter-spacing: 1px;
  font-size: 0.93em;
}

.blocked-badge {
  color: #ec4646;
  font-weight: bold;
  margin-left: 0.7rem;
  letter-spacing: 1px;
  font-size: 0.93em;
}

.quick-actions .quick-form {
  display: flex;
  gap: 0.7rem;
  align-items: center;
  margin-top: 0.6rem;
}

.quick-actions input {
  flex: 1 1 0%;
  min-width: 180px;
  font-size: 1em;
  padding: 0.43rem 0.9rem;
  border: 1.5px solid #42b983;
  border-radius: 5px;
  background: #23272f;
  color: #fafafa;
  outline: none;
  transition: border 0.16s;
}

.quick-actions input:focus {
  border: 1.5px solid #42b983;
  background: #262b34;
  color: #fff;
}

.error {
  color: #ec4646;
  margin-bottom: 0.7rem;
  font-weight: 700;
  letter-spacing: 0.5px;
}
</style>