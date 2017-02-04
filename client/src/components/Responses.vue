<template>
<div class="container">
    <div class="column is-3">
        <div class="notification is-success" id="ok-login">
          Ok, logged in 
        </div>
    </div>
    <div class="column is-3">
        <div class="notification is-danger" id="wrong-cred">
          Wrong credentials 
        </div>
    </div>
    <div class="column is-3">
        <div class="notification is-success" id="ok-created">
          Ok, user created 
        </div>
    </div>
    <div class="column is-3">
        <div class="notification is-danger" id="user-taken">
          User taken
        </div>
    </div>
    <div class="column is-3">
        <div class="notification is-danger" id="invalid-form">
          One or more fields in form are malformed
        </div>
    </div>
    <div class="column is-3">
        <div class="notification is-danger" id="inter-err">
          Something went wrong 
        </div>
    </div>
</div>
</template>

<script>
export default {
  name: 'responses',
  data () {
  	return {name:'responses'}
  }
}
</script>

<style scoped>
.container{
  margin-top: 0;
}
.notification{
  display: none;
  text-align: center;
  font-size: 1.2rem;
}
</style>