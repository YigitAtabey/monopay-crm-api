<template>
  <div class="user-panel">
    <h2>Kullanıcı Paneli</h2>

    <!-- Profil Bilgileri -->
    <section class="profile">
      <h3>
        Profilim
        <span :class="['role-badge', user.role]">
          {{ user.role === 'admin' ? '👑 Admin' : (user.role === 'blocked' ? '⛔ Engelli' : 'User') }}
        </span>
      </h3>
      <p><strong>İsim:</strong> {{ user.name }}</p>
      <p><strong>E-posta:</strong> {{ user.email }}</p>
    </section>

    <!-- Kendi Görevlerim -->
    <section class="my-tasks">
      <h3>Görevlerim</h3>
      <p v-if="taskError" class="error">{{ taskError }}</p>
      <ul class="tasks-ul">
        <li v-for="t in tasks" :key="t.id">
          <span class="task-title">{{ t.title }}</span>
          <button class="delete-btn" @click="deleteTask(t.id)">
            <span class="delete-icon">🗑</span> Sil
          </button>
        </li>
      </ul>
      <p v-if="!taskError && tasks.length === 0" class="empty">Henüz hiç görev yok.</p>
      <div class="new-task">
        <input
          v-model="newTaskTitle"
          placeholder="Yeni görev başlığı"
          @keyup.enter="addTask"
        />
        <button class="add-btn" @click="addTask" :disabled="!newTaskTitle.trim()">
          ➕ Ekle
        </button>
      </div>
    </section>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import axios from 'axios'
import { useAuthStore } from '@/store/auth'

const auth = useAuthStore()
// Farklı backend alan adları için uyumlu isimlendirme (name/firstName/lastName)
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
    taskError.value = 'Görevler alınırken hata oluştu.'
  }
}

const addTask = async () => {
  if (!newTaskTitle.value.trim()) return
  try {
    await axios.post('/tasks', { title: newTaskTitle.value.trim() })
    newTaskTitle.value = ''
    await fetchTasks()
  } catch {
    taskError.value = 'Görev eklenirken hata oluştu.'
  }
}

const deleteTask = async (id) => {
  try {
    await axios.delete(`/tasks/${id}`)
    tasks.value = tasks.value.filter(t => t.id !== id)
  } catch {
    taskError.value = 'Görev silinirken hata oluştu.'
  }
}

onMounted(fetchTasks)
</script>

<style scoped>
.user-panel {
  max-width: 700px;
  margin: 2rem auto;
  padding: 1.6rem 1.5rem 1.5rem 1.5rem;
  background: #23272f;
  border-radius: 16px;
  box-shadow: 0 8px 32px 0 rgba(60,60,90,0.07), 0 2px 4px rgba(50,50,100,0.06);
  color: #f7f7f7;
}
h2 {
  text-align: center;
  margin-bottom: 2rem;
  color: #42b983;
  font-weight: 800;
  letter-spacing: 1px;
  text-shadow: 0 2px 10px rgba(66,185,131,0.05);
}
.profile, .my-tasks {
  margin-bottom: 2rem;
  background: rgba(40,44,52,0.93);
  border-radius: 12px;
  padding: 1.5rem 1.2rem;
  box-shadow: 0 1px 4px rgba(66,185,131,0.03);
}
.profile h3 {
  display: flex;
  align-items: center;
  gap: 1rem;
  font-size: 1.25em;
  margin-bottom: 1.1rem;
}
.role-badge {
  display: inline-block;
  padding: 0.2em 0.9em;
  border-radius: 12px;
  font-size: 0.98em;
  font-weight: 700;
  letter-spacing: 1px;
  margin-left: 0.7rem;
}
.role-badge.admin {
  background: linear-gradient(90deg, #42b983 60%, #2c3e50 100%);
  color: #fff;
  box-shadow: 0 1px 4px rgba(66,185,131,0.13);
}
.role-badge.blocked {
  background: linear-gradient(90deg, #ec4646 60%, #2c3e50 100%);
  color: #fff;
  box-shadow: 0 1px 4px rgba(236,70,70,0.10);
}
.role-badge.user {
  background: #3a3f4a;
  color: #eafff2;
}
.my-tasks {
  margin-top: 2rem;
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
  padding: 0.7rem 0.4rem;
  border-bottom: 1px solid #31343e;
  font-size: 1.07em;
  transition: background 0.16s;
}
.task-title {
  font-weight: 600;
  color: #fafafa;
  flex: 1;
}
.delete-btn {
  background: linear-gradient(90deg, #ec4646 60%, #2c3e50 100%);
  color: #fff;
  border: none;
  border-radius: 5px;
  padding: 0.36rem 0.88rem;
  font-weight: 600;
  font-size: 1em;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.2em;
  box-shadow: 0 1px 4px rgba(236,70,70,0.10);
  transition: background 0.18s, color 0.18s;
}
.delete-btn:hover {
  background: #d73030;
  color: #fff;
  transform: translateY(-1px) scale(1.03);
}
.delete-icon {
  font-size: 1.13em;
}
.empty {
  color: #b9b9b9;
  text-align: center;
  font-size: 1.08em;
  font-style: italic;
  margin: 2.2rem 0 1rem 0;
  letter-spacing: 0.2px;
}
.new-task {
  margin-top: 1.5rem;
  display: flex;
  gap: 0.7rem;
  align-items: center;
  justify-content: stretch;
}
.new-task input {
  flex: 1;
  font-size: 1em;
  padding: 0.5rem 1rem;
  border: 1.5px solid #42b983;
  border-radius: 5px;
  background: #23272f;
  color: #fafafa;
  outline: none;
  transition: border 0.16s;
}
.new-task input:focus {
  border: 1.5px solid #43e39b;
  background: #262b34;
  color: #fff;
}
.add-btn {
  background: linear-gradient(90deg, #42b983 60%, #2c3e50 100%);
  color: #fff;
  border: none;
  border-radius: 5px;
  font-weight: 700;
  font-size: 1.02em;
  padding: 0.55rem 1.25rem;
  cursor: pointer;
  transition: background 0.18s, color 0.18s;
  box-shadow: 0 1px 4px rgba(66,185,131,0.13);
}
.add-btn:disabled {
  background: #3a3a3a;
  color: #aaa;
  cursor: not-allowed;
  opacity: 0.68;
}
.add-btn:not(:disabled):hover {
  background: #369f6e;
  color: #fff;
  transform: translateY(-1px) scale(1.01);
}
.error {
  color: #ec4646;
  margin-bottom: 1rem;
  text-align: center;
  font-weight: 700;
  letter-spacing: 0.5px;
}
</style>