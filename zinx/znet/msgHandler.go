package znet

import (
	"fmt"
	"strconv"
	"ziface"
)

/*
	消息管理模块
*/

type MsgHandle struct{
	// 存放每个MsgID所对应的处理方法
	Apis map[uint32] ziface.IRouter
}

func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32] ziface.IRouter),
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


