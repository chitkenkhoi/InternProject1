<template>
  <div class="LayoutPage">
    <SideBar class="SideBar" />
    <Header class="Header"/>
    <div class="mid"></div>
    <main class="LayoutPage_main">
      <router-view class="bodyforpages"></router-view>
    </main>
  </div>
</template>

<script>
import SideBar from "./SideBar.vue";
import Header from "./Header.vue";

export default {
  name: "LayoutPage",
  components: { SideBar, Header},
  mounted(){
    const email = localStorage.getItem("c_user");
    const password = localStorage.getItem("password");
    if(email&&password){
      return
    }else{
      this.$router.push({name:"Login"})
    }


  }
};
</script>

<style scoped>
.SideBar {
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  width: 210px;
  height: auto;
  color: whitesmoke;

  background-color: rgb(30, 25, 57);
  border: none;
  padding: 30px;
}
.Header {
  position: fixed;
  top: 0;
  left: 270px;

  width: 1700px;
  height: 75px;
  color: black;
  background-color: white;
  border-bottom: 0.1px solid lightgrey;
  text-align: center;
}
.bodyforpages {
  position: fixed;
  top: 126px;
  left: 300px;
  height: 650px;
  width: 190vh;
  /* border: 5px solid black; */
  /* overflow-y: auto; */
}
.mid {
  position: fixed;
  top: 76px;
  left: 270px;
  background-color: whitesmoke;
  height: 50px;
  border-bottom: 1px solid lightgrey;
  width: 1500px;
  display:flex;
  align-items: center;
  
}
</style>