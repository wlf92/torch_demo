package main

import (
	"fmt"
	"log"
	"net/url"
	"sync/atomic"
	"time"
	"torch_demo/assets/pbcli"

	"github.com/gorilla/websocket"
	"github.com/wlf92/torch/packet"
	"google.golang.org/protobuf/proto"
)

func LoginReq() []byte {
	req := &pbcli.LoginReq{ChannelId: 1}
	bts, _ := proto.Marshal(req)
	return packet.Pack(&packet.Message{Route: uint32(pbcli.Msg_Id_LoginReq), Buffer: bts})
}

var finishCount int32 = 0

type Client struct {
	ccc *websocket.Conn

	chsRecv chan []byte
	chsSend chan []byte
}

func (wsc *Client) dial() {
	var err error

	u := url.URL{Scheme: "ws", Host: "localhost:8444", Path: "/"}
	wsc.ccc, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("dial success")

	go wsc.recvLoop()
	go wsc.sendRoop()

	wsc.chsSend <- LoginReq()
}

func (wsc *Client) unpack(datas []byte) error {
	atomic.AddInt32(&finishCount, 1)

	// reader := bytes.NewReader(datas)

	// var length int16
	// var mid uint32
	// bts := make([]byte, length)

	// binary.Read(reader, binary.BigEndian, &length)
	// binary.Read(reader, binary.BigEndian, &mid)
	// binary.Read(reader, binary.BigEndian, &bts)
	// fmt.Printf("%d %+v", mid, bts)

	return nil
}

func (wsc *Client) recvLoop() {
	for {
		_, message, err := wsc.ccc.ReadMessage()
		if err != nil {
			if err.Error() != websocket.ErrReadLimit.Error() {
				log.Println("read:", err)
			}
			break
		}
		wsc.unpack(message)
		wsc.chsSend <- LoginReq()
	}
}

func (wsc *Client) sendRoop() {
	for {
		msg := <-wsc.chsSend
		err := wsc.ccc.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			log.Println("write:", err)
			continue
		}
	}
}

func main() {
	for i := 0; i < 1; i++ {
		ct := Client{
			chsSend: make(chan []byte, 10),
			chsRecv: make(chan []byte, 10),
		}
		ct.dial()
	}

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(atomic.LoadInt32(&finishCount))
			atomic.StoreInt32(&finishCount, 0)
		}
	}()

	time.Sleep(time.Minute)
}
