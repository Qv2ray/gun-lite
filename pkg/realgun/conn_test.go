package realgun

import (
	"io"
	"testing"

	"golang.org/x/net/context"
)

func Test(t *testing.T) {
	client := NewGunClientWithContext(context.TODO(), &Config{
		RemoteAddr: "127.0.0.1:23333",
		Cleartext:  true,
	})
	conn, err := client.DialConn()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	_, err = io.Copy(conn, conn)
	panic(err)
}
