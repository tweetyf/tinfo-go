<script setup>
import Footer from './components/Footer.vue'
import Home from './pages/Home.vue'
import Tools from './pages/Tools.vue'
import About from './pages/About.vue'
import API from './pages/API.vue'

</script>

<script>

export default {
  components: {
    Footer,
    Home,
    Tools,
    About,
    API,
  },
  data() {
    return {
      currentTab: 'Home',
      tabs: ['Home', 'Tools', 'API', 'About']
    }
  }
}
</script>

<template>
  <header>
    <nav class="w3-bar w3-collapse w3-blue" id="mySidebar"><br>
      <div class="w3-bar">
        <button v-for="tab in tabs" :key="tab"
          :class="['w3-bar-item w3-button w3-hover-none w3-border-blue w3-bottombar w3-hover-border-white w3-padding', { active: currentTab === tab }]"
          @click="currentTab = tab">
          {{ tab }}
        </button>
        <a href="https://lzspot.info" onclick="" target="_blank"
          class="w3-bar-item w3-button w3-hover-none w3-border-blue w3-bottombar w3-hover-border-white w3-padding">BLOG</a>
      </div>
    </nav>
  </header>

  <main>
    <component :is="currentTab" class="w3-container"></component>
  </main>

  <footer>
    <Footer/>
  </footer>
</template>

<style>
@import 'assets/w3.css';


#app {
  max-width: 1280px;
  margin: 0 auto;
  padding: 2rem;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: normal;
}

header {
  line-height: 1.5;
}

main {
  min-height: 600px;
  background-color:white;
}

body,
h1,
h2,
h3,
h4,
h5,
h6 {
  font-family: Arial, Helvetica, sans-serif;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
