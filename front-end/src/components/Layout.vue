<script setup>
import { RouterView } from 'vue-router';
import ThemeSwitch from './commmon/themeSwitch.vue';
import router from '@/router';

const openLink = (link) => {
  window.open(link, '_blank')
}
</script>

<template>
  <div class="h-screen overflow-y-auto bg-background theme-transition">
    <div class="h-screen overflow-y-auto grid-background theme-transition">
      <div class="h-12 bg-[hsl(var(--header-bg))] shadow-sm border-b border-[hsl(var(--border))] place-items-center flex">
        <div class="font-bold text-xl flex-grow text-center title bg-clip-text bg-gradient-to-r from-purple-500 via-red-500 to-pink-500 text-transparent cursor-pointer" @click="router.push('/')">FastCV</div>
        <div class="flex flex-nowrap gap-2 px-4">
          <ThemeSwitch/>
          <a href="https://github.com/Flyinsky2004/FastCV" class="hover:text-[hsl(var(--primary))] transition-colors">
            <svg class="w-6 h-6 text-[hsl(var(--icon-color))]" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
              <path fill-rule="evenodd"
                    d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
                    clip-rule="evenodd"></path>
            </svg>
          </a>
        </div>
      </div>
      <RouterView/>
    </div>
  </div>
</template> 