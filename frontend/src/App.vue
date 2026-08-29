<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios'
import type { Entry, NewEntry } from './types';
import EntryForm from './components/EntryForm.vue';
import EntryList from './components/EntryList.vue';
import FilterButton from './components/FilterButton.vue';

const message = ref("Watched Seen Read");
const entries = ref<Entry[]>([]);
const filter = ref("all");

const fetchEntries = async () => {
  try {
    const response = await axios.get<Entry[]>(
      'http://localhost:8080/api/records'
    )

    entries.value = response.data
  } catch (error) {
    console.error('Unable to load entries:', error)
  }
}

const addEntry = async (entry: NewEntry) => {
  try {
    await axios.post(
      'http://localhost:8080/api/records',
      entry
    )

    await fetchEntries()
  } catch (error) {
    console.error('Unable to add entry:', error)
  }
}

function removeEntry(id: number) { // rework to delete from db
  const index = entries.value.findIndex((entry) => entry.id === id);
  if (index !== -1) {
    entries.value.splice(index, 1);
  }
}

const filteredEntries = computed(() => {
  switch (filter.value) {
    case 'books':
      return entries.value.filter((entry) => entry.record_type === 'book')

    case 'movies':
      return entries.value.filter((entry) => entry.record_type === 'movie')

    case 'shows':
      return entries.value.filter((entry) => entry.record_type === 'show')

    case 'other':
      return entries.value.filter((entry) => entry.record_type === 'other')

    default:
      return entries.value
  }
})

function setFilter(value: string) {
  filter.value = value;
}

onMounted(fetchEntries)
</script>

<template>
  <main>
    <h1>{{ message }}</h1>
    <EntryForm @add-entry="addEntry" />
    <h3 v-if="!entries.length">Add an entry to get started.</h3>
    <h3 v-else-if="entries.length === 1">{{ entries.length }} entry found</h3>
    <h3 v-else>{{ entries.length }} entries found</h3>
    <div v-if="entries.length" class="button-container">
      <FilterButton  :currentFilter="filter" filter="all" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="books" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="movies" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="shows" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="other" @set-filter="setFilter" />
    </div>
    <EntryList :entries @remove-entry="removeEntry" />
  </main>
</template>

<style>
main {
  max-width: 800px;
  margin: 1rem auto;
}

.button-container {
  display: flex;
  justify-content: end;
  gap: 0.5rem;
}
</style>