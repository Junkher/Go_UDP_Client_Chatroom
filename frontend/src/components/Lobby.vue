<script  setup>
import { ref } from 'vue'
import {reactive} from 'vue'
import { onMounted } from 'vue'
import {
  Right,
  Sort,
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
  Aim,
  ToiletPaper,
  Close,
  Plus,
  User,
} from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const createRoom = () => {
  ElMessageBox.prompt('请输入房间名', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputPattern:  /^[^@, ]{1,15}$/,
    inputErrorMessage: '房间名长度1~15，且不能包含,@以及空格',
  })
    .then(({ value }) => {
      window.go.main.client.CreateRoom(value).then(res => {
       if(res) {
       ElMessage({
            type: 'success',
          message: `${value}创建成功`,
          duration: 2000,
         })
       } else {
            ElMessage({
            type: 'warning',
          message: '房间名重复',
          duration: 1500,
         })
       }

      })

    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消输入',
      })
    })
}

const nick = () =>{
     ElMessageBox.prompt('请输入昵称', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
         inputPattern:  /^[a-zA-Z0-9_\u4e00-\u9fa5]{1,10}$/,
            inputErrorMessage: '昵称长度1~10, 可包含中文、数字、字母以及下划线',
  })
    .then(({ value }) => {
      window.go.main.client.Nick(value).then(

       ElMessage({
            type: 'success',
          message: `我记住了你的昵称：${value}`,
          duration: 2000,
       
         })
     )
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消输入',
      })
    })
}


onMounted(
  () => {

  window.runtime.EventsOn("lobbyMsg", (msg)=>{
        data.lobbyMsg = msg
        // data.listFiles = false
  })
  // window.runtime.EventsOn("ls", (res)=>{
  //       var fileNames = res.split(",");
  //       for(let item of fileNames){
  //               data.tableData.push(
  //                 {
  //                   fileName: item 
  //                 }
  //               )
  //       }
  //       data.listFiles = true
  //       data.inMsg = "别看我，看表!"
  // })
  window.runtime.EventsOn("roomList", (res)=>{
       console.log(res)
      var roomList = res.split(",");
      data.tableData = []
      for(let room of roomList){
                console.log(room)
                var s = room.split("@")
                console.log(s[0])
                console.log(s[1])
                data.tableData.push(
                  {
                    roomName: s[0],
                    roomMember: s[1],
                  }
                )
      }})
 

  window.runtime.EventsOn("connect", (res)=>{
      console.log(res)
      data.isConnect = res
  })

  }
)

const data = reactive({
  name: "",
  resultText: "Please enter port below 👇",
  port: "0.0.0.0:8000",
  lobbyMsg: "Why not connect?",
  // inMsg: "Why not connect?",
  tableData: [{roomName:"", roomMember: ""}],
  isConnect: false,
})



function connect() {
    // data.isConnect =true
    window.go.main.client.Connect(data.port).then(() => {
         console.log(data.port)
        //  this.port = ""
        // window.go.main.client.GetRoomList()
       })
}

function disconnect() {
  //  data.isConnect = false
   window.go.main.client.Disconnect()
}


const enterRoom = (row) => {
  console.log(row)
  console.log("/join " + row.roomName)
  window.go.main.client.EnterRoom(row.roomName)
}



</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        
      <div class="input1" style="margin-bottom: 30px">
        <el-row >
        <el-col :span="10" :offset="7">
              <el-input
            v-model="data.port"
            placeholder="输入IP地址:端口">
            <template #append>
              <!-- <el-button-group >
    <el-button type="primary" :icon="Sort" style="padding: 9px" />
    <el-button type="danger" :icon="Sort"  style="padding: 9px"/>
  </el-button-group> -->
              <el-button  v-if="!data.isConnect" :icon="Aim"  style="padding: 10px" @click="connect" />
              <el-button  v-if="data.isConnect" :icon="Close"  style="padding: 10px" @click="disconnect" />
            </template>
        </el-input>
         
        </el-col>
        <el-col :span="3" :offset="4">
           <el-button v-if="data.isConnect"  :icon="User" circle @click="nick"/>
        </el-col>
      </el-row>
    </div>

                 {{data.lobbyMsg}}
      
    
      </el-header>
          
          
          <el-main style="magin-top: 60px">
            
    
         <el-table v-if="data.isConnect" :data="data.tableData" style="width: 100%" class="my-class">

    <el-table-column label="房间名" >
      <template #default="scope">
           {{ scope.row.roomName }}
      </template>
    </el-table-column>

    <el-table-column label="人数" width="150">
      <template #default="scope">
           {{ scope.row.roomMember }}
      </template>
    </el-table-column>

    <el-table-column label="进入" width="100">
      <template #default="scope">
        <el-button :icon="Right" style="padding: 9px" @click="enterRoom(scope.row)" ></el-button>
  
      </template>
    </el-table-column>
  </el-table>

          </el-main>
          <el-footer v-if="data.isConnect">
               <el-button :icon="Plus"  @click="createRoom" >创建房间</el-button>
          </el-footer>
    </el-container>
  </div>
</template>

<style scoped >




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

.el-header {
  font-size: 17px;
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

  padding: 30px;
  padding-top: 60px;
}


.el-input__inner {
    color: #e2ebf0;
}
/* .el-input__wrapper.is-focus {
   box-shadow:#d09123
} */

/* 这个能生效，.root又不能生效 */
.el-input {
  --el-input-bg-color: transparant;
  --el-input-text-color: #e2ebf0;
  --el-input-border-color:#61718d;
  --el-input-focus-border-color: #e2ebf0;
}


.msginput {
  margin-left: 20px;
  margin-right: 20px;
  
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

.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}

</style>

