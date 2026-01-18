<template>
  <div class="min-h-screen bg-black flex font-sans text-zinc-400 selection:bg-blue-500/30">
    
    <aside class="fixed inset-y-0 left-0 z-50 w-64 bg-[#0c0c0e] border-r border-white/5 transition-transform duration-300 flex flex-col hidden md:flex">
      
      <div class="h-24 flex items-center px-8 border-b border-white/5">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-xl bg-gradient-to-tr from-[#0071e3] to-indigo-600 flex items-center justify-center shadow-lg shadow-blue-900/20">
            <svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <div>
            <h1 class="text-lg font-bold text-white tracking-tight">Corvi</h1>
            <p class="text-[10px] text-zinc-500 uppercase tracking-widest font-semibold">System</p>
          </div>
        </div>
      </div>

      <nav class="flex-1 px-4 py-8 space-y-2 overflow-y-auto custom-scrollbar">
        
        <p class="px-4 text-xs font-semibold text-zinc-600 uppercase tracking-wider mb-4">Menu Principal</p>

        <router-link 
          to="/" 
          class="flex items-center gap-3 px-4 py-3 rounded-xl transition-all duration-200 group relative overflow-hidden"
          active-class="!bg-[#0071e3] !text-white shadow-lg shadow-blue-900/20"
          :class="$route.path === '/' ? 'text-white' : 'hover:bg-white/5 hover:text-zinc-200'"
        >
          <svg class="w-5 h-5 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
          <span class="font-medium">Dashboard</span>
        </router-link>

        <router-link 
          to="/items" 
          class="flex items-center gap-3 px-4 py-3 rounded-xl transition-all duration-200 group hover:bg-white/5 hover:text-zinc-200"
          active-class="!bg-[#0071e3] !text-white shadow-lg shadow-blue-900/20"
        >
          <svg class="w-5 h-5 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
          <span class="font-medium">Estoque</span>
        </router-link>

        <router-link 
          to="/users" 
          class="flex items-center gap-3 px-4 py-3 rounded-xl transition-all duration-200 group hover:bg-white/5 hover:text-zinc-200"
          active-class="!bg-[#0071e3] !text-white shadow-lg shadow-blue-900/20"
        >
          <svg class="w-5 h-5 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <span class="font-medium">Usuários</span>
        </router-link>
        
        <router-link 
          to="/about" 
          class="flex items-center gap-3 px-4 py-3 rounded-xl transition-all duration-200 group hover:bg-white/5 hover:text-zinc-200"
          active-class="!bg-[#0071e3] !text-white shadow-lg shadow-blue-900/20"
        >
          <svg class="w-5 h-5 transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span class="font-medium">Sobre</span>
        </router-link>

      </nav>

      <div class="p-4 border-t border-white/5">
        <button 
          @click="handleLogout"
          class="flex items-center gap-3 w-full p-3 rounded-xl bg-white/5 border border-white/5 hover:bg-red-500/10 hover:border-red-500/20 hover:text-red-400 transition-all cursor-pointer group"
        >
          <div class="w-9 h-9 rounded-lg bg-[#1c1c1e] flex items-center justify-center text-zinc-400 group-hover:text-red-400 transition-colors">
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
          </div>
          <div class="flex-1 text-left">
            <p class="text-sm font-medium text-zinc-200 group-hover:text-red-400 transition-colors">Sair do Sistema</p>
            <p class="text-[10px] text-zinc-500">Encerrar sessão</p>
          </div>
        </button>
      </div>
    </aside>

    <main class="flex-1 md:ml-64 relative min-h-screen transition-all duration-300 flex flex-col">
      
      <div class="md:hidden h-16 bg-[#0c0c0e] border-b border-white/5 flex items-center justify-between px-4 sticky top-0 z-40">
        <div class="flex items-center gap-2">
          <div class="w-8 h-8 rounded-lg bg-gradient-to-tr from-[#0071e3] to-indigo-600 flex items-center justify-center">
             <svg class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <span class="font-bold text-white">Corvi</span>
        </div>
        <button class="text-zinc-400 hover:text-white">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" /></svg>
        </button>
      </div>

      <slot />
      
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  if (authStore.logout) {
    authStore.logout()
  }
  router.push('/login')
}
</script>

<style scoped>
/* Scrollbar fina para o menu */
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #27272a;
  border-radius: 10px;
}
</style>