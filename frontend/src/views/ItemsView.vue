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

      <div v-else-if="itemsStore.items.length > 0">
        <!-- Filtros -->
        <div class="bg-white rounded-2xl border border-slate-200 p-4 mb-6 shadow-sm">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="relative">
              <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Buscar por nome ou categoria..."
                class="w-full pl-10 pr-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl focus:ring-2 focus:ring-purple-500/20 focus:border-purple-500 transition-all text-slate-700 placeholder-slate-400"
              />
            </div>
            
            <div class="relative">
              <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z" />
              </svg>
              <select
                v-model="categoryFilter"
                class="w-full pl-10 pr-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl focus:ring-2 focus:ring-purple-500/20 focus:border-purple-500 transition-all text-slate-700 appearance-none cursor-pointer"
              >
                <option value="">Todas as categorias</option>
                <option value="Periféricos">Periféricos</option>
                <option value="Hardware">Hardware</option>
                <option value="Cabos e Adaptadores">Cabos e Adaptadores</option>
                <option value="Monitores">Monitores</option>
                <option value="Computadores">Computadores</option>
                <option value="Notebooks">Notebooks</option>
                <option value="Impressoras">Impressoras</option>
                <option value="Rede">Rede</option>
                <option value="Armazenamento">Armazenamento</option>
                <option value="Outros">Outros</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Lista de Itens -->
        <div v-if="items.length > 0" class="space-y-3">
        <div
          v-for="item in items"
          :key="item.id"
          class="group bg-white rounded-xl border-2 border-slate-200 p-4 hover:shadow-lg hover:shadow-purple-500/10 hover:border-purple-500 transition-all duration-300 flex items-center gap-4"
        >
          <!-- Ícone -->
          <div class="p-3 bg-slate-50 rounded-xl group-hover:bg-purple-50 transition-colors flex-shrink-0">
            <svg class="w-6 h-6 text-slate-400 group-hover:text-purple-500 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>

          <!-- Informações -->
          <div class="flex-1 min-w-0">
            <h3 class="text-lg font-bold text-slate-800 truncate mb-1">{{ item.name }}</h3>
            <p v-if="item.description" class="text-sm text-slate-600 mb-2 line-clamp-1">{{ item.description }}</p>
            <div class="flex items-center gap-4 text-sm flex-wrap">
              <div :class="getCategoryColor(item.category)" class="flex items-center gap-2 px-2.5 py-1 rounded-lg font-medium">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                </svg>
                <span>{{ item.category }}</span>
              </div>
              <div v-if="item.patrimony" class="flex items-center gap-2 text-purple-600">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
                <span class="font-medium">{{ item.patrimony }}</span>
              </div>
              <div v-if="item.value" class="flex items-center gap-2 text-green-600">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span class="font-bold">R$ {{ ((item.value || 0) * (item.quantity || 0)).toFixed(2).replace('.', ',') }}</span>
              </div>
              <div v-if="item.assigned_to" class="flex items-center gap-2 text-indigo-600">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
                <span class="font-medium">{{ item.assigned_to }}</span>
              </div>
            </div>
          </div>

          <!-- Quantidade -->
          <div class="flex-shrink-0">
            <div :class="[
              'px-4 py-2 rounded-lg text-sm font-bold border-2',
              item.quantity < 5 
                ? 'bg-red-50 text-red-600 border-red-200' 
                : 'bg-emerald-50 text-emerald-600 border-emerald-200'
            ]">
              {{ item.quantity }} unidades
            </div>
          </div>

          <!-- Ações -->
          <div class="flex gap-2 flex-shrink-0">
            <button
              @click="startEdit(item)"
              class="p-2.5 text-slate-600 bg-slate-50 rounded-lg hover:bg-indigo-50 hover:text-indigo-600 transition-all"
              title="Editar"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
            </button>
            <button
              @click="confirmDelete(item)"
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

        <div v-else class="text-center py-12 bg-slate-50 rounded-xl border border-slate-200">
          <svg class="w-12 h-12 text-slate-300 mx-auto mb-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <p class="text-slate-500">Nenhum item encontrado com os filtros aplicados</p>
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
          
          <div class="bg-white rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-hidden transform transition-all flex flex-col">
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

            <div class="p-6 overflow-y-auto flex-1">
              <form @submit.prevent="handleSubmit" class="space-y-5">
                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Nome do Item</label>
                  <input
                    v-model="form.name"
                    @input="capitalizeFirstLetter"
                    type="text"
                    required
                    placeholder="Ex: Teclado Mecânico"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                  />
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Categoria</label>
                  <div class="relative">
                    <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-slate-400 pointer-events-none" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                    </svg>
                    <input
                      v-model="categorySearch"
                      @input="searchCategories"
                      @focus="showCategorySuggestions = true"
                      type="text"
                      required
                      placeholder="Digite ou selecione uma categoria"
                      class="w-full pl-10 pr-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                    />
                    
                    <!-- Dropdown de categorias -->
                    <div v-if="showCategorySuggestions && filteredCategories.length > 0" class="absolute z-10 w-full mt-1 bg-white border border-slate-300 rounded-lg shadow-lg max-h-60 overflow-y-auto">
                      <button
                        v-for="category in filteredCategories"
                        :key="category"
                        type="button"
                        @click="selectCategory(category)"
                        class="w-full px-4 py-2.5 text-left hover:bg-indigo-50 transition-colors flex items-center gap-2 border-b border-slate-100 last:border-b-0"
                      >
                        <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                        </svg>
                        <span class="text-slate-700">{{ category }}</span>
                      </button>
                    </div>
                  </div>
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Observações</label>
                  <textarea
                    v-model="form.description"
                    rows="3"
                    :placeholder="getDescriptionPlaceholder()"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400 resize-none"
                  ></textarea>
                  <p class="text-xs text-slate-500 mt-1">{{ getDescriptionHint() }}</p>
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Patrimônio</label>
                  <input
                    v-model="form.patrimony"
                    type="text"
                    maxlength="5"
                    @input="formatPatrimony"
                    placeholder="Ex: 12345 (opcional)"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                  />
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Valor (R$)</label>
                  <div class="relative">
                    <span class="absolute left-3 top-1/2 transform -translate-y-1/2 text-slate-500 font-medium">R$</span>
                    <input
                      v-model="form.value"
                      type="text"
                      @input="formatValue"
                      placeholder="0,00 (opcional)"
                      class="w-full pl-12 pr-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                    />
                  </div>
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Vincular a Usuário</label>
                  <div class="relative">
                    <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                    <input
                      v-model="userSearch"
                      @input="searchUsers"
                      @focus="showUserSuggestions = true"
                      type="text"
                      placeholder="Digite o nome do usuário (opcional)"
                      class="w-full pl-10 pr-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700 placeholder-slate-400"
                    />
                    
                    <!-- Dropdown de sugestões -->
                    <div v-if="showUserSuggestions && filteredUsers.length > 0" class="absolute z-10 w-full mt-1 bg-white border border-slate-300 rounded-lg shadow-lg max-h-48 overflow-y-auto">
                      <button
                        v-for="user in filteredUsers"
                        :key="user.id"
                        type="button"
                        @click="selectUser(user)"
                        class="w-full px-4 py-2.5 text-left hover:bg-indigo-50 transition-colors border-b border-slate-100 last:border-b-0"
                      >
                        <div class="flex items-center gap-2">
                          <svg class="w-4 h-4 text-indigo-600 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                          </svg>
                          <div class="flex-1 min-w-0">
                            <div class="text-slate-800 font-medium">{{ user.name }}</div>
                            <div v-if="user.email || user.department" class="text-xs text-slate-500 flex items-center gap-2">
                              <span v-if="user.email">{{ user.email }}</span>
                              <span v-if="user.email && user.department">•</span>
                              <span v-if="user.department">{{ user.department }}</span>
                            </div>
                          </div>
                        </div>
                      </button>
                    </div>
                  </div>
                  <p class="text-xs text-slate-500 mt-1">Digite para buscar usuários disponíveis</p>
                </div>

                <div>
                  <label class="block text-sm font-semibold text-slate-700 mb-2">Quantidade</label>
                  <input
                    v-model.number="form.quantity"
                    type="number"
                    required
                    min="0"
                    class="w-full px-4 py-2.5 bg-white border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all text-slate-700"
                  />
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
              <p class="text-slate-600 text-center mb-6">
                Tem certeza que deseja excluir <span class="font-semibold text-slate-800">"{{ itemToDelete?.name }}"</span> do estoque?
              </p>
              
              <div class="flex gap-3">
                <button
                  @click="cancelDelete"
                  class="flex-1 px-4 py-2.5 border-2 border-slate-300 text-slate-700 font-medium rounded-xl hover:bg-slate-50 transition-all"
                >
                  Cancelar
                </button>
                <button
                  @click="executeDelete"
                  :disabled="itemsStore.loading"
                  class="flex-1 px-4 py-2.5 bg-red-600 text-white font-medium rounded-xl hover:bg-red-700 focus:ring-4 focus:ring-red-500/30 disabled:opacity-70 transition-all flex items-center justify-center"
                >
                  <svg v-if="itemsStore.loading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  {{ itemsStore.loading ? 'Excluindo...' : 'Sim, Excluir' }}
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

// Estado para autocomplete de usuários
const userSearch = ref('')
const showUserSuggestions = ref(false)

// Estado para autocomplete de categorias
const categorySearch = ref('')
const showCategorySuggestions = ref(false)
const availableCategories = ref([
  'Periféricos',
  'Hardware',
  'Cabos e Adaptadores',
  'Monitores',
  'Computadores',
  'Notebooks',
  'Impressoras',
  'Rede',
  'Armazenamento',
  'Outros'
])

const filteredCategories = computed(() => {
  if (!categorySearch.value) return availableCategories.value
  const search = categorySearch.value.toLowerCase()
  return availableCategories.value.filter(category => 
    category.toLowerCase().includes(search)
  )
})

const searchCategories = () => {
  showCategorySuggestions.value = true
  form.value.category = categorySearch.value
}

const selectCategory = (category) => {
  categorySearch.value = category
  form.value.category = category
  showCategorySuggestions.value = false
}

const filteredUsers = computed(() => {
  if (!userSearch.value) return []
  const search = userSearch.value.toLowerCase()
  return usersStore.users.filter(user => 
    user.name.toLowerCase().includes(search) ||
    (user.email && user.email.toLowerCase().includes(search))
  )
})

const searchUsers = () => {
  showUserSuggestions.value = userSearch.value.length > 0
}

const selectUser = (user) => {
  userSearch.value = user.name
  form.value.assigned_to = user.name
  showUserSuggestions.value = false
}

const getCategoryColor = (category) => {
  const colors = {
    'Periféricos': 'bg-blue-50 text-blue-700 border border-blue-200',
    'Hardware': 'bg-purple-50 text-purple-700 border border-purple-200',
    'Cabos e Adaptadores': 'bg-yellow-50 text-yellow-700 border border-yellow-200',
    'Monitores': 'bg-cyan-50 text-cyan-700 border border-cyan-200',
    'Computadores': 'bg-indigo-50 text-indigo-700 border border-indigo-200',
    'Notebooks': 'bg-pink-50 text-pink-700 border border-pink-200',
    'Impressoras': 'bg-orange-50 text-orange-700 border border-orange-200',
    'Rede': 'bg-green-50 text-green-700 border border-green-200',
    'Armazenamento': 'bg-teal-50 text-teal-700 border border-teal-200',
    'Outros': 'bg-slate-50 text-slate-700 border border-slate-200'
  }
  return colors[category] || 'bg-slate-50 text-slate-700 border border-slate-200'
}

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
  return placeholders[form.value.category] || 'Adicione detalhes sobre este item (opcional)'
}

const getDescriptionHint = () => {
  const hints = {
    'Cabos e Adaptadores': 'Informe tipo, comprimento e especificações',
    'Monitores': 'Tamanho, resolução e conexões disponíveis',
    'Hardware': 'Especificações técnicas e modelo',
    'Periféricos': 'Tipo de conexão e características',
    'Computadores': 'Marca, modelo e configuração',
    'Notebooks': 'Marca, tamanho da tela e especificações',
    'Impressoras': 'Tipo, velocidade e recursos',
    'Rede': 'Capacidade, velocidade e portas',
    'Armazenamento': 'Capacidade, interface e velocidade',
    'Outros': 'Qualquer informação relevante'
  }
  return hints[form.value.category] || 'Adicione informações úteis sobre o item'
}

const formatPatrimony = (event) => {
  // Permite apenas números
  let value = event.target.value.replace(/\D/g, '')
  // Limita a 5 dígitos
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
  
  // Converte para número e formata com 2 casas decimais
  const numValue = parseInt(value) / 100
  form.value.value = numValue.toFixed(2).replace('.', ',')
}

const items = computed(() => {
  let filtered = itemsStore.items
  
  // Filtro por busca
  if (searchQuery.value) {
    filtered = filtered.filter(item => 
      item.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      item.category.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  
  // Filtro por categoria
  if (categoryFilter.value) {
    filtered = filtered.filter(item => item.category === categoryFilter.value)
  }
  
  return filtered
})

const categories = computed(() => {
  const cats = itemsStore.items.map(item => item.category).filter(Boolean)
  return [...new Set(cats)]
})

onMounted(() => {
  itemsStore.fetchItems()
  usersStore.fetchUsers()
})

// Actions
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
  }, 200) // Pequeno delay para a animação terminar
}

const resetForm = () => {
  form.value = {
    name: '',
    category: '',
    description: '',
    patrimony: '',
    assigned_to: '',
    value: '',
    quantity: 0
  }
  userSearch.value = ''
  categorySearch.value = ''
}

const handleSubmit = async () => {
  let success = false
  
  // Converter valor de string formatada para número
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