<template>
  <div class="logout-page">
    <div v-if="!wasAdmin" class="logout-message">
      <p>Çıkış yapılıyor… Lütfen bekleyin.</p>
    </div>
    <div v-else class="logout-message admin-logout">
      <span class="crown">👑</span>
      <p><strong>Yönetici olarak çıkış yapıyorsunuz…</strong></p>
      <p class="admin-hint">Tüm yönetim oturumları güvenle sonlandırılıyor.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const auth      = useAuthStore()
const router    = useRouter()
const wasAdmin  = ref(false)

// Kullanıcı çıkışı yapılmadan önce rolü sakla
if (auth.user && auth.user.role === 'admin') {
  wasAdmin.value = true
}

onMounted(async () => {
  // 1) Store’u temizle
  auth.logout()

  // 2) login sayfasına yönlendir (replace: geri tuşunda /logout kalmasın)
  await router.replace('/login')

  // 3) Kısa bir bekleme -> tam sayfa yenilemesi
  await new Promise(res => setTimeout(res, 400))
  window.location.reload()
})
</script>

<style scoped>
.logout-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 150px);
}

.logout-message {
  font-size: 1.2rem;
  color: #444;
  text-align: center;
  padding: 2.5rem 2rem;
  background: #f8f8f8;
  border-radius: 12px;
  box-shadow: 0 2px 16px rgba(60,60,90,0.08);
}

.admin-logout {
  background: linear-gradient(100deg, #42b983 65%, #2c3e50 100%);
  color: #fff;
  box-shadow: 0 4px 32px 0 rgba(66,185,131,0.07);
  border: 2.5px solid #42b983;
  position: relative;
  padding-top: 2.5rem;
  padding-bottom: 2.2rem;
  animation: admin-pop 0.7s cubic-bezier(.21,.68,.41,1.12);
}

@keyframes admin-pop {
  0%   { transform: scale(0.7) translateY(70px); opacity: 0; }
  65%  { transform: scale(1.07) translateY(-8px); opacity: 1; }
  100% { transform: scale(1) translateY(0); opacity: 1; }
}

.crown {
  font-size: 2.7rem;
  display: block;
  margin-bottom: 0.7rem;
  filter: drop-shadow(0 2px 8px rgba(66,185,131,0.23));
}

.admin-hint {
  margin-top: 0.5rem;
  font-size: 1.05rem;
  color: #e1f6ee;
  opacity: 0.84;
}
</style>