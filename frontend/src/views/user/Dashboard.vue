<template>
  <div class="dashboard">
    <h2>Kullanıcı Paneli</h2>
    <p v-if="error" class="error">{{ error }}</p>
    <ul class="user-list">
      <li v-for="u in users" :key="u.id">
        {{ u.name }} ({{ u.email }})
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const users = ref([])
const error = ref('')

const fetchUsers = async () => {
  try {
    const res = await axios.get('/users')
    users.value = res.data
    error.value = ''
  } catch {
    error.value = 'Kullanıcılar alınırken hata oluştu'
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.dashboard {
  max-width: 600px;
  margin: auto;
  padding: 1rem;
}
.user-list { list-style: none; padding: 0; }
.user-list li {
  padding: 0.5rem 0;
  border-bottom: 1px solid #ddd;
}
.error { color: red; margin-bottom: 0.5rem; }
</style>
