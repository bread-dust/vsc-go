package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	M sync.Map // 键值对
}

// 判断用户是否存存在map中
func (b *Ban) IsIn(user string)bool{
	fmt.Printf("%s 进来了\n",user)
	// Load 方法返回两个值，一个是如果能拿到的 key 的 value
    // 还有一个是否能够拿到这个值的 bool 结果
	// v,ok := b.M.Load(user)
	v,ok :=b.M.LoadOrStore(user,time.Now()) // Load 查询对应的key 的值
	if !ok{
		// 如果没有，可以真长访问
		fmt.Printf("名单里没有%s,可以访问",user)
		// 将用户名存到Ban list 中
		// b.M.Store(user,time.Now())
		return false
	}
	// 如果有，则判断用户的时间距离现在是否已经超过了 180 秒，也就是3分钟
	if time.Now().Second() - v.(time.Time).Second() >180{
	// 超过则可以继续访问
		fmt.Printf("时间为:%d-%d\n",v.(time.Time).Second(),time.Now().Second())
	// 同时重新存入时间
		b.M.Store(user,time.Now())
	}
	// 否则不能访问
	fmt.Printf("名单里有%s,拒绝访问\n",user)
	return true
}

func IpVisit(){
	var success int64= 0
	ban := new(Ban)
	wg := sync.WaitGroup{} // 保证程序运行完成
	for i:=0;i<2;i++{ // 每个user 连续访问两次
		for j:=0;j<10;j++{ //人数预定10个
			wg.Add(1)
			go func(c int){
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d",c)
				if ban.IsIn(ip){
					// 原子操作增加计数器，用来统计人数
					atomic.AddInt64(&success,1)
				}
			}(j)
		}
	}
	wg.Wait()
	fmt.Println("success:",success)
}

type Ban2 struct {
	visitIPs map[string]time.Time
	lock      sync.Mutex
}

func NewBan(ctx context.Context) *Ban2 {
	o := &Ban2{visitIPs: make(map[string]time.Time)}
	// 每隔1分钟检查ban列表里的ip值有没有超过1分钟
	// 超过1分钟从进制列表中删除，允许访问
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Since(v) >= time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}
func (o *Ban2) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	// ip在禁止列表里
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	// ip不在禁止列表里，允许访问，将ip存入列表
	o.visitIPs[ip] = time.Now()
	return false
}

func main2() {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)

	wait := &sync.WaitGroup{}

	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}

	}
	wait.Wait()

	fmt.Println("success:", success)
}