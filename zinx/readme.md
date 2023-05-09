## v0.1基础server
### 方法
1. 启动服务器
    - 基本的服务器开发 1. 创建Addr，2. 创建lisetener 3. 处理客户端基本的业务功能，回显功能
2. 运行服务器
    - 调用start()方法，调用之后做阻塞处理，在之间可以做一个扩展功能
3. 停止服务器

## V0.2 简单的链接封装和业务绑定
### 链接模块
- 方法
  - 启动链接Start()
  - 停止链接Stop()
  - 获取当前链接的conn对象(套接字)
  - 得到链接ID
  - 得到客户端链接的地址和端口
  - 发送数据的方法Send()

- 属性
  - socket TCP套接字
  - 链接ID
  - 链接状态是否关闭
  - 与当前链接绑定的业务处理方法API
  - 等待被告知关闭的channel

## V0.3 基础route模块
### Request 请求封装 
- 将链接和数据绑在一起
- 属性
  - 链接Connection
  - 请求数据
- 方法
  - 得到当前链接
  - 得到当前数据
### Route模块
- 抽象的IRouter：
  - PreHandle 处理业务之前的方法
  - Handle 处理业务的方法
  - PostHandle 处理业务之后的方法
- 实现的BaseRouter
- 集成Router模块
  - IServer 添加路由功能
  - Server 类添加Router成员
  - Connection 类绑定一个Router成员
  - Connection  调用已经注册的Router处理业务

## v0.4 全局配置
- 全局配置模块
  - init -> 读取配置文件到Config结构体中
  - 将硬代码用Config结构体中的数据替换


## V0.5消息封装
### 定义一个消息的结构体
- 属性
  - DataLen uint32
  - ID uint32
  - Data []byte
### 解决粘包问题
### 封包拆包模块
- 针对Message进行TLV格式的封装
  - 写Message的长度
  - 写Message的ID
  - 写Message的内容
- 针对Message进行TLV格式的拆包
  - 先读取固定长度的head->消息内容的长度和消息的类型
  - 再根据消息内容的长度，再进行一次读写，从conn中读取内容
### 消息封装机制集成到Zinx中
- 将Message添加到Request中
- 修改链接读取数据的机制，将之前的byte读取改为TLV读取
- 给链接提供发包机制：将发送的消息打包再发送 

## V0.6 多路由模式
### 消息管理模块
- 属性
  - 消息ID和对应的路由关系-map
- 方法
  - 根据消息ID得到路由
  - 添加路由到map中

### 集成到Zinx框架中
- 在server模块的Route属性替换成MsgHandle
- 将server之前的AddRouter替换车鞥当调用MsgHandle的AddRouter方法
- 在connection模块中Route属性替换成MsgHandle
- Connection的之前调度Router业务，改为调度MsgHandle方法
  
## V0.7 读写分离
### reader和writer通信的channel

### 添加一个writer goroutine

### reader之前直接发送客户端，改成发送给Channel

### 启动reader和writer一起工作


## V0.8 消息队列及多任务
### 创建一个消息队列
- MsgHandler消息管理模块
  - 属性
    - 消息队列
    - worker工作池的数量
### 创建多任务worker的工作池
- 创建一个worker的工作池
  - 根据WorkerPoolSize创建worker
  - 每个worker都有一个go去处理消息队列中的消息
    - 阻塞等待与当前worker绑定的channel的消息
    - 一旦有消息，就执行当前消息所绑定的路由业务DomsgHandle()
    - 
### 将之前的发送消息改为发送给消息队列和worker池处理
- 定义一个方法，将消息发送给消息队列，
  - 保证每个worker分配的request是平均的
- 将消息发送个对应的worker的消息队列
### 将消息队列机制集成到Zinx框架中
- 开启饼干调用消息队列WorkerPool
- 将从客户端处理的消息，发送给当前的worker工作池处理

## V0.9 链接管理模块
### 创建一个链接管理模块
- 属性
  - 已经创建的链接集合
  - 读写锁

- 方法
  - 添加链接
  - 删除链接
  - 根据链接ID得到链接
  - 得到当前链接的总数
  - 清理并终止所有链接

### 将链接管理模块集成到Zinx框架中
- 将链接管理模块添加到Server中
- 每次成功与客户端建立链接之后，将链接添加到链接管理模块中
- 判断当前的链接总数，如果超过最大链接数，就将最早创建的链接关闭
- 每次与客户端断开链接之后，将链接从链接管理模块中删除
- 
### 链接管理模块添加一个创建链接之后/销毁链接之前索要处理的方法
- 添加的属性
  - 该Server创建之后自动调用的钩子函数，OnConnStart
  - 该Server销毁之前自动调用的钩子函数,OnConnStop
- 添加的方法
  - 注册OnConnStart钩子函数
  - 注册OnConnStop钩子函数
  - 调用OnConnStart钩子函数
  - 调用OnConnStop钩子函数
  - 
### 使用ZinxV0.9框架开发
- 注册OnConnStart钩子函数
- 注册OnConnStop钩子函数

## V1.0 链接属性
### 为链接添加配置属性的功能
- 新增属性
  - 链接属性集合--map
  - 保护链接的锁
- 新增方法
  - 设置链接属性
  - 获取链接属性
  - 移除链接属性
  