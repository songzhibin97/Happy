/******
** @创建时间 : 2020/8/19 14:01
** @作者 : SongZhiBin
******/
package websocket

import "github.com/gorilla/websocket"

// 定义了一些接口
type (
	// Management:Websocket管理对象包含的function
	IManagement interface {
		Start()                               // 初始化操作并启动守护线程处理一些内容
		GSend(message []byte, ignore IClient) // 向连接websocket的管道chan写入数据
		GetClients() map[int64]IClient
		GetBroadcast() chan []byte
		GetRegister() chan IClient
		GetUnregister() chan IClient
		Close()
	}
	// IClient:客户端定义
	IClient interface {
		// 主要实现读写分离
		Read(IManagement)  // 读入协程
		Write(IManagement) // 写入协程
		GetID() int64      // 获取唯一标识
		GetSocket() *websocket.Conn
		GetSendChan() chan []byte
	}
)
