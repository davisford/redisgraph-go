package examples

import (
	"testing"

	rg "github.com/davisford/redisgraph-go"
	"github.com/gomodule/redigo/redis"
)

func TestExample(t *testing.T) {
	conn, _ := redis.Dial("tcp", "0.0.0.0:6379")
	defer conn.Close()

	conn.Do("FLUSHALL")
	graph := rg.Graph{}.New("social", conn)

	john := rg.Node{
		Label: "person",
		Properties: map[string]interface{}{
			"name":   "John Doe",
			"age":    33,
			"gender": "male",
			"status": "single",
		},
	}
	graph.PutNode(rg.CREATE, &john)

	japan := rg.Node{
		Label: "country",
		Properties: map[string]interface{}{
			"name": "Japan",
		},
	}
	graph.PutNode(rg.CREATE, &japan)

	edge := rg.Edge{
		Source:      &john,
		Relation:    "visited",
		Destination: &japan,
	}
	graph.PutEdge(rg.CREATE, &edge)

	graph.Commit()

	query := `MATCH (p:person)-[v:visited]->(c:country)
		   RETURN p.name, p.age, v.purpose, c.name`
	rs, _ := graph.Query(query)

	rs.PrettyPrint()
}
