package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	protol_dowloader "github.com/ledgerwatch/erigon-lib/gointerfaces/downloader"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("addr", "127.0.0.1:9093", "dowloader server")
)

func main() {
	conn, errConn := grpc.Dial(*serverAddr, grpc.WithInsecure())
	defer conn.Close()
	if errConn != nil {
		fmt.Printf("[---errConn---]:")
		fmt.Println(errConn)
		return
	}
	downloaderClient := protol_dowloader.NewDownloaderClient(conn)
	fmt.Println("Connected to downloader")
	fmt.Println("Starting to query download progress")
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	stats, errResp := downloaderClient.Stats(ctx, &protol_dowloader.StatsRequest{})
	if errResp != nil {
		fmt.Printf("[---errResp---]:")
		fmt.Println(errResp)
		return
	}
	fmt.Printf("Get response form downloader @ %s\n", *serverAddr)
	fmt.Printf("Curr progress: %.2f\n", stats.Progress)
}