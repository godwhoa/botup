<template>
	<div class="container">
        <div class="column is-3">
          <label class="label">Name</label>
          <p class="control">
            <input class="input is-medium mont" id="name" type="text" placeholder="Name" :value="card.name">
          </p>
        </div>
        <div class="column is-3">
          <label class="label">Server</label>
          <p class="control">
            <input class="input is-medium mont" id="name" type="text" placeholder="Server" :value="card.server">
          </p>
        </div>
        <div class="column is-3">
          <label class="label">Channel</label>
          <p class="control">
            <input class="input is-medium mont" id="name" type="text" placeholder="Channel" :value="card.channel">
          </p>
        </div>
        <div class="column is-3">
          <p class="control has-icon">
	          <a class="button is-danger is-medium mont" id="alive" v-if="card.alive">
	              <span class="icon">
	                <i class="fa fa-times"></i>
	              </span>
	              <span>Kill</span>
	          </a>
	          <a class="button is-primary is-medium mont" id="alive" v-else>
	              <span class="icon">
	                <i class="fa fa-times"></i>
	              </span>
	              <span>Awake</span>
	          </a>
          </p>
        </div>
	</div>
</template>

<script>
export default {
  name: 'edit',
  data(){
  	return{
        card:{
          bid:1,
          name:"pacbot",
          server:"chat.freenode.net:7373",
          channel:"##oodnet",
          alive:true,
          plugins:['git','timein']
        }
  	}
  },
  updated(){
    window.bus.$emit('route-change','/edit')
  },
  created(){
    window.bus.$emit('route-change','/edit')
    window.bus.$on('logout',(page)=>{
      // do ajax
      // redirect to login
      this.$router.push('login')
    })
  }
}
</script>

<style scoped>
.label {
	font-size: 1.3rem;
}
</style>