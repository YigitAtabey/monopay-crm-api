import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: JSON.parse(localStorage.getItem('user') || 'null')
  }),
  getters: {
    isLoggedIn: (state) => !!state.token,
    // Role tabanlı sistem: admin kontrolü role stringiyle yapılır
    isAdmin: (state) => state.user?.role === 'admin',
    isBlocked: (state) => state.user?.role === 'blocked'
  },
  actions: {
    // Login işlemi: token al, sonra profil verisini çek
    async login(credentials) {
      // 1) Kimlik doğrulama: token al
      const res = await axios.post('/login', credentials)
      const token = res.data.token
      const user = res.data.user // Login response'da user objesi var!

      // 2) Token'ı state'e ve localStorage'a kaydet, header ayarla
      this.token = token
      localStorage.setItem('token', token)
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

      // 3) User objesini state'e ve localStorage'a kaydet
      this.user = user
      localStorage.setItem('user', JSON.stringify(user))

      return user
    },

    // Logout: tüm verileri temizle
    logout() {
      this.token = null
      this.user  = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      delete axios.defaults.headers.common['Authorization']
    }
  }
})