<template>
    <div class="task-management">
      <h2>Görev Yönetimi</h2>
      <p v-if="error" class="error">{{ error }}</p>
      <ul class="task-list">
        <li v-for="t in tasks" :key="t.id">
          {{ t.title }}
          <button @click="deleteTask(t.id)">Sil</button>
        </li>
      </ul>
      <div class="new-task">
        <input v-model="newTaskTitle" placeholder="Yeni görev başlığı" />
        <button @click="addTask">Ekle</button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  
  const tasks = ref([])
  const error = ref('')
  const newTaskTitle = ref('')
  
  const fetchTasks = async () => {
    try {
      const res = await axios.get('/tasks')
      tasks.value = res.data
      error.value = ''
    } catch {
      error.value = 'Görevler alınırken hata oluştu'
    }
  }
  
  const addTask = async () => {
    if (!newTaskTitle.value.trim()) return
    try {
      await axios.post('/tasks', { title: newTaskTitle.value })
      newTaskTitle.value = ''
      await fetchTasks()
    } catch {
      error.value = 'Görev eklenirken hata oluştu'
    }
  }
  
  const deleteTask = async (id) => {
    try {
      await axios.delete(`/tasks/${id}`)
      tasks.value = tasks.value.filter(t => t.id !== id)
    } catch {
      error.value = 'Görev silinirken hata oluştu'
    }
  }
  
  onMounted(fetchTasks)
  </script>
  
  <style scoped>
  .task-management {
    max-width: 600px;
    margin: auto;
    padding: 1rem;
  }
  .task-list { list-style: none; padding: 0; }
  .task-list li {
    display: flex;
    justify-content: space-between;
    padding: 0.5rem 0;
    border-bottom: 1px solid #ddd;
  }
  .new-task {
    margin-top: 1rem;
    display: flex;
    gap: 0.5rem;
  }
  .error { color: red; margin-bottom: 0.5rem; }
  button { cursor: pointer; }
  </style>
  