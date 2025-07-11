<template>
    <div class="task-list">
      <h2>Görev Listesi</h2>
      <p v-if="error" class="error">{{ error }}</p>
      <ul>
        <li v-for="t in tasks" :key="t.id">
          {{ t.title }}
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  
  const tasks = ref([])
  const error = ref('')
  
  const fetchTasks = async () => {
    try {
      const res = await axios.get('/tasks')
      tasks.value = res.data
      error.value = ''
    } catch {
      error.value = 'Görevler alınırken hata oluştu'
    }
  }
  
  onMounted(fetchTasks)
  </script>
  
  <style scoped>
  .task-list {
    max-width: 600px;
    margin: auto;
  }
  ul {
    list-style: none;
    padding: 0;
  }
  li {
    padding: 0.5rem 0;
    border-bottom: 1px solid #ddd;
  }
  .error {
    color: red;
    margin-bottom: 1rem;
  }
  </style>
  