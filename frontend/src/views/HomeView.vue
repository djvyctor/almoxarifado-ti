<template>
  <AppLayout>
    <div class="min-h-screen bg-black text-white font-sans selection:bg-blue-500/30 pb-12">
      
      <div class="fixed top-0 left-1/2 -translate-x-1/2 w-[1000px] h-[400px] bg-blue-900/10 rounded-full blur-[120px] pointer-events-none"></div>

      <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        
        <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10">
          <div>
            <h1 class="text-3xl font-semibold text-white tracking-tight">Visão Geral</h1>
            <p class="text-zinc-400 mt-1 text-sm">Resumo do seu inventário em tempo real.</p>
          </div>
          
          <div class="flex items-center gap-2 px-3 py-1.5 rounded-full bg-[#1c1c1e] border border-white/5 shadow-sm">
            <span class="relative flex h-2 w-2">
              <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-500 opacity-75"></span>
              <span class="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
            </span>
            <span class="text-xs font-medium text-zinc-300">Sistema Online</span>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
          
          <div class="bg-[#1c1c1e] p-6 rounded-3xl border border-white/5 hover:bg-[#252528] transition-colors duration-300 flex flex-col justify-between h-40 group">
            <div class="flex justify-between items-start">
              <div class="w-10 h-10 rounded-full bg-blue-500/10 flex items-center justify-center text-blue-500 group-hover:bg-blue-500 group-hover:text-white transition-all duration-300">
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" /></svg>
              </div>
              <span class="text-xs font-medium text-zinc-500 bg-white/5 px-2 py-1 rounded-md">Total</span>
            </div>
            <div>
              <h3 class="text-3xl font-semibold text-white tracking-tight">{{ totalItems }}</h3>
              <p class="text-sm text-zinc-500">Itens cadastrados</p>
            </div>
          </div>

          <div class="bg-[#1c1c1e] p-6 rounded-3xl border border-white/5 hover:bg-[#252528] transition-colors duration-300 flex flex-col justify-between h-40 group">
            <div class="flex justify-between items-start">
              <div class="w-10 h-10 rounded-full bg-orange-500/10 flex items-center justify-center text-orange-500 group-hover:bg-orange-500 group-hover:text-white transition-all duration-300">
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
              </div>
              <span class="text-xs font-medium text-zinc-500 bg-white/5 px-2 py-1 rounded-md">Atenção</span>
            </div>
            <div>
              <h3 class="text-3xl font-semibold text-white tracking-tight">{{ lowStockItems }}</h3>
              <p class="text-sm text-zinc-500">Itens com estoque baixo</p>
            </div>
          </div>

          <div class="bg-[#1c1c1e] p-6 rounded-3xl border border-white/5 hover:bg-[#252528] transition-colors duration-300 flex flex-col justify-between h-40 group">
            <div class="flex justify-between items-start">
              <div class="w-10 h-10 rounded-full bg-emerald-500/10 flex items-center justify-center text-emerald-500 group-hover:bg-emerald-500 group-hover:text-white transition-all duration-300">
                <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
              </div>
              <span class="text-xs font-medium text-zinc-500 bg-white/5 px-2 py-1 rounded-md">Patrimônio</span>
            </div>
            <div>
              <h3 class="text-3xl font-semibold text-white tracking-tight truncate">{{ formatCurrency(totalValue) }}</h3>
              <p class="text-sm text-zinc-500">Valor total estimado</p>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-4 gap-4">
          
          <div class="lg:col-span-3 bg-[#1c1c1e] rounded-3xl p-8 border border-white/5">
            <div class="flex items-center justify-between mb-8">
              <h3 class="text-lg font-semibold text-white">Categorias</h3>
              <button class="text-xs text-blue-400 hover:text-blue-300 font-medium transition-colors">Ver todas</button>
            </div>

            <div v-if="categoryStats.length > 0" class="grid grid-cols-1 sm:grid-cols-2 gap-x-12 gap-y-6">
              <div v-for="stat in categoryStats" :key="stat.category" class="group">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center gap-3">
                    <div :class="getCategoryColorDot(stat.category)" class="w-2 h-2 rounded-full"></div>
                    <span class="text-sm text-zinc-300 font-medium">{{ stat.category }}</span>
                  </div>
                  <span class="text-sm text-zinc-500">{{ stat.count }}</span>
                </div>
                <div class="relative h-1 w-full bg-[#2c2c2e] rounded-full overflow-hidden">
                  <div 
                    :class="getCategoryColorBar(stat.category)"
                    :style="{ width: stat.percentage + '%' }"
                    class="absolute h-full rounded-full transition-all duration-1000 ease-out"
                  ></div>
                </div>
              </div>
            </div>

            <div v-else class="flex flex-col items-center justify-center py-12 text-zinc-600">
              <svg class="w-12 h-12 mb-3 opacity-20" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M4 12h16m-7 6h7" /></svg>
              <p class="text-sm">Nenhum dado para exibir</p>
            </div>
          </div>

          <div class="flex flex-col gap-4">
            
            <router-link to="/items" class="bg-[#0071e3] hover:bg-[#0077ED] rounded-3xl p-6 flex flex-col items-start justify-between h-48 transition-all duration-300 shadow-lg shadow-blue-900/20 group relative overflow-hidden">
               <div class="absolute -right-4 -top-4 w-24 h-24 bg-white/10 rounded-full blur-2xl group-hover:bg-white/20 transition-all"></div>
               
               <div class="w-10 h-10 bg-white/20 backdrop-blur-sm rounded-full flex items-center justify-center">
                 <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" /></svg>
               </div>
               
               <div>
                 <h3 class="text-xl font-bold text-white mb-1">Acessar Inventário</h3>
                 <p class="text-sm text-blue-100">Gerenciar itens e ativos</p>
               </div>
            </router-link>

            <div class="flex-1 bg-[#1c1c1e] rounded-3xl p-6 border border-white/5 flex flex-col justify-center gap-4 opacity-60">
               <div class="flex items-center gap-3 text-zinc-500 cursor-not-allowed">
                  <div class="w-8 h-8 rounded-full bg-[#2c2c2e] flex items-center justify-center">
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" /></svg>
                  </div>
                  <span class="text-sm font-medium">Relatórios</span>
               </div>
               <div class="h-px bg-white/5 w-full"></div>
               <div class="flex items-center gap-3 text-zinc-500 cursor-not-allowed">
                  <div class="w-8 h-8 rounded-full bg-[#2c2c2e] flex items-center justify-center">
                    <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/></svg>
                  </div>
                  <span class="text-sm font-medium">Configurações</span>
               </div>
            </div>

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

onMounted(() => {
  itemsStore.fetchItems()
})

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

// Cores Sólidas e Foscas (Sem brilho Neon)
const getCategoryColorDot = (category) => {
  const colors = {
    'Periféricos': 'bg-blue-500',
    'Hardware': 'bg-purple-500',
    'Cabos e Adaptadores': 'bg-yellow-500',
    'Monitores': 'bg-cyan-500',
    'Computadores': 'bg-indigo-500',
    'Notebooks': 'bg-pink-500',
    'Impressoras': 'bg-orange-500',
    'Rede': 'bg-green-500',
    'Armazenamento': 'bg-teal-500',
    'Outros': 'bg-zinc-500'
  }
  return colors[category] || 'bg-zinc-500'
}

// Barras de progresso com cor sólida
const getCategoryColorBar = (category) => {
  const colors = {
    'Periféricos': 'bg-blue-500',
    'Hardware': 'bg-purple-500',
    'Cabos e Adaptadores': 'bg-yellow-500',
    'Monitores': 'bg-cyan-500',
    'Computadores': 'bg-indigo-500',
    'Notebooks': 'bg-pink-500',
    'Impressoras': 'bg-orange-500',
    'Rede': 'bg-green-500',
    'Armazenamento': 'bg-teal-500',
    'Outros': 'bg-zinc-500'
  }
  return colors[category] || 'bg-zinc-500'
}
</script>