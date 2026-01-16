import { defineStore } from 'pinia'
import api from '@/services/api'

export const useUsersStore = defineStore('users', {
  state: () => ({
    users: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchUsers() {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.get('/users')
        this.users = response.data
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao carregar usuários'
      } finally {
        this.loading = false
      }
    },

    async createUser(user) {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.post('/users', user)
        this.users.push(response.data)
        await this.fetchUsers()
        return true
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao criar usuário'
        return false
      } finally {
        this.loading = false
      }
    },

    async updateUser(id, user) {
      this.loading = true
      this.error = null
      
      try {
        await api.put(`/users/${id}`, user)
        await this.fetchUsers()
        return true
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao atualizar usuário'
        return false
      } finally {
        this.loading = false
      }
    },

    async getLinkedItemsCount(userName) {
      try {
        const response = await api.get(`/users/linked-items?name=${encodeURIComponent(userName)}`)
        return response.data.count
      } catch (error) {
        console.error('Erro ao buscar itens vinculados:', error)
        return 0
      }
    },

    async deleteUser(id) {
      this.loading = true
      this.error = null
      
      try {
        await api.delete(`/users/${id}`)
        this.users = this.users.filter(u => u.id !== id)
        return true
      } catch (error) {
        this.error = error.response?.data?.error || 'Erro ao deletar usuário'
        return false
      } finally {
        this.loading = false
      }
    }
  }
})
