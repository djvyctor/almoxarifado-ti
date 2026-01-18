<template>
  <AppLayout>
    <div class="min-h-screen bg-black text-white font-sans selection:bg-blue-500/30 pb-20 relative overflow-hidden">
      
      <div class="fixed top-0 left-1/2 -translate-x-1/2 w-[1000px] h-[400px] bg-purple-900/10 rounded-full blur-[120px] pointer-events-none z-0"></div>

      <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        
        <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10">
          <div>
            <h1 class="text-3xl font-semibold text-white tracking-tight">Usuários</h1>
            <p class="text-zinc-400 mt-1 text-sm">Gestão de colaboradores e acessos ao sistema.</p>
          </div>
          
          <button
            @click="openModal()"
            class="inline-flex items-center px-5 py-2.5 rounded-full text-white text-sm font-medium bg-[#0071e3] hover:bg-[#0077ED] transition-all duration-300 shadow-lg shadow-blue-900/20 hover:scale-105 active:scale-95"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Novo Usuário
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

        <div v-if="usersStore.error" class="mb-6 p-4 bg-red-500/10 border border-red-500/20 rounded-2xl flex items-center text-red-400 shadow-sm backdrop-blur-sm">
          <svg class="w-5 h-5 mr-3 shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
          </svg>
          {{ usersStore.error }}
        </div>

        <div v-if="usersStore.loading && usersStore.users.length === 0" class="py-24 text-center">
          <div class="inline-block animate-spin rounded-full h-10 w-10 border-4 border-white/10 border-t-purple-500"></div>
          <p class="mt-4 text-zinc-500">Carregando usuários...</p>
        </div>

        <div v-else-if="usersStore.users.length > 0">
          
          <div class="bg-[#1c1c1e] rounded-3xl border border-white/5 p-4 mb-8 shadow-2xl">
             <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="relative group">
                <svg class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-zinc-500 group-focus-within:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
                <input
                  v-model="searchQuery"
                  type="text"
                  placeholder="Buscar usuário..."
                  class="w-full pl-12 pr-4 py-3 bg-[#2c2c2e] border-none rounded-2xl text-white placeholder-zinc-600 focus:ring-2 focus:ring-purple-500/50 transition-all"
                />
              </div>

              <div class="relative group">
                <svg class="absolute left-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-zinc-500 group-focus-within:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
                </svg>
                <select
                  v-model="departmentFilter"
                  class="w-full pl-12 pr-10 py-3 bg-[#2c2c2e] border-none rounded-2xl text-white appearance-none cursor-pointer focus:ring-2 focus:ring-purple-500/50 transition-all"
                >
                  <option value="" class="bg-[#2c2c2e]">Todos os departamentos</option>
                  <option v-for="dept in availableDepartments" :key="dept" :value="dept" class="bg-[#2c2c2e]">{{ dept }}</option>
                </select>
                <div class="absolute right-4 top-1/2 transform -translate-y-1/2 pointer-events-none">
                  <svg class="w-4 h-4 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
                </div>
              </div>
             </div>
          </div>

          <div class="grid grid-cols-1 gap-3">
            <div
              v-for="user in filteredUsers"
              :key="user.id"
              class="group bg-[#1c1c1e] rounded-2xl border border-white/5 p-4 hover:bg-[#252528] transition-all duration-300 flex items-center gap-5 relative overflow-hidden"
            >
              <div class="absolute left-0 top-0 bottom-0 w-1 bg-purple-500 opacity-0 group-hover:opacity-100 transition-opacity"></div>

              <div class="p-3 bg-white/5 rounded-2xl group-hover:bg-purple-500/10 group-hover:text-purple-400 transition-colors shrink-0 text-zinc-400">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>

              <div class="flex-1 min-w-0">
                <div class="flex items-center justify-between md:justify-start gap-4 mb-1">
                  <h3 class="text-base font-semibold text-white truncate">{{ user.name }}</h3>
                </div>
                
                <div class="flex items-center gap-4 text-xs flex-wrap">
                  <div v-if="user.email" class="flex items-center gap-1.5 text-zinc-500">
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                    <span>{{ user.email }}</span>
                  </div>

                  <div v-if="user.department" :class="getDepartmentColor(user.department)" class="flex items-center gap-1.5 px-2.5 py-1 rounded-md font-medium border border-transparent">
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                    </svg>
                    <span>{{ user.department }}</span>
                  </div>
                </div>
              </div>

              <div class="flex gap-2 shrink-0 opacity-100 sm:opacity-0 sm:group-hover:opacity-100 transition-opacity">
                <button
                  @click="startEdit(user)"
                  class="p-2 text-zinc-400 hover:text-white hover:bg-white/10 rounded-lg transition-all"
                  title="Editar"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  @click="confirmDelete(user)"
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
        </div>

        <div v-else class="text-center py-20 bg-[#1c1c1e] rounded-3xl border border-white/5 border-dashed">
          <div class="w-16 h-16 bg-white/5 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-zinc-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-medium text-white">Nenhum usuário</h3>
          <p class="text-zinc-500 mb-6">Não encontramos usuários com os filtros atuais.</p>
          <button
            @click="openModal()"
            class="text-blue-400 font-medium hover:text-blue-300 hover:underline"
          >
            Criar novo usuário &rarr;
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
            
            <div class="bg-[#1c1c1e] border border-white/10 rounded-2xl shadow-2xl w-full max-w-md overflow-hidden flex flex-col">
              <div class="px-6 py-5 border-b border-white/10 flex justify-between items-center bg-[#242426]">
                <h3 class="text-lg font-bold text-white">
                  {{ editingUser ? 'Editar Usuário' : 'Novo Usuário' }}
                </h3>
                <button @click="closeModal" class="text-zinc-500 hover:text-white transition-colors">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>

              <div class="p-8">
                <form @submit.prevent="handleSubmit" class="space-y-6">
                  <div>
                    <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Nome Completo</label>
                    <input
                      v-model="form.name"
                      @input="capitalizeFirstLetter"
                      type="text"
                      required
                      placeholder="Ex: João Silva"
                      class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                    />
                  </div>

                  <div>
                    <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Email (Opcional)</label>
                    <input
                      v-model="form.email"
                      type="email"
                      placeholder="Ex: joao@empresa.com"
                      class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                    />
                  </div>

                  <div>
                    <label class="block text-xs font-semibold text-zinc-400 uppercase tracking-wider mb-2">Departamento</label>
                    <div class="relative">
                      <input
                        v-model="departmentSearch"
                        @input="searchDepartments"
                        @focus="showDepartmentSuggestions = true"
                        type="text"
                        placeholder="Selecionar..."
                        class="w-full px-4 py-3 bg-[#2c2c2e] border border-transparent rounded-xl focus:ring-2 focus:ring-blue-500/50 focus:bg-[#323234] transition-all text-white placeholder-zinc-600 outline-none"
                      />
                      
                      <div v-if="showDepartmentSuggestions && filteredDepartments.length > 0" class="absolute z-10 w-full mt-1 bg-[#2c2c2e] border border-white/10 rounded-xl shadow-xl max-h-48 overflow-y-auto custom-scrollbar">
                        <button
                          v-for="dept in filteredDepartments"
                          :key="dept"
                          type="button"
                          @click="selectDepartment(dept)"
                          class="w-full px-4 py-2.5 text-left hover:bg-blue-500 hover:text-white transition-colors border-b border-white/5 last:border-b-0 text-zinc-300"
                        >
                          {{ dept }}
                        </button>
                      </div>
                    </div>
                  </div>

                  <div class="pt-4 flex gap-3">
                    <button
                      type="button"
                      @click="closeModal"
                      class="flex-1 px-4 py-3 border border-white/10 text-zinc-300 font-medium rounded-xl hover:bg-white/5 transition-colors"
                    >
                      Cancelar
                    </button>
                    <button
                      type="submit"
                      :disabled="usersStore.loading"
                      class="flex-1 px-4 py-3 bg-[#0071e3] text-white font-medium rounded-xl hover:bg-[#0077ED] disabled:opacity-70 transition-all shadow-lg shadow-blue-900/20"
                    >
                      {{ usersStore.loading ? 'Salvando...' : (editingUser ? 'Salvar' : 'Criar') }}
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
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <h3 class="text-lg font-bold text-white mb-2">Excluir Usuário?</h3>
              <p class="text-zinc-400 mb-6 text-sm">
                Tem certeza que deseja remover <span class="text-white font-medium">"{{ userToDelete?.name }}"</span>?
              </p>
              
              <div v-if="linkedItemsCount !== null && linkedItemsCount > 0" class="mb-6 p-3 bg-amber-500/10 border border-amber-500/20 rounded-lg">
                <div class="flex items-center gap-2 text-amber-400 justify-center text-sm">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                  </svg>
                  <span class="font-semibold">{{ linkedItemsCount }} itens vinculados serão excluídos!</span>
                </div>
              </div>
              
              <div v-if="loadingCount" class="mb-6 text-center text-zinc-500 text-sm">
                <div class="inline-block animate-spin rounded-full h-4 w-4 border-2 border-zinc-500 border-t-white mr-2"></div>
                Verificando vínculos...
              </div>
              
              <div class="flex gap-3">
                <button
                  @click="cancelDelete"
                  class="flex-1 px-4 py-2.5 border border-white/10 text-zinc-300 font-medium rounded-xl hover:bg-white/5 transition-all"
                >
                  Cancelar
                </button>
                <button
                  @click="executeDelete"
                  :disabled="usersStore.loading"
                  class="flex-1 px-4 py-2.5 bg-red-500/10 text-red-500 border border-red-500/20 font-medium rounded-xl hover:bg-red-500 hover:text-white transition-all"
                >
                  {{ usersStore.loading ? '...' : 'Excluir' }}
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
import { useUsersStore } from '@/stores/users'
import AppLayout from '@/components/AppLayout.vue'

const usersStore = useUsersStore()

const showModal = ref(false)
const showDeleteModal = ref(false)
const editingUser = ref(null)
const userToDelete = ref(null)
const successMessage = ref('')
const searchQuery = ref('')
const departmentFilter = ref('')
const linkedItemsCount = ref(null)
const loadingCount = ref(false)

const form = ref({
  name: '',
  email: '',
  department: ''
})

// Estados de Autocomplete
const departmentSearch = ref('')
const showDepartmentSuggestions = ref(false)
const availableDepartments = ref([
  'Administrativo', 'RH', 'TI', 'Telemarketing', 'Gerentes',
  'Lider', 'Segurança', 'Marketing', 'Faxina', 'Operacional',
  'Estrategico', 'Recepção'
])

const filteredDepartments = computed(() => {
  if (!departmentSearch.value) return availableDepartments.value
  const search = departmentSearch.value.toLowerCase()
  return availableDepartments.value.filter(dept => 
    dept.toLowerCase().includes(search)
  )
})

const filteredUsers = computed(() => {
  let filtered = usersStore.users
  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    filtered = filtered.filter(user => 
      user.name.toLowerCase().includes(search) ||
      (user.email && user.email.toLowerCase().includes(search)) ||
      (user.department && user.department.toLowerCase().includes(search))
    )
  }
  if (departmentFilter.value) {
    filtered = filtered.filter(user => 
      user.department && user.department === departmentFilter.value
    )
  }
  return filtered
})

// Cores para os departamentos (Estilo Glass)
const getDepartmentColor = (dept) => {
  const colors = {
    'TI': 'bg-blue-500/10 text-blue-400 border-blue-500/20',
    'Administrativo': 'bg-zinc-500/10 text-zinc-400 border-zinc-500/20',
    'RH': 'bg-purple-500/10 text-purple-400 border-purple-500/20',
    'Marketing': 'bg-pink-500/10 text-pink-400 border-pink-500/20',
    'Gerentes': 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20',
    'Operacional': 'bg-orange-500/10 text-orange-400 border-orange-500/20',
  }
  return colors[dept] || 'bg-zinc-500/10 text-zinc-400 border-zinc-500/20'
}

onMounted(() => {
  usersStore.fetchUsers()
})

const openModal = () => {
  editingUser.value = null
  resetForm()
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  setTimeout(() => {
    editingUser.value = null
    resetForm()
  }, 200)
}

const resetForm = () => {
  form.value = { name: '', email: '', department: '' }
  departmentSearch.value = ''
  showDepartmentSuggestions.value = false
}

const capitalizeFirstLetter = () => {
  const value = form.value.name
  if (value.length > 0) {
    form.value.name = value.charAt(0).toUpperCase() + value.slice(1)
  }
}

const searchDepartments = () => {
  showDepartmentSuggestions.value = true
  form.value.department = departmentSearch.value
}

const selectDepartment = (dept) => {
  departmentSearch.value = dept
  form.value.department = dept
  showDepartmentSuggestions.value = false
}

const handleSubmit = async () => {
  let success = false
  if (editingUser.value) {
    success = await usersStore.updateUser(editingUser.value.id, form.value)
    if (success) showSuccess('Usuário atualizado com sucesso!')
  } else {
    success = await usersStore.createUser(form.value)
    if (success) showSuccess('Usuário criado com sucesso!')
  }
  if (success) closeModal()
}

const startEdit = (user) => {
  editingUser.value = user
  form.value = {
    name: user.name,
    email: user.email || '',
    department: user.department || ''
  }
  departmentSearch.value = user.department || ''
  showModal.value = true
}

const confirmDelete = async (user) => {
  userToDelete.value = user
  linkedItemsCount.value = null
  loadingCount.value = true
  showDeleteModal.value = true
  
  const count = await usersStore.getLinkedItemsCount(user.name)
  linkedItemsCount.value = count
  loadingCount.value = false
}

const cancelDelete = () => {
  showDeleteModal.value = false
  setTimeout(() => {
    userToDelete.value = null
    linkedItemsCount.value = null
    loadingCount.value = false
  }, 200)
}

const executeDelete = async () => {
  if (!userToDelete.value) return
  const success = await usersStore.deleteUser(userToDelete.value.id)
  if (success) {
    showSuccess('Usuário removido do sistema.')
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