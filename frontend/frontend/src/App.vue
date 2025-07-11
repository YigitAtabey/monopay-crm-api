<!-- src/App.vue -->
<template>
  <div id="app">
    <!-- NavBar bileşenini buraya ekliyoruz -->
    <NavBar />

    <!-- Router view: sayfa içeriği buraya gelecek -->
    <main class="main-container">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import NavBar from '@/components/NavBar.vue'
</script>

<style>
/* Global ayarlar */
body {
  margin: 0;
  background: #f0f2f5;
  font-family: 'Segoe UI', sans-serif;
}

/* App container */
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* Ana içerik alanı */
.main-container {
  flex: 1;
  max-width: 900px;
  margin: 2rem auto;
  padding: 1rem;
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}
</style>
