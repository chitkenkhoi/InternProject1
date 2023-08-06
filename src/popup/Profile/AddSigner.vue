<template>
  <Popup v-on:turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div class="text">署名者を追加</div>
        <span class="material-symbols-outlined" @click="turnoff">close</span>
      </div>
      <div class="submitarea">
        <div class="h2">歌手</div>
        <input type="text" v-model="data" @keyup.enter = "turnoff" />
        <button class="btn" @click="turnoff">変更する</button>
      </div>
    </div>
    </Popup>
  </template>
  
  <script>
  import Popup from "../Popup.vue";
  import { ref } from "vue";
  export default {
    name: "EditAddress",
    components:{Popup},
    setup(_, context) {
      const data = ref("");
      const turnoff = () => {
        context.emit("no-pop-up");
      };
      const savechange = () => {
        if(data.value === "") return
        else context.emit("data-change", data.value);
      };
  
      return { turnoff, data, savechange };
    },
  };
  </script>
  
  <style lang="scss" scoped>
  .container {
    margin-top: 105px;
    background-color: white;
    display: flex;
    border-radius: 5px;
    height: 291px;
    width: 500px;
    flex-direction: column;
    color: black;
    align-items: center;
    justify-content: space-evenly;
    .h {
      width: 400px;
      display: flex;
      justify-content: space-between;
      .text {
        color: rgb(72, 61, 139);
        font-size:16px;
      }
      .material-symbols-outlined {
        cursor: pointer;
        &:hover {
          opacity: 0.5;
        }
      }
    }
    .submitarea {
      display: flex;
      flex-direction: column;
      .h2{
        color: rgb(72, 61, 139);
        font-size:12px;
        margin-bottom:5px;
      }
      input {
        width: 379px;
        border: 1px solid gray;
        border-radius: 5px;
        height: 20px;
        margin-bottom: 20px;
        padding: 10px
      }
      .btn {
        background-color: rgb(247, 84, 25);
        color: white;
        width: 400px;
        border: none;
        border-radius: 5px;
        font-size:13px;
        height: 40px;
        cursor: pointer;
        &:hover {
          background-color: rgb(220,20,60);
        }
      }
    }
  }
  </style>