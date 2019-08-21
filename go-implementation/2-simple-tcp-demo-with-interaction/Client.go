package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var (
		err  error
		conn net.Conn
		buf  []byte
		n    int
	)

	conn, err = net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("dial错误", err)
		return
	}
	defer conn.Close()

	//go func(conn net.Conn) {
	_, err = conn.Write([]byte("hello,i am client"))
	if err != nil {
		fmt.Println("写入错误", err)
		return
	}

	buf = make([]byte, 4096)
	go func() {
		buf2 := make([]byte, 4096)

		for {
			fmt.Println(1234)
			n, err = os.Stdin.Read(buf2)
			fmt.Println(string(buf2))

			if err != nil {
				fmt.Println("读取键盘输入错误", err)
				return
			}
			conn.Write(buf2[:n])
		}
	}()

	for {
		n, err = conn.Read(buf)

		if err != nil {
			fmt.Println("读取错误", err)
			return
		}
		fmt.Println("收到服务器信息:", string(buf[:n]))
	}

	for {
		time.Sleep(1 * time.Second)
	}

	//}(conn)

}
