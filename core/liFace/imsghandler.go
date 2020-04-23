package liFace

/*
	消息管理抽象层
 */
type IMsgHandle interface{
	DoMsgHandler(request IRequest, respond IMessage)       //马上以非阻塞方式处理消息
	AddRouter(router IRouter)            //为消息添加具体的处理逻辑
	StartWorkerPool()                    //启动worker工作池
	StopWorkerPool()                     //关闭worker工作池
	SendMsgToTaskQueue(request IRequest) //将消息交给TaskQueue,由worker进行处理
}