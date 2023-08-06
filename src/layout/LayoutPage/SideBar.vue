<template>
  <aside>
    <img class = "Logo" :src = "logo"/>
    <div class = mid >
      <img class = "logo" :src = "logo2" />
      <div class = "name">株式会社C-mind</div>
    </div>
    <Button class = "buttonAdd">
      <span class="material-symbols-outlined">add</span>
      <div class = "text1">契約書を作成する</div>
      
    </Button>
    <div class = "menu" ><menu-item v-for= "item in menu" v-bind:key = "item.id" :todoProp="item" v-on:set = "set" v-on:toggle ="toggle" :showProp = "showchildren"/></div>
    <div class = "space">
      <div>利用規約</div>
      <div>特定商取引法に基づく表記</div>
      <div>プライパシーポリシー</div>
      <div>運営会社</div>
      <div>よくあるご質問</div>
      <div>お問い合わせ</div>
    </div>
    <div class = "iconlist">
      <a href = "https://www.youtube.com/"><img :src = "ytbicon"/></a>
      <a href = "https://line.me/en/"><img :src = "lineicon"/></a>
      <a href = "https://www.facebook.com/"><img :src = "fbicon"/></a>
      <a href = "https://twitter.com/"><img :src = "twticon"/></a>
    </div>
    <img class = "last" :src = "foot">
  </aside>
</template>

<script>
import Button from '@/components/Button.vue'
import foot from '../../../public/text/textfooter.png'
import logo from '../../../public/logo/Mask Group 36.png'
import logo2 from '../../../public/logo/logo.png'
import ytbicon from '../../../public/icon/ytb.png'
import fbicon from '../../../public/icon/fb.png'
import lineicon from '../../../public/icon/line.png'
import twticon from '../../../public/icon/twt.png'
import {ref} from 'vue'
import MenuItem from './MenuItem.vue'
export default {
name: "SideBar",
components: {MenuItem,Button},
setup(){
  const showchildren = ref(false)
  const menu = [
                {tittle: "ホーム", id: 1,path :"/dashboard"},
                {tittle: "契約書の管理", id: 2,path :"/folder"},
                {tittle: "雛形の管理", id: 3,path :"/template"},
                {tittle: "アカウントの管理", id: 4,path :"/management"},
                {tittle: "顧客一覧", id: 5,path :"/customer"},
              ]
  const set = ()=>{
    showchildren.value = false
  }
  const toggle = ()=>{
    showchildren.value = !showchildren.value

  }
  return {menu,set,toggle,showchildren,logo,logo2,ytbicon,fbicon,lineicon,twticon,foot}
}

}
</script>

<style lang = "scss" scoped>
aside{
  display: flex;
  flex-direction: column;
  font-size:14px;
  .mid{
    // position: absolute;
    // top:70px;
    // left: 0; 
    background-color:black;
    height:52px;
    width: 270px;
    margin: 0 -30px;
    margin-top:25px;
    display:flex;
    align-items: center;
    .logo{
      height:40px;
      width:40px;
      border-radius:50%;
      margin-left:30px
    }
    .name{
      color: lightgray;
      margin-left:10px;
      font-size:14px;
      opacity:50%;
    }
  }
  .Logo{
    width :150px;
    height: 25px;
    margin-top:-5px;
  }
  .buttonAdd{
    display: flex;
    margin-top:30px;
    
		align-items: center;
		text-decoration: none;
		transition: 0.2s ease-in-out;
		// padding: 0.5rem 1rem;
    border:none;
    cursor: pointer;
    border-radius: 5px;
    width:210px;
    height:50px;
    
    .material-symbols-outlined{
      color:white;
      margin-left: 5px;
      width:20px;
      height:20px;
    }
    .text1{
      margin-left: 15px;
      color: white;
      font-size: 14px;
    }
  }

  .menu{
    display: flex;
    flex-direction: column;
    margin: 0 -30px;
    transition: 0.5s ;
    margin-top:10px;
    
  }
  .space{
    // flex: 1 1 0;
    display:flex;
    flex-direction: column;
    // position:absolute;
    // margin-top:520px;
    height: 200px;
    // border:5px solid black;
    justify-content: space-evenly;
    font-size:12px;
  }
  .iconlist{
    display: flex;
    justify-content: space-between;
    color: white;
    position:absolute;
    bottom: 10px;
    width:200px;
    height:40px;
  }
  .last{
    position:absolute;
    bottom:10px;
  }
}

</style>