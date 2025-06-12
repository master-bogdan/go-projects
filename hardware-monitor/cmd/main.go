package main

import (
	"context"
	"fmt"
	"hardware-monitor/internal/hardware"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type server struct {
	subscribeMessageBuffer int
	subscribersMutex       sync.Mutex
	subscribers            map[*subscriber]struct{}
	mux                    http.ServeMux
}

type subscriber struct {
	msgs chan []byte
}

func NewServer() *server {
	s := &server{
		subscribeMessageBuffer: 10,
		subscribers:            make(map[*subscriber]struct{}),
	}

	s.mux.Handle("/", http.FileServer(http.Dir("./htmx")))
	s.mux.HandleFunc("/ws", s.subscriberHandler)
	return s
}

func (s *server) subscriberHandler(writer http.ResponseWriter, req *http.Request) {
	err := s.subscribe(req.Context(), writer, req)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *server) addSubscriber(subscriber *subscriber) {
	s.subscribersMutex.Lock()
	s.subscribers[subscriber] = struct{}{}
	s.subscribersMutex.Unlock()
	fmt.Println("Added subscriber", subscriber)
}

func (s *server) subscribe(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) error {
	var c *websocket.Conn
	subscriber := &subscriber{
		msgs: make(chan []byte, s.subscribeMessageBuffer),
	}
	s.addSubscriber(subscriber)

	c, err := websocket.Accept(writer, req, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)

	for {
		select {
		case msg := <-subscriber.msgs:
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			err := c.Write(ctx, websocket.MessageText, msg)

			if err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) broadcast(msg []byte) {
	s.subscribersMutex.Lock()
	for subscriber := range s.subscribers {
		subscriber.msgs <- msg
	}
	s.subscribersMutex.Unlock()
}

func main() {
	fmt.Println("Starting system monitor..")

	srv := NewServer()
	go func(s *server) {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				fmt.Println(err)
				continue
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				fmt.Println(err)
				continue
			}

			cpuSection, err := hardware.GetCpuSection()
			if err != nil {
				fmt.Println(err)
				continue
			}

			timeStamp := time.Now().Format("2006-01-02 15:04:05")
			msg := []byte(`
      <div hx-swap-oob="innerHTML:#update-timestamp">
        <p><i style="color: green" class="fa fa-circle"></i> ` + timeStamp + `</p>
      </div>
      <div hx-swap-oob="innerHTML:#system-data">` + systemSection + `</div>
      <div hx-swap-oob="innerHTML:#cpu-data">` + cpuSection + `</div>
      <div hx-swap-oob="innerHTML:#disk-data">` + diskSection + `</div>`)

			s.broadcast([]byte(msg))

			time.Sleep(3 * time.Second)
		}
	}(srv)

	err := http.ListenAndServe(":8080", &srv.mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
