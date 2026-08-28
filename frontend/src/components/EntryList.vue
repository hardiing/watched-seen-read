<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const emits = defineEmits(['removeEntry'])

type DBRecord = {
    id: number
    record_date: string
    record_title: string
    record_type: string
}

const records = ref<DBRecord[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

const fetchEntries = async () => {
  try {
    const response = await axios.get<DBRecord[]>(
      'http://localhost:8080/api/records'
    )
    records.value = response.data
  } catch (err) {
    console.error(err)
    error.value = 'Unable to load records.'
  } finally {
    loading.value = false
  }
}

onMounted(fetchEntries)

</script>

<template>
    <TransitionGroup name="list" tag="div" class="task-list">
        <article v-for="record in records" class="entry" :key="record.id">
            {{ record.record_date }} <p>//</p>
            {{ record.record_title }} <p>//</p>
            {{ record.record_type }}
            <button @click="emits('removeEntry', record.id)" class="outline">Remove</button>
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