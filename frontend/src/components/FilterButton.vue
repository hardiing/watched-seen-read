<script lang="ts" setup>

const props = defineProps(['filter', 'currentFilter'])

const emits = defineEmits(['setFilter'])
</script>

<template>
    <button class="secondary" @click="emits('setFilter', props.filter)" :class="{contrast: props.currentFilter === props.filter}">{{ props.filter }}</button>
</template>

<style scoped>
button {
    text-transform: capitalize;
}
</style>