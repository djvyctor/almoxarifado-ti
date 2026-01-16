<template>
  <AppLayout>
    <div class="items-page">
      <div class="header-section">
        <h2>Gerenciar Itens</h2>
        <button @click="showAddForm = true" class="btn-primary">
          ➕ Adicionar Item
        </button>
      </div>

      <!-- Mensagens -->
      <div v-if="itemsStore.error" class="alert alert-error">
        {{ itemsStore.error }}
      </div>

      <div v-if="successMessage" class="alert alert-success">
        {{ successMessage }}
      </div>

      <!-- Formulário de Adicionar/Editar -->
      <div v-if="showAddForm || editingItem" class="form-modal">
        <div class="modal-content">
          <h3>{{ editingItem ? 'Editar Item' : 'Adicionar Item' }}</h3>
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label>Nome</label>
              <input v-model="form.nome" type="text" required />
            </div>

            <div class="form-group">
              <label>Descrição</label>
              <textarea v-model="form.descricao" rows="3"></textarea>
            </div>

            <div class="form-group">
              <label>Quantidade</label>
              <input v-model.number="form.quantidade" type="number" required min="0" />
            </div>

            <div class="form-group">
              <label>Localização</label>
              <input v-model="form.localizacao" type="text" />
            </div>

            <div class="form-actions">
              <button type="button" @click="cancelForm" class="btn-secondary">
                Cancelar
              </button>
              <button type="submit" class="btn-primary" :disabled="itemsStore.loading">
                {{ itemsStore.loading ? 'Salvando...' : 'Salvar' }}
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Loading -->
      <div v-if="itemsStore.loading && items.length === 0" class="loading">
        Carregando itens...
      </div>

      <!-- Lista de Itens -->
      <div v-else class="items-grid">
        <div v-for="item in items" :key="item.id" class="item-card">
          <div class="item-header">
            <h3>{{ item.nome }}</h3>
            <span class="item-quantity">Qtd: {{ item.quantidade }}</span>
          </div>
          
          <p class="item-description">{{ item.descricao || 'Sem descrição' }}</p>
          
          <p class="item-location">
            <span>📍</span> {{ item.localizacao || 'Não especificado' }}
          </p>

          <div class="item-actions">
            <button @click="startEdit(item)" class="btn-edit">
              ✏️ Editar
            </button>
            <button @click="confirmDelete(item)" class="btn-delete">
              🗑️ Deletar
            </button>
          </div>
        </div>

        <div v-if="items.length === 0" class="empty-state">
          <p>Nenhum item cadastrado</p>
          <p class="empty-hint">Clique em "Adicionar Item" para começar</p>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useItemsStore } from '@/stores/items'
import AppLayout from '@/components/AppLayout.vue'

const itemsStore = useItemsStore()

const showAddForm = ref(false)
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

const handleSubmit = async () => {
  let success = false
  
  if (editingItem.value) {
    success = await itemsStore.updateItem(editingItem.value.id, form.value)
    if (success) {
      successMessage.value = 'Item atualizado com sucesso!'
    }
  } else {
    success = await itemsStore.createItem(form.value)
    if (success) {
      successMessage.value = 'Item criado com sucesso!'
    }
  }

  if (success) {
    cancelForm()
    setTimeout(() => {
      successMessage.value = ''
    }, 3000)
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
}

const cancelForm = () => {
  showAddForm.value = false
  editingItem.value = null
  form.value = {
    nome: '',
    descricao: '',
    quantidade: 0,
    localizacao: ''
  }
}

const confirmDelete = async (item) => {
  if (confirm(`Tem certeza que deseja deletar "${item.nome}"?`)) {
    const success = await itemsStore.deleteItem(item.id)
    if (success) {
      successMessage.value = 'Item deletado com sucesso!'
      setTimeout(() => {
        successMessage.value = ''
      }, 3000)
    }
  }
}
</script>

<style scoped>
.items-page {
  max-width: 1200px;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

h2 {
  color: #2c3e50;
  margin: 0;
}

.btn-primary {
  padding: 0.75rem 1.5rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
}

.btn-primary:hover:not(:disabled) {
  background-color: #2980b9;
}

.btn-primary:disabled {
  background-color: #95a5a6;
  cursor: not-allowed;
}

.btn-secondary {
  padding: 0.75rem 1.5rem;
  background-color: #95a5a6;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

.btn-secondary:hover {
  background-color: #7f8c8d;
}

.alert {
  padding: 1rem;
  border-radius: 4px;
  margin-bottom: 1rem;
}

.alert-error {
  background-color: #ffebee;
  color: #c62828;
  border-left: 4px solid #c62828;
}

.alert-success {
  background-color: #e8f5e9;
  color: #2e7d32;
  border-left: 4px solid #2e7d32;
}

.form-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-content h3 {
  margin-bottom: 1.5rem;
  color: #2c3e50;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-weight: 500;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  font-family: inherit;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #3498db;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 2rem;
}

.loading {
  text-align: center;
  padding: 2rem;
  color: #7f8c8d;
  background: white;
  border-radius: 8px;
}

.items-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
}

.item-card {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  transition: transform 0.2s, box-shadow 0.2s;
}

.item-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.15);
}

.item-header {
  display: flex;
  justify-content: space-between;
  align-items: start;
  margin-bottom: 1rem;
}

.item-header h3 {
  color: #2c3e50;
  margin: 0;
  font-size: 1.25rem;
}

.item-quantity {
  background-color: #3498db;
  color: white;
  padding: 0.25rem 0.75rem;
  border-radius: 12px;
  font-size: 0.875rem;
  font-weight: 500;
}

.item-description {
  color: #7f8c8d;
  margin-bottom: 1rem;
  min-height: 3em;
}

.item-location {
  color: #95a5a6;
  font-size: 0.9rem;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.item-actions {
  display: flex;
  gap: 0.5rem;
  padding-top: 1rem;
  border-top: 1px solid #ecf0f1;
}

.btn-edit,
.btn-delete {
  flex: 1;
  padding: 0.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: background-color 0.2s;
}

.btn-edit {
  background-color: #f39c12;
  color: white;
}

.btn-edit:hover {
  background-color: #e67e22;
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
}

.btn-delete:hover {
  background-color: #c0392b;
}

.empty-state {
  grid-column: 1 / -1;
  text-align: center;
  padding: 3rem;
  background: white;
  border-radius: 8px;
  color: #95a5a6;
}

.empty-hint {
  margin-top: 0.5rem;
  font-size: 0.9rem;
}
</style>
