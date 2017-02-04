<template>
<div class="pure-container">
  <topbar></topbar>
  <router-view></router-view>
</div>
</template>

<script>
  import Topbar from './components/Topbar.vue'

  export default {
    name: 'app',
    data () {
      return {name:'app'}
    },
    components:{
      'topbar':Topbar
    }
  }
</script>

<style>
  .nav {
    height: 4rem;
  }

  .brand-botup {
    font-family: 'Montserrat', sans-serif;
    font-size: 2.2rem;
    color: #7a7a7a;
  }

  .brand-botup span {
    color: #b5b5b5;
  }

  .page-title {
    font-family: 'Montserrat', sans-serif;
    font-size: 2.2rem;
    color: #363636;
  }

  .brand-right {
    font-family: 'Montserrat', sans-serif;
    font-size: 1.7rem;
    color: #9eb7a7;
  }
</style>