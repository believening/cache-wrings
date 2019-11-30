package server

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/believening/cache-wrings/cache"
)

var (
	address = "127.0.0.1:8088"

	invalidRequest = errors.New("invalid request")
	internalError  = errors.New("internal error")
)

// Server 后端
type Server struct {
	cache.Cache
}

// New Server
func New(c cache.Cache) *Server {
	return &Server{c}
}

// Run 运行后端
func (s *Server) Run() {
	l, err := net.Listen("tcp", address)
	defer l.Close()
	if err != nil {
		log.Fatalf("server listen at %s failed: %v\n", address, err)
		return
	}
	for {
		conn, err := l.Accept()
		log.Printf("conn: %v", conn)
		if err != nil {
			log.Fatalf("listener accept failed: %v\n", err)
		}
		go s.serve(conn)
	}
}

// tcp conn 复用持续读取时，一次读取多少内容需要明确的
// 采用
func (s *Server) serve(conn net.Conn) {
	log.Printf("serveing %s", conn.RemoteAddr())
	in := bufio.NewReader(conn)
	defer conn.Close()
	for {
		data, _, err := in.ReadLine()
		log.Println(string(data), len(data))
		if err != nil {
			log.Printf("read data form client failed: %v\n", err)
			// writeResponse(nil, internalError, conn)
			return
		}
		if len(data) == 0 {
			log.Printf("read nil form client\n")
			// writeResponse(nil, invalidRequest, conn)
			continue
		}
		opt := data[0]
		switch opt {
		case 'G', 'g': // Glen key
			err = s.handleGet(conn, data[1:len(data)])
		case 'P', 'p': // Plen len keyvalue
			err = s.handlePut(conn, data[1:len(data)])
		case 'D', 'd': // Dlen key
			err = s.handleDel(conn, data[1:len(data)])
		case 'S', 's': // S
			err = s.handleStat(conn)
		default:
			log.Printf("invalid oprate %s\n", string(opt))
		}
		if err != nil {
			log.Printf("close conn with internal error: %v\n", err)
			return
		}
	}
}

// handleGet
//  data: len key
//  resp: len value | len -error
func (s *Server) handleGet(conn net.Conn, data []byte) error {
	sp := bytes.Index(data, []byte{' '})
	if sp < 0 {
		// writeResponse(nil, invalidRequest, conn)
		return fmt.Errorf("invalid request %s", string(data))
	}
	len, err := strconv.Atoi(string(data[:sp]))
	if err != nil {
		// writeResponse(nil, invalidRequest, conn)
		return fmt.Errorf("parse key len %s failed: %v", string(data), err)
	}
	sp++
	// key := data[sp:sp+len]
	value, err := s.Cache.Get(string(data[sp : sp+len]))
	return writeResponse(value, err, conn)
}

func (s *Server) handlePut(conn net.Conn, data []byte) error {
	llkv := bytes.SplitN(data, []byte{' '}, 3)
	if llkv == nil || len(llkv) != 3 {
		// writeResponse(nil, invalidRequest, conn)
		return fmt.Errorf("invalid format of len len kv: %s\n", string(data))
	}
	klen, err := strconv.Atoi(string(llkv[0]))
	if err != nil {
		// writeResponse(nil, internalError, conn)
		return fmt.Errorf("parse key len %s faild: %v\n", string(data), err)
	}
	vlen, err := strconv.Atoi(string(llkv[1]))
	if err != nil {
		// writeResponse(nil, internalError, conn)
		return fmt.Errorf("parse key len %s faild: %v\n", string(data), err)
	}
	key := string(llkv[2][:klen])
	value := llkv[2][klen : klen+vlen]
	return writeResponse(nil, s.Cache.Set(key, value), conn)
}

func (s *Server) handleDel(conn net.Conn, data []byte) error {
	sp := bytes.Index(data, []byte{' '})
	if sp < 0 {
		// writeResponse(nil, invalidRequest, conn)
		return fmt.Errorf("invalid request %s", string(data))
	}
	len, err := strconv.Atoi(string(data[:sp]))
	if err != nil {
		// writeResponse(nil, invalidRequest, conn)
		return fmt.Errorf("parse key len %s failed: %v", string(data), err)
	}
	sp++
	// key := data[sp:sp+len]
	err = s.Cache.Del(string(data[sp : sp+len]))
	return writeResponse(nil, err, conn)
}

func (s *Server) handleStat(conn net.Conn) error {
	stat := s.Cache.GetStat()
	resp := fmt.Sprintf("hold %d pair key", stat.Cnt)
	return writeResponse([]byte(resp), nil, conn)
}

func writeResponse(value []byte, e error, conn net.Conn) error {
	var resp string
	if e != nil {
		errmsg := e.Error()
		resp = fmt.Sprintf("%d %s", len(errmsg), errmsg)
	} else {
		resp = fmt.Sprintf("%d %s", len(value), value)
	}
	_, err := conn.Write([]byte(resp))
	return err
}
