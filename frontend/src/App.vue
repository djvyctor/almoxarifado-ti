<template>
  <router-view v-slot="{ Component, route }">
    <transition :name="transitionName" mode="out-in">
      <component :is="Component" :key="route.path" />
    </transition>
  </router-view>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const transitionName = ref('')

// Monitora a mudança de rota
watch(
  () => route.path,
  (to, from) => {
    // Se a página de destino (to) OU a de origem (from) for '/login'
    // Aplicamos a animação 'page-fade'.
    if (to === '/login' || from === '/login') {
      transitionName.value = 'page-fade'
    } else {
      // Caso contrário (navegando dentro do sistema), sem animação.
      transitionName.value = ''
    }
  }
)
</script>

<style>
/* 1. Reset Global */
html, body {
  background-color: #000000;
  color: #ffffff;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

/* 2. Scrollbar Dark */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-track {
  background: #09090b;
}
::-webkit-scrollbar-thumb {
  background: #27272a;
  border-radius: 4px;
  border: 2px solid #09090b;
}
::-webkit-scrollbar-thumb:hover {
  background: #3f3f46;
}

/* 3. Animação de Entrada/Saída (SÓ funciona quando transitionName for 'page-fade') */
.page-fade-enter-active,
.page-fade-leave-active {
  transition: opacity 0.4s ease, transform 0.4s ease; /* Aumentei um pouquinho para ficar mais cinematográfico no login */
}

.page-fade-enter-from {
  opacity: 0;
  transform: scale(0.98); /* Leve zoom */
}

.page-fade-leave-to {
  opacity: 0;
  transform: scale(1.02);
}

/* 4. Remove outline feio */
input:focus, textarea:focus, select:focus {
  outline: none;
}
</style>