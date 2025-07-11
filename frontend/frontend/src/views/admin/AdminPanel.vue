<template>
    <div class="admin-panel">
      <h2>Admin Paneli</h2>
  
      <!-- 1) Kullanıcı Listesi ve Yönetim -->
      <section class="user-management">
        <h3>Kullanıcıları Yönet</h3>
        <p v-if="userError" class="error">{{ userError }}</p>
        <ul>
          <li v-for="u in users" :key="u.id">
            {{ u.name }} ({{ u.email }})
            <button @click="makeAdmin(u.email)">Yönetici Yap</button>
            <button @click="blockUser(u.id)">Engelle</button>
          </li>
        </ul>
      </section>
  
      <!-- 2) E-posta ile Hızlı İşlemler -->
      <section class="quick-actions">
        <h3>Toplu İşlem</h3>
        <input v-model="emailAction" placeholder="E-posta gir" />
        <button @click="makeAdmin(emailAction)">Yönetici Yap</button>
        <button @click="blockByEmail(emailAction)">Engelle</button>
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
      const res = await axios.get('/admin/users')
      users.value = res.data
      userError.value = ''
    } catch {
      userError.value = 'Kullanıcılar alınırken hata oluştu'
    }
  }
  
  // Yönetici yap (email ile)
  const makeAdmin = async (email) => {
    if (!email.trim()) return
    try {
      await axios.put('/admin/make', { email })
      await fetchUsers()
    } catch {
      userError.value = 'Yönetici yapma başarısız'
    }
  }
  
  // Kullanıcıyı engelle (ID ile)
  const blockUser = async (id) => {
    try {
      await axios.put(`/admin/block/${id}`)
      users.value = users.value.filter(u => u.id !== id)
    } catch {
      userError.value = 'Engelleme başarısız'
    }
  }
  
  // E-posta ile engelle
  const blockByEmail = async (email) => {
    if (!email.trim()) return
    try {
      await axios.put('/admin/block', { email })
      await fetchUsers()
    } catch {
      userError.value = 'Engelleme başarısız'
    }
  }
  
  onMounted(fetchUsers)
  </script>
  
  <style scoped>
  .admin-panel { max-width: 700px; margin: auto; padding: 1rem; }
  .user-management, .quick-actions {
    margin-bottom: 2rem;
    padding: 1rem; border: 1px solid #ddd; border-radius: 4px;
  }
  .user-management ul { list-style: none; padding: 0; }
  .user-management li {
    display: flex; align-items: center;
    justify-content: space-between;
    padding: 0.5rem 0; border-bottom: 1px solid #eee;
  }
  .quick-actions input { width: 60%; padding: 0.4rem; }
  button { margin-left: 0.5rem; cursor: pointer; }
  .error { color: red; margin-bottom: 0.5rem; }
  </style>
  