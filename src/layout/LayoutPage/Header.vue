<template>
  <div class = "header">
    <div class ="search">
      <span class="material-symbols-outlined">search</span>
      <div class = "text">契約書の検索</div>
    </div>
    <div class = "space"></div>
    <Button class = "btn">
      <span class="material-symbols-outlined">diamond</span>
      <div class ="text" id="text">プランを変更</div>
    </Button>
    <div class = "profile" @click = "logout">
    <img :src = "avatar"/>
    <div class = "username">{{ name }}</div>
    </div>
  </div>
</template>
<script>
import Button from '@/components/Button.vue'
import HTTP from '@/http-common'
import {ref} from 'vue'
export default {
name: "HeaderComponent",
components: {Button},
setup(){
  const avatar = ref("")
  const name = ref("")
  return {avatar,name}
},
async created(){
  const email = localStorage.getItem('c_user')
  if (!email) return
  const response = await HTTP.get(`getUser/${email}`) 
  console.log(response)
  this.avatar = response.data.data.avatar
  this.name = response.data.data.first_name + ' ' + response.data.data.last_name;
},
methods:{
  logout(){
    const a = confirm("Are you sure you want to log out?")
    if (a)
    {   localStorage.clear();
        this.$router.push({name: 'Login'})
    }
  }
}

}

</script>

<style lang = "scss" scoped>
.header{
  display:flex;
  align-items: center;
  font-size:14px;
  .search{
    display:flex;
    align-items: center;
    color: rgb(247,84,25);
    margin-left:23px;
  }
  .space{
    width: 50%;
  }
  .btn{
    display:flex;
    align-items: center;
    justify-content: space-evenly;
    color:white;
    height: 40px;
    width: 135px;
    border-radius: 3px;
    border: none;
    transition: 0.2s ease-in-out;
    font-size:12px;
    margin-left: -70px;
    cursor: pointer;
    .text{
      font-weight: 400;
    }
    #text{
      font-size:12px;
    }

  }
  .profile{
    display:flex;
    align-items:center;
    justify-content: space-evenly;
    margin-left: 30px;
    cursor:pointer;
    img{
      border-radius: 50%;
      height:50px;
      width: 40px;
    }
    .username{
      margin-left:15px;
      color:rgb(72,61,139);
    }
  }

}

</style>