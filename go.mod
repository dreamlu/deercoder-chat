module deercoder-chat

go 1.12

require (
	github.com/dreamlu/go-tool v1.0.7
	github.com/dreamlu/go.uuid v1.2.0
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.1
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/consul/api v1.1.0
	github.com/micro/go-micro v1.6.0
	github.com/nats-io/nats-server/v2 v2.0.0 // indirect
	github.com/prometheus/common v0.2.0
	github.com/ugorji/go v1.1.5-pre // indirect
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
