package main

import (
	_ "github.com/go-chassis/go-chassis/examples/demo/server/schemas"
	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis/bootstrap"
	_ "github.com/go-chassis/go-chassis/config-center"
	"github.com/go-chassis/go-chassis/core/lager"
	_ "github.com/huaweicse/auth/adaptor/gochassis"
	"net/http"
	rf "github.com/go-chassis/go-chassis/server/restful"

)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder
func main() {
	//start all server you register in server/schemas.
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

	b.Write([]byte("get name: " + id))
}


func (s *RestFulMessage) URLPatterns() []rf.Route {
	return []rf.Route{
		{http.MethodGet, "/saymessage/{name}", "Saymessage"},
	}
}