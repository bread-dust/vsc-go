package main

import (
	"fmt"
	"sync"
	"time"
)

/*
type sp interface {
    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}
*/

type Map struct {
	c   map[string]*entry // 键值对
	rmx sync.RWMutex      // 读写锁
}

type entry struct {
	ch      chan struct{} // Get监听通道
	value   interface{}   // 结果是否有效
	isExist bool          // 是否存在
}

// Put 
func (m *Map) Put(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	// entry 已存在
	if e,ok:=m.c[key];ok{
		e.value = val //更新value
		e.isExist = true //更新isExist
		close(e.ch)    //关闭通道
	}else{
		// entry 不存在
		e := &entry{
			ch:      make(chan struct{}), //向get放发送信号
			value:   val,
			isExist: true,
		}
		m.c[key]=e // 插入entry到map中
		close(e.ch) // 关闭通道
	}
}

// Rd 
func (m *Map) Rd(key string, timeout time.Duration) interface{}{
	m.rmx.Lock()
	// key 对应的 entry 存在
	if e,ok:=m.c[key];ok&&e.isExist{
		m.rmx.Unlock()
		return e.value // 更新value
	}else if !ok{ // key 对应的 entry 不存在
		e = &entry{ // 创建空的entry返回
			ch:      make(chan struct{}), 
			value:   nil,
			isExist: false,
		}
		m.c[key]=e
		m.rmx.Unlock()
		fmt.Println("协程阻塞->",key)
		select {
		case <-e.ch:
			return e.value // 读取为空
		case <-time.After(timeout):
			fmt.Println("协程超时->",key)
			return nil
		}
	}else{ //ok is true and isExist is false
		// key 对应的 entry 存在但为空
		m.rmx.Unlock()
		fmt.Println("协程阻塞->",key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时->",key)
			return nil
		}
	}
}