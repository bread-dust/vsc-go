package main

import (
	"google.golang.org/grpc"
)

// 负载均衡源码解析
func main() {
grpc.Dial("",grpc.WithDefaultServiceConfig(`{"localBalancingPolicy": "round:robin"}`))
}

// round-robin 策略实现
//google.golang.org/grpc/roundronbin/roundrobin.go 
// build() 和 pick()
func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
	subConnsLen := uint32(len(p.subConns))
	nextIndex := atomic.AddUint32(&p.next, 1)

	sc := p.subConns[nextIndex%subConnsLen]
	return balancer.PickResult{SubConn: sc}, nil
}

// 调用 ronbin

 // newCCBalancerWrapper() -> switch to 和 go watcher
 // watcher —> 监听updatechannel 中的消息并读取 -> 根据消息类型不同执行不同操作

 // newCCResolverWrapper()
// Build -> consul_resolver实现 -> papulateEndpoint -> updateState() ->  resloverWrapper 实现 -> updateResolverState() -> applyServceConfigAndBalancer() 根据字符串解析 -> switchTo -> handleSwitchTo() -> banlancer.Get() -> balancer.SwitchTo() -> builder.Build -> rrPicker{} 
//