<template>
    <transition name="fade" appear>
    <div class="modal-overlay" @click = "turnoff"></div>
    </transition>
    <div class="content">
            <slot/>
        </div>
</template>
<script>
export default {
    name:"PopUp",
    methods:{
        turnoff(){
            this.$emit('turn-off')
        }
    }
}
</script>

<style lang = "scss" scoped>
.modal-overlay {
  position: fixed;
  display:flex;
  justify-content: center;
  align-items: center;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 98;
  background-color: rgba(0, 0, 0, 0.3);
}
.fade-enter-active {
  transition: opacity 0.5s ease;
}
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.content{
        position:fixed;
        top :150px;
        left:550px;
        z-index: 99;
    }

</style>