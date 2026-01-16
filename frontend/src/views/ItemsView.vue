<template>
  <AppLayout>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-8">
        <div>
          <h1 class="text-3xl font-bold text-slate-800 tracking-tight">Estoque</h1>
          <p class="text-slate-500 mt-1">Gerencie seu inventário de forma eficiente</p>
        </div>
        
        <div class="flex gap-3">
          <button
            @click="openModal()"
            class="inline-flex items-center px-5 py-2.5 rounded-xl text-white font-medium bg-gradient-to-r from-purple-600 to-indigo-600 hover:from-purple-700 hover:to-indigo-700 shadow-lg shadow-purple-500/20 transition-all hover:scale-105 active:scale-95"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Novo Item
          </button>
        </div>
      </div>

      <Transition
        enter-active-class="transition duration-300 ease-out"
        enter-from-class="transform -translate-y-2 opacity-0"
        enter-to-class="transform translate-y-0 opacity-100"
        leave-active-class="transition duration-200 ease-in"
        leave-from-class="transform translate-y-0 opacity-100"
        leave-to-class="transform -translate-y-2 opacity-0"
      >
        <div v-if="successMessage" class="mb-6 p-4 bg-green-50 border border-green-200 rounded-xl flex items-center text-green-700 shadow-sm">
          <svg class="w-5 h-5 mr-3 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
          </svg>
          {{ successMessage }}
        </div>
      </Transition>

      <div v-if="itemsStore.error" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-xl flex items-center text-red-700 shadow-sm">
        <svg class="w-5 h-5 mr-3 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
        </svg>
        {{ itemsStore.error }}
      </div>

      <div v-if="itemsStore.loading && items.length === 0" class="py-20 text-center">
        <div class="inline-block animate-spin rounded-full h-10 w-10 border-4 border-purple-100 border-t-purple-600"></div>
        <p class="mt-4 text-slate-500">Carregando inventário...</p>
      </div>

      <div v-else-if="items.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <div
          v-for="item in items"
          :key="item.id"
          class="group bg-white rounded-2xl border border-slate-200 p-5 hover:shadow-xl hover:shadow-slate-200/50 hover:border-purple-200 transition-all duration-300 flex flex-col"
        >
          <div class="flex justify-between items-start mb-4">
            <div class="p-2 bg-slate-50 rounded-lg group-hover:bg-purple-50 transition-colors">
              <svg class="w-6 h-6 text-slate-400 group-hover:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            <div :class="[
              'px-3 py-1 rounded-full text-xs font-bold border',
              item.quantidade < 5 
                ? 'bg-red-50 text-red-600 border-red-100' 
                : 'bg-emerald-50 text-emerald-600 border-emerald-100'
            ]">
              {{ item.quantidade }} un
            </div>
          </div>

          <h3 class="text-lg font-bold text-slate-800 mb-1 line-clamp-1" :title="item.nome">{{ item.nome }}</h3>
          <p class="text-sm text-slate-500 mb-4 line-clamp-2 min-h-[2.5rem]">{{ item.descricao || 'Sem descrição definida.' }}</p>
          
          <div class="flex items-center text-xs text-slate-400 font-medium mb-6">
            <svg class="w-4 h-4 mr-1.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            {{ item.localizacao || 'Sem local' }}
          </div>

          <div class="mt-auto flex gap-2 pt-4 border-t border-slate-100">
            <button
              @click="startEdit(item)"
              class="flex-1 flex items-center justify-center py-2 px-3 text-sm font-medium text-slate-600 bg-slate-50 rounded-lg hover:bg-white hover:text-indigo-600 hover:shadow border border-transparent hover:border-slate-200 transition-all"
            >
              Editar
            </button>
            <button
              @click="confirmDelete(item)"
              class="flex items-center justify-center p-2 text-slate-400 hover:text-red-500 hover:bg-red-50 rounded-lg transition-colors"
              title="Excluir"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-20 bg-white rounded-3xl border border-dashed border-slate-300">
        <div class="w-16 h-16 bg-slate-50 rounded-full flex items-center justify-center mx-auto mb-4">
          <svg class="w-8 h-8 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-slate-900">Nenhum item encontrado</h3>
        <p class="text-slate-500 mb-6">Comece adicionando itens ao seu almoxarifado.</p>
        <button
          @click="openModal()"
          class="text-indigo-600 font-medium hover:text-indigo-700 hover:underline"
        >
          Criar primeiro item &rarr;
        </button>
      </div>

      <Transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showModal" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="closeModal">
          
          <div class="bg-white rounded-2xl shadow-2xl w-full max-w-md overflow-hidden transform transition-all">
            <div class="px-6 py-4 border-b border-slate-100 flex justify-between items-center bg-slate-50/50">
              <h3 class="text-lg font-bold text-slate-800">
                {{ editingItem ? 'Editar Item' : 'Novo Item' }}
              </h3>
              <button @click="closeModal" class="text-slate-400 hover:text-slate-600">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>

            <div class="p-6">
              <form @submit.prevent="handleSubmit" class="space-y-5">
                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Nome do Item</label>
                  <input
                    v-model="form.nome"
                    type="text"
                    required
                    placeholder="Ex: Teclado Mecânico"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                  />
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Descrição</label>
                  <textarea
                    v-model="form.descricao"
                    rows="3"
                    placeholder="Detalhes sobre o item..."
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400 resize-none"
                  ></textarea>
                </div>

                <div class="grid grid-cols-2 gap-4">
                  <div>
                    <label class="block text-sm font-semibold text-slate-700 mb-2">Quantidade</label>
                    <input
                      v-model.number="form.quantidade"
                      type="number"
                      required
                      min="0"
                      class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-semibold text-slate-700 mb-2">Localização</label>
                    <input
                      v-model="form.localizacao"
                      type="text"
                      placeholder="Ex: Armário A2"
                      class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700"
                    />
                  </div>
                </div>

                <div class="pt-4 flex gap-3">
                  <button
                    type="button"
                    @click="closeModal"
                    class="flex-1 px-4 py-2.5 border border-slate-300 text-slate-700 font-medium rounded-xl hover:bg-slate-50 transition-colors"
                  >
                    Cancelar
                  </button>
                  <button
                    type="submit"
                    :disabled="itemsStore.loading"
                    class="flex-1 px-4 py-2.5 bg-indigo-600 text-white font-medium rounded-xl hover:bg-indigo-700 focus:ring-4 focus:ring-indigo-500/30 disabled:opacity-70 transition-all flex items-center justify-center"
                  >
                    <svg v-if="itemsStore.loading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    {{ itemsStore.loading ? 'Salvando...' : (editingItem ? 'Salvar Alterações' : 'Criar Item') }}
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </Transition>

    </div>
  </AppLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useItemsStore } from '@/stores/items'
import AppLayout from '@/components/AppLayout.vue'

const itemsStore = useItemsStore()

// State
const showModal = ref(false)
const editingItem = ref(null)
const successMessage = ref('')

const form = ref({
  nome: '',
  descricao: '',
  quantidade: 0,
  localizacao: ''
})

const items = computed(() => itemsStore.items)

onMounted(() => {
  itemsStore.fetchItems()
})

// Actions
const openModal = () => {
  editingItem.value = null
  resetForm()
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  setTimeout(() => {
    editingItem.value = null
    resetForm()
  }, 200) // Pequeno delay para a animação terminar
}

const resetForm = () => {
  form.value = {
    nome: '',
    descricao: '',
    quantidade: 0,
    localizacao: ''
  }
}

const handleSubmit = async () => {
  let success = false
  
  if (editingItem.value) {
    success = await itemsStore.updateItem(editingItem.value.id, form.value)
    if (success) showSuccess('Item atualizado com sucesso!')
  } else {
    success = await itemsStore.createItem(form.value)
    if (success) showSuccess('Item criado com sucesso!')
  }

  if (success) {
    closeModal()
  }
}

const startEdit = (item) => {
  editingItem.value = item
  form.value = {
    nome: item.nome,
    descricao: item.descricao,
    quantidade: item.quantidade,
    localizacao: item.localizacao
  }
  showModal.value = true
}

const confirmDelete = async (item) => {
  // Nota: Idealmente, use um Modal customizado aqui em vez de confirm() nativo para manter a UI consistente.
  if (confirm(`Tem certeza que deseja excluir "${item.nome}" do estoque?`)) {
    const success = await itemsStore.deleteItem(item.id)
    if (success) {
      showSuccess('Item removido do estoque.')
    }
  }
}

const showSuccess = (msg) => {
  successMessage.value = msg
  setTimeout(() => {
    successMessage.value = ''
  }, 3000)
}
</script>