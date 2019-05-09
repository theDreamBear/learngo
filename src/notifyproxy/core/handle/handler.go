package handle

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gomcpack/mcpack"
	"gomcpack/npc"
	"notifyproxy/core/define"
	"notifyproxy/core/mysql"
)

type Handler struct {
}

func (th *Handler) Serve(rw npc.ResponseWriter, re *npc.Request) {
	//ParseHeader(re)
	buf := make([]byte, 40960)
	len1, _ := re.Body.Read(buf)
	//fmt.Println(buf[:len1])
	var obj = &define.NotifyNmqOnlineInfo{}
	//fmt.Println(buf[:len1])
	if err := mcpack.Unmarshal(buf[:len1], obj); err != nil {
		panic(err)
	} else {
		//fmt.Println()
		fmt.Printf("_______________________new request___________________________\n")
		fmt.Printf("_product: %v\n", obj.Product_)
		fmt.Printf("Topic: %v\n", obj.Topic)
		fmt.Printf("Cmd: %v\n", obj.Cmd)
		fmt.Printf("Idc: %v\n", obj.Idc)
		fmt.Printf("Cluster: %v\n", obj.Cluster)
		fmt.Printf("TopicGroupName: %v\n", obj.TopicGroupName)
		fmt.Printf("CommitTime: %v\n", obj.CommitTime)
		fmt.Printf("CommitTimeUs: %v\n", obj.CommitTimeUs)
		fmt.Printf("TransId: %v\n", obj.TransId)
		fmt.Printf("LogId: %v\n", obj.LogId)
		fmt.Printf("ClientIp: %v\n", obj.ClientIp)
		fmt.Printf("Provider: %v\n", obj.Provider)
		fmt.Printf("Cmdno: %v\n", obj.Cmdno)
		fmt.Printf("Uid: %v\n", obj.Uid)
		fmt.Printf("Exral: %v\n", obj.Extra)
		fmt.Printf("Stime: %v\n", obj.Stime)
		fmt.Printf("Utime: %v\n", obj.Utime)
		fmt.Printf("_______________________end___________________________\n")
		go mysql.InsertToDb(obj, rw)
	}
}
