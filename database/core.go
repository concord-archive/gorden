package database

import (
	"github.com/gocql/gocql"
)

func setup_cassandra() *gocql.ClusterConfig {
	client := gocql.NewCluster("localhost")
	return client
}
