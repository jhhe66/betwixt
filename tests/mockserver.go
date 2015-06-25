package tests

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/go-commons/network"
	"github.com/zubairhamed/canopus"
)

func NewMockServer() betwixt.Server {
	return &MockServer{
		stats: 		&MockServerStatistics{},
		httpServer: network.NewDefaultHttpServer("8080"),
	}
}

type MockServer struct {
	stats betwixt.ServerStatistics
	httpServer *network.HttpServer
	coapServer *canopus.CoapServer
}

func (server *MockServer) Start() {

}

func (server *MockServer) UseRegistry(reg betwixt.Registry) {

}

func (server *MockServer) On(e betwixt.EventType, fn betwixt.FnEvent) {

}

func (server *MockServer) GetClients() map[string]betwixt.RegisteredClient {
	return make(map[string]betwixt.RegisteredClient)
}

func (server *MockServer) GetStats() betwixt.ServerStatistics {
	return server.stats
}

func (server *MockServer) GetHttpServer() (*network.HttpServer) {
	return server.httpServer
}

func (server *MockServer) GetCoapServer() (*canopus.CoapServer) {
	return server.coapServer
}


func (server *MockServer) GetClient(id string) betwixt.RegisteredClient {
	return nil
}
