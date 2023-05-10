package main

import "strings"
// www.leetcode-cn.com/problems/replace-words/
func ReplaceWords(dictionary []string, sentence string) string {
	root := Constructor()
	for _, s := range dictionary {
		root.Insert(s)
	}

	words := strings.Split(sentence," ")
	for i,word := range words{
		words[i] = root.Replace(word)
	}
	return strings.Join(words," ")
}

type Trie struct{
	next [26]*Trie
	end bool
}

func Constructor()(_ Trie) {return}
func (root *Trie)Insert(s string){
	cur := root 
	for _,cha := range s{
		path := cha - 'a'
		if cur.next[path] == nil{
			cur.next[path] = &Trie{}
		}
		cur = cur.next[path]
	}
	cur.end = true
}

func (root *Trie)Replace(s string)string{
	cur := root 
	for i,cha := range s{
		path := cha - 'a'
		if cur.next[path] == nil{
			return s
		}
		cur = cur.next[path]
		if cur.end{
			return s[:i+1]
		}
	}
	return s
}
