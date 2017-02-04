<template>
  <div class="pure-container">
    <div class="container">
      <div class="column is-3">
        <p class="control has-icon">
          <input class="input is-medium mont" type="email"  id="email" placeholder="Email">
          <i class="fa fa-envelope"></i>
        </p>
      </div>
      <div class="column is-3">
        <p class="control has-icon">
          <input class="input is-medium mont" type="text"  id="user" placeholder="Username">
          <i class="fa fa-user"></i>
        </p>
      </div>
      <div class="column is-3">
        <p class="control has-icon">
          <input class="input is-medium mont" type="password"  id="pass" placeholder="Password">
          <i class="fa fa-lock"></i>
        </p>
      </div>
      <div class="column is-3">
        <p class="control has-icon login">
          <a class="button is-primary is-medium mont" id="register">
            <span class="icon">
              <i class="fa fa-user-plus"></i>
            </span>
            <span>Register</span>
          </a>
        </p>
      </div>
        <responses></responses>
    </div>
  </div>
</template>

<script>
import Responses from './Responses.vue'

export default {
  name: 'register',
  data () {
    return {name:'register'}
  },
  updated(){
      window.bus.$emit('route-change','/register')
  },
  created(){
      window.bus.$emit('route-change','/register')
  },
  components:{
    'responses':Responses
  }
}
</script>

<style>

/* Horizontally center forms */
.container > .column {
  margin-left: auto;
  margin-right:auto;
}

/* Vertically center forms */
.container {
  margin-top: 8%;
}

/* Horizontally center login button */
.login {
  text-align: center;
}

/* Login button colors */
.login > .button{
  background-color: #5b5a5a;
}

.login > .button:hover, .login > .button:focus{
  background-color: #363636;
}
</style>