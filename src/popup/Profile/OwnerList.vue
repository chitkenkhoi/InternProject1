<template>
  <div class = "item" :class = "{'first':eleProp.id == 1}">
    <img :src = "ava"/>
    <div class = "NameEmail">
        <div class = "name">{{ eleProp.name }}</div>
        <div class = "email">{{ eleProp.email }}</div>
    </div>
    <div class ="CompanyName">
        <div class = "inner">{{ eleProp.company }}</div></div>
  </div>
</template>

<script>
import ava from '@/layout/LayoutPage/img/avatar.jpg'
export default {
name: "OwnerList",
props: ['eleProp'],
setup(){return {ava}}
}
</script>

<style lang = "scss" scoped>
.item{
    margin-left:15px;
    height:65px;
    width:470px;
    display:flex;
    align-items: center;
    border-top:1px solid lightgrey;

    img{
        height:30px;
        width:30px;
        border-radius:50%;
        margin-left:10px;
    }
    .NameEmail{
        margin-left:20px;
        display:flex;
        flex-direction: column;
        width:240px;
        height:42px;
        justify-content: space-between;
    }
    .CompanyName{
        margin-left:10px;
        width:150px;
        height:30px;
        border-radius:30px;
        border: 1px solid lightgrey;
        display:flex;
        justify-content: center;
        align-items: center;
    }
}
.first{
    border-top:none;
}
</style>