<script lang="ts" setup>
import { reactive } from 'vue';
import type { NewEntry } from '../types';

const emit = defineEmits<{
  addEntry: [entry: NewEntry]
}>()

const entry = reactive<NewEntry>({
    record_date: '',
    record_title: '',
    record_type: ''
})

const formSubmitted = () => {
  emit('addEntry', {
    record_date: entry.record_date,
    record_title: entry.record_title,
    record_type: entry.record_type
  })

  entry.record_date = ''
  entry.record_title = ''
  entry.record_type = ''
}
</script>

<template>
    <form @submit.prevent="formSubmitted">
        <!-- Date Input -->
      <div class="form-group">
        <label for="entryDate">Date (YYYY-MM-DD)</label>
        <input 
          type="text" 
          id="entryDate" 
          v-model.trim="entry.record_date" 
          required 
        />
      </div>

       <!-- Title Input -->
      <div class="form-group">
        <label for="entryTitle">Title</label>
        <input 
          type="text" 
          id="entryTitle" 
          v-model.trim="entry.record_title" 
          required 
        />
      </div>

      <!-- Select Dropdown -->
      <div class="form-group">
        <label for="entryType">Type:</label>
        <select id="entryType" v-model="entry.record_type">
          <option value="book">Book</option>
          <option value="movie">Movie</option>
          <option value="show">TV Show</option>
          <option value="other">Other</option>
        </select>
      </div>
        <div class="button-container">
        <button>Add</button>
        </div>
    </form>
    <!-- Previewing Saved State -->
    <!-- <pre v-if="submittedData">{{ submittedData }}</pre> -->
</template>