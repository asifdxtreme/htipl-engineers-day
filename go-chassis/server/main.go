package main

import (
	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	_ "github.com/go-chassis/go-chassis/config-center"
	"github.com/go-chassis/go-chassis/core/lager"
	"net/http"
	rf "github.com/go-chassis/go-chassis/server/restful"

)

func main() {
	if err := chassis.Init(); err != nil {
		lager.Logger.Errorf("Init failed: %s", err)
		return
	}
	chassis.Run()
}

func init() {
	chassis.RegisterSchema("rest", &RestFulMessage{})
}

type RestFulMessage struct {
}

func (r *RestFulMessage) Saymessage(b *rf.Context) {
	id := b.ReadPathParameter("name")

	b.Write([]byte("Welcome to : " + id))
}


func (s *RestFulMessage) URLPatterns() []rf.Route {
	return []rf.Route{
		{http.MethodGet, "/saymessage/{name}", "Saymessage"},
	}
}
