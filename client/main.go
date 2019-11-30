package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var (
	address = "127.0.0.1:8088"
)

func main() {
	conn, err := net.Dial("tcp", address)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		log.Fatalf("dial server %s failed: %v", address, err)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.TrimSpace(scanner.Text())
		sendrequest(data, conn)
	}
}

func sendrequest(data string, conn net.Conn) {
	opts := strings.SplitN(data, " ", 3)
	opt := opts[0]
	var req string
	switch opt {
	case "d":
		key := opts[1]
		req = fmt.Sprintf("%s%d %s\n", opt, len(key), key)
	case "g":
		key := opts[1]
		req = fmt.Sprintf("%s%d %s\n", opt, len(key), key)
	case "p":
		key, value := opts[1], opts[2]
		req = fmt.Sprintf("%s%d %d %s%s\n", opt, len(key), len(value), key, value)
	case "s":
		req = fmt.Sprintf("%s\n", opt)
	default:
		log.Println("error opt: ", data)
		return
	}
	_, err := conn.Write([]byte(req))
	if err != nil {
		return
	}
	log.Print("write:", req)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln("read from server faild: ", err)
		return
	}
	log.Print(string(buf[:n]))
}
