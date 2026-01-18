<template>
  <AppLayout>
    <div class="min-h-screen bg-black text-white font-sans selection:bg-blue-500/30 relative overflow-hidden flex flex-col items-center justify-center">
      
      <div class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-indigo-900/20 rounded-full blur-[120px] pointer-events-none z-0"></div>

      <div class="relative z-10 max-w-2xl w-full px-6">
        
        <div class="bg-[#1c1c1e] border border-white/5 rounded-[2.5rem] p-12 text-center shadow-2xl relative overflow-hidden group">
          
          <div class="absolute top-0 left-0 right-0 h-1 bg-gradient-to-r from-transparent via-white/10 to-transparent opacity-50"></div>

          <div class="w-24 h-24 bg-gradient-to-tr from-[#0071e3] to-indigo-600 rounded-3xl flex items-center justify-center mx-auto mb-8 shadow-lg shadow-blue-900/30 group-hover:scale-105 transition-transform duration-500">
            <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>

          <h1 class="text-4xl font-bold text-white mb-2 tracking-tight">Corvi System</h1>
          <p class="text-lg text-zinc-400 mb-8 font-light">Gestão Inteligente de Almoxarifado & TI</p>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8 text-left">
            <div class="bg-[#2c2c2e] p-4 rounded-2xl border border-white/5">
              <p class="text-xs text-zinc-500 uppercase tracking-wider font-semibold mb-1">Versão</p>
              <p class="text-white font-medium">1.0.0 (Beta)</p>
            </div>
            <div class="bg-[#2c2c2e] p-4 rounded-2xl border border-white/5">
              <p class="text-xs text-zinc-500 uppercase tracking-wider font-semibold mb-1">Desenvolvedor</p>
              <p class="text-white font-medium">Victor Medeiros</p>
            </div>
          </div>

          <div class="pt-6 border-t border-white/5">
            <p class="text-sm text-zinc-500 mb-4">Tecnologias Utilizadas</p>
            <div class="flex flex-wrap justify-center gap-2">
              <span class="px-3 py-1 rounded-full bg-emerald-500/10 text-emerald-400 text-xs font-medium border border-emerald-500/20">Vue.js 3</span>
              <span class="px-3 py-1 rounded-full bg-cyan-500/10 text-cyan-400 text-xs font-medium border border-cyan-500/20">Tailwind CSS</span>
              <span class="px-3 py-1 rounded-full bg-blue-500/10 text-blue-400 text-xs font-medium border border-blue-500/20">Go (Gin)</span>
              <span class="px-3 py-1 rounded-full bg-yellow-500/10 text-yellow-400 text-xs font-medium border border-yellow-500/20">Vite</span>
            </div>
          </div>

        </div>

        <div class="text-center mt-8">
          <p class="text-zinc-600 text-sm">© 2026 Corvi Inc. Todos os direitos reservados.</p>
        </div>

      </div>
    </div>
  </AppLayout>
</template>

<script setup>
import AppLayout from '@/components/AppLayout.vue'
</script>