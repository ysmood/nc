package main

import (
	"io"
	"net"
	"os"

	kit "github.com/ysmood/gokit"
)

func main() {
	app := kit.TasksNew("nc", "a simple netcat tool")
	app.Version("v0.0.1")

	kit.Tasks().App(app).Add(
		kit.Task("serve", "run as server for both tcp and udp on the same port").
			Init(func(cmd kit.TaskCmd) func() {
				cmd.Default()
				addr := cmd.Arg("address", "the host and port address").Default(":8080").String()
				return func() {
					serve(*addr)
				}
			}),
		kit.Task("send", "send tcp or udp package").
			Init(func(cmd kit.TaskCmd) func() {
				addr := cmd.Arg("address", "the host and port address").Default(":8080").String()
				protocol := cmd.Flag("protocol", "protocol to use").Short('p').Default("tcp").String()
				return func() {
					send(*protocol, *addr)
				}
			}),
	).Do()
}

func serve(addr string) {
	tcpL, err := net.Listen("tcp", addr)
	kit.E(err)
	defer func() { kit.E(tcpL.Close()) }()

	udpL, err := net.ListenPacket("udp", addr)
	kit.E(err)
	defer func() { kit.E(udpL.Close()) }()

	kit.All(func() {
		conn, err := tcpL.Accept()
		kit.E(err)

		go func() { kit.E(io.Copy(conn, os.Stdin)) }()
		kit.E(io.Copy(os.Stdout, conn))
	}, func() {
		for {
			buf := make([]byte, 30*1024)
			n, _, err := udpL.ReadFrom(buf)
			kit.E(err)
			kit.E(os.Stdout.Write(buf[:n]))
		}
	})
}

func send(protocol, addr string) {
	conn, err := net.Dial(protocol, addr)
	kit.E(err)

	go func() { kit.E(io.Copy(conn, os.Stdin)) }()
	kit.E(io.Copy(os.Stdout, conn))
}
