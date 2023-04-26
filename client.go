package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

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

	wsc.send(4, []byte{})
}

func (wsc *Client) send(mid uint32, bts []byte) {
	wsc.chsSend <- wsc.pack(mid, bts)
}

func (wsc *Client) pack(mid uint32, bts []byte) []byte {
	var buf bytes.Buffer
	length := len(bts) + 4 + 2
	buf.Grow(length)

	binary.Write(&buf, binary.BigEndian, int16(length))
	binary.Write(&buf, binary.BigEndian, mid)
	binary.Write(&buf, binary.BigEndian, bts)
	return buf.Bytes()
}

func (wsc *Client) unpack(datas []byte) error {
	fmt.Println(datas)

	reader := bytes.NewReader(datas)

	var length int16
	var mid uint32
	bts := make([]byte, length)

	binary.Read(reader, binary.BigEndian, &length)
	binary.Read(reader, binary.BigEndian, &mid)
	binary.Read(reader, binary.BigEndian, &bts)
	fmt.Printf("%d %+v", mid, bts)

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
	ct := Client{
		chsSend: make(chan []byte, 10),
		chsRecv: make(chan []byte, 10),
	}
	ct.dial()

	time.Sleep(time.Minute)
}
