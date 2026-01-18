<template>
  <AppLayout>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-8">
        <div>
          <h1 class="text-3xl font-bold text-slate-800 tracking-tight">Usuários</h1>
          <p class="text-slate-500 mt-1">Gerencie os usuários do sistema</p>
        </div>
        
        <div class="flex gap-3">
          <button
            @click="openModal()"
            class="inline-flex items-center px-5 py-2.5 rounded-xl text-white font-medium bg-linear-to-r from-purple-600 to-indigo-600 hover:from-purple-700 hover:to-indigo-700 shadow-lg shadow-purple-500/20 transition-all hover:scale-105 active:scale-95"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Novo Usuário
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
          <svg class="w-5 h-5 mr-3 shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
          </svg>
          {{ successMessage }}
        </div>
      </Transition>

      <div v-if="usersStore.error" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-xl flex items-center text-red-700 shadow-sm">
        <svg class="w-5 h-5 mr-3 shrink-0" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
        </svg>
        {{ usersStore.error }}
      </div>

      <div v-if="usersStore.loading && usersStore.users.length === 0" class="py-20 text-center">
        <div class="inline-block animate-spin rounded-full h-10 w-10 border-4 border-purple-100 border-t-purple-600"></div>
        <p class="mt-4 text-slate-500">Carregando usuários...</p>
      </div>

      <div v-else-if="usersStore.users.length > 0">
        <div class="bg-white rounded-2xl border border-slate-200 p-4 mb-6 shadow-sm space-y-3">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Buscar usuário por nome, email ou departamento..."
            class="w-full px-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl focus:ring-2 focus:ring-purple-500/20 focus:border-purple-500 transition-all text-slate-700 placeholder-slate-400"
          />
          <div class="flex items-center gap-3">
            <label class="text-sm font-medium text-slate-600">Filtrar por Setor:</label>
            <select
              v-model="departmentFilter"
              class="px-4 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:ring-2 focus:ring-purple-500/20 focus:border-purple-500 transition-all text-slate-700"
            >
              <option value="">Todos</option>
              <option v-for="dept in availableDepartments" :key="dept" :value="dept">{{ dept }}</option>
            </select>
          </div>
        </div>

        <div class="space-y-4">
          <div
            v-for="user in filteredUsers"
            :key="user.id"
            class="group bg-white rounded-xl border-2 border-slate-200 p-4 hover:shadow-lg hover:shadow-purple-500/10 hover:border-purple-500 transition-all duration-300 flex items-center gap-4"
          >
            <div class="p-3 bg-slate-50 rounded-xl group-hover:bg-purple-50 transition-colors shrink-0">
              <svg class="w-6 h-6 text-slate-400 group-hover:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </div>

            <div class="flex-1 min-w-0">
              <h3 class="text-lg font-bold text-slate-800 truncate mb-1">{{ user.name }}</h3>
              <div class="flex items-center gap-4 text-sm flex-wrap">
                <div v-if="user.email" class="flex items-center gap-2 text-slate-500">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                  <span>{{ user.email }}</span>
                </div>
                <div v-if="user.department" class="flex items-center gap-2 text-indigo-600 px-2.5 py-1 bg-indigo-50 rounded-lg font-medium">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                  </svg>
                  <span>{{ user.department }}</span>
                </div>
              </div>
            </div>

            <div class="flex gap-2 shrink-0">
              <button
                @click="startEdit(user)"
                class="p-2.5 text-slate-600 bg-slate-50 rounded-lg hover:bg-indigo-50 hover:text-indigo-600 transition-all"
                title="Editar"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>
              <button
                @click="confirmDelete(user)"
                class="p-2.5 text-slate-400 hover:text-red-500 hover:bg-red-50 rounded-lg transition-all"
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

      <div v-else class="text-center py-20 bg-white rounded-3xl border border-dashed border-slate-300">
        <div class="w-16 h-16 bg-slate-50 rounded-full flex items-center justify-center mx-auto mb-4">
          <svg class="w-8 h-8 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-slate-900">Nenhum usuário encontrado</h3>
        <p class="text-slate-500 mb-6">Comece adicionando usuários ao sistema.</p>
        <button
          @click="openModal()"
          class="text-indigo-600 font-medium hover:text-indigo-700 hover:underline"
        >
          Criar primeiro usuário &rarr;
        </button>
      </div>

      <!-- Modal de Criar/Editar -->
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
                {{ editingUser ? 'Editar Usuário' : 'Novo Usuário' }}
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
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Nome Completo</label>
                  <input
                    v-model="form.name"
                    @input="capitalizeFirstLetter"
                    type="text"
                    required
                    placeholder="Ex: João Silva"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                  />
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Email (Opcional)</label>
                  <input
                    v-model="form.email"
                    type="email"
                    placeholder="Ex: joao.silva@empresa.com"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                  />
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Departamento (Opcional)</label>
                  <div class="relative">
                    <input
                      v-model="departmentSearch"
                      @input="searchDepartments"
                      @focus="showDepartmentSuggestions = true"
                      type="text"
                      placeholder="Digite ou selecione o departamento"
                      class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                    />
                    
                    <!-- Dropdown de sugestões de departamentos -->
                    <div v-if="showDepartmentSuggestions && filteredDepartments.length > 0" class="absolute z-10 w-full mt-1 bg-white border border-slate-300 rounded-lg shadow-lg max-h-48 overflow-y-auto">
                      <button
                        v-for="dept in filteredDepartments"
                        :key="dept"
                        type="button"
                        @click="selectDepartment(dept)"
                        class="w-full px-4 py-2.5 text-left hover:bg-indigo-50 transition-colors border-b border-slate-100 last:border-b-0 flex items-center gap-2"
                      >
                        <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4" />
                        </svg>
                        <span class="text-slate-700">{{ dept }}</span>
                      </button>
                    </div>
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
                    :disabled="usersStore.loading"
                    class="flex-1 px-4 py-2.5 bg-indigo-600 text-white font-medium rounded-xl hover:bg-indigo-700 focus:ring-4 focus:ring-indigo-500/30 disabled:opacity-70 transition-all flex items-center justify-center"
                  >
                    <svg v-if="usersStore.loading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    {{ usersStore.loading ? 'Salvando...' : (editingUser ? 'Salvar Alterações' : 'Criar Usuário') }}
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </Transition>

      <!-- Modal de Confirmação de Exclusão -->
      <Transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showDeleteModal" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm z-50 flex items-center justify-center p-4" @click.self="cancelDelete">
          <div class="bg-white rounded-2xl shadow-2xl w-full max-w-md overflow-hidden transform transition-all animate-scale-in">
            <div class="p-6">
              <div class="flex items-center justify-center w-12 h-12 mx-auto mb-4 rounded-full bg-red-100">
                <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-slate-800 text-center mb-2">Confirmar Exclusão</h3>
              <p class="text-slate-600 text-center mb-2">
                Tem certeza que deseja excluir <span class="font-semibold text-slate-800">"{{ userToDelete?.name }}"</span>?
              </p>
              
              <div v-if="linkedItemsCount !== null" class="mb-6 p-3 bg-amber-50 border border-amber-200 rounded-lg">
                <div class="flex items-center gap-2 text-amber-800 justify-center">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                  </svg>
                  <span class="font-semibold">{{ linkedItemsCount }} {{ linkedItemsCount === 1 ? 'item vinculado' : 'itens vinculados' }} serão excluídos!</span>
                </div>
              </div>
              
              <div v-if="loadingCount" class="mb-6 text-center text-slate-500 text-sm">
                <div class="inline-block animate-spin rounded-full h-4 w-4 border-2 border-slate-300 border-t-slate-600 mr-2"></div>
                Verificando itens vinculados...
              </div>
              
              <div class="flex gap-3">
                <button
                  @click="cancelDelete"
                  class="flex-1 px-4 py-2.5 border-2 border-slate-300 text-slate-700 font-medium rounded-xl hover:bg-slate-50 transition-all"
                >
                  Cancelar
                </button>
                <button
                  @click="executeDelete"
                  :disabled="usersStore.loading"
                  class="flex-1 px-4 py-2.5 bg-red-600 text-white font-medium rounded-xl hover:bg-red-700 focus:ring-4 focus:ring-red-500/30 disabled:opacity-70 transition-all flex items-center justify-center"
                >
                  <svg v-if="usersStore.loading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  {{ usersStore.loading ? 'Excluindo...' : 'Sim, Excluir' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </Transition>

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

// Estado para autocomplete de departamentos
const departmentSearch = ref('')
const showDepartmentSuggestions = ref(false)
const availableDepartments = ref([
  'Administrativo',
  'RH',
  'TI',
  'Telemarketing',
  'Gerentes',
  'Lider',
  'Segurança',
  'Marketing',
  'Faxina',
  'Operacional',
  'Estrategico',
  'Recepção'
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
  
  // Filtro por busca de texto
  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    filtered = filtered.filter(user => 
      user.name.toLowerCase().includes(search) ||
      (user.email && user.email.toLowerCase().includes(search)) ||
      (user.department && user.department.toLowerCase().includes(search))
    )
  }
  
  // Filtro por departamento
  if (departmentFilter.value) {
    filtered = filtered.filter(user => 
      user.department && user.department === departmentFilter.value
    )
  }
  
  return filtered
})

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
  form.value = {
    name: '',
    email: '',
    department: ''
  }
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

  if (success) {
    closeModal()
  }
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
  
  // Buscar contagem de itens vinculados
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
