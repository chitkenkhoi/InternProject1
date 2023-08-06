<template>
    <Popup v-on:turn-off="turnoff">
      <div class="container">
        <div class="h">
          <div class="text">メンバーを削除</div>
          <span class="material-symbols-outlined" @click="turnoff">close</span>
        </div>
        <div class="textbox">
          <div class= line1><div class = "name_email">{{ memberProp.name }} ({{ memberProp.email }})</div>を削除してもよろしいですか。</div>
          <div>メンバーは完全にアカウントにアクセスできなくなります。</div>
        </div>
        <button class="btn" @click="del">削除する</button>
      </div>
    </Popup>
  </template>
    
    <script>
  import Popup from "../Popup.vue";
  export default {
    name: "MemDelete",
    components: { Popup },
    props: ['memberProp'],
    setup(_, context) {
      const turnoff = () => {
        context.emit("no-pop-up");
      };
      const del = () => {
        context.emit("to-delete");
      };
      return { turnoff, del };
    },
  };
  </script>
    
    <style lang="scss" scoped>
  .container {
    margin-top: 105px;
    background-color: white;
    display: flex;
    border-radius: 5px;
    height: 268px;
    width: 500px;
    flex-direction: column;
    color: black;
    align-items: center;
    padding: 10px;
    justify-content: space-evenly;
    .h {
      width: 400px;
      display: flex;
      justify-content: space-between;
      .text {
        color: rgb(72, 61, 139);
        font-size: 16px;
        font-family: 'Noto Sans Jp';
        font-weight: 500;
      }
      .material-symbols-outlined {
        cursor: pointer;
        &:hover {
          opacity: 0.5;
        }
      }
    }
    .textbox {
      display: flex;
      flex-direction: column;
      align-items: center;
      font-size: 14px;
      color: #666666;
      font-family: 'Noto Sans Jp';
      font-weight: 400;
      .line1{
        display:flex;
        flex-direction: row;
        .name_email{
            color: #17B182;
        }
      }
    }
    .btn {
      background-color: rgb(247, 84, 25);
      color: white;
      width: 400px;
      border: none;
      border-radius: 5px;
      height: 40px;
      cursor: pointer;
      font-family: 'Noto Sans Jp';
      font-weight: 500;
      &:hover {
        background-color: rgb(220, 20, 60);
      }
    }
  }
  </style>