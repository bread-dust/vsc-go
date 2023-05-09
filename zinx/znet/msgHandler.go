package znet

import (
	"fmt"
	"strconv"
	"ziface"
	"settings"
)

/*
	消息管理模块
*/

type MsgHandle struct{
	// 存放每个MsgID所对应的处理方法
	Apis map[uint32] ziface.IRouter

	// 负责Worker取任务的消息队列
	TaskQueue []chan ziface.IRequest	
	// 业余工作Worker工作池的worker数量
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32] ziface.IRouter),
		WorkerPoolSize: settings.GlobalObject.WorkerPoolSize, // 从全局配置中获取
		TaskQueue: make([]chan ziface.IRequest, settings.GlobalObject.WorkerPoolSize),
	}
}

// 调度、执行对应的Router消息处理方法
func (mh *MsgHandle)DoMsgHandler(request ziface.IRequest){
	// 从request中找到msgID
	handler,ok:=mh.Apis[request.GetMsgId()]
	if !ok{
		fmt.Println("api msgId = ",request.GetMsgId(),"is not FOUND!")
		return
	}

	// 根据msgID调度对应router业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// 为消息添加具体的处理逻辑
func (mh *MsgHandle)AddRouter(msgId uint32,router ziface.IRouter){
	// 判断当前msg绑定的API处理方法是否已经存在

	if _,ok := mh.Apis[msgId];ok{
		// id已经注册了
		panic("repeat api,msgId = " + strconv.Itoa(int(msgId)))
	}

	// 添加msg与API的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api MsgID = ",msgId,"succ!")
}

// 启动一个Worker工作池(开启工作池的动作只能发生一次，一个zinx框架只能有一个worker工作池)
func (mh *MsgHandle)StartWorkerPool(){
	// 根据workerPoolSize 分别开启Worker，每个Worker用一个go来承载
	for i:=0;i<int(mh.WorkerPoolSize);i++{
		// 一个worker被启动
		// 1 当前的worker对应的channel消息队列 开辟空间 第0个worker 就用第0个channel
		mh.TaskQueue[i] = make(chan ziface.IRequest,settings.GlobalObject.MaxWorkerTaskLen)
		// 2 启动当前的Worker，阻塞等待消息从channel传递进来
		go mh.StartOneWorker(i,mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandle)StartOneWorker(workerID int,taskQueue chan ziface.IRequest){
	fmt.Println("Worker ID = ",workerID,"is started...")
	// 不断的阻塞等待对应消息队列的消息
	for{
		select{
		// 如果有消息过来，出列的就是一个客户端的Request，执行当前Request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}

}
