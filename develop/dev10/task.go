package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// go run task.go opennet.ru 80
func main() {
	var wg sync.WaitGroup

	addr := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]
	fullAddr := fmt.Sprintf("%s:%s", addr, port)

	timeout := flag.Int("timeout", 10, "timeout")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	eofChannel := make(chan bool, 2)

	conn, err := net.DialTimeout("tcp", fullAddr, time.Second*time.Duration(*timeout))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	sender := bufio.NewReader(os.Stdin)

	wg.Add(3)
	go func() {
		defer wg.Done()

		for {
			err := listen(*reader)
			if err != nil {
				conn.Close()
				log.Println(err)
				eofChannel <- true
				return
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			err := send(*sender, conn)
			if err != nil {
				conn.Close()
				log.Println(err)
				eofChannel <- true
				return
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-c:
				conn.Close()
				log.Println("Connection closed")
				return

			case <-eofChannel:
				conn.Close()
				log.Println("EOF")
				return

			default:
				continue
			}
		}
	}()

	wg.Wait()
}

func listen(reader bufio.Reader) error {
	data, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(os.Stdout, data)
	return err
}

func send(sender bufio.Reader, conn net.Conn) error {
	data, err := sender.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(data))
	return err
}
