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
  background: #23272f;
  font-family: 'Segoe UI', 'Roboto', 'Arial', sans-serif;
  color: #f7f7f7;
}

/* App container */
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: #23272f;
}

/* Ana içerik alanı */
.main-container {
  flex: 1;
  max-width: 900px;
  margin: 2.5rem auto 1.5rem auto;
  padding: 2rem 1.5rem 1.5rem 1.5rem;
  background: #23272f;
  border-radius: 16px;
  box-shadow: 0 8px 32px 0 rgba(60,60,90,0.10), 0 2px 4px rgba(50,50,100,0.07);
  color: #f7f7f7;
  transition: background 0.23s, color 0.18s;
}

/* Scrollbar görünümü */
::-webkit-scrollbar {
  width: 8px;
  background: #191c22;
}
::-webkit-scrollbar-thumb {
  background: #353946;
  border-radius: 8px;
}

/* Responsive */
@media (max-width: 980px) {
  .main-container {
    max-width: 98vw;
    padding: 1.1rem 0.3rem 1.1rem 0.3rem;
  }
}

/* Karanlık tema linkleri */
a {
  color: #42b983;
  text-decoration: none;
  transition: color 0.15s;
}
a:hover {
  color: #2fa072;
  text-decoration: underline;
}
</style>