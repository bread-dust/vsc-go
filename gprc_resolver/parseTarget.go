package main

// Dial调用创建一个clientConn对象，将target赋值给对象里的target
cc := &ClientConn{
	target: target,
	...
}

// ParseTarget 解析这个target
func ParseTarget(target string) (ret resolver.vartarget){
	var ok bool
	ret.Scheme,ret.Endpoint,ok=split2(target,"://")
	if !ok{
		return resolver.Target{Endpoint:target}
	}
	ret.Authority,ret.Endpoint,ok=split2(ret.Endpoint,"/")
	if !ok{
		return resolver.Target{Endpoint:target}
	}
	return ret
}

// 执行完tartget执行这一句
resolverBuilder :=cc.getResolver(cc.ParseTarget.Scheme)

// getResolver 获得一个resolver的builder
func (cc *ClientConn) getResolver (scheme string)resolver.Builder{
	for _,rb:=range.cc.dopts.resolvers{
		// Dial() 中已指定解析器(withResolver(&dlwsolverbuilder{}))
		if scheme == rb.Scheme(){ //Scheme() 已自定义返回dlw
			return rb
		}
	}
	// Dial()未指定解析器参数，调用Get(scheme) 查找
	return resolver.Get(scheme)
}

// Get(scheme) 查找是否存在scheme对应的resolver的builder
func Get(scheme string)Builder{
	// m这个map 返回一个builer，在Register()函数插入builer
	if b,ok:=m[scheme];ok{
		return b
	}
	return nil
}

// Register 插入scheme 对应的builer
// passthrough 是gRPC默认的解析器，当target=ip:port,什么都不做，
func Register (b Builder){
	m[b.Scheme()] = b
}

//passthrough 解析器调用 Register
func init(){
	resolver.Register(&passthroughBuilder{})
}

// 如果getResolver获得的builder是nil scheme设置成默认的scheme，即：ip:port
resolverBuilder := cc.getResolver(cc.ParseTarget.Scheme) 
if resolverBuilder == nil{
	channelz.Infof(cc.chanelzID,"scheme nit regissterd,fallback to default sckeme",cc.ParseTarget.Scheme)
	cc.ParseTarget = resolver.Target{
		// defaultScheme = "passthrough"
		Scheme:resolver.GetfaultScheme(),
		Endpoint:target,
	}
	resolverBuilder == cc.getResolver(cc.ParseTarget.Scheme)
	if resolverBuilder == nil{
		retrun nil,fmt.Errorf("could not get resolver for default scheme")
	}
}

// 使用builder构建一个resolver

rWrapper,err:= newCCResolverWrapper(cc,resolverBuilder)
if err!=nil{
	return nil,fmt.Errorf("failed to buid resolber:%v",err)
}
cc.mu.Lock()
cc.resolverWrapper=rWrapper
cc.Unlock()

// newCCRResolverWrapper 创建一个reolver
func newCCRResolverWrapper(cc *ClientConn,rb resolver.Builder)(*newCCRResolverWrapper,error){
	...
	ccr := &ccResolverWrapper{
		cc :cc ,
		...
	}
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver ,err = rb.Build(cc.parsedtarget,ccr,rbo)
}
