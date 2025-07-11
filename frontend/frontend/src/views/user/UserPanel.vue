<template>
    <div class="user-panel">
      <h2>Kullanıcı Paneli</h2>
  
      <!-- Profil Bilgileri -->
      <section class="profile">
        <h3>Profilim
          <span :class="['role-badge', user.role]">
            {{ user.role === 'admin' ? 'Admin' : 'User' }}
          </span>
        </h3>
        <p><strong>İsim:</strong> {{ user.name }}</p>
        <p><strong>E-posta:</strong> {{ user.email }}</p>
      </section>
  
      <!-- Kendi Görevlerim -->
      <section class="my-tasks">
        <h3>Görevlerim</h3>
        <p v-if="taskError" class="error">{{ taskError }}</p>
        <ul>
          <li v-for="t in tasks" :key="t.id">
            {{ t.title }}
            <button @click="deleteTask(t.id)">Sil</button>
          </li>
        </ul>
        <div class="new-task">
          <input v-model="newTaskTitle" placeholder="Yeni görev başlığı" />
          <button @click="addTask">Ekle</button>
        </div>
      </section>
    </div>
  </template>
  
  <script setup>
  import { computed, ref, onMounted } from 'vue'
  import axios from 'axios'
  import { useAuthStore } from '@/store/auth'
  
  const auth = useAuthStore()
  // Backend'den dönen farklı alan isimlerini (firstName/lastName veya name) birleştiriyoruz
  const user = computed(() => {
    const u = auth.user || {}
    const fullName = u.name || [u.firstName, u.lastName].filter(Boolean).join(' ')
    return {
      name: fullName || '-',
      email: u.email || '',
      role: u.role || ''
    }
  })
  
  const tasks = ref([])
  const newTaskTitle = ref('')
  const taskError = ref('')
  
  const fetchTasks = async () => {
    try {
      const res = await axios.get('/tasks')
      tasks.value = res.data
      taskError.value = ''
    } catch {
      taskError.value = 'Görevler alınırken hata oluştu'
    }
  }
  
  const addTask = async () => {
    if (!newTaskTitle.value.trim()) return
    try {
      await axios.post('/tasks', { title: newTaskTitle.value })
      newTaskTitle.value = ''
      await fetchTasks()
    } catch {
      taskError.value = 'Görev eklenirken hata oluştu'
    }
  }
  
  const deleteTask = async (id) => {
    try {
      await axios.delete(`/tasks/${id}`)
      tasks.value = tasks.value.filter(t => t.id !== id)
    } catch {
      taskError.value = 'Görev silinirken hata oluştu'
    }
  }
  
  onMounted(fetchTasks)
  </script>
  
  <style scoped>
  .user-panel { max-width: 700px; margin: auto; padding: 1rem; }
  .profile, .my-tasks {
    margin-bottom: 2rem;
    padding: 1rem; border: 1px solid #ddd; border-radius: 4px;
  }
  .my-tasks ul { list-style: none; padding: 0; }
  .my-tasks li {
    display: flex; justify-content: space-between;
    padding: 0.5rem 0; border-bottom: 1px solid #eee;
  }
  .new-task { display: flex; gap: 0.5rem; margin-top: 1rem; }
  button { cursor: pointer; }
  .error { color: red; margin-bottom: 0.5rem; }
  </style>
  