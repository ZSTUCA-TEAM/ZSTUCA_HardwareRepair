import { ref, reactive } from 'vue'
import { defineStore } from 'pinia'

export const useBackstageStore = defineStore('backstage', () => {
  const jwt = ref("")
  const username = ref("")
  return { jwt, username }
}, { persist: false })
