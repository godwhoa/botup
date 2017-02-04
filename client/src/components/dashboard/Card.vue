<template>
<div class="card">
  <header class="card-header">
    <p class="card-header-title center">{{ name }}<span v-if="alive" class="fa fa-circle status" title="dead"></span><span v-else class="fa fa-circle status" title="alive" style="color:#cb4b16;"></span></p>
  </header>
  <div class="card-content center">
    <div class="content center">
      <span class="card-content-title">Server</span><br>
      <span class="card-content-subtitle">{{server}}</span><br>

      <span class="card-content-title">Channel</span><br>
      <span class="card-content-subtitle">{{channel}}</span><br>

      <span class="card-content-title">Plugins</span><br>
      <span v-for="plugin in plugins">
        <span class="fa fa-git plugin" v-if="plugin == 'git'"></span>  
        <span class="fa fa-clock-o plugin" v-if="plugin == 'timein'"></span>
      </span>
    </div>
  </div>
  <footer class="card-footer">
    <router-link :to="'/dashboard/'+bid.toString()" class="card-footer-item">Manage</router-link>
  </footer>
</div>
</template>

<script>
export default {
  name: 'card',
  props: ['name','server','channel','alive','plugins','bid'],
  data(){
  	return{}
  }
}
</script>

<style scoped>
.card{
  margin: auto;
}
.center {
  text-align: center;
}

.card-header-title{
  display:inline-block;
  font-size: 1.4rem;
  color: #0a0a0a;
}

.status {
  padding-left: 0.4rem;
  padding-top: 0.7rem;
  font-size:0.8rem;
  color:#23d160;
}

.card-content-title {
  font-family: 'Montserrat', sans-serif;
  font-size: 1.4rem;
  color: #0a0a0a;
  font-weight: 300;
}

.card-content-subtitle {
  font-family: 'Montserrat', sans-serif;
  font-size: 1.1rem;
  color: #7a7a7a;
}

.plugin{
  padding-left: 5px;
  font-size: 1.1rem;
  color: #7a7a7a;
}
</style>