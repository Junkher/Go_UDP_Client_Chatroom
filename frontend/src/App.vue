<script setup>
// import Client from './components/Client.vue'
import Lobby from './components/Lobby.vue'
import Room from './components/Room.vue'
import {reactive} from 'vue'
import { onMounted} from 'vue'

const clientData = reactive({
   inRoom: false
})

// 根据inRoom来显示不同组件
onMounted(
  () => {
  window.runtime.EventsOn("inroom", (res)=>{
        clientData.inRoom = res
      })
  }
)


</script>

<template>
     <div v-if="clientData.inRoom">
        <Room/>
    </div>
    <div v-if="!clientData.inRoom">
        <Lobby/>
    </div>
</template>

<style>

:root {
  background: #202124;
  border-radius: 6px;
  /* --el-color-primary: #b2cd4f !important; */
}



#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
