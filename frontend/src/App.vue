<script lang="ts" setup>
import { ref } from 'vue';
import type { Entry } from './types';
import EntryForm from './components/EntryForm.vue';
import EntryList from './components/EntryList.vue';
import FilterButton from './components/FilterButton.vue';

const message = ref("Watched Seen Read");
const entries = ref<Entry[]>([]);
const filter = ref("all");
const entry = ref<Entry | null>(null);
const handleFormSubmission = (data: Entry): void => {
  entry.value = data
};

/* const filteredEntries = computed(() => {
  switch(filter.value) {
    case "all":
      return entries.value;
    case "books":
      return entries.value.filter((entry) => entry.book);
    case "movies":
      return entries.value.filter((entry) => entry.movie);
    case "events":
      return entries.value.filter((entry) => entry.event);
  }
  return entries.value;
}) */

function addEntry(data: Entry) {
  entries.value.push({
    date: data.date,
    title: data.title,
    type: data.type,
  });
}

function removeEntry(id: number) {
  const index = entries.value.findIndex((entry) => entry.id === id);
  if (index !== -1) {
    entries.value.splice(index, 1);
  }
}

function setFilter(value: string) {
  filter.value = value;
}
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
      <FilterButton :currentFilter="filter" filter="events" @set-filter="setFilter" />
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