/**
 * @Author: sheng
 * @Email: 396039228@qq.com
 * @Date: 2022/5/28 3:48 下午
 * @Desc: 连接管理器
 */

package ws

import (
	"github.com/bitini111/mvta/log"
	"github.com/bitini111/mvta/network"
	"sync"

	"github.com/gorilla/websocket"
)

type connMgr struct {
	mu        sync.Mutex                      // 连接读写锁
	id        int64                           // 连接ID
	pool      sync.Pool                       // 连接池
	conns     map[*websocket.Conn]*serverConn // 连接集合
	TaskQueue []chan *chRead                  //Worker负责取任务的消息队列
	server    *server                         // 服务器
}

func newConnMgr(server *server) *connMgr {
	it := &connMgr{
		server:    server,
		conns:     make(map[*websocket.Conn]*serverConn),
		pool:      sync.Pool{New: func() interface{} { return &serverConn{} }},
		TaskQueue: make([]chan *chRead, server.opts.maxWorkSize),
	}
	it.startWorkPool()
	return it
}

// 关闭连接
func (cm *connMgr) close() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for _, conn := range cm.conns {
		_ = conn.Close(false)
	}
}

// 分配连接
func (cm *connMgr) allocate(c *websocket.Conn) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if len(cm.conns) >= cm.server.opts.maxConnNum {
		return network.ErrTooManyConnection
	}

	cm.id++
	conn := cm.pool.Get().(*serverConn)
	conn.init(c, cm)
	cm.conns[c] = conn

	return nil
}

// 回收连接
func (cm *connMgr) recycle(conn *serverConn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	delete(cm.conns, conn.conn)
	cm.pool.Put(conn)
}

// 打开线程池
func (cm *connMgr) startWorkPool() {

	for i := 0; i < cm.server.opts.maxWorkSize; i++ {
		// A worker is started
		// Allocate space for the corresponding task queue for the current worker
		// (给当前worker对应的任务队列开辟空间)
		cm.TaskQueue[i] = make(chan *chRead, cm.server.opts.maxTaskLen)
		j := i
		// Start the current worker, blocking and waiting for messages to be passed in the corresponding task queue
		// (启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来)
		go cm.startOneWorker(j, cm.TaskQueue[j])
	}
}

func (cm *connMgr) startOneWorker(workerID int, taskQueue chan *chRead) {
	log.Infof("Worker ID = %d is started.", workerID)
	for {
		select {
		case request := <-taskQueue:

			if cm.server.receiveHandler != nil {
				log.Infof("Add ConnID=%d,workerID=%d", request.conn.ID(), workerID)
				cm.server.receiveHandler(request.conn, request.msg, request.msgType)
			}

		}
	}
}

// 将消息交给TaskQueue,由worker进行处理
func (cm *connMgr) SendMsgToTaskQueue(request *chRead) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//得到需要处理此条连接的workerID

	workerID := int(request.conn.ID()) % cm.server.opts.maxWorkSize
	//fmt.Println("Add ConnID=", request.GetConnection().GetConnID()," request msgID=", request.GetMsgID(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	cm.TaskQueue[workerID] <- request
}
