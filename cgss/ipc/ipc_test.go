package ipc

import (
	"encoding/json"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, request string) *Response {
	resp := Response{"Echo", "Echo" + request}
	return &resp
}
func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("ECHO", "From Client1")
	resp2, _ := client2.Call("ECHO", "From Client2")

	var result1 []byte
	var result2 []byte
	result1, _ = json.Marshal(resp1)
	result2, _ = json.Marshal(resp2)
	if string(result1) != "ECHO: From Clent1" || string(result2) != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", string(result1), "resp2:", string(result2))
	}
	client1.Close()
	client2.Close()
}
