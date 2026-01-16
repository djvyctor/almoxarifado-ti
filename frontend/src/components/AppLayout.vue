<template>
  <div class="app-layout">
    <header class="header">
      <div class="header-content">
        <h1>Almoxarifado TI</h1>
        <div class="user-menu">
          <span>Olá, Admin</span>
          <button @click="handleLogout" class="logout-btn">Sair</button>
        </div>
      </div>
    </header>

    <div class="main-container">
      <aside class="sidebar">
        <nav>
          <router-link to="/" class="nav-item">
            <span>🏠</span> Home
          </router-link>
          <router-link to="/items" class="nav-item">
            <span>📦</span> Itens
          </router-link>
        </nav>
      </aside>

      <main class="content">
        <slot></slot>
      </main>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.app-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: #2c3e50;
  color: white;
  padding: 1rem 2rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.header-content {
  max-width: 1400px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header h1 {
  margin: 0;
  font-size: 1.5rem;
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logout-btn {
  padding: 0.5rem 1rem;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
}

.logout-btn:hover {
  background-color: #c0392b;
}

.main-container {
  flex: 1;
  display: flex;
  max-width: 1400px;
  width: 100%;
  margin: 0 auto;
}

.sidebar {
  width: 250px;
  background-color: #34495e;
  padding: 2rem 0;
}

.sidebar nav {
  display: flex;
  flex-direction: column;
}

.nav-item {
  padding: 1rem 2rem;
  color: #ecf0f1;
  text-decoration: none;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  transition: background-color 0.2s;
}

.nav-item:hover {
  background-color: #2c3e50;
}

.nav-item.router-link-active {
  background-color: #2c3e50;
  border-left: 4px solid #3498db;
}

.content {
  flex: 1;
  padding: 2rem;
  background-color: #ecf0f1;
}
</style>