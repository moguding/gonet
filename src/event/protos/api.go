package protos

import "misc/packet"

var Code map[string]int16 = map[string]int16 {
	"ping_req":0,	// PING
	"ping_ack":1000,	// 返回请求数值
	"add_req":1,	// 添加一个cooldown请求
	"add_ack":1001,	// 返回这个cooldown编号
	"cancel_req":2,	// 取消一个cooldown请求
	"cancel_ack":1002,	// 返回0
}

var RCode map[int16]string = map[int16]string {
	0:"ping_req",
	1000:"ping_ack",
	1:"add_req",
	1001:"add_ack",
	2:"cancel_req",
	1002:"cancel_ack",
}

var ProtoHandler map[uint16]func(*packet.Packet) ([]byte, error) = map[uint16]func(*packet.Packet)([]byte, error){
	0:P_ping_req,
	1:P_add_req,
	2:P_cancel_req,
}