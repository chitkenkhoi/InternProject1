<template>
  <Popup v-on:turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div class="text">すべてのメンバーに二段階認証を要求する</div>
        <span class="material-symbols-outlined" @click="turnOff">close</span>
      </div>
      <div class="textbox">
        <div>すべてのメンバーに二段階認証設定を要求しますか？</div>
        <div>次回メンバーがログインした際に</div>
        <div>二段階認証を設定する必要があります。</div>
      </div>
      <button class="btn" @click="toSwitch">変更する</button>
    </div>
  </Popup>
</template>
  
  <script>
import Popup from "../Popup.vue";
export default {
  name: "ConfirmOn",
  components: { Popup },
  setup(_, context) {
    const turnOff = () => {
      context.emit("no-pop-up");
    };
    const toSwitch = () => {
      context.emit("to-switch");
    };
    return { turnOff, toSwitch };
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
  }
  .btn {
    background-color: rgb(247, 84, 25);
    color: white;
    width: 400px;
    border: none;
    border-radius: 5px;
    height: 40px;
    cursor: pointer;
    &:hover {
      background-color: rgb(220, 20, 60);
    }
  }
}
</style>