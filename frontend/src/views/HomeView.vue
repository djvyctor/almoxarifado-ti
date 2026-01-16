<template>
  <AppLayout>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      
      <div class="mb-10 flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 class="text-3xl font-bold text-slate-800 tracking-tight">
            Olá, Administrador 👋
          </h1>
          <p class="text-slate-500 mt-1">Aqui está o resumo do seu almoxarifado hoje.</p>
        </div>
        <div class="hidden md:flex items-center gap-2 text-sm text-slate-400 bg-white px-4 py-2 rounded-full shadow-sm border border-slate-100">
          <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
          Sistema Operacional
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
        
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-slate-100 hover:shadow-lg transition-all duration-300 group">
          <div class="flex justify-between items-start">
            <div>
              <p class="text-sm font-medium text-slate-500 mb-1">Total de Itens</p>
              <h3 class="text-3xl font-bold text-slate-800">{{ totalItems }}</h3>
            </div>
            <div class="p-3 bg-indigo-50 rounded-xl text-indigo-600 group-hover:bg-indigo-600 group-hover:text-white transition-colors">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
          </div>
          <div class="mt-4 flex items-center text-xs text-slate-400">
            <span class="text-emerald-500 font-medium flex items-center mr-2">
              <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/></svg>
              Atualizado
            </span>
            agora mesmo
          </div>
        </div>

        <div class="bg-white rounded-2xl p-6 shadow-sm border border-slate-100 hover:shadow-lg transition-all duration-300 group">
          <div class="flex justify-between items-start">
            <div>
              <p class="text-sm font-medium text-slate-500 mb-1">Estoque Baixo</p>
              <h3 class="text-3xl font-bold text-slate-800">{{ lowStockItems }}</h3>
            </div>
            <div class="p-3 bg-orange-50 rounded-xl text-orange-600 group-hover:bg-orange-500 group-hover:text-white transition-colors">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
          </div>
          <div class="mt-4 text-xs text-slate-400">
            <span v-if="lowStockItems > 0" class="text-orange-500 font-medium">Requer atenção</span>
            <span v-else class="text-emerald-500 font-medium">Tudo certo</span>
             nos itens abaixo de 5 un.
          </div>
        </div>

        <div class="bg-white rounded-2xl p-6 shadow-sm border border-slate-100 hover:shadow-lg transition-all duration-300 group">
          <div class="flex justify-between items-start">
            <div>
              <p class="text-sm font-medium text-slate-500 mb-1">Valor Estimado</p>
              <h3 class="text-3xl font-bold text-slate-800">---</h3>
            </div>
            <div class="p-3 bg-emerald-50 rounded-xl text-emerald-600 group-hover:bg-emerald-600 group-hover:text-white transition-colors">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
          <div class="mt-4 text-xs text-slate-400">
            Funcionalidade em breve
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <div class="lg:col-span-2">
          <router-link 
            to="/items"
            class="relative overflow-hidden block bg-gradient-to-br from-indigo-600 to-purple-700 rounded-3xl p-8 shadow-xl hover:shadow-2xl hover:scale-[1.01] transition-all duration-300 group"
          >
            <div class="absolute top-0 right-0 -mr-20 -mt-20 w-64 h-64 rounded-full bg-white/10 blur-3xl group-hover:bg-white/20 transition-all"></div>
            <div class="absolute bottom-0 left-0 -ml-20 -mb-20 w-64 h-64 rounded-full bg-black/10 blur-3xl"></div>

            <div class="relative z-10 flex flex-col md:flex-row items-start md:items-center justify-between gap-6">
              <div>
                <h2 class="text-2xl font-bold text-white mb-2">Gerenciar Inventário</h2>
                <p class="text-indigo-100 max-w-md">Adicione novos equipamentos, atualize quantidades ou altere localizações do seu almoxarifado.</p>
                <div class="mt-6 inline-flex items-center px-4 py-2 bg-white/20 backdrop-blur-md rounded-lg text-white font-medium group-hover:bg-white group-hover:text-indigo-600 transition-colors">
                  Acessar Itens
                  <svg class="w-4 h-4 ml-2 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/></svg>
                </div>
              </div>
              <div class="bg-white/10 p-4 rounded-2xl backdrop-blur-sm border border-white/10">
                <svg class="w-16 h-16 text-white/80" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
                </svg>
              </div>
            </div>
          </router-link>
        </div>

        <div class="bg-white rounded-3xl p-6 shadow-sm border border-slate-100 flex flex-col">
          <h3 class="font-bold text-slate-800 mb-4">Acesso Rápido</h3>
          <div class="space-y-3 flex-1">
            <button disabled class="w-full flex items-center p-3 rounded-xl bg-slate-50 text-slate-400 cursor-not-allowed border border-transparent">
              <div class="w-8 h-8 rounded-lg bg-slate-200 flex items-center justify-center mr-3">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/></svg>
              </div>
              <div class="text-left">
                <span class="block text-sm font-semibold">Relatórios</span>
                <span class="text-xs">Em desenvolvimento</span>
              </div>
            </button>

            <button disabled class="w-full flex items-center p-3 rounded-xl bg-slate-50 text-slate-400 cursor-not-allowed border border-transparent">
              <div class="w-8 h-8 rounded-lg bg-slate-200 flex items-center justify-center mr-3">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/></svg>
              </div>
              <div class="text-left">
                <span class="block text-sm font-semibold">Configurações</span>
                <span class="text-xs">Em breve</span>
              </div>
            </button>
          </div>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useItemsStore } from '@/stores/items'
import AppLayout from '@/components/AppLayout.vue'

const itemsStore = useItemsStore()

// Carregar itens ao montar para ter os números atualizados
onMounted(() => {
  itemsStore.fetchItems()
})

// Cálculos automáticos baseados na Store
const totalItems = computed(() => itemsStore.items.length)

const lowStockItems = computed(() => {
  return itemsStore.items.filter(item => item.quantidade < 5).length
})
</script>