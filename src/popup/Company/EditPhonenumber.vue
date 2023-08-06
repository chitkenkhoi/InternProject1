<template>
  <Popup v-on:turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div class="text">電話番号を変更する</div>
        <span class="material-symbols-outlined" @click="turnoff">close</span>
      </div>
      <div class="submitarea">
        <div class="h2">電話番号</div>
        <input
          type="text"
          placeholder="+81"
          v-model="data"
          @keyup.enter="savechange"
        />
        <div class = "alert" v-if = "alert">有効な値を入れてください</div>
        <button class="btn" @click="savechange">変更する</button>
      </div>
    </div>
  </Popup>
</template>

<script>
import Popup from "../Popup.vue";
import { ref } from "vue";
export default {
  name: "EditPhone",
  components: { Popup },
  props: ["popprop"],
  setup(props, context) {
    const alert = ref(false)
    const data = ref(props.popprop.data);
    const turnoff = () => {
      context.emit("no-pop-up");
    };
    const savechange = () => {
      if (data.value !== "") context.emit("data-change", data.value);
      else alert.value = true;
    };

    return { turnoff, savechange, data ,alert};
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
      font-size: 16px;
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
    .alert{
      position: absolute;
      font-size:10px;
      top:290px;
    }
    .h2 {
      color: rgb(72, 61, 139);
      font-size: 12px;
      margin-bottom: 5px;
    }
    input {
      width: 379px;
      border: 1px solid gray;
      border-radius: 5px;
      height: 20px;
      margin-bottom: 20px;
      padding: 10px;
    }
    .btn {
      background-color: rgb(247, 84, 25);
      color: white;
      width: 400px;
      border: none;
      border-radius: 5px;
      font-size: 13px;
      height: 40px;
      cursor: pointer;
      &:hover {
        background-color: rgb(220, 20, 60);
      }
    }
  }
}
</style>