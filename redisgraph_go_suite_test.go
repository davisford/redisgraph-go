package redisgraph_test

import (
	"os"
	"testing"

	"github.com/gomodule/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const LocalRedisURL = "0.0.0.0:6379"
const GraphName = "__rgtest"

var (
	// the graph to use in test suite
	g *Graph
	// redis connection
	conn redis.Conn
)

var _ = BeforeSuite(func() {
	url := os.Getenv("REDIS_TEST_URL")
	if url == "" {
		url = LocalRedisURL
	}
	conn, err := redis.Dial("tcp", url)
	Ω(err).To(BeNil(), "can't connect to redis for test suite")
	g = Graph{}.New(GraphName, conn)
})

var _ = AfterSuite(func() {
	if conn != nil {
		conn.Close()
	}
})

func TestRedisgraphGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisgraphGo Suite")
}
