package websocket

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

//  ClientManager:
type ClientManager struct {
	Clients    map[int64]IClient // map根据UID获取对应的Client
	Broadcast  chan []byte       // 发送给已经注册的所有客户端进行消息广播
	Register   chan IClient      //  注册
	Unregister chan IClient      // 取消注册 删除map对应信息
	IsClose    bool
}

// Client:客户端
type Client struct {
	ID      int64           // 唯一标识
	Socket  *websocket.Conn // websocket句柄
	Send    chan []byte     // 发送的数据包
	IsClose bool
}

// Message is an object for websocket message which is mapped to json type
type Message struct {
	Sender    int64  `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var M *ClientManager

func init() {
	M = NewManager(2)
	M.Start()
}

// Manager:初始化对象
func NewManager(buf int) *ClientManager {
	return &ClientManager{
		Broadcast:  make(chan []byte, buf),
		Register:   make(chan IClient, buf),
		Unregister: make(chan IClient, buf),
		Clients:    make(map[int64]IClient),
	}
}

// Start:开启服务并加入到进程
func (manager *ClientManager) Start() {
	go func() {
		// 哨兵
		for {
			if manager.IsClose {
				return
			}
			select {
			case conn := <-manager.Register:
				// 注册:
				// 1.判断在Clients是否已经注册过
				c, ok := manager.Clients[conn.GetID()]
				if ok {
					// 表示存在 进行上一个websocket还存在 关闭
					_ = c.GetSocket().Close()
				}
				// 赋值
				manager.Clients[conn.GetID()] = conn
				zap.L().Info("Register", zap.Any("conn", conn.GetID()))
				continue
			case conn := <-manager.Unregister:
				// 1.判断在Clients是否已经注册过
				c, ok := manager.Clients[conn.GetID()]
				if !ok {
					continue
				}
				close(c.GetSendChan())
				_ = c.GetSocket().Close()
				// 删除
				zap.L().Info("Unregister", zap.Any("conn", conn.GetID()))
				delete(manager.Clients, conn.GetID())
				continue
			case message := <-manager.Broadcast:
				// 遍历map向客户端发送消息
				for _, c := range manager.Clients {
					c.GetSendChan() <- message
				}
			}
		}
	}()
}

func (manager *ClientManager) Close() {
	// 资源回收
	for _, v := range manager.Clients {
		_ = v.GetSocket().Close()
	}
	close(manager.Broadcast)
	close(manager.Unregister)
	close(manager.Register)
	manager.IsClose = true
}

// GSend:广播消息
func (manager *ClientManager) GSend(message []byte, ignore IClient) {
	for _, conn := range manager.Clients {
		if conn.GetID() != ignore.GetID() {
			conn.GetSendChan() <- message
		}
	}
}

// GetClients:获取Clients
func (manager *ClientManager) GetClients() map[int64]IClient {
	return manager.Clients
}

// GetBroadcast:获取Broadcast
func (manager *ClientManager) GetBroadcast() chan []byte {
	return manager.Broadcast
}

// GetRegister:获取Broadcast
func (manager *ClientManager) GetRegister() chan IClient {
	return manager.Register
}

func (manager *ClientManager) GetUnregister() chan IClient {
	return manager.Unregister
}

// Read:阅读消息
func (c *Client) Read(m IManagement) {
	defer func() {
		m.GetUnregister() <- c
		c.IsClose = true
	}()
	for {
		if c.IsClose {
			return
		}
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			m.GetUnregister() <- c
			break
		}
		// todo 业务逻辑 接收到ws消息进一系列处理
		// 这个地方解析定义好的message
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		zap.L().Info("收取成功", zap.Any("message", string(message)))
		// 发送给全局
		m.GetBroadcast() <- jsonMessage
	}
}

// Write:写入消息
func (c *Client) Write(m IManagement) {
	defer func() {
		m.GetUnregister() <- c
		c.IsClose = true
	}()
	for {
		if c.IsClose {
			return
		}
		select {
		case message, ok := <-c.Send:
			// todo 接收到某消息要转发 业务处理
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
			zap.L().Info("写入成功", zap.Any("message", string(message)))
		}
	}
}

// GetID:获取唯一标识
func (c *Client) GetID() int64 {
	return c.ID
}

// GetSocket:获取socket句柄
func (c *Client) GetSocket() *websocket.Conn {
	return c.Socket
}

// GetSendChan:获取发送消息channel
func (c *Client) GetSendChan() chan []byte {
	return c.Send
}

// WsPage is a websocket handler
func WsPage(c *gin.Context) {
	// gin先处理
	// todo 业务逻辑
	// 处理完升级为websocket
	uid, _ := strconv.Atoi(c.DefaultQuery("uid", "1"))

	// change the request to websocket model
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("websocket Upgrader", zap.Error(err))
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &Client{ID: int64(uid), Socket: conn, Send: make(chan []byte)}

	M.Register <- client

	// 开启读写线程
	go client.Read(M)
	go client.Write(M)
}
