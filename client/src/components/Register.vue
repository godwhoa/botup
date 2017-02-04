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
          <a class="button is-primary is-medium mont" id="register" @click="doRegister">
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
import dom from '../utils/dom.js'

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
  },
  methods:{
    doRegister(){
      const email = document.getElementById("email").value
      const user = document.getElementById("user").value
      const password = document.getElementById("pass").value

      this.$http.post('/api/user/register', {'email': email,'user':user,'pass':password}).then(response => {
        switch (response.body){
          case "OK_USER_CREATED":
            dom.show("#ok-created")
            setTimeout(()=>{
              this.$router.push("login")
            },1000)
            break;
          case "ERR_USER_TAKEN":
            dom.show("#user-taken")
            break;
          case "ERR_FIELDS_MISSING":
            dom.show("#invalid-form")
            break;
          case "ERR_INTERNAL":
            dom.show("#inter-err")
            break;
          default:
            dom.show("#inter-err")
        }
        setTimeout(()=>{
          dom.hide(".notification")
        },2000)
      },err => {console.log(err)});
    }
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