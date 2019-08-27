package server

import (
	"gtihub.com/believening/cache-wrings/cache"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTP server error 已经被定义好的
// var (
// 	// ErrInvalidKeyLenth 键长度非法
// 	ErrInvalidKeyLenth = errors.New("server: invalid length of the key")
// 	// ErrUnsupportMethod 不支持的 HTTP method
// 	ErrUnsupportMethod = errors.New("server: unsupport method")
// )

// Server 后端
type Server struct {
	cache.Cache
}

// New Server
func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

func (s *Server) statHandler() http.Handler {
	return &statHandler{s}
}

type cacheHandler struct {
	*Server
}

// 定义规则：method /cache/<key>\r\n<value>
func (ch *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// strings.Split() 会在每一个分隔符处都切割，若分隔符在字符串最开始，则会增加一个空字符
	// fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	// output: ["" "man " "plan " "canal panama"]
	k := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(k) <= 0 {
		w.WriteHeader(http.StatusBadRequest) // 400 error
	}
	switch r.Method {
	case http.MethodPut:
		// 读取 v
		v, _ := ioutil.ReadAll(r.Body)
		if err := ch.Set(k, v); err != nil {
			log.Panicln(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	case http.MethodGet:
		v, err := ch.Get(k)
		if err != nil {
			log.Panicln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(v) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(v)
		return
	case http.MethodDelete:
		if err := ch.Del(k); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

type statHandler struct {
	*Server
}

// 规则： GET /stat/
func (sh *statHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 序列化数据
		b, err := json.Marshal(sh.GetStat())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// Run 运行后端
func (s *Server) Run() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/stat", s.statHandler())
	http.ListenAndServe("127.0.0.1:55555", nil)
}
