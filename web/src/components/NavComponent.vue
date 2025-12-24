<script setup lang="ts">
import { ref } from 'vue'
import { Button, Drawer } from 'primevue'
import { RouterLink } from 'vue-router'

const visible = ref(false)
</script>

<template>
	<header>
		<Drawer v-model:visible="visible">
			<template #container="{ closeCallback }">
				<ul @click="closeCallback">
					<li><RouterLink to="/">Home</RouterLink></li>
					<li><RouterLink to="/article">Article</RouterLink></li>
				</ul>
				<div class="button-position">
					<Button
						type="button"
						@click="closeCallback"
						icon="pi pi-times"
						variant="outlined"
					></Button>
				</div>
			</template>
		</Drawer>
	</header>
	<div class="button-position">
		<Button icon="pi pi-bars" @click="visible = true" />
	</div>
</template>

<style scoped></style>
