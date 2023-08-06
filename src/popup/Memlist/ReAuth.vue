<template>
    <Popup v-on:turn-off="turnoff">
      <div class="container">
        <div class="h">
          <div class="text">二段階認証をリセット</div>
          <span class="material-symbols-outlined" @click="turnoff">close</span>
        </div>
        <div class="textbox">
          <div>このメンバーの二段階認証の設定をリセットしますか？</div>
          <div>次回ログインした際に</div>
          <div>二段階認証の再設定が要求されます。</div>
        </div>
        <button class="btn" @click="reset">リセットする </button>
      </div>
    </Popup>
  </template>
    
    <script>
  import Popup from "../Popup.vue";
  export default {
    name: "ConfirmOn",
    components: { Popup },
    setup(_, context) {
      const turnoff = () => {
        context.emit("no-pop-up");
      };
      const reset = () => {
        context.emit("to-reset");
      };
      return { turnoff, reset };
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