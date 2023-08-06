<template>
    <Popup v-on:turn-off="turnoff">
      <div class="box">
        <div class="h">
          <div class="text">企業ロゴを変更</div>
          <span class="material-symbols-outlined" @click="turnoff">close</span>
        </div>
        <div class="imgarea">
          <img :src="NewImage || ImagePreview" id="img-preview" />
          <input
            accept="image/*"
            type="file"
            id="file-input"
            @change="previewFiles($event)"
          />
          <label for="file-input">
            <div class="cap1">
              <span class="material-symbols-outlined"> photo_camera </span>
              <div class="text">画像のアップロード</div>
            </div>
          </label>
  
          <div class="cap2" v-if="NewImage" @click="back">
            <span class="material-symbols-outlined"> delete </span>
            <div class="text">ロゴを削除する</div>
          </div>
        </div>
        <button class="btn" @click="savechanges">変更する</button>
      </div>
    </Popup>
  </template>
  
  <script>
  import { ref } from "vue";
  import Popup from "../Popup.vue";
  
  export default {
    name: "UpdateAvatar",
    props: ['srcProp'],
    components: { Popup },
    setup(props, context) {
      const file = ref(false)
      const turnoff = () => {
        context.emit("no-pop-up");
      };
      const ImagePreview = ref(props.srcProp);
      var NewImage = ref("");
  
      return { turnoff, NewImage, ImagePreview ,file};
    },
    methods: {
      previewFiles(event) {
        if (event.target.files[0]) {
          this.file = event.target.files[0]
          var src = URL.createObjectURL(event.target.files[0]);
          this.NewImage = src;
        }
      },
      back() {
        this.NewImage = "";
        this.file = false
      },
      savechanges(){
      if (this.file){
        this.$emit("image-change",this.file)
      }
      else
      {
        this.$emit("no-pop-up")
      }}
    },
  };
  </script>
  
  <style lang = "scss" scoped>
  .box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-evenly;
    background-color: white;
    height: 500px;
    width: 500px;
    border-radius: 5px;
    padding: 10px;
    .h {
      color: rgb(72, 61, 139);
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      width: 400px;
      align-items: none;
      .material-symbols-outlined {
        color: black;
        cursor: pointer;
        &:hover {
          opacity: 0.5;
        }
      }
    }
    .imgarea {
      background-color: rgb(255, 239, 213);
      border: 3px dashed rgb(247, 84, 25);
      color: rgb(247, 84, 25);
      width: 400px;
      height: 300px;
      display: flex;
      flex-direction: column;
      align-items: center;
      img {
        border-radius: 50%;
        height: 75px;
        width: 75px;
        margin-top: 100px;
      }
      #file-input {
        display: none;
      }
      .cap1 {
        color: rgb(247, 84, 25);
        display: flex;
        margin-top: 20px;
        align-items: center;
        cursor: pointer;
        &:hover {
          color: black;
        }
  
        .text {
          margin-left: 20px;
        }
      }
      .cap1 {
        color: rgb(247, 84, 25);
        display: flex;
        margin-top: 20px;
        align-items: center;
  
        .text {
          margin-left: 20px;
        }
      }
      .cap2 {
        color: rgb(247, 84, 25);
        display: flex;
        margin-top: 20px;
        align-items: center;
        cursor: pointer;
        &:hover {
          color: black;
        }
  
        .text {
          margin-left: 20px;
        }
      }
    }
    .btn {
      background-color: rgb(247, 84, 25);
      color: white;
      border-radius: 5px;
      width: 400px;
      height: 50px;
      transition: 0.2s ease-in-out;
      cursor: pointer;
      &:hover {
        background-color: rgb(220, 20, 60);
      }
      border:none;
    }
  }
  </style>