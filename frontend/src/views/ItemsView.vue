<template>
  <AppLayout>
    <div class="min-h-screen bg-black text-white font-sans selection:bg-blue-500/30 pb-20 relative overflow-hidden">
      
      <div class="fixed top-0 left-1/2 -translate-x-1/2 w-[1000px] h-[400px] bg-blue-900/10 rounded-full blur-[120px] pointer-events-none z-0"></div>

      <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        
        <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10">
          <div>
            <h1 class="text-3xl font-semibold text-white tracking-tight">Estoque</h1>
            <p class="text-zinc-400 mt-1 text-sm">Gerenciamento completo de ativos e periféricos.</p>
          </div>
          
          <button
            @click="openModal()"
            class="inline-flex items-center px-5 py-2.5 rounded-full text-white text-sm font-medium bg-[#0071e3] hover:bg-[#0077ED] transition-all duration-300 shadow-lg shadow-blue-900/20 hover:scale-105 active:scale-95"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Novo Item
          </button>
        </div>

        <Transition
          enter-active-class="transition duration-300 ease-out"
          enter-from-class="transform -translate-y-2 opacity-0"
          enter-to-class="transform translate-y-0 opacity-100"
          leave-active-class="transition duration-200 ease-in"
          leave-from-class="transform translate-y-0 opacity-100"
          leave-to-class="transform -translate-y-2 opacity-0"
        >
          <div v-if="successMessage" class="mb-6 p-4 bg-emerald-500/10 border border-emerald-500/20 rounded-2xl flex items-center text-emerald-400 shadow-sm backdrop-blur-sm">
            <svg class="w-5 h-5 mr-3 shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
            </svg>
            {{ successMessage }}
          </div>
        </Transition>

        <div v-if="itemsStore.error" class="mb-6 p-4 bg-red-500/10 border border-red-500/20 rounded-2xl flex items-center text-red-400 shadow-sm backdrop-blur-sm">
          <svg class="w-5 h-5 mr-3 shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
          </svg>
          {{ itemsStore.error }}
        </div>

        <div v-if="itemsStore.loading && items.length === 0" class="py-24 text-center">
          <div class="inline-block animate-spin rounded-full h-10 w-10 border-4 border-white/10 border-t-blue-500"></div>
          <p class="mt-4 text-zinc-500">Sincronizando inventário...</p>
        </div>

        <div v-else-if="itemsStore.items.length > 0">
          
          <div class="bg-[#1c1c1e] rounded-3xl border border-white/5 p-4 mb-8 shadow-2xl">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="relative group">
                <svg class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-zinc-500 group-focus-within:text-blue-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <input
                  v-model="searchQuery"
                  type="text"
                  placeholder="Buscar item..."
                  class="w-full pl-12 pr-4 py-3 bg-[#2c2c2e] border-none rounded-2xl text-white placeholder-zinc-600 focus:ring-2 focus:ring-blue-500/50 transition-all"
                />
              </div>
              
              <div class="relative group">
                <svg class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-zinc-500 group-focus-within:text-blue-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
                </svg>
                <select
                  v-model="categoryFilter"
                  class="w-full pl-12 pr-10 py-3 bg-[#2c2c2e] border-none rounded-2xl text-white appearance-none cursor-pointer focus:ring-2 focus:ring-blue-500/50 transition-all"
                >
                  <option value="" class="bg-[#2c2c2e]">Todas as categorias</option>
                  <option v-for="cat in availableCategories" :key="cat" :value="cat" class="bg-[#2c2c2e]">{{ cat }}</option>
                </select>
                <div class="absolute right-4 top-1/2 transform -translate-y-1/2 pointer-events-none">
                  <svg class="w-4 h-4 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
                </div>
              </div>
            </div>
          </div>

          <div v-if="items.length > 0" class="space-y-3">
            <div
              v-for="item in items"
              :key="item.id"
              class="group bg-[#1c1c1e] rounded-2xl border border-white/5 p-4 hover:bg-[#252528] transition-all duration-300 flex items-center gap-5 relative overflow-hidden"
            >
              <div class="absolute left-0 top-0 bottom-0 w-1 bg-blue-500 opacity-0 group-hover:opacity-100 transition-opacity"></div>

              <div class="p-3 bg-white/5 rounded-2xl group-hover:bg-blue-500/10 group-hover:text-blue-400 transition-colors shrink-0 text-zinc-400">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>

              <div class="flex-1 min-w-0">
                <div class="flex items-center gap-3 mb-1">
                  <h3 class="text-base font-semibold text-white truncate">{{ item.name }}</h3>
                  <div v-if="item.quantity < 5" class="w-2 h-2 rounded-full bg-red-500 animate-pulse" title="Estoque Baixo"></div>
                </div>
                
                <p v-if="item.description" class="text-sm text-zinc-500 mb-3 line-clamp-1">{{ item.description }}</p>
                
                <div class="flex items-center gap-3 text-xs flex-wrap">
                  <div :class="getCategoryColor(item.category)" class="flex items-center gap-1.5 px-2.5 py-1 rounded-md font-medium border">
                    <span>{{ item.category }}</span>
                  </div>

                  <div v-if="item.patrimony" class="flex items-center gap-1.5 text-zinc-400 bg-white/5 px-2.5 py-1 rounded-md border border-white/5">
                    <span class="text-zinc-600">#</span>
                    <span>{{ item.patrimony }}</span>
                  </div>

                  <div v-if="item.value" class="flex items-center gap-1.5 text-emerald-400 bg-emerald-500/5 px-2.5 py-1 rounded-md border border-emerald-500/10">
                    <span>R$ {{ ((item.value || 0) * (item.quantity || 0)).toFixed(2).replace('.', ',') }}</span>
                  </div>

                  <div v-if="item.assigned_to" class="flex items-center gap-1.5 text-indigo-400 bg-indigo-500/5 px-2.5 py-1 rounded-md border border-indigo-500/10">
                    <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" /></svg>
                    <span>{{ item.assigned_to }}</span>
                  </div>
                </div>
              </div>

              <div class="shrink-0 text-right mr-2">
                <div class="text-lg font-bold" :class="item.quantity < 5 ? 'text-red-400' : 'text-white'">
                  {{ item.quantity }}
                </div>
                <div class="text-[10px] text-zinc-500 uppercase tracking-wider font-medium">Unid</div>
              </div>

              <div class="flex gap-2 shrink-0 opacity-100 sm:opacity-0 sm:group-hover:opacity-100 transition-opacity">
                <button
                  @click="startEdit(item)"
                  class="p-2 text-zinc-400 hover:text-white hover:bg-white/10 rounded-lg transition-all"
                  title="Editar"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  @click="confirmDelete(item)"
                  class="p-2 text-zinc-400 hover:text-red-400 hover:bg-red-500/10 rounded-lg transition-all"
                  title="Excluir"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>

          <div v-else class="text-center py-20 bg-[#1c1c1e] rounded-3xl border border-white/5 border-dashed">
            <svg class="w-16 h-16 text-zinc-700 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
            <p class="text-zinc-500">Nenhum item encontrado com os filtros aplicados</p>
          </div>
        </div>

        <div v-else class="text-center py-24 bg-[#1c1c1e] rounded-[2.5rem] border border-white/5 border-dashed">
          <div class="w-20 h-20 bg-white/5 rounded-full flex items-center justify-center mx-auto mb-6">
            <svg class="w-10 h-10 text-zinc-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <h3 class="text-xl font-semibold text-white mb-2">Inventário vazio</h3>
          <p class="text-zinc-500 mb-8 max-w-sm mx-auto">Comece a organizar seu patrimônio adicionando o primeiro item ao sistema.</p>
          <button
            @click="openModal()"
            class="text-blue-400 font-medium hover:text-blue-300 hover:underline"
          >
            Adicionar item agora &rarr;
          </button>
        </div>

        <Transition
          enter-active-class="transition duration-200 ease-out"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition duration-150 ease-in"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div v-if="showModal" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="closeModal">
            
            <div class="bg-[#1c1c1e] border border-white/10 rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-hidden transform transition-all flex flex-col">
              <div class="px-6 py-5 border-b border-white/10 flex justify-between items-center bg-[#242426]">
                <h3 class="text-lg font-bold text-white">
                  {{ editingItem ? 'Editar Item' : 'Novo Item' }}
                </h3>
                <button @click="closeModal" class="text-zinc-500 hover:text-white transition-colors">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>

              <div class="p-8 overflow-y-auto flex-1 custom-scrollbar">
                <form @submit.prevent="handleSubmit" class="space-y-6">
                  <div>
                    <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Nome do Item</label>
                    <input
                      v-model="form.name"
                      @input="capitalizeFirstLetter"
                      type="text"
                      required
                      placeholder="Ex: Teclado Mecânico"
                      class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                    />
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div>
                      <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Categoria</label>
                      <div class="relative">
                        <input
                          v-model="categorySearch"
                          @input="searchCategories"
                          @focus="showCategorySuggestions = true"
                          type="text"
                          required
                          placeholder="Selecionar..."
                          class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                        />
                        
                        <div v-if="showCategorySuggestions && filteredCategories.length > 0" class="absolute z-10 w-full mt-1 bg-[#2c2c2e] border border-white/10 rounded-xl shadow-xl max-h-48 overflow-y-auto">
                          <button
                            v-for="category in filteredCategories"
                            :key="category"
                            type="button"
                            @click="selectCategory(category)"
                            class="w-full px-4 py-2.5 text-left hover:bg-blue-500 hover:text-white transition-colors flex items-center gap-2 text-zinc-300"
                          >
                            <span>{{ category }}</span>
                          </button>
                        </div>
                      </div>
                    </div>

                    <div>
                      <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Quantidade</label>
                      <input
                        v-model.number="form.quantity"
                        type="number"
                        required
                        min="0"
                        class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                      />
                    </div>
                  </div>

                  <div>
                    <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Observações</label>
                    <textarea
                      v-model="form.description"
                      rows="3"
                      :placeholder="getDescriptionPlaceholder()"
                      class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none resize-none"
                    ></textarea>
                    <p class="text-xs text-zinc-600 mt-2">{{ getDescriptionHint() }}</p>
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                     <div>
                      <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Patrimônio</label>
                      <input
                        v-model="form.patrimony"
                        type="text"
                        maxlength="5"
                        @input="formatPatrimony"
                        placeholder="00000"
                        class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                      />
                    </div>
                    
                    <div>
                      <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Valor (R$)</label>
                      <div class="relative">
                        <span class="absolute left-4 top-1/2 transform -translate-y-1/2 text-zinc-500 font-medium">R$</span>
                        <input
                          v-model="form.value"
                          type="text"
                          @input="formatValue"
                          placeholder="0,00"
                          class="w-full pl-10 pr-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                        />
                      </div>
                    </div>
                  </div>

                  <div>
                    <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Vincular a Usuário</label>
                    <div class="relative">
                      <input
                        v-model="userSearch"
                        @input="searchUsers"
                        @focus="showUserSuggestions = true"
                        type="text"
                        placeholder="Buscar usuário..."
                        class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                      />
                       <div v-if="showUserSuggestions && filteredUsers.length > 0" class="absolute z-10 w-full mt-1 bg-[#2c2c2e] border border-white/10 rounded-xl shadow-xl max-h-48 overflow-y-auto">
                        <button
                          v-for="user in filteredUsers"
                          :key="user.id"
                          type="button"
                          @click="selectUser(user)"
                          class="w-full px-4 py-3 text-left hover:bg-blue-500 hover:text-white transition-colors border-b border-white/5 last:border-b-0"
                        >
                          <div class="text-sm font-medium text-white">{{ user.name }}</div>
                          <div class="text-xs text-zinc-500" :class="{'text-blue-200': false}">{{ user.email }}</div>
                        </button>
                      </div>
                    </div>
                  </div>

                  <div class="pt-6 flex gap-3">
                    <button
                      type="button"
                      @click="closeModal"
                      class="flex-1 px-4 py-3 border border-white/10 text-zinc-300 font-medium rounded-xl hover:bg-white/5 transition-colors"
                    >
                      Cancelar
                    </button>
                    <button
                      type="submit"
                      :disabled="itemsStore.loading"
                      class="flex-1 px-4 py-3 bg-[#0071e3] text-white font-medium rounded-xl hover:bg-[#0077ED] disabled:opacity-70 transition-all shadow-lg shadow-blue-900/20"
                    >
                      {{ itemsStore.loading ? 'Salvando...' : (editingItem ? 'Salvar Alterações' : 'Criar Item') }}
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </Transition>

        <Transition
          enter-active-class="transition duration-200 ease-out"
          enter-from-class="opacity-0 scale-95"
          enter-to-class="opacity-100 scale-100"
          leave-active-class="transition duration-150 ease-in"
          leave-from-class="opacity-100 scale-100"
          leave-to-class="opacity-0 scale-95"
        >
          <div v-if="showDeleteModal" class="fixed inset-0 bg-black/80 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="cancelDelete">
            <div class="bg-[#1c1c1e] border border-white/10 rounded-2xl shadow-2xl w-full max-w-sm overflow-hidden p-6 text-center">
              <div class="w-14 h-14 bg-red-500/10 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-7 h-7 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </div>
              <h3 class="text-lg font-bold text-white mb-2">Excluir Item?</h3>
              <p class="text-zinc-400 mb-6 text-sm">
                Tem certeza que deseja remover <span class="text-white font-medium">"{{ itemToDelete?.name }}"</span>? Essa ação não pode ser desfeita.
              </p>
              
              <div class="flex gap-3">
                <button
                  @click="cancelDelete"
                  class="flex-1 px-4 py-2.5 border border-white/10 text-zinc-300 font-medium rounded-xl hover:bg-white/5 transition-all"
                >
                  Cancelar
                </button>
                <button
                  @click="executeDelete"
                  :disabled="itemsStore.loading"
                  class="flex-1 px-4 py-2.5 bg-red-500/10 text-red-500 border border-red-500/20 font-medium rounded-xl hover:bg-red-500 hover:text-white transition-all"
                >
                  {{ itemsStore.loading ? '...' : 'Excluir' }}
                </button>
              </div>
            </div>
          </div>
        </Transition>

      </div>
    </div>
  </AppLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useItemsStore } from '@/stores/items'
import { useUsersStore } from '@/stores/users'
import AppLayout from '@/components/AppLayout.vue'

const itemsStore = useItemsStore()
const usersStore = useUsersStore()

// State
const showModal = ref(false)
const showDeleteModal = ref(false)
const editingItem = ref(null)
const itemToDelete = ref(null)
const successMessage = ref('')
const searchQuery = ref('')
const categoryFilter = ref('')

const form = ref({
  name: '',
  category: '',
  description: '',
  patrimony: '',
  assigned_to: '',
  value: '',
  quantity: 0
})

// Estados de Autocomplete
const userSearch = ref('')
const showUserSuggestions = ref(false)
const categorySearch = ref('')
const showCategorySuggestions = ref(false)
const availableCategories = ref([
  'Periféricos', 'Hardware', 'Cabos e Adaptadores', 'Monitores',
  'Computadores', 'Notebooks', 'Impressoras', 'Rede', 'Armazenamento', 'Outros'
])

// Computed Properties
const filteredCategories = computed(() => {
  if (!categorySearch.value) return availableCategories.value
  const search = categorySearch.value.toLowerCase()
  return availableCategories.value.filter(category => 
    category.toLowerCase().includes(search)
  )
})

const filteredUsers = computed(() => {
  if (!userSearch.value) return []
  const search = userSearch.value.toLowerCase()
  return usersStore.users.filter(user => 
    user.name.toLowerCase().includes(search) ||
    (user.email && user.email.toLowerCase().includes(search))
  )
})

// Nova Lógica de Cores para o Dark Mode (Translúcidos)
const getCategoryColor = (category) => {
  const colors = {
    'Periféricos': 'bg-blue-500/10 text-blue-400 border-blue-500/20',
    'Hardware': 'bg-purple-500/10 text-purple-400 border-purple-500/20',
    'Cabos e Adaptadores': 'bg-yellow-500/10 text-yellow-400 border-yellow-500/20',
    'Monitores': 'bg-cyan-500/10 text-cyan-400 border-cyan-500/20',
    'Computadores': 'bg-indigo-500/10 text-indigo-400 border-indigo-500/20',
    'Notebooks': 'bg-pink-500/10 text-pink-400 border-pink-500/20',
    'Impressoras': 'bg-orange-500/10 text-orange-400 border-orange-500/20',
    'Rede': 'bg-green-500/10 text-green-400 border-green-500/20',
    'Armazenamento': 'bg-teal-500/10 text-teal-400 border-teal-500/20',
    'Outros': 'bg-zinc-500/10 text-zinc-400 border-zinc-500/20'
  }
  return colors[category] || 'bg-zinc-500/10 text-zinc-400 border-zinc-500/20'
}

// Lógica de Placeholder Dinâmico
const getDescriptionPlaceholder = () => {
  const placeholders = {
    'Cabos e Adaptadores': 'Ex: RJ45, 10 metros, Cat6',
    'Monitores': 'Ex: 24 polegadas, Full HD, HDMI',
    'Hardware': 'Ex: Processador i5, 8GB RAM',
    'Periféricos': 'Ex: USB, Wireless, RGB',
    'Computadores': 'Ex: Dell OptiPlex, i7, 16GB',
    'Notebooks': 'Ex: Lenovo ThinkPad, 15.6", SSD 512GB',
    'Impressoras': 'Ex: Laser, Colorida, Duplex',
    'Rede': 'Ex: Switch 24 portas, Gigabit',
    'Armazenamento': 'Ex: SSD 1TB, SATA III',
    'Outros': 'Descreva os detalhes do item'
  }
  return placeholders[form.value.category] || 'Detalhes do item...'
}

const getDescriptionHint = () => {
  const hints = {
    'Cabos e Adaptadores': 'Tipo, comprimento e spec.',
    'Monitores': 'Tamanho, resolução e portas.',
    'Hardware': 'Modelo e detalhes técnicos.',
    'Periféricos': 'Conexão e layout.',
    'Computadores': 'Configuração completa.',
    'Notebooks': 'Modelo e specs.',
    'Impressoras': 'Tipo e suprimentos.',
    'Rede': 'Velocidade e portas.',
    'Armazenamento': 'Interface e capacidade.',
    'Outros': 'Info relevante.'
  }
  return hints[form.value.category] || 'Adicione informações relevantes.'
}

// Helpers de Formatação e Busca
const searchCategories = () => {
  showCategorySuggestions.value = true
  form.value.category = categorySearch.value
}

const selectCategory = (category) => {
  categorySearch.value = category
  form.value.category = category
  showCategorySuggestions.value = false
}

const searchUsers = () => {
  showUserSuggestions.value = userSearch.value.length > 0
}

const selectUser = (user) => {
  userSearch.value = user.name
  form.value.assigned_to = user.name
  showUserSuggestions.value = false
}

const formatPatrimony = (event) => {
  let value = event.target.value.replace(/\D/g, '')
  value = value.slice(0, 5)
  form.value.patrimony = value
}

const capitalizeFirstLetter = (event) => {
  const value = event.target.value
  if (value.length > 0) {
    form.value.name = value.charAt(0).toUpperCase() + value.slice(1)
  }
}

const formatValue = (event) => {
  let value = event.target.value.replace(/\D/g, '')
  if (value === '') {
    form.value.value = ''
    return
  }
  const numValue = parseInt(value) / 100
  form.value.value = numValue.toFixed(2).replace('.', ',')
}

const items = computed(() => {
  let filtered = itemsStore.items
  if (searchQuery.value) {
    filtered = filtered.filter(item => 
      item.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      item.category.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  if (categoryFilter.value) {
    filtered = filtered.filter(item => item.category === categoryFilter.value)
  }
  return filtered
})

onMounted(() => {
  itemsStore.fetchItems()
  usersStore.fetchUsers()
})

// Modal Actions
const openModal = () => {
  editingItem.value = null
  resetForm()
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  showUserSuggestions.value = false
  showCategorySuggestions.value = false
  setTimeout(() => {
    editingItem.value = null
    resetForm()
  }, 200)
}

const resetForm = () => {
  form.value = {
    name: '', category: '', description: '', patrimony: '',
    assigned_to: '', value: '', quantity: 0
  }
  userSearch.value = ''
  categorySearch.value = ''
}

const handleSubmit = async () => {
  let success = false
  const submitData = {
    ...form.value,
    value: form.value.value ? parseFloat(form.value.value.replace(',', '.')) : null
  }
  
  if (editingItem.value) {
    success = await itemsStore.updateItem(editingItem.value.id, submitData)
    if (success) showSuccess('Item atualizado com sucesso!')
  } else {
    success = await itemsStore.createItem(submitData)
    if (success) showSuccess('Item criado com sucesso!')
  }

  if (success) {
    closeModal()
  }
}

const startEdit = (item) => {
  editingItem.value = item
  form.value = {
    name: item.name,
    category: item.category,
    description: item.description || '',
    patrimony: item.patrimony || '',
    assigned_to: item.assigned_to || '',
    value: item.value ? item.value.toFixed(2).replace('.', ',') : '',
    quantity: item.quantity
  }
  userSearch.value = item.assigned_to || ''
  categorySearch.value = item.category || ''
  showModal.value = true
}

const confirmDelete = (item) => {
  itemToDelete.value = item
  showDeleteModal.value = true
}

const cancelDelete = () => {
  showDeleteModal.value = false
  setTimeout(() => {
    itemToDelete.value = null
  }, 200)
}

const executeDelete = async () => {
  if (!itemToDelete.value) return
  const success = await itemsStore.deleteItem(itemToDelete.value.id)
  if (success) {
    showSuccess('Item removido do estoque.')
    cancelDelete()
  }
}

const showSuccess = (msg) => {
  successMessage.value = msg
  setTimeout(() => {
    successMessage.value = ''
  }, 3000)
}
</script>

<style scoped>
/* Scrollbar para o modal dark */
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: #1c1c1e;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #3f3f46;
  border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #52525b;
}
</style>