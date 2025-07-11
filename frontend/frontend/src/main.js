import { createApp }    from 'vue'
import App              from './App.vue'
import router           from './router'
import { createPinia }  from 'pinia'
import axios            from 'axios'

// Bütün axios çağrıları /api ile başlar
axios.defaults.baseURL = '/api'

const savedToken = localStorage.getItem('token')
if (savedToken) {
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + savedToken
}

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
