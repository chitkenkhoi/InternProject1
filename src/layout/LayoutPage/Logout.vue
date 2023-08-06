<template>
    <button @click = "logout">Logout</button>
</template>

<script>
export default {
name : "LogOut",
methods:{
    logout(){
    const a = confirm("Are you sure you want to log out?")
    if (a)
    {
        this.$router.push({name: 'Login'})
    }
  }
}
}
</script>

<style lang = "scss" scoped>
button{
    background-color:rgb(255,69,0);
    color:white;
    height: 40px;
    width: 120px;
    border-radius: 3px;
    border: none;
    transition: 0.2s ease-in-out;
    cursor: pointer;
        &:hover {
				background-color: rgb(255,165,0);
				color: black;
				}
            }
</style>