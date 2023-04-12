/*
@author:Deng.l.w
@version:1.20
@date:2023-03-05 10:06
@file:snowflake.go
*/

package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	// 时间因子从那一年开始
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return err
}

func GenID() int64 {
	return node.Generate().Int64()

}
