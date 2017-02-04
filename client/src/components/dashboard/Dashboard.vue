<template>
<div class="container">
<div class="columns" v-for="column in columns">
  <div class="column is-4" v-for="card in column">
    <card :name="card.name" :server="card.server"
    :channel="card.channel" :alive="card.alive" :plugins="card.plugins" :bid="card.bid"></card>
  </div>
</div>
</div>  
</template>

<script>
import Card from "./Card.vue"

export default {
  name: 'dashboard',
  data(){
  	return{
      columns:[
          [
            {
              bid:1,
              name:"pacbot",
              server:"chat.freenode.net:7373",
              channel:"##oodnet",
              alive:true,
              plugins:['git','timein']
            },
            {
              bid:2,
              name:"pacbot",
              server:"chat.freenode.net:7373",
              channel:"##oodnet",
              alive:true,
              plugins:['git','timein']
            },
            {
              bid:3,
              name:"pacbot",
              server:"chat.freenode.net:7373",
              channel:"##oodnet",
              alive:true,
              plugins:['git','timein']
            }
          ],
          [
            {
              bid:4,
              name:"pacbot",
              server:"chat.freenode.net:7373",
              channel:"##oodnet",
              alive:true,
              plugins:['git','timein']
            },
            {
              bid:5,
              name:"pacbot",
              server:"chat.freenode.net:7373",
              channel:"##oodnet",
              alive:true,
              plugins:['git','timein']
            },
            {
              bid:6,
              name:"pacbot",
              server:"chat.freenode.net:7373",
              channel:"##oodnet",
              alive:true,
              plugins:['git','timein']
            }
          ]
      ]
    }
  },
  updated(){
    window.bus.$emit('route-change','/dashboard')
  },
  created(){
    window.bus.$emit('route-change','/dashboard')
    window.bus.$on('logout',(page)=>{
      // do ajax
      // redirect to login
      this.$router.push('login')
    })
  },
  components:{
    'card':Card
  }
}
</script>