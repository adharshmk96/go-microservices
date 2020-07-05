package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to the cluster
	cluster = gocql.NewCluster("cassandra-server-srv")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	// session, err := cluster.CreateSession()
	// if err != nil {
	// 	panic("Erro cr connecting to db")
	// }
	// fmt.Println("Connection successful")
	// defer session.Close()
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
