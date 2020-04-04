package cassandra

import (
	"log"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster:
	log.Println("initializing cassandra session")
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

// GetSession - create a new cassandra session object
func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
