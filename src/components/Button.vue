<template>
  <button class="button">
    <slot />
  </button>
</template>

<script>
export default {
  name: "ButtonCommon",
};
</script>

<style lang = "scss" scoped>
.button {
  border: none;
  display: flex;
  // justify-content: center;
  align-items: center;
  background-color: #d64d10;
  cursor: pointer;
  border-radius: 4px;
  font-size: 14px;
  font-family: "Noto Sans JP";
  font-weight: 500;
  &:hover {
    background-color: rgb(220, 20, 60);
  }
}
</style>