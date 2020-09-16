package tcpserver

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/believening/cache-wrings/cache"
	"github.com/believening/cache-wrings/server"
)

var (
	port = ":12346"
	typ  = "tcp"
)

type tcpServer struct {
	cache.Cache
}

func New(c cache.Cache) server.Server {
	return &tcpServer{c}
}

func init() {
	server.Register(typ, New)
}

func (s *tcpServer) Run() {
	log.Println("serve tcp at: ", port)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.process(c)
	}
}

func (s *tcpServer) process(c net.Conn) {
	defer c.Close()
	r := bufio.NewReader(c)
	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Printf("close connection due to : %v\n", err)
				return
			}
		}
		switch op {
		case 'S':
			err = s.set(c, r)
		case 'G':
			err = s.get(c, r)
		case 'D':
			err = s.del(c, r)
		default:
			log.Fatalln("close connection due to invalid operation: ", op)
			return
		}
		if err != nil {
			log.Printf("close connection due to : %v\n", err)
			return
		}
	}
}

func (s *tcpServer) get(conn net.Conn, r *bufio.Reader) error {
	k, err := s.readKey(r)
	if err != nil {
		return err
	}
	v, err := s.Get(k)
	return sendResponse(v, err, conn)
}

func (s *tcpServer) del(conn net.Conn, r *bufio.Reader) error {
	k, err := s.readKey(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Del(k), conn)
}

func (s *tcpServer) set(conn net.Conn, r *bufio.Reader) error {
	k, v, err := s.readKeyAndValue(r)
	if err != nil {
		return err
	}
	return sendResponse(nil, s.Set(k, v), conn)
}

func (s *tcpServer) readKey(r *bufio.Reader) (string, error) {
	l, err := readLen(r)
	if err != nil {
		return "", err
	}
	k := make([]byte, l)
	_, err = io.ReadAtLeast(r, k, l)
	if err != nil {
		return "", err
	}
	return string(k), nil
}

func (s *tcpServer) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	kl, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	vl, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	k := make([]byte, kl)
	_, err = io.ReadAtLeast(r, k, kl)
	if err != nil {
		return "", nil, err
	}
	v := make([]byte, vl)
	_, err = io.ReadAtLeast(r, v, vl)
	if err != nil {
		return "", nil, err
	}
	return string(k), v, nil
}

func readLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return l, nil
}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errMsg := err.Error()
		_, e := conn.Write([]byte(fmt.Sprintf("-%d %s", len(errMsg), errMsg)))
		return e
	}
	_, e := conn.Write(append([]byte(fmt.Sprintf("%d ", len(value))), value...))
	return e
}
