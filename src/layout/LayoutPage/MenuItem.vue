<template>
  <div class="item">
    <router-link
      :to="todoProp.path"
      :class="[todoProp.id === 4 ? 'button-special' : 'button']"
      @click="toggle(todoProp.id)"
    >
      <span class="text">
        {{ todoProp.tittle }}
      </span>
      <span
        class="material-icons"
        v-if="todoProp.id === 4"
        :class="{ 'rotate': rotate }"
      >
        chevron_right
      </span>
    </router-link>
    <span v-if="todoProp.id === 4 && showProp" class="dropdown">
      <router-link to="/management/company" class="btn">
        <span class="textdrop">企業の詳細</span></router-link
      >
      <router-link to="/management/member" class="btn">
        <span class="textdrop">メンバー一覧</span></router-link
      >
      <router-link to="/management/error" class="btn">
        <span class="textdrop">グループ</span></router-link
      >
      <router-link to="/management/error" class="btn">
        <span class="textdrop">請求書</span></router-link
      >
    </span>
  </div>
</template>

<script>
import { ref } from "vue";
export default {
  name: "MenuItem",
  props: ["todoProp", "showProp"],
  setup(_, context) {
    const rotate = ref(false);
    const toggle = (id) => {
      if (id != 4) {
        context.emit("set");
        rotate.value = false;
        console.log(rotate.value);
      } else {
        context.emit("toggle");
        rotate.value = !rotate.value;
      }
    };
    return { toggle, rotate };
  },
};
</script>

<style lang = "scss" scoped>
.item {
  display: flex;
  flex-direction: column;
  height: 40;
  .rotate{
    transform: rotate(90deg);
  }
  
  .button {
    text-decoration: none;
    transition: 0.1s ease-in-out;
    padding: 0.5rem 1rem;
    .text {
      color: white;
      transition: 0.1s ease-in-out;
      margin-left: 15px;
    }

    &:hover {
      background-color: rgb(119, 136, 153);
      color: white;
      border-left: 3px solid rgb(247, 84, 25);
    }

    &.router-link-exact-active {
      background-color: rgb(119, 136, 153);
      border-left: 3px solid rgb(247, 84, 25);

      .text {
        color: rgb(247, 84, 25);
      }
    }
  }
  .button-special {
    display: flex;
    justify-content: space-between;
    align-items: center;
    text-decoration: none;
    transition: 0.1s ease-in-out;
    padding: 0.5rem 1rem;
    .text {
      color: white;
      transition: 0.1s ease-in-out;
      margin-left: 15px;
    }
    .material-icons {
      color: white;
    }
    .rotate{
      transform: rotate(90deg);
    }
    &:hover {
      background-color: rgb(119, 136, 153);
      color: white;
    }

    &.router-link-active {
      .text {
        color: rgb(247, 84, 25);
      }
      .material-icons {
        color: rgb(247, 84, 25);
      }
    }
  }
  .dropdown {
    display: flex;
    flex-direction: column;
    .btn {
      text-decoration: none;
      transition: 0.1s ease-in-out;
      padding: 0.5rem 1rem;
      .textdrop {
        color: white;
        transition: 0.1s ease-in-out;
        margin-left: 30px;
      }

      &:hover {
        background-color: rgb(119, 136, 153);
        color: white;
      }

      &.router-link-exact-active {
        background-color: rgb(119, 136, 153);
        border-left: 3px solid rgb(247, 84, 25);

        .textdrop {
          color: white;
        }
      }
    }
  }
}
</style>