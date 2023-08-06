<template>
  <div class = "footer">
    <div class = "item" v-for="item in menu" :key="item.id" :class = "{'-mark': item.id===show}" @click = "scrollMark(item.id)"><div class="text">{{ item.tittle }}</div></div>
  </div>
</template>

<script>
import {ref} from 'vue'
export default {
name: "FooterComponent",
setup() {
    const menu = [
      { tittle: "企業情報の設定", id: 1 },
      { tittle: "セキュリティ設定", id: 2 },
      { tittle: "メール通知設定", id: 3 },
      { tittle: "アラート通知先の設定", id: 4 },
    ];
    const show = ref(1);
    return { menu, show };
  },
methods:{
  scrollMark(id){
    this.$emit('scroll',id)
    this.show = id
  }
}

}
</script>

<style lang = "scss" scoped>
.footer{
  display:flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  margin-right:30px;
  width:220px;
  height:190px;
  .item{
    display:flex;
    flex-direction: row;
    align-items:center;
    height:20px;
    width: 220px;
    padding:10px;
    color: rgb(72,61,139);
    text-align: left;
    cursor:pointer;
    &:hover{
      color: rgb(247,84,25);
    }
    border-radius:1px;
    font-size:14px;

    
  }
  .-mark{
      background-color:rgb(255,239,213);
      color:rgb(247,84,25);
    }
}

</style>