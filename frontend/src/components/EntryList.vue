<script lang="ts" setup>

const props = defineProps(['entries']);

const emits = defineEmits(['removeEntry'])

</script>

<template>
    <TransitionGroup name="list" tag="div" class="task-list">
        <article v-for="entry in entries" class="entry" :key="entry.entryTitle">
            {{ entry.entryDate }}
            {{ entry.entryTitle }}
            {{ entry.entryType }}
            <button @click="emits('removeEntry', entry.id)" class="outline">Remove</button>
        </article>
    </TransitionGroup>
</template>

<style>
.task-list {
    margin-top: 1rem;
}

.entry {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.list-enter-active,
.list-leave-active {
    transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
    opacity: 0;
    transform: translateX(300px);
}
</style>