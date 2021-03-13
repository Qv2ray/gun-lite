package main

import (
	"errors"
	"flag"
	"github.com/Qv2ray/gun-lite/pkg/realgun"
	"golang.org/x/net/context"
	"io"
	"log"
	"net"
)

var (
	RemoteAddr = flag.String("remote", "", "remote gun server address")
	LocalAddr = flag.String("local", "", "local server endpoint")
	ServerName = flag.String("sni", "", "(optional) server name indication")
	ServiceName = flag.String("service", "", "(optional) custom service name")
	Cleartext = flag.Bool("cleartext", false, "(optional) use unsafe h2c")
)

func init() {
	flag.Parse()
}

func main() {
	if *RemoteAddr == "" {
		log.Fatalf("need remote address")
	}
	if *LocalAddr == "" {
		log.Fatal("need local endpoint")
	}
	listen, err := net.Listen("tcp", *LocalAddr)
	if err != nil {
		log.Fatalf("failed to listen tcp %v: %v", *LocalAddr, err)
	}

	client := realgun.NewGunClientWithContext(context.TODO(), &realgun.Config{
		RemoteAddr:  *RemoteAddr,
		ServerName:  *ServerName,
		ServiceName: *ServiceName,
		Cleartext:   *Cleartext,
	})

	for {
		localConn, err := listen.Accept()
		if err != nil {
			log.Printf("accept local failed: %v", err)
			continue
		}
		go func() {
			defer localConn.Close()
			remoteConn, err := client.DialConn()
			if err != nil {
				log.Printf("dial remote failed: %v", err)
			}


			go func() {
				defer remoteConn.Close()
				n, e := io.Copy(localConn, remoteConn)
				if e != nil && !errors.Is(e, net.ErrClosed) {
					log.Printf("copy from remote to local failed: %v", e)
				}
				log.Printf("copied %d bytes from remote to local", n)
			}()

			n, e := io.Copy(remoteConn, localConn)
			if e != nil && !errors.Is(e, net.ErrClosed) {
				log.Printf("copy from local to remote failed: %v", e)
			}
			log.Printf("copied %d bytes from local to remote", n)
		}()

	}
}