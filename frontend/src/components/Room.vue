<script  setup>
import {ref, reactive} from 'vue'
import { onMounted, onUpdated, onUnmounted } from 'vue'
import {
  Sort,
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
  Download,
  Back
} from '@element-plus/icons-vue'
import { Timer } from '@element-plus/icons-vue'
// import { isDark } from '~/composables/dark'
// import {Greet} from '../../wailsjs/go/main/App'

const count = ref(2)

onMounted(
  () => {
  window.runtime.EventsOn("roomMsg", (roomMsg)=>{
        console.log(roomMsg)
        roomdata.msgList.push(roomMsg)
  })

  window.runtime.EventsOn("roomName", (roomName)=>{
        roomdata.name = roomName
  })

  window.runtime.EventsOn("userList", (res)=>{
         console.log(res)
         var userList = res.split(",")
         roomdata.userList = []
         for (let user of userList) {
           roomdata.userList.push(user)
         }
  })
  }

)

const roomdata = reactive({
  name: "UDPyyds",
  inMsg: "Why not connect?",
  msgList: [],
  userList: []
})



function send(){
      window.go.main.client.Broadcast(roomdata.outMsg).then(res => {
        //  this.title = res
        //  roomdata.inMsg = res
          roomdata.outMsg = ""
       })
      // window.runtime.EventsEmit("send", roomdata.outMsg)
 }

 function quitRoom() {
      window.go.main.client.QuitRoom().then(()=>{}
      )

 }


const handleRequest = (row) => {
  console.log(row)
  console.log("req " + row.fileName)
  window.go.main.App.Handle("req " + row.fileName)
}


</script>

<template>
    
    <div class="home-Warp">

    <el-container class="main_container">

      <el-header>
         <el-row >
           <el-col :span="2" :offset="0">
              <el-button :icon="Back" size="small" style="padding: 7px" @click="quitRoom" ></el-button>
           </el-col>
            
           <el-col :span="10" :offset="6">
             <div >
                  {{roomdata.name}}
             </div>
           </el-col>
         </el-row>
         
        
         
      </el-header>

      <el-container>
        <el-aside width="70px" >
          <div>
        <el-popover v-for="item in roomdata.userList" :key="item"
    placement="right"
    :width="100"
    trigger="hover"
  >
    <template #reference>
       <el-avatar > {{item.substring(0,2)}} </el-avatar>
    </template>
      我是{{item}}
  </el-popover>
     
          </div>
          </el-aside>
      <el-container>

          <el-main >
              <el-scrollbar max-height="500px">
    <p v-for="item in roomdata.msgList" :key="item" class="scrollbar-demo-item">
      {{ item }}
    </p>
  </el-scrollbar>
            <!-- <div id="inMsg" class="result">{{ roomdata.inMsg }}</div> -->
          </el-main>


          <el-footer>
               
             <div id="input2" class="msginput">
      <!-- <input id="name" v-model="roomdata.outMsg" autocomplete="on" class="input" type="text" @pressEnter="send"/> -->
      <el-input v-model="roomdata.outMsg" placeholder="Please input" clearable  @keyup.enter.native="send" />
      <!-- <button class="btn" @click="send">Send</button> -->
    </div>

          </el-footer>
       
        </el-container>
    
      </el-container>
   
    </el-container>
  </div>


    


</template>

<style scoped >

/* .homeWrap {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.main_container {
  height: 100%;
} */
.el-container{

height:100%

}
.el-main {
  position: relative;
}

.el-header {
  margin-top: 20px;
  font-size: 25px;
  /* font-family:Cambria, Cochin, Georgia, Times, 'Times New Roman', serif; */
  color: rgb(187, 193, 192);
}

.el-footer {
  position: absolute;
  margin-top: 30px;
  bottom: 0px;
  width: 400px;
}

.scrollbar-demo-item {
  display: flex;
  align-items: center;
  /* justify-content: center; */
  height: 50px;
  margin: 10px;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
  padding-left: 10px;
}

.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}

.main {
  height: 100%;
  width: 100%;
  background: #202124;
  border-radius: 6px;
  border: 0; 
}

.el-button:focus  {
    color: #cfd9df ;
    border-color: #cfd9df;
    background-color: transparent;
    outline: 0;
}
.el-button:hover {
    color: #cfd9df ;
    border-color: #cfd9df;
    background-color: #606d93
}

.el-button {
    --el-button-bg-color: transparent; 
    --el-button-hover-text-color: var(--el-color-primary);
    --el-button-hover-bg-color: var(--el-color-primary-light-9);
    --el-button-hover-border-color: var(--el-color-primary-light-7);
    --el-button-active-text-color: transparent;
    --el-button-active-border-color:  transparent;
    --el-button-active-bg-color: transparent;
}

.el-input__wrapper {
   background-color: transparent;
   
}
.el-main {
  /* color:  transparent; */
}


.el-input__inner {
    color: #e2ebf0;
}
/* .el-input__wrapper.is-focus {
   box-shadow:#d09123
} */
.el-table {
    --el-table-border-color: var(--el-border-color-lighter);
    --el-table-border: 1px solid var(--el-table-border-color);
    --el-table-text-color: var(--el-text-color-regular);
    --el-table-header-text-color: var(--el-text-color-secondary);
    --el-table-row-hover-bg-color: var(--el-fill-color-light);
    --el-table-current-row-bg-color: var(--el-color-primary-light-9);
    --el-table-header-bg-color: #47479400;
    --el-table-fixed-box-shadow: var(--el-box-shadow-light);
    --el-table-bg-color: transparent;
    --el-table-tr-bg-color: transparent;
    --el-table-expanded-cell-bg-color: var(--el-fill-color-blank);
    --el-table-fixed-left-column: inset 10px 0 10px -10px rgba(0, 0, 0, 0.15);
    --el-table-fixed-right-column: inset -10px 0 10px -10px rgba(0, 0, 0, 0.15);
}

/* 这个能生效，.root又不能生效 */
.el-input {
  --el-input-bg-color: transparant;
  --el-input-text-color: #e2ebf0;
  --el-input-border-color:#61718d;
  --el-input-focus-border-color: #e2ebf0;
}


.msginput {
  margin-left: 0;
  margin-right: 0;
  
}


.result {
  height: 20px;
  line-height: 20px;
  margin-bottom: 40px;
  margin-top: 120px;
  background: #202124;
  font-size: 30px;
  /* margin: 1.5rem auto; */
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 30px 0 0 20px;
  padding: 0 0px;
  cursor: pointer;
  font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
