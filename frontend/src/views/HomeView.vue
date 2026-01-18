<template>
  <AppLayout>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      
      <div class="mb-10 flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 class="text-3xl font-bold text-slate-800 tracking-tight">
            Olá, Administrador 👋
          </h1>
          <p class="text-slate-500 mt-1">Aqui está o resumo do seu sistema hoje.</p>
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
              <h3 class="text-3xl font-bold text-slate-800">{{ formatCurrency(totalValue) }}</h3>
            </div>
            <div class="p-3 bg-emerald-50 rounded-xl text-emerald-600 group-hover:bg-emerald-600 group-hover:text-white transition-colors">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
          <div class="mt-4 text-xs text-slate-400">
            Soma de todos os itens em estoque
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <!-- Itens por Categoria - Card Power BI Style -->
        <div class="lg:col-span-2 bg-white rounded-3xl p-8 shadow-sm border border-slate-100">
          <div class="flex items-center justify-between mb-6">
            <div>
              <h3 class="text-xl font-bold text-slate-800">Itens por Categoria</h3>
              <p class="text-sm text-slate-500 mt-1">Distribuição do inventário</p>
            </div>
            <div class="p-3 bg-indigo-50 rounded-xl">
              <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
            </div>
          </div>

          <div v-if="categoryStats.length > 0" class="space-y-4">
            <div v-for="stat in categoryStats" :key="stat.category" class="group">
              <div class="flex items-center justify-between mb-2">
                <div class="flex items-center gap-2">
                  <div :class="getCategoryColorDot(stat.category)" class="w-3 h-3 rounded-full"></div>
                  <span class="text-sm font-semibold text-slate-700">{{ stat.category }}</span>
                </div>
                <div class="flex items-center gap-3">
                  <span class="text-xs text-slate-500">{{ stat.percentage }}%</span>
                  <span class="text-sm font-bold text-slate-800">{{ stat.count }} un.</span>
                </div>
              </div>
              <div class="relative h-3 bg-slate-100 rounded-full overflow-hidden">
                <div 
                  :class="getCategoryColorBar(stat.category)"
                  :style="{ width: stat.percentage + '%' }"
                  class="absolute h-full rounded-full transition-all duration-1000 ease-out group-hover:opacity-90"
                ></div>
              </div>
            </div>
          </div>

          <div v-else class="text-center py-12">
            <svg class="w-16 h-16 text-slate-300 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            <p class="text-slate-500 font-medium">Nenhum item cadastrado ainda</p>
            <p class="text-sm text-slate-400 mt-1">Adicione itens para visualizar a distribuição</p>
          </div>
        </div>

        <!-- Card de Gerenciar Inventário movido para baixo -->
        <div class="bg-white rounded-3xl p-6 shadow-sm border border-slate-100 flex flex-col">
          <h3 class="font-bold text-slate-800 mb-4">Acesso Rápido</h3>
          <div class="space-y-3 flex-1">
            <router-link 
              to="/items"
              class="w-full flex items-center p-3 rounded-xl bg-linear-to-r from-indigo-500 to-purple-600 text-white hover:shadow-lg transition-all hover:scale-[1.02]"
            >
              <div class="w-8 h-8 rounded-lg bg-white/20 flex items-center justify-center mr-3">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>
              <div class="text-left">
                <span class="block text-sm font-semibold">Inventário</span>
                <span class="text-xs opacity-90">Gerenciar itens</span>
              </div>
              <svg class="w-4 h-4 ml-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </router-link>

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
const totalItems = computed(() => {
  return itemsStore.items.reduce((sum, item) => sum + (item.quantity || 0), 0)
})

const lowStockItems = computed(() => {
  return itemsStore.items.filter(item => item.quantity < 5).length
})

const totalValue = computed(() => {
  return itemsStore.items.reduce((sum, item) => {
    const itemValue = (item.value || 0) * (item.quantity || 0)
    return sum + itemValue
  }, 0)
})

const formatCurrency = (value) => {
  return new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL'
  }).format(value)
}

// Estatísticas por categoria
const categoryStats = computed(() => {
  const categoryCount = {}
  
  itemsStore.items.forEach(item => {
    if (item.category) {
      categoryCount[item.category] = (categoryCount[item.category] || 0) + (item.quantity || 0)
    }
  })
  
  const total = totalItems.value
  
  return Object.entries(categoryCount)
    .map(([category, count]) => ({
      category,
      count,
      percentage: total > 0 ? Math.round((count / total) * 100) : 0
    }))
    .sort((a, b) => b.count - a.count)
})

const getCategoryColorDot = (category) => {
  const colors = {
    'Periféricos': 'bg-blue-600',
    'Hardware': 'bg-purple-600',
    'Cabos e Adaptadores': 'bg-yellow-600',
    'Monitores': 'bg-cyan-600',
    'Computadores': 'bg-indigo-600',
    'Notebooks': 'bg-pink-600',
    'Impressoras': 'bg-orange-600',
    'Rede': 'bg-green-600',
    'Armazenamento': 'bg-teal-600',
    'Outros': 'bg-slate-600'
  }
  return colors[category] || 'bg-slate-600'
}

const getCategoryColorBar = (category) => {
  const colors = {
    'Periféricos': 'bg-gradient-to-r from-blue-400 to-blue-600',
    'Hardware': 'bg-gradient-to-r from-purple-400 to-purple-600',
    'Cabos e Adaptadores': 'bg-gradient-to-r from-yellow-400 to-yellow-600',
    'Monitores': 'bg-gradient-to-r from-cyan-400 to-cyan-600',
    'Computadores': 'bg-gradient-to-r from-indigo-400 to-indigo-600',
    'Notebooks': 'bg-gradient-to-r from-pink-400 to-pink-600',
    'Impressoras': 'bg-gradient-to-r from-orange-400 to-orange-600',
    'Rede': 'bg-gradient-to-r from-green-400 to-green-600',
    'Armazenamento': 'bg-gradient-to-r from-teal-400 to-teal-600',
    'Outros': 'bg-gradient-to-r from-slate-400 to-slate-600'
  }
  return colors[category] || 'bg-gradient-to-r from-slate-400 to-slate-600'
}
</script>