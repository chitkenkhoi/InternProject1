<template>
  <Popup @turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div>オーナー変更</div>
        <span class="material-symbols-outlined" @click="turnoff">close</span>
      </div>
      <div class="boxtext">
        <div>自分の所有権をメンバーに譲渡することができます。</div>
        <div>所有権を譲渡した後は、正会員となります。</div>
        <div>その際、一部の権限が失われます。</div>
      </div>
      <div class="boxsearch">
        <input
          type="text"
          placeholder="メールアドレス、または既存のメンバーの氏名を入力する"
        />
        <div class="list">
          <owner-list v-for="ele in OwnerProp" :key="ele.id" :eleProp="ele" />
        </div>
      </div>
      <div class="btn" @click="turnoff">
        <div class = "text">送信する</div>
      </div>
    </div>
  </Popup>
</template>

<script>
import Popup from "../Popup.vue";
import OwnerList from "./OwnerList.vue";
export default {
  name: "ChangeOwner",
  props: ["OwnerProp"],
  components: { Popup, OwnerList },
  setup() {},
  methods: {
    turnoff() {
      this.$emit("no-pop-up");
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  position: fixed;
  background-color: white;
  z-index: 99;
  top: 40px;
  left: 500px;
  width: 600px;
  height: 716px;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
  .h {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 500px;
    height: 30px;
    margin-top: 50px;
    color:#271B5A;
    font-size:16px;
    font-family: 'Noto Sans JP';
    font-weight: 500;
    .material-symbols-outlined{
        color:black;
        cursor: pointer;
      &:hover {
        opacity: 0.5;
      }
    
    }
  }
  .boxtext {
    width: 500px;
    height: 76px;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 50px;
    font-size: 14px;
    font-family: "Noto Sans JP";
    font-weight: 400;
    justify-content: space-between;
  }
  .boxsearch {
    margin-top: 50px;
    width: 500px;
    height: 350px;
    display: flex;
    flex-direction: column;
    border-radius: 5px;
    border: 2px solid lightgrey;
    input {
      width: 496px;
      height: 50px;
      border: none;
      border-radius: 5px;
    }
    input:focus {
      outline: none;
    }
    .list {
      border-top: 2px solid lightgray;
      width: 500px;
      height: 300px;
      overflow-y: auto;
      overflow-x: hidden;
      // display: flex;
      // flex-direction: column;
      // align-items: center;
      // flex-wrap:nowrap;
    }

  }
  .btn{
    width:500px;
    height:50px;
    margin-top:20px;
    background-color: #D64D10;
    display:flex;
    align-items: center;
    justify-content: center;
    border-radius:5px;

    cursor:pointer;
    &:hover{
        background-color: rgb(220, 20, 60);
    }
    .text{
        color:white;
        font-size:14px;
        font-family:'Noto Sans JP';
        font-weight: 500px;
    }
  }
}
</style>