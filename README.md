# README


使用Wails框架(go后端，vue前端)搭建的简单的chatroom客户端，根据UDP_Server_Chatroom定制，即服务器提供的指令封装成能被前端调用的函数。


## Live Development

在当前项目路径的命令行输入`wails dev`就会进入开发模式，支持前端的热部署

## Building

命令 `wails build`会将后端和前端打包成一个可执行的文件，也就是客户端的应用。

## Feature

- 查看房间列表

输入服务器的地址，点击按钮可以获取服务器的房间信息列表

<img src=img/Pasted%20image%2020220509000403.png width=48%> <img src=img/Pasted%20image%2020220508235230.png width=48%>

- 修改昵称

可以多次修改，在房间中将会显示用户的昵称，如果用户没有修改昵称，则默认显示IP:port

<img src=img/Pasted%20image%2020220508235408.png width=48%>   <img src=img/Pasted%20image%2020220508235518.png width=48%>

- 创建房间


点击创建房间，并输入房间名，就能够创建房间并同时自动进入该房间，打开另一个客户端就能看到房间列表已经更新

<img src=img/Pasted%20image%2020220509000008.png width=48%>   <img src=img/Pasted%20image%2020220508235706.png width=48%><img src=img/Pasted%20image%2020220509000807.png width=48%>   <img src=img/Pasted%20image%2020220509000904.png width=48%>

- 进入房间

点击房间的进入按钮便可进入指定房间，并向房间中的所有用户广播进入信息，再打开一个客户端就可以看到房间的信息也发生了变化

<img src=img/Pasted%20image%2020220509001025.png width=48%> <img src=img/Pasted%20image%2020220509001052.png width=48%>
<img src=img/Pasted%20image%2020220509001123.png width=48%> <img src=img/Pasted%20image%2020220509001238.png width=48%>

- 查看房间用户列表

在房间的左边侧栏显示用户的昵称，鼠标悬停可以看到完整昵称

<img src=img/Pasted%20image%2020220509001357.png width=48%>

- 多人聊天

在输入框输入要发送的信息，并且回车就能够发给房间中的所有人，

<img src=img/Pasted%20image%2020220509001534.png width=48%>

- 离开房间

点击左上角的离开按钮退出房间，并且会通知房间内的所有人，用户列表也会实时更新，如果该用户是房间中的最后一人，则该房间会被删除

<img src=img/Pasted%20image%2020220509001649.png width=48%>   <img src=img/Pasted%20image%2020220509001805.png width=48%>


## Wails

[wails官方文档](https://wails.io/)

零基础开始学go和wails，个人觉得以下两个概念比较关键，理解了后就能相对快速上手，虽然但是，官方文档有点难理解，缺少一些相关示例，可能是我菜了。

- Binding Methods

简单来说就是绑定在后端的结构体中实现的公有方法，让前端能够直接通过`window.go.main.<structName>.<methodName>`来调用

>This will bind all public methods in our App struct (it will never bind the startup and shutdown methods).

- [Events](https://wails.io/docs/reference/runtime/events)

EventsOn和EventsEmit应该是成对出现的。
通过EventsOn来监听事件，如果该事件被EventsEmit，就能触发对应的回调函数，并且能够传递数据。

> The Wails runtime provides a unified events system, where events can be emitted or received by either Go or Javascript. Optionally, data may be passed with the events. Listeners will receive the data in the local data types.




