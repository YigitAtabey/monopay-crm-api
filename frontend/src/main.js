import { createApp }    from 'vue'
import App              from './App.vue'
import router           from './router'
import { createPinia }  from 'pinia'
import axios            from 'axios'

// Bütün axios çağrıları /api ile başlar
axios.defaults.baseURL = '/api'

// Token localStorage'da varsa header'a ekle
const savedToken = localStorage.getItem('token')
if (savedToken) {
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + savedToken
}

// (Opsiyonel) Tüm axios isteklerinde token süresi bitmişse logout ve yönlendirme
// axios.interceptors.response.use(
//   response => response,
//   err => {
//     if (err.response && err.response.status === 401) {
//       localStorage.removeItem('token')
//       window.location = '/login'
//     }
//     return Promise.reject(err)
//   }
// )

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')