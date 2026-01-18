<template>
  <div class="fixed inset-0 w-full h-full bg-black flex flex-col items-center justify-center overflow-hidden font-sans selection:bg-blue-500/30 selection:text-blue-200">
    
    <div class="absolute top-1/4 left-1/2 -translate-x-1/2 w-[600px] h-[600px] bg-blue-600/10 rounded-full blur-[120px] pointer-events-none"></div>

    <transition name="fade" mode="out-in">
      <div v-if="!showLoginForm" class="flex flex-col items-center justify-center z-10 w-full max-w-4xl px-6 text-center" key="landing">
        
        <div class="relative w-72 h-72 mb-10 flex items-center justify-center">
          
          <div class="absolute inset-0 border border-white/5 rounded-full animate-spin-slow-reverse"></div>
          <div class="absolute inset-8 border border-white/10 rounded-full animate-pulse-slow"></div>
          
          <div class="absolute top-0 right-12 w-14 h-14 bg-[#1c1c1e] border border-white/10 rounded-2xl shadow-2xl shadow-green-900/20 animate-float-slow flex items-center justify-center group hover:scale-110 transition-transform cursor-default">
             <div class="absolute inset-0 bg-green-500/10 rounded-2xl blur-md"></div>
             <svg class="w-7 h-7 text-green-400 relative z-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
             </svg>
          </div>

          <div class="absolute bottom-6 left-2 w-14 h-14 bg-[#1c1c1e] border border-white/10 rounded-2xl shadow-2xl shadow-blue-900/20 animate-float-delayed flex items-center justify-center group hover:scale-110 transition-transform cursor-default">
             <div class="absolute inset-0 bg-blue-500/10 rounded-2xl blur-md"></div>
             <svg class="w-7 h-7 text-blue-400 relative z-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4v1m6 11h2m-6 0h-2v4h-4v-2h4v-2h2v4m6-4v2h2v-2h-2zm1-5V4a1 1 0 00-1-1h-3.28a1 1 0 00-.948.683l-1.38 3.448a1 1 0 00.948 1.352h3.28a1 1 0 001-1zM10 4H6a1 1 0 00-1 1v3.28a1 1 0 00.683.948l3.448 1.38a1 1 0 001.352-.948V5a1 1 0 00-1-1z" />
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h6v6H3V3zm12 0h6v6h-6V3zm0 12h6v6h-6v-6zM3 15h6v6H3v-6z" />
             </svg>
          </div>

          <div class="absolute bottom-12 right-0 w-12 h-12 bg-[#1c1c1e] border border-white/10 rounded-2xl shadow-2xl shadow-orange-900/20 animate-float flex items-center justify-center group hover:scale-110 transition-transform cursor-default">
             <div class="absolute inset-0 bg-orange-500/10 rounded-2xl blur-md"></div>
             <svg class="w-6 h-6 text-orange-400 relative z-10" fill="none" viewBox="0 0 24 24" stroke="currentColor">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
             </svg>
          </div>

          <div class="relative w-36 h-36 rounded-[2rem] bg-gradient-to-br from-slate-800 to-black border-4 border-[#2c2c2e] shadow-2xl flex items-center justify-center z-20">
            <div class="absolute inset-0 bg-gradient-to-tr from-blue-500/20 to-transparent rounded-[2rem]"></div>
            <svg class="w-20 h-20 text-white drop-shadow-[0_0_15px_rgba(255,255,255,0.3)]" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
            <div class="absolute -top-2 -right-2 w-8 h-8 bg-red-500 rounded-full border-4 border-black flex items-center justify-center">
              <span class="text-[10px] font-bold text-white">3</span>
            </div>
          </div>
        </div>

        <h1 class="text-6xl font-bold text-white mb-8 tracking-tight">Corvi.com</h1>

        <button 
          @click="openLoginForm"
          class="bg-white text-black px-8 py-3 rounded-full font-medium text-lg hover:scale-105 transition-transform duration-300 focus:ring-4 focus:ring-white/30 outline-none mb-12 shadow-[0_0_20px_rgba(255,255,255,0.2)]"
        >
          Iniciar sessão
        </button>

        <h2 class="text-3xl md:text-4xl font-semibold text-white/90 leading-tight max-w-2xl mx-auto">
          Tenha total controle sobre seu <span class="text-blue-400">Estoque</span>, <span class="text-green-400">Patrimônio</span> e usuários do almoxarifado.
        </h2>
      </div>

      <div v-else class="relative z-10 w-full max-w-[380px] px-4" key="login">
        
        <button @click="closeLoginForm" class="absolute -top-12 left-4 text-white/50 hover:text-white transition-colors flex items-center gap-2 text-sm cursor-pointer">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" /></svg>
          Voltar
        </button>

        <div class="bg-[#1c1c1e]/90 backdrop-blur-md border border-white/10 rounded-2xl shadow-2xl p-8 pt-10 text-center">
          
          <div class="mb-6 flex justify-center">
            <svg class="h-10 w-10 text-white animate-pulse-slow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          
          <h3 class="text-xl font-semibold text-white mb-1">Bem-vindo de volta</h3>
          <p class="text-white/40 text-sm mb-6">Acesse o painel de controle Corvi.</p>

          <form @submit.prevent="handleLogin" class="space-y-4">
            <div class="relative group text-left">
              <input 
                id="email" v-model="email" type="email" required 
                class="w-full bg-[#2c2c2e] border border-transparent focus:border-blue-500/50 rounded-lg px-4 py-3.5 text-white placeholder-white/30 focus:ring-2 focus:ring-blue-500/20 focus:bg-[#3a3a3c] transition-all outline-none"
                placeholder="Email corporativo"
              />
            </div>

            <div class="relative group text-left">
              <input 
                id="password" v-model="password" :type="showPassword ? 'text' : 'password'" required 
                class="w-full bg-[#2c2c2e] border border-transparent focus:border-blue-500/50 rounded-lg px-4 py-3.5 text-white placeholder-white/30 focus:ring-2 focus:ring-blue-500/20 focus:bg-[#3a3a3c] transition-all outline-none"
                placeholder="Senha"
              />
               <button type="button" @click="showPassword = !showPassword" class="absolute right-3 top-3.5 text-white/30 hover:text-white cursor-pointer">
                 <svg v-if="showPassword" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" /></svg>
                 <svg v-else class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" /></svg>
               </button>
            </div>

            <div class="pt-2">
              <button type="button" @click="handleLogin" class="w-full bg-[#0A84FF] hover:bg-[#0077ED] text-white font-medium py-3 rounded-lg transition-colors flex items-center justify-center gap-2 shadow-lg shadow-blue-500/20 cursor-pointer">
                 <svg v-if="authStore.loading" class="animate-spin h-5 w-5 text-white" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
                 <span class="text-[17px] pr-2" v-else>Entrar no Sistema</span>
              </button>
            </div>
          </form>

           <div class="mt-6 border-t border-white/10 pt-4 flex justify-between text-xs">
              <a href="#" class="text-blue-400 hover:underline">Esqueceu a senha?</a>
              <a href="#" class="text-blue-400 hover:underline">Criar conta</a>
           </div>

        </div>
      </div>
    </transition>

    <div class="absolute bottom-6 text-center w-full">
       <p class="text-white/20 text-xs">Copyright © 2026 Corvi.com.br</p>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const showLoginForm = ref(false)

const openLoginForm = () => {
  showLoginForm.value = true
  window.history.pushState({ showLogin: true }, '')
}

const closeLoginForm = () => {
  showLoginForm.value = false
}

const handlePopState = (event) => {
  if (showLoginForm.value) {
    showLoginForm.value = false
  }
}

onMounted(() => {
  window.addEventListener('popstate', handlePopState)
})

onUnmounted(() => {
  window.removeEventListener('popstate', handlePopState)
})

const email = ref('')
const password = ref('')
const showPassword = ref(false)

const handleLogin = async () => {
  if (authStore.error) authStore.error = null
  const success = await authStore.login(email.value, password.value)
  if (success) {
    router.push('/')
  }
}
</script>

<style scoped>
/* Transição entre Landing Page e Login Form */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease, transform 0.5s ease;
}

.fade-enter-from {
  opacity: 0;
  transform: scale(0.95);
}
.fade-leave-to {
  opacity: 0;
  transform: scale(1.05);
}

/* Animações Flutuantes */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
}

@keyframes float-slow {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-15px) rotate(2deg); }
}

@keyframes float-delayed {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(10px); }
}

@keyframes spin-slow-reverse {
  from { transform: rotate(360deg); }
  to { transform: rotate(0deg); }
}

.animate-float {
  animation: float 3s ease-in-out infinite;
}
.animate-float-slow {
  animation: float-slow 5s ease-in-out infinite;
}
.animate-float-delayed {
  animation: float-delayed 4s ease-in-out infinite;
}
.animate-spin-slow-reverse {
  animation: spin-slow-reverse 20s linear infinite;
}

.animate-pulse-slow {
  animation: pulse 8s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
</style>