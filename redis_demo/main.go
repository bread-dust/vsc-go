package main

import (
	"fmt"

	"github.com/go-redis/redis"
)


func RdbConn() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.56.136:6379",
		Password: "dengliwei", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, //连接池大小
	})
	return rdb
}

func Key(rdb *redis.Client){
	rdb.Set("student", "value", 0)
	fmt.Println(rdb.Exists("student").Result())
	// rdb.Expire("student", 10) // 设置存活时间
	setnx:=rdb.SetNX("student", "value", 0) // 设置不存在的key
	fmt.Println(setnx.Result()) 
}

func Hash(rdb *redis.Client){

	rdb.HSet("person", "name", "dengliwei") // 设置hash
	rdb.HSet("person", "age", 115)
	hash:=rdb.HGet("person", "name") // 获取hash
	fmt.Println(hash.Result()) 

	rdb.HMSet("person2", map[string]interface{}{"name": "zhangsan", "age": 18}) // 设置多个hash

	name,_:=rdb.HGet("person","name").Result() // 获取hash
	nage,_:=rdb.HMGet("person2", "name", "age").Result() // 获取多个hash

	fmt.Println(name)
	fmt.Println(nage)

	all,_:=rdb.HVals("person2").Result() // 获取所有的值
	fmt.Println(all)

	gall,_:=rdb.HGetAll("person2").Result() // 获取所有的值
	fmt.Println(gall)

	// rdb.HDel("person2", "name") // 删除hash
}

func List(rdb *redis.Client){
	rdb.LPush("list1","0","1") // 头插
	rdb.LPush("list1","deng","li","wei")
	lrange:= rdb.LRange("list1",0,2)
	fmt.Println(lrange.Result())

	rdb.RPush("list1","e0","e1") // 尾插
	llen := rdb.LLen("list1")
	fmt.Println(llen.Result())

	rdb.LRem("list1",10,"f0") // 删除10个f0
}

func Set(rdb *redis.Client){
	rdb.SAdd("key1","0","1","2")
	rdb.SAdd("key2","2","3","4")
	scard:=rdb.SCard("key1") // 集合的元素个数
	fmt.Println(scard.Result())
	smembers:=rdb.SMembers("key1") // 返回集合中的所有成员
	fmt.Println(smembers.Result())
}

func Sorted(rdb *redis.Client){
	rdb.ZAdd("zset1",redis.Z{
		Score: 1,
		Member: "a",
	},redis.Z{
		Score: 2,
		Member: "b",
	},redis.Z{
		Score: 3,
		Member: "c",
	})

	zmem:=rdb.ZRange("zset1",0,2) // 返回有序集中，指定区间内的成员
	sli:=make([]string,2)
	zmem.ScanSlice(&sli)
	fmt.Println(sli)

	zws:=rdb.ZRangeWithScores("zset1",0,2)
	zws.Result()

	zscore := rdb.ZScore("zset1","a") // 返回有序集中，成员的分数值
	fmt.Println(zscore.Result()) 

}
func main() {
	rdb := RdbConn()
	Key(rdb)
	Hash(rdb)
	List(rdb)
	Set(rdb)
	Sorted(rdb)


	defer rdb.Close()
}