//go:build !windows && !plan9

package syslog

import (
	"fmt"
	"github.com/go-packagist/logger"
	"github.com/go-packagist/monolog"
	"github.com/stretchr/testify/assert"
	"net"
	"sync"
	"testing"
	"time"
)

var result = make(chan string, 10)

type server struct {
	conn net.PacketConn
}

func newServer() *server {
	conn, _ := net.ListenPacket("udp", ":30732")

	return &server{
		conn: conn,
	}
}

func (s *server) start() {
	buffer := make([]byte, 1024)

	go func() {
		for {
			n, _, err := s.conn.ReadFrom(buffer)
			if err != nil {
				panic(err)
			}

			result <- string(buffer[:n])
		}
	}()
}

func init() {
	newServer().start()
}

func createMonolog() *monolog.Logger {
	return monolog.NewLogger("test", monolog.WithHandler(
		NewHandler("go-packagist",
			WithLevel(logger.Debug),
			WithNetwork("udp"),
			// WithRaddr("192.168.8.92:30732"),
			WithRaddr(":30732"),
		),
	))
}

func TestHandler_Local(t *testing.T) {
	h := NewHandler("go-packagist", WithLevel(logger.Info))

	assert.False(t, h.Handle(nil))
	assert.True(t, h.Handle(&monolog.Record{
		Level:   logger.Info,
		Message: "test message",
		Channel: "go-packagist",
		Time:    time.Now(),
	}))
}

func TestHandler_Udp(t *testing.T) {
	m := createMonolog()
	defer m.Close()

	wg := sync.WaitGroup{}
	defer wg.Wait()
	go func() {
		for {
			select {
			case r := <-result:
				fmt.Print(r)
				wg.Done()
			}
		}
	}()

	wg.Add(8)
	m.Emergency("test emergency")
	m.Alert("test alert")
	m.Critical("test critical")
	m.Error("test error")
	m.Warning("test warning")
	m.Notice("test notice")
	m.Info("test info")
	m.Debug("test debug")
}

func BenchmarkHandler_Udp(b *testing.B) {
	m := createMonolog()
	defer m.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Info("test message")
		}
	})
}
