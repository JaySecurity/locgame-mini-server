package network

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	// pprof allows you to profile the server using http.
	_ "net/http/pprof"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/prometheus/client_golang/prometheus"
	"locgame-mini-server/pkg/log"
)

// Verbose turns on debug mode.
var Verbose = false

type Server struct {
	OnDisconnectClient func(client *Client)
	OnConnectClient    func(client *Client)
	MethodValidateFunc func(methodID uint16) bool

	handler ServerHandler

	mutex       sync.Mutex
	connections map[*Client]struct{}
	inShutdown  *AtomicBool

	connectionsGauge prometheus.Gauge
}

// NewWebSocketHandler allows you to create a new web socket server instance.
func NewWebSocketHandler() *Server {
	setRLimit()
	rand.Seed(time.Now().Unix())

	s := &Server{
		inShutdown: new(AtomicBool),
	}

	s.connectionsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "loc_connections",
		Help: "The total number of current connections",
	})

	_ = prometheus.Register(s.connectionsGauge)

	return s
}

func (s *Server) SetServerHandler(handler ServerHandler) {
	s.handler = handler
	s.MethodValidateFunc = handler.Validate
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.inShutdown.IsSet() {
		http.Error(w, "Server in shutdown", http.StatusServiceUnavailable)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	conn, _, _, _ := ws.UpgradeHTTP(r, w)

	newClient := NewClient(
		&Stream{
			Conn:         conn,
			ValidateFunc: s.MethodValidateFunc,
		},
		r.Header.Get("X-Client-Geo-Region")) // TODO Get GeoRegion from LoadBalancer

	if s.OnConnectClient != nil {
		s.OnConnectClient(newClient)
	}
	newClient.OnCloseConnection = func(client *Client) {
		s.removeClient(client)
	}

	s.addClient(newClient)

	// Release http handler for GC
	func() {
		err := newClient.Handle(s.handler)
		if err != nil {
			if wsErr, ok := err.(wsutil.ClosedError); ok {
				switch wsErr.Code {
				case ws.StatusGoingAway:
					return
				case ws.StatusNormalClosure, ws.StatusNoStatusRcvd:
					return
				}
			}

			if err == io.EOF ||
				strings.Contains(err.Error(), "operation timed out") ||
				(newClient.Stream.isClosed && strings.Contains(err.Error(), "use of closed network connection")) ||
				strings.Contains(err.Error(), "connection reset by peer") {
				return
			}

			log.Error(err)
			return
		}
	}()
}

func (s *Server) GetClients() []*Client {
	s.mutex.Lock()
	clients := make([]*Client, len(s.connections))
	idx := 0
	for client := range s.connections {
		clients[idx] = client
		idx++
	}
	s.mutex.Unlock()
	return clients
}

func (s *Server) addClient(newClient *Client) {
	s.mutex.Lock()
	if s.connections == nil {
		s.connections = make(map[*Client]struct{})
	}
	s.connections[newClient] = struct{}{}
	s.connectionsGauge.Set(float64(len(s.connections)))
	s.mutex.Unlock()
}

func (s *Server) removeClient(client *Client) {
	if s.OnDisconnectClient != nil {
		s.OnDisconnectClient(client)
	}
	s.mutex.Lock()
	delete(s.connections, client)
	s.connectionsGauge.Set(float64(len(s.connections)))
	s.mutex.Unlock()
}

func (s *Server) AllowConnect(value bool) {
	if !value {
		s.inShutdown.Set()
	} else {
		s.inShutdown.UnSet()
	}
}

// Shutdown allows you to stop the server.
func (s *Server) Shutdown(ctx context.Context) error {
	log.Debug("Stopping the server...")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer func() {
		ticker.Stop()
	}()
	for {
		s.mutex.Lock()
		connections := s.connections
		s.mutex.Unlock()
		for client := range connections {
			client.Shutdown()
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			s.mutex.Lock()
			count := len(s.connections)
			s.mutex.Unlock()

			if count > 0 {
				log.Debugf("Waiting to close %v connections.", len(s.connections))
			}
		}

		s.mutex.Lock()
		count := len(s.connections)
		s.mutex.Unlock()

		if count == 0 {
			return nil
		}
	}
}

func setRLimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatal("Error getting the current limit Rlimit", err)
	}

	rLimit.Cur = 1024000
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Error("Rlimit set error:", err)
		return
	}

	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)

	if err == nil {
		log.Debug("The file descriptor limit is set:", rLimit.Cur)
	} else {
		log.Error("Rlimit set error:", err)
	}
}
