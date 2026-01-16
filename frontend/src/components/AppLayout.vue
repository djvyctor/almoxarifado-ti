<template>
  <div class="min-h-screen bg-slate-50 flex">
    <aside class="hidden md:flex flex-col w-64 bg-slate-900 text-white fixed h-full z-20">
      <div class="h-16 flex items-center px-6 border-b border-slate-800">
        <div class="w-8 h-8 bg-gradient-to-br from-purple-500 to-pink-500 rounded-lg mr-3 flex items-center justify-center">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
        </div>
        <span class="font-bold text-lg tracking-wide">Almoxarifado</span>
      </div>

      <nav class="flex-1 py-6 px-3 space-y-1">
        <router-link 
          to="/" 
          class="flex items-center px-4 py-3 rounded-xl transition-all group"
          active-class="bg-purple-600 text-white shadow-lg shadow-purple-500/30"
          :class="$route.path === '/' ? '' : 'text-slate-400 hover:bg-slate-800 hover:text-white'"
        >
          <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          Dashboard
        </router-link>

        <router-link 
          to="/items" 
          class="flex items-center px-4 py-3 rounded-xl transition-all group"
          active-class="bg-purple-600 text-white shadow-lg shadow-purple-500/30"
          :class="$route.path.includes('/items') ? '' : 'text-slate-400 hover:bg-slate-800 hover:text-white'"
        >
          <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          Estoque
        </router-link>
      </nav>

      <div class="p-4 border-t border-slate-800">
        <button 
          @click="handleLogout"
          class="flex items-center w-full px-4 py-2 text-slate-400 hover:text-white hover:bg-slate-800 rounded-lg transition-colors"
        >
          <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          Sair do Sistema
        </button>
      </div>
    </aside>

    <main class="flex-1 md:ml-64 transition-all">
      <div class="md:hidden h-16 bg-slate-900 text-white flex items-center justify-between px-4">
        <span class="font-bold">Almoxarifado</span>
        </div>

      <slot></slot>
      
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.logout() // Certifique-se que sua store tem essa action
  router.push('/login')
}
</script>