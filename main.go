package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Printf("order allow,deny\n")
	fmt.Printf("allow from all\n")
	filename := "host.txt"

	fp, err := os.Open(filename)
	if err != nil {
		panic("not file")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// ここで一行ずつ処理
		ip := net.ParseIP(scanner.Text())
		ipv4 := ip.To4()
		fmt.Printf("SetEnvIF X-Forwarded-For ^%d\\.%d\\.%d\\.%d DenyIP\n", ipv4[0], ipv4[1], ipv4[2], ipv4[3])    // 8.8.8.8
		fmt.Printf("REMOTE_ADDR X-Forwarded-For ^%d\\.%d\\.%d\\.%d DenyIP\n", ipv4[0], ipv4[1], ipv4[2], ipv4[3]) // 8.8.8.8
	}

	if err = scanner.Err(); err != nil {
		panic("not line")
	}

	fmt.Printf("deny from env=DenyIP")
}
