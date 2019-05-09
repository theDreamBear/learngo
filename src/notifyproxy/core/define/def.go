package define

type PushNmqOnlineInfo struct {
	Product        string `mcpack:"_product"`
	Topic          string `mcpack:"_topic"`
	Cmd            string `mcpack:"_cmd"`
	Idc            string `mcpack:"_idc"`
	Cluster        string `mcpack:"_cluster"`
	TopicGroupName string `mcpack:"_topic_group_name"`
	CommitTime     uint32 `mcpack:"_commit_time"`
	CommitTimeUs   uint32 `mcpack:"_commit_time_us"`
	TransId        uint64 `mcpack:"_transid"`
	LogId          uint64 `mcpack:"_log_id"`
	ClientIp       string `mcpack:"_client_ip"`
	Provider       string `mcpack:"_provider"`
	Uid            uint64 `mcpack:"uid"`
	Data           string `mcpack:"data"`
}

type NotifyNmqOnlineInfo struct {
	Product_       string `mcpack:"_product"`
	Topic          string `mcpack:"_topic"`
	Cmd            string `mcpack:"_cmd"`
	Idc            string `mcpack:"_idc"`
	Cluster        string `mcpack:"_cluster"`
	TopicGroupName string `mcpack:"_topic_group_name"`
	CommitTime     uint32 `mcpack:"_commit_time"`
	CommitTimeUs   uint32 `mcpack:"_commit_time_us"`
	TransId        uint64 `mcpack:"_transid"`
	LogId          uint64 `mcpack:"_log_id"`
	ClientIp       string `mcpack:"_client_ip"`
	Provider       string `mcpack:"_provider"`
	Cmdno          uint32 `mcpack:"cmdno"`
	Uid            uint64 `mcpack:"uid"`
	Product        string `mcpack:"product"`
	Extra          string `mcpack:"extra"`
	Stime          int32  `mcpack:"stime"`
	Utime          int32  `mcpack:"utime"`
}

type NmqResponse struct {
	ErrorNo  int32  `mcpack:"_error_no"`
	ErrorMsg string `mcpack:"_error_msg"`
	TransId  uint64 `mcpack:"_transid"`
}
