package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	// cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	// connect to the cluster
	cluster := gocql.NewCluster("cassandra-server-srv")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	// session, err := cluster.CreateSession()
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	// fmt.Println("Connection successful")
	// defer session.Close()
}

func GetSession() *gocql.Session {
	return session
}
