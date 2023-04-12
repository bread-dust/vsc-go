package main

// option 模式

// ServiceConfig 定义一个服务结构体
type ServiceConfig struct{
	A string
	B string 
	C int
}

// NewServiceConifg 创建一个ServiceConfig
func NewServiceConifg(a,b string,c int) *ServiceConfig{
	return &ServiceConfig{
		A: "",
		B: "",
		C: 0,
	}
}

// A和B 必须传，C可选
func NewServiceConifg2(a,b string,c ...int) *ServiceConfig{
	valueC := defaultValueC
	if len(c)>0{value C =c[0]}
	return &ServiceConfig{
		A: "",
		B: "",
		C: 0,
	}
}

// option模式
type FuncServiceConfigOption func(*ServiceConfig)
func NewServiceConifg(a,b string,options ...FuncServiceConfigOption) *ServiceConfig{
	sc := &ServiceConfig{
		A: "",
		B: "",
		C: defaultValueC,
	}

	// 针对传进来的Fsco参数做处理
	for _,opt := range options {
		opt(sc)
	}
	return sc 
}

// 定制专用的配置方法
func With(c int) FuncServiceConfigOption {  
	return func (sc *ServiceConfig)  { //修改FuncServiceConfigOption中的字段
		sc.C =c
	}
}


//进阶
type config struct{
	name string
	age int
}

// 接口隐藏具体实现，接口对外可见
// name = "deng"
type ConfigOption interface{
	apply(*config) 
}

type funcOption struct{
	f func(*config) // 处理器字段,处理config结构体
}

func (f funcOption) apply(cfg *config){ //实现方法
	f.f(cfg) //实现初始化foucOption
}

for WithConfigName(name string) *ConfigOption{
	return NewfuncOption(func (cfg *config){
		cfg.name = name // 得到一个处理器,
	})
}

func NewfuncOption(f func(*config)) funcOption  {
	return funcOption{f:f} //得到一个包含处理器的结构体，去实现apply方法 
}
/*
	x funcOption{
		f :func(cfg *config){cfg.name="deng"}
	}
*/


func NewConfig(age int,opts ...ConfigOption)*config{
	cfg:= &config{age:age}
	for _,opt := range opts{
		opt.apply(cfg) //x.apply(cfg),   x.func(cfg)
	}
	return cfg
}


func main() {
	NewServiceConifg(1,2,With(3))
	test4(With(3))
	NewConfig(1,WithConfigName("deng"))
}




type test1 struct{
	a int
	b int
}


type Test2 interface{
	apply(*test1)
}

type test3 struct{
	f func(*test)
}

func (t3 test3)apply(cfg *test1){
	t3.f(cfg)
}


// 处理器的具体实现，返回一个test2 ，即对test1的处理器
func With(b int) Test2 {
	return func(t *test1) {
		t.a = b
	}
}

func Newtest3(cpu func(*test))test3{
	return test3{
		f: cpu,
	}
}


func test4(options ...test2)*test1{
	aa = &test1{
		a: 0,
	}
	// 对aa调用test2 处理器 
	for _,option := range options{
		option(aa)
	}
	return aa
}