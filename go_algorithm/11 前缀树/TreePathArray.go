package main

type Trie struct{
	next [26]*Trie
	pass int
	end int
}

func Constructor()(_ Trie){return }

func (root *Trie) Insert(s string){
	root.pass++
	cur := root
	for _,cha := range s{
		path := cha - 'a'
		if cur.next[path] == nil{
			cur.next[path]=&Trie{}
		}
		cur=cur.next[path]
		cur.pass++

	}
		cur.end++
}

func (root *Trie) CountWorksEqualTo(s string) int{
	cur := root
	for _,cha := range s{
		cur := cur.next[cha-'a']
		if cur == nil{
			return 0
		}
	}
	return cur.end
}

func (root *Trie) CountWorkdsStartingWith(s string) int{
	cur := root
	for _,cha := range s{
		cur = cur.next[cha-'a']
		if cur==nil{
			return 0
		}
	}
	return cur.pass
}

func (root *Trie) Erase1(s string){
	if root.CountWorksEqualTo(s)>0{
		root.pass--
		cur := root
		for _,cha := range s{
			cur = cur.next[cha-'a']
			cur.pass--
		}
		cur.end--
	}
}

// 删除改进
func (root *Trie) Erase2(s string){
	if root.CountWorksEqualTo(s)>0{
		root.pass--
		cur := root
		for _,cha := range s{
			path := cha - 'a'
			next := cur.next[path]
			next.pass--
			if next.pass == 0{
				cur.next[path]=nil
				return
			}else{
				cur = next
			}
		}
		cur.end--
	}
}
