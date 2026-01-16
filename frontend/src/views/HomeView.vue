<template>
  <AppLayout>
    <div class="home">
      <h2>Bem-vindo ao Sistema de Almoxarifado TI</h2>
      <p>Use o menu lateral para navegar</p>
    </div>
  </AppLayout>
</template>

<script setup>
import AppLayout from '@/components/AppLayout.vue'
</script>

<style scoped>
.home {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

h2 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

p {
  color: #7f8c8d;
}
</style>
