<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios'
import type { Entry, NewEntry, EntryFilter } from './types';
import EntryForm from './components/EntryForm.vue';
import EntryList from './components/EntryList.vue';
import FilterButton from './components/FilterButton.vue';

const message = ref("Watched Seen Read");
const entries = ref<Entry[]>([]);
const filter = ref<EntryFilter>("all");

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

const removeEntry = async (id: number) => {
  try {
    await axios.delete(
      'http://localhost:8080/api/records/' + id,
    )

    await fetchEntries()
  } catch (error) {
    console.error('Unable to delete entry:', error)
  }
}

const filteredEntries = computed(() => {
  switch (filter.value) {
    case 'book':
      return entries.value.filter((entry) => entry.record_type === 'book')

    case 'movie':
      return entries.value.filter((entry) => entry.record_type === 'movie')

    case 'show':
      return entries.value.filter((entry) => entry.record_type === 'show')

    case 'other':
      return entries.value.filter((entry) => entry.record_type === 'other')

    default:
      return entries.value
  }
})

function setFilter(value: EntryFilter) {
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
      <FilterButton :currentFilter="filter" filter="book" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="movie" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="show" @set-filter="setFilter" />
      <FilterButton :currentFilter="filter" filter="other" @set-filter="setFilter" />
    </div>
    <EntryList :entries="filteredEntries" @remove-entry="removeEntry" />
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