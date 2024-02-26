<template>
  <div>
    {{ myName }}
  </div>
</template>

<script>
export default {
  name: 'My',
}
</script>

<script setup>
import { ref } from 'vue'
const myName = ref('name')
</script>
