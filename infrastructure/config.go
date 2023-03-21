package infrastructure

import (
	"github.com/Jhooomn/concurrent-eda-pio/sub/infrastructure/broker"
	"github.com/Jhooomn/concurrent-eda-pio/sub/infrastructure/server"
)

func SetUp() {
	go broker.SetUp()
	server.SetUp() // last line of code
}
