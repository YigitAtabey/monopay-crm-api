<template>
  <div class="task-list">
    <h2>Görev Listesi</h2>
    <p v-if="error" class="error">{{ error }}</p>
    <ul class="tasks-ul">
      <li v-for="t in tasks" :key="t.id" :class="{'completed': t.completed}">
        <span class="task-title">{{ t.title }}</span>
        <span v-if="t.completed" class="task-status">✔ Tamamlandı</span>
        <span v-else class="task-status pending">⏳ Devam Ediyor</span>
      </li>
    </ul>
    <p v-if="!error && tasks.length === 0" class="empty">Henüz hiç görev yok.</p>
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
    // Eğer backend'den completed alanı gelmiyorsa, hepsi false olsun
    tasks.value = (res.data || []).map(t => ({
      ...t,
      completed: t.completed ?? false
    }))
    error.value = ''
  } catch {
    error.value = 'Görevler alınırken hata oluştu.'
    tasks.value = []
  }
}

onMounted(fetchTasks)
</script>

<style scoped>
.task-list {
  max-width: 600px;
  margin: 2rem auto;
  padding: 1.5rem;
  background: #23272f;
  border-radius: 16px;
  box-shadow: 0 8px 32px 0 rgba(60,60,90,0.07), 0 2px 4px rgba(50,50,100,0.06);
  color: #f7f7f7;
}
h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #42b983;
  font-weight: 800;
  letter-spacing: 1px;
  text-shadow: 0 2px 10px rgba(66,185,131,0.05);
}
.tasks-ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
.tasks-ul li {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  padding: 0.7rem 0.6rem;
  border-bottom: 1px solid #31343e;
  font-size: 1.09em;
  transition: background 0.16s;
}
.tasks-ul li.completed {
  background: linear-gradient(90deg, #42b98322 10%, #23272f 70%);
  color: #42b983;
  text-decoration: line-through;
}
.task-title {
  font-weight: 600;
  color: #fafafa;
}
.task-status {
  font-size: 0.97em;
  font-weight: 700;
  padding: 0.18em 0.7em;
  border-radius: 7px;
  margin-left: 0.7rem;
}
.task-status {
  background: #23272f;
  color: #42b983;
}
.tasks-ul li.completed .task-status {
  background: #42b983;
  color: #23272f;
}
.task-status.pending {
  background: #ec4646;
  color: #fff;
}
.error {
  color: #ec4646;
  margin-bottom: 1rem;
  text-align: center;
  font-weight: 700;
  letter-spacing: 0.5px;
}
.empty {
  text-align: center;
  color: #b9b9b9;
  margin-top: 2.2rem;
  font-size: 1.08em;
  font-style: italic;
  letter-spacing: 0.2px;
}
</style>