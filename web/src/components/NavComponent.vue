<script setup lang="ts">
import { ref } from 'vue'
import { Button, Drawer } from 'primevue'

const visible = ref(false)
</script>

<template>
	<header>
		<Drawer v-model:visible="visible">
			<template #container="{ closeCallback }">
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
