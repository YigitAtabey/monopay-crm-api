<!-- src/views/user/Logout.vue -->
<template>
  <div class="logout-page">
    <p>Çıkış yapılıyor… Lütfen bekleyin.</p>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const auth   = useAuthStore()
const router = useRouter()

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
  font-size: 1.2rem;
  color: #444;
}
</style>
