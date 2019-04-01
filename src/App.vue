<template>
  <div id="app">
    <Recommender/>
  </div>
</template>

<script>
import Recommender from './components/Recommender.vue'

export default {
  name: 'app',
  components: {
    Recommender
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>
