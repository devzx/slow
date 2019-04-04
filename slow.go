package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	waitTime := flag.Int("wait", 1, "time to wait in seconds between printing output")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			time.Sleep(time.Duration(*waitTime) * time.Second)
		}
	}()

	<-c
}
