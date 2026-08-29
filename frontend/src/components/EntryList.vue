<script lang="ts" setup>
import { ref } from 'vue'
import type { Entry } from '../types';

const props = defineProps<{
  entries: Entry[]
}>()

const emits = defineEmits<{
  removeEntry: [id: number]
}>()

</script>

<template>
    <TransitionGroup name="list" tag="div" class="task-list">
        <article v-for="entry in props.entries" class="entry" :key="entry.id">
            {{ entry.record_date }} <p class="entry">//</p>
            {{ entry.record_title }} <p class="entry">//</p>
            {{ entry.record_type }}
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