<script lang="ts" setup>
import { reactive, ref } from 'vue';
import axios from 'axios';

const emit = defineEmits(['formSubmitted', 'addEntry']);

const newEntry = ref("");
const error = ref("");
const Entry = reactive({
    entryDate: '',
    entryTitle: '',
    entryType: ''
})
const submittedData = ref(Entry)

//function formSubmitted() {
    //submittedData.value = { ...Entry }
    //emit("addEntry", { ...Entry })
    /* if (newEntry.value.trim()) {
        emit("addEntry", newEntry.value.trim());
        newEntry.value = "";
    } else {
        error.value = "Entry cannot be empty!"
    } */
//}

const formSubmitted = async () => {
  try {
    const response = await axios.post('http://localhost:8080/api/records', {
      record_date: Entry.entryDate,
      record_title: Entry.entryTitle,
      record_type: Entry.entryType
    })
    //message.value = response.data.message
    Entry.entryDate = ''
    Entry.entryTitle = ''
    Entry.entryType = ''
  } catch (error: any) {
    console.log("status:", error.response?.status)
    console.log("data:", error.response?.data)
    console.log("error:", error)
    //message.value = 'Error saving data: ' + (error.response?.data || error.message)
  }
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
          v-model.trim="Entry.entryDate" 
          required 
        />
      </div>

       <!-- Title Input -->
      <div class="form-group">
        <label for="entryTitle">Title</label>
        <input 
          type="text" 
          id="entryTitle" 
          v-model.trim="Entry.entryTitle" 
          required 
        />
      </div>

      <!-- Select Dropdown -->
      <div class="form-group">
        <label for="entryType">Type:</label>
        <select id="entryType" v-model="Entry.entryType">
          <option value="book">Book</option>
          <option value="movie">Movie</option>
          <option value="show">TV Show</option>
          <option value="other">Other</option>
        </select>
      </div>
        <!-- <label>
        Entry Date
        <input 
            v-model="newEntry"
            name="newEntry"
            :aria-invalid="!!error || undefined"
            @input="error = ''"
        >
        <small v-if="error" id="invalid-helper">
            {{ error }}
        </small>
        </label>
        <label>
        Entry Title
        <input 
            v-model="newEntry"
            name="newEntry"
            :aria-invalid="!!error || undefined"
            @input="error = ''"
        >
        <small v-if="error" id="invalid-helper">
            {{ error }}
        </small>
        </label>
        <label>
        Entry Type
        <input 
            v-model="newEntry"
            name="newEntry"
            :aria-invalid="!!error || undefined"
            @input="error = ''"
        >
        <small v-if="error" id="invalid-helper">
            {{ error }}
        </small>
        </label> -->
        <div class="button-container">
        <button>Add</button>
        </div>
    </form>
    <!-- Previewing Saved State -->
    <pre v-if="submittedData">{{ submittedData }}</pre>
</template>