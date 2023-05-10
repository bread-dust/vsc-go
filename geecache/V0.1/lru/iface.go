package lru

type Value interface {
	Len() int // 返回值所占用的字节数
}
