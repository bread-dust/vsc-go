package main


func (root *Trie)FindMaximumXOR(num int) int {
	cur := root
	ans :=0
	for i:=30;i>=0;i--{
		path := 0
	
	if (num&(1<<i))!=0{
		path =1
	}
	want := path ^ 1

	if cur.next[want] == nil{
		want ^= 1
	}

	ans |= (want ^ path) << i

	cur = cur.next[want]
	}
	return ans
}


type Trie struct {
	next [2]*Trie
}


func Constructor() (_ Trie){return}

func (root *Trie)Insert(num int){
	cur := root
	for i:=30;i>=0;i--{
		path:=0
		if (num&(1<<i))!=0{
			path = 1
		}
		if cur.next[path] ==nil{
			cur.next[path] = &Trie{}
		}
		cur=cur.next[path]
	}
}

