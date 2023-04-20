/*
@author:Deng.l.w
@version:1.20
@date:2023-02-18 12:07
@file:Treemap.go
*/

package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

// 有序表，底层是红黑树
func TreeMap(){
	treepMap:=treemap.NewWithIntComparator() // empty (keys are of int type)
	treepMap.Put(3,"a")
	treepMap.Put(1,"b")

	fmt.Println(treepMap.Get(3))
	fmt.Println(treepMap.Get(1))

	treepMap.Put(1,"c")
	
	fmt.Println(treepMap.Get(1))

	


}


