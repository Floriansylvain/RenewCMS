<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavComponent from './components/NavComponent.vue'
</script>

<template>
	<nav-component />
	<main>
		<router-view />
	</main>
</template>

<style scoped></style>
