package main

import (
	"github.com/micro/micro/cmd"
)
// main.go
func init() {
	//token := &token.Token{}
	//token.InitConfig("127.0.0.1:8500", "micro", "config", "jwt-key", "key")

	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithName("auth"),
	//	plugin.WithHandler(
	//		auth.JWTAuthWrapper(),
	//	),
	//))
}
const name = "API gateway"

func main() {

	cmd.Init()
}
