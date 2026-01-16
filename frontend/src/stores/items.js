import { defineStore } from 'pinia'
import api from '@/services/api'

export const useItemsStore = defineStore('items', {
  state: () => ({
    items: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchItems() {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.get('/items')
        this.items = response.data
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao carregar itens'
      } finally {
        this.loading = false
      }
    },

    async createItem(item) {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.post('/items', item)
        this.items.push(response.data)
        return true
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao criar item'
        return false
      } finally {
        this.loading = false
      }
    },

    async updateItem(id, item) {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.put(`/items/${id}`, item)
        const index = this.items.findIndex(i => i.id === id)
        if (index !== -1) {
          this.items[index] = response.data
        }
        return true
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao atualizar item'
        return false
      } finally {
        this.loading = false
      }
    },

    async deleteItem(id) {
      this.loading = true
      this.error = null
      
      try {
        await api.delete(`/items/${id}`)
        this.items = this.items.filter(i => i.id !== id)
        return true
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao deletar item'
        return false
      } finally {
        this.loading = false
      }
    }
  }
})
