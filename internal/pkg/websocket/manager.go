package websocket

import (
	"sync"
)

type Manager struct {
	clients    map[string]*Client // 全部客户端
	register   chan *Client       // 注册客户端
	unregister chan *Client       // 注销客户端
	mutex      sync.RWMutex       // 读写互斥锁
}

func init() {
	go manager.Listener()
}

var manager = &Manager{
	clients:    map[string]*Client{},
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

// GetClient 获取客户端
func (manager *Manager) GetClient(clientId string) *Client {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	return manager.clients[clientId]
}

// GetClients 获取全部客户端
func (manager *Manager) GetClients() map[string]*Client {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	return manager.clients
}

// Listener 监听器
func (manager *Manager) Listener() {
	for {
		select {
		case client := <-manager.register:
			manager.clients[client.GetClientId()] = client
		case client := <-manager.unregister:
			delete(manager.clients, client.GetClientId())
		}
	}
}

// Register 注册客户端
func (manager *Manager) Register(client *Client) {
	manager.register <- client
}

// Unregister 注销客户端
func (manager *Manager) Unregister(client *Client) {
	manager.unregister <- client
}
