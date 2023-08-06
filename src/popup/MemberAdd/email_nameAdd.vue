<template>
  <Popup class="popup" @turn-off="turnoff">
    <div class="container">
      <div class="h">
        <div class="tittle">メンバーを招待する</div>
        <div class="style">
          <div class="text">{{ role_JP }}</div>
        </div>
        <span class="material-symbols-outlined" id="x" @click="turnoff"
          >close</span
        >
      </div>
      <div class="boxtext">
        <div class="tittle">メンバーリスト</div>
        <div class="btnAdd" @click="add">
          <span class="material-symbols-outlined"> add </span>
          <div>メンバーを招待する</div>
        </div>
      </div>
      <div class="AddingList">
        <div class="AddingItem" v-for="mem in state.AddingList" :key="mem.id">
          <div class="box">{{ mem.email }}</div>
          <div class="box" id="box2">{{ mem.name }}</div>
          <div class="del" @click="del(mem.id)">
            <span class="material-symbols-outlined"> close </span>
          </div>
        </div>
      </div>

      <div class="AddingInput">
        <input
          ref = "ref_email"
          type="text"
          v-model="state.email"
          placeholder="メールアドレス"
          @keyup.enter="focus_next"
        />
        <input
          ref = "ref_name"
          type="text"
          v-model="state.name"
          placeholder="名前"
          @keyup.enter="add"
        />
        <div class="delete">
          <span class="material-symbols-outlined"> close </span>
        </div>
      </div>
      <div class="controll">
        <div class="back" @click="back">
          <span class="material-symbols-outlined"> keyboard_backspace </span>
          <div class="t" id="t1">本会員権限選択へ戻る</div>
        </div>
        <div
          class="save"
          @click="save"
          :class="{ isActive: state.AddingList.length !== 0 }"
        >
          <div class="t" id="t2">次へ</div>
          <span class="material-symbols-outlined"> keyboard_backspace </span>
        </div>
      </div>
      <div class="space"></div>
    </div>
  </Popup>
</template>
    
    <script>
import Popup from "../Popup.vue";
import { reactive, computed } from "vue";
export default {
  name: "EmailNameAdd",
  props: ["role"],
  components: { Popup },
  setup(props) {
    const state = reactive({
      AddingList: [],
      email: "",
      name: "",
      alert: false,
    });
    const role_JP = computed(() => {
      var a = "";
      switch (props.role) {
        case "Owner":
          a = "オーナー";
          break;
        case "Manager":
          a = "マネージャー";
          break;
        case "Member":
          a = "メンバー";
          break;
      }
      return a;
    });
    return { state,role_JP };
  },
  methods: {
    back() {
      this.$emit("back");
    },
    save() {
      if (this.state.AddingList.length === 0) return;

      this.$emit("save", this.state.AddingList);
      this.state.AddingList = [];
      this.state.email = "";
      this.state.name = "";
    },
    turnoff() {
      this.$emit("no-pop-up");
    },
    next() {
      this.$emit("next");
    },
    clean_data(search) {
      search = search.trim(); // *
      search = search.split(" ");
      var a = [];
      for (let i = 0; i < search.length; i++) {
        if (search[i] !== "") {
          a = a.concat([search[i]]); // clean data from "   Le    Quang Khoi   " to "Le Quang Khoi"
        }
      }
      search = "";
      for (let i = 0; i < a.length; i++) {
        search = search + " " + a[i];
      }
      search = search.trim();
      return search;
    },
    focus_next(){
      
      this.$refs.ref_name.focus()
    },
    add() {
      if (!this.state.email || !this.state.name) {
        this.state.alert = true;
      } else {
        if (
          /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(this.state.email)
        ) {
          this.state.name = this.clean_data(this.state.name);
          var id = this.state.AddingList.length + 1;
          const member = {
            id: id,
            email: this.state.email,
            name: this.state.name,
          };
          this.state.AddingList.push(member);
          this.state.email = "";
          this.state.name = "";
          this.state.alert = false;
          this.$refs.ref_email.focus()
        } else {
          this.state.alert = true;
        }
      }
    },
    del(id) {
      this.state.AddingList = this.state.AddingList.filter(
        (ele) => ele.id !== id
      );
      var x = id - 1;
      var size = this.state.AddingList.length - 1;
      while (x <= size) {
        this.state.AddingList[x].id -= 1;
        x++;
      }
    },
  },
};
</script>
    
<style lang="scss" scoped>
.container {
  background-color: white;
  z-index: 99;
  width: 600px;
  border-radius: 5px;
  margin-top: -40px;
  display: flex;
  flex-direction: column;
  align-items: center;
  .h {
    display: flex;
    width: 500px;
    height: 30px;
    align-items: center;
    justify-content: space-between;
    margin-top: 30px;
    .tittle {
      font-family: "Noto Sans JP";
      font-weight: 400;
      font-size: 16px;
      color: #271b5a;
    }
    .style {
      margin-left: -130px;
      padding:5px;
      // width: 70px;
      height: 20px;
      display: flex;
      align-items: center;
      color: #d64d10;
      background-color: #f8f2e9;
      font-size: 12px;
      border-radius: 30px;
      .text {
        margin-left: 5px;
        margin-top: 3px;
      }
    }
    #x {
      cursor: pointer;
      &:hover {
        opacity: 50%;
      }
    }
  }
  .boxtext {
    display: flex;
    justify-content: space-between;
    margin-top: 49px;
    width: 500px;
    height: 35px;
    border-bottom: 3px solid black;
    .tittle {
      font-family: "Noto Sans JP";
      font-weight: 400;
      font-size: 16px;
      color: #271b5a;
    }
    .btnAdd {
      display: flex;
      color: #d64d10;
      font-size: 14px;
      font-family: "Noto Sans JP";
      font-weight: 400;
      margin-right: 15px;
      cursor: pointer;
    }
  }
  .AddingList {
    display: flex;
    flex-direction: column;
    position: relative;
    left: 25px;
    width: 550px;
    max-height: 400px;
    overflow-y: auto;

    .AddingItem {
      margin-top: 10px;
      display: flex;
      flex-direction: row;
      align-items: center;
      width: 500px;
      .box {
        width: 200px;
        height: 18px;
        padding: 15px;
        border-radius: 4px;
        border: solid 1px #aaaaaa;
      }
      #box2 {
        margin-left: 7px;
      }
      .del {
        width: 20px;
        height: 20px;
        border-radius: 50%;
        border: 2px solid #d00d41;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-left: 7px;
        .material-symbols-outlined {
          font-size: 18px;
          color: #d00d41;
        }
        cursor: pointer;
      }
    }
  }
  .AddingInput {
    margin-top: 10px;
    display: flex;
    flex-direction: row;
    height: 52px;
    width: 500px;
    justify-content: space-between;
    align-items: center;
    input {
      width: 200px;
      height: 18px;
      padding: 15px;
      border-radius: 4px;
      border: solid 1px #aaaaaa;
      :focus {
        outline: none;
      }
    }
    .delete {
      width: 20px;
      height: 20px;
      border-radius: 50%;
      border: solid 2px #aaaaaa;
      color: #aaaaaa;
      display: flex;
      justify-content: center;
      align-items: center;

      .material-symbols-outlined {
        font-size: 18px;
      }
    }
  }
  .controll {
    margin-top: 21px;
    width: 500px;
    height: 20px;
    display: flex;
    justify-content: space-between;
    .back {
      display: flex;
      color: #d64d10;
      .t {
        font-size: 14px;
        font-family: "Noto Sans JP";
        font-weight: 400;
        cursor: pointer;
      }
    }
    .save {
      display: flex;
      color: #aaaaaa;
      .material-symbols-outlined {
        margin-top: 24px;
        color: #aaaaaa;
        transform: rotate(180deg);
      }
    }
    .isActive {
      color: #d64d10;
      cursor: pointer;
      .material-symbols-outlined {
        color: #d64d10;
      }
    }
  }
  .space {
    height: 50px;
  }
}
</style>