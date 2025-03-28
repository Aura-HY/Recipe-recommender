<template>
  <div>
    <h1 @click="login()">Login</h1>
    <h1 @click="signup()">signup</h1>
  </div>
</template>

<script>
export default {
  name: "login",
  methods: {
    login() {
      window.location.href = this.getSigninUrl();
    },
    signup() {
      window.location.href = this.getSignupUrl();
    },
  },
};
</script>

<style></style>
