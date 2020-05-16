package cassandra

import (
	"log"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster:
	log.Println("initializing cassandra session")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

// GetSession - create a new cassandra session object
func GetSession() *gocql.Session {
	return session
}
