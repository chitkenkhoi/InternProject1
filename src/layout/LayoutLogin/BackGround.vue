<template>
  <div class="Container">
    <div class="header">
      <img :src="logo" />
      <div class="backToHome">
        <span class="material-symbols-outlined"> arrow_back </span>
        <div class="text_BTH">ホームに戻る</div>
      </div>
    </div>
    <slot/>
    <div class="footer">
      <div class="Sign">Digital Sign ©2021</div>
      <div class="textFooter">利用規約</div>
      <div class="textFooter">特定商取引法に基づく表記</div>
      <div class="textFooter">プライパシーポリシー</div>
      <div class="textFooter">運営会社</div>
      <div class="textFooter">よくあるご質問</div>
      <div class="textFooter">お問い合わせ</div>
      <a href="https://www.youtube.com/"><img :src="ytbicon" /></a>
      <a href="https://line.me/en/"><img :src="lineicon" /></a>
      <a href="https://www.facebook.com/"><img :src="fbicon" /></a>
      <a href="https://twitter.com/"><img :src="twticon" /></a>
    </div>
  </div>
</template>

<script>
import logo from "../../../public/logo/Mask Group 36.png";
import ytbicon from "../../../public/icon/ytb.png";
import fbicon from "../../../public/icon/fb.png";
import lineicon from "../../../public/icon/line.png";
import twticon from "../../../public/icon/twt.png";

export default {
  name: "BackGround",
  setup() {
    return { logo, ytbicon, fbicon, lineicon, twticon };
  },
};
</script>

<style lang = "scss" scoped>
.Container {
  position: relative;
  background: linear-gradient(170deg, #201a39 0 50%, #272450 50% 100%);
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  // overflow-y: auto;
  width:100vw;
  height:1000px;
  display:flex;
  flex-direction: column;
  .header {
    position: absolute;
    top: 0;
    left: 0;
    width: 1500px;
    height: 80px;
    display: flex;
    align-items: center;
    img {
      height: 30px;
      width: 170px;
      position: absolute;
      top: 25px;
      left: 150px;
    }
    .backToHome {
      color: white;
      font-size: 14px;
      display: flex;
      width: 118px;
      height: 20px;
      position: absolute;
      
      left: 1232px;
      top: 30px;
      cursor:pointer;
      .material-symbols-outlined {
        width: 20px;
        height: 20px;
      }
      .text_BTH {
        margin-left: 10px;
        font-family: "Noto Sans JP";
        font-weight: 400;
      }
    }
  }
  .footer {
    display: flex;
    position: absolute;
    width: 1500px;
    height: 80px;
    bottom:0;
    left:0;
    align-items: center;

    .Sign{
        width: 116px;
        // height:20px;
        font-size:12px;
        margin-left: 150px;
        color:#AAAAAA;
        margin-top:3px;
    }
    .textFooter{
        font-size:12px;
        font-family: 'Noto Sans JP';
        font-weight: 400;
        margin-left:50px;
        color:white;
    }
    a{
        width:20px;
        height:20px;
        margin-left:30px;
    }
  }
}
</style>