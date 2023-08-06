<template>
  <Popup @turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div>メンバーの招待</div>
        <span class="material-symbols-outlined" @click="turnoff">close</span>
      </div>
      <div class="boxtext">
        <div>招待するメンバーの権限を選んでください。</div>
        <div>選択した権限は後から変更することも可能です。</div>
      </div>
      <div class="boxrole">
        <div class="role">
          <div class="box">
            <div class = "head">オーナー</div>
            <div class = "firstline">本サービス上の利用者コンテンツを使用でき、</div>
            <div>全ての権限を保有する利用者を意味します。</div>
          </div>
          <span class="material-symbols-outlined" @click = "next('Owner')"> chevron_right </span>
        </div>
        <div class="role">
          <div class="box">
            <div class = "head">マネージャー</div>
            <div class = "firstline">本サービス上の利用者コンテンツを使用でき、</div>
            <div>メンバーを追加・削除できる利用者を意味します。</div>
          </div>
          <span class="material-symbols-outlined" @click = "next('Manager')"> chevron_right </span>
        </div>
        <div class="role">
          <div class="box">
            <div class = "head">メンバー</div>
            <div class = "firstline">本サービス上の利用者コンテンツを使用できる利用者を意味します。</div>
          </div>
          <span class="material-symbols-outlined" @click = "next('Member')"> chevron_right </span>
        </div>
      </div>
    </div>
  </Popup>
</template>
  
  <script>
import Popup from "../Popup.vue";
export default {
  name: "RoleAdd",
  components: { Popup },
  setup() {},
  methods: {
    turnoff() {
      this.$emit("no-pop-up");
    },
    next(data){
      this.$emit('next',data)
    }
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
  height: 621px;
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
    color: #271b5a;
    font-size: 16px;
    font-family: "Noto Sans JP";
    font-weight: 500;
    .material-symbols-outlined {
      color: black;
      cursor: pointer;
      &:hover {
        opacity: 0.5;
      }
    }
  }
  .boxtext {
    width: 500px;
    height: 48px;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 50px;
    font-size: 14px;
    font-family: "Noto Sans JP";
    font-weight: 400;
    justify-content: space-between;
    color:#666666;
  }
  .boxrole{
    margin-top:39px;
    width: 500px;
    height: 373px;
    display:flex;
    flex-direction: column;
    justify-content: space-evenly;
    .role{
      width:500px;
      height:107px;
      border:2px lightgrey solid;
      border-radius:5px;
      display:flex;
      justify-content: space-evenly;
      align-items: center;
      .material-symbols-outlined {
      color: black;
      cursor: pointer;
      &:hover {
        opacity: 0.5;
      }
    }
      .box{
        font-size:12px;
        width:420px;
        height: 60px;
        display:flex;
        flex-direction: column;
        font-family: 'Noto Sans JP';
        font-weight: 400;
        color: #666666;
        .head{
          color:#D64D10;
          font-size:14px;
          font-family: 'Noto Sans JP';
          font-weight: 500;
        }
        .firstline{
          margin-top:5px;
        }
      }
    }
  }
}
</style>