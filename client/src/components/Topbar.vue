<template>
  <nav class="nav">
    <div class="nav-left">
      <router-link class="nav-item" to="/dashboard" v-if="page == 'edit'">
        <h1 class="title brand-right"><i class="fa fa-arrow-left"></i></h1>
      </router-link>
      <router-link class="nav-item" to="/" v-else-if="page != 'home'">
        <h1 class="title brand-botup"><span>B</span>otup</h1>
      </router-link>
    </div>

    <div class="nav-center">
    	<!-- Home -->
	    <router-link class="nav-item" to="/" v-if="page === 'home'">
	      <h1 class="title brand-botup"><span>B</span>otup</h1>
	    </router-link>
	    <a class="nav-item" href="#" v-else>
	      <h1 class="title page-title">{{ title }}</h1>
	    </a>
    </div>
    
    <div class="nav-right">
      <!-- Login -->
      <router-link class="nav-item" to="/register" v-if="page === 'login'">
        <h1 class="title brand-right">Register<i class="fa fa-user-plus"></i></h1>
      </router-link>
      <!-- Register -->
      <router-link class="nav-item" to="/login" v-else-if="page === 'register'">
          <h1 class="title brand-right">Login<i class="fa fa-sign-in"></i></h1>
      </router-link>
	  <!-- Dashboard -->
      <a class="nav-item" @click="emitLogout" v-else-if="page === 'dashboard'">
          <h1 class="title brand-right">Logout<i class="fa fa-sign-out"></i></h1>
      </a>
      <!-- Other -->
      <a class="nav-item" href="#" v-else>
      </a>
    </div>
  </nav>
</template>

<script>
	export default {
	  name: 'topbar',
	  data () {
	    return {name:'topbar',title:'',page:''}
	  },
      created(){
        window.bus.$on('route-change',(route) => {
          this.title = route.charAt(1).toUpperCase() + route.slice(2) // capitalize first char
          this.page = route.slice(1)
        })
      },
      methods:{
      	emitLogout(){
      		window.bus.$emit('logout','dashboard')
      	}
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
