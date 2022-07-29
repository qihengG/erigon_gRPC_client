# erigon_gRPC_client
Test connection to erigon through gRPC

1. Start an erigon node
./build/bin/erigon --datadir goerli --chain goerli --metrics --metrics.expensive --downloader.api.addr 127.0.0.1:9093

2. Start a separate downloader
./build/bin/downloader --downloader.api.addr 127.0.0.1:9093

3. Run `go run downloaderClient.go`

Sample Output:
Connected to downloader
Starting to query download progress
Get response form downloader @ 127.0.0.1:9093
Curr progress: 0.97

Reference:

gRPC tutorial https://grpc.io/docs/languages/go/basics/

erigon downloader log https://github.com/ledgerwatch/erigon/blob/554d27468e248cd9fd97ff50b46d028cc941aea5/eth/stagedsync/stage_headers.go#L1361 - where the stats info is extracted

erigon google protobuf definition https://github.com/ledgerwatch/erigon-lib/tree/main/interfaces

erigon docker-compose.yml: https://github.com/ledgerwatch/erigon/blob/devel/docker-compose.yml - for flags reference



