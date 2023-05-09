消息代码放在codec 子目录下

// 利用类型转换确保接口被实现，这样 IDE 和编译期间就可以检查，而不是等到使用的时候。
var _ Codec = (*GobCodec)(nil)