<template>
<div class="container">
	<router-link to="/login" class="title brand-right routes">Login<i class="fa fa-sign-in"></i></router-link>
	<router-link to="/register" class="title brand-right routes">Register<i class="fa fa-user-plus"></i></router-link>
	<router-link to="/dashboard" class="title brand-right routes">Dashboard<i class="fa fa-dashboard"></i></router-link>
</div>
</template>

<script>
export default {
  name: 'home',
  data(){
  	return{'loggedin':false}
  },
  updated(){
    window.bus.$emit('route-change','/home')
  },
  created(){
    window.bus.$emit('route-change','/home')
  }
}
</script>

<style scoped>
	.container{
		text-align: center;
	}
	.routes{
		display:block;
		font-size:1.8rem;
	}
	i{
		margin-left: 4px;
	}
</style>