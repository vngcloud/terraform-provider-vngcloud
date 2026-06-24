package vdbv2

import (
	"log"
	"sync"
)

// kafkaClusterMutexes serializes mutating operations targeting the same Kafka
// cluster (create/update/delete of topics or users). Terraform invokes resource
// CRUD funcs concurrently by default, but the upstream API treats a cluster as
// a single-writer entity — N requests arriving at once leave the cluster in a
// transient state and collide. Holding this mutex from before the API call
// until the cluster has returned to ACTIVE makes those operations effectively
// sequential per cluster, while still permitting parallelism across different
// clusters.
var kafkaClusterMutexes sync.Map // map[string]*sync.Mutex

// lockKafkaCluster acquires the per-cluster mutex and returns an unlock func.
// Typical use: `defer lockKafkaCluster(clusterID)()`.
func lockKafkaCluster(clusterID string) func() {
	v, _ := kafkaClusterMutexes.LoadOrStore(clusterID, &sync.Mutex{})
	mu := v.(*sync.Mutex)
	log.Printf("[DEBUG] acquiring kafka cluster mutex: %s", clusterID)
	mu.Lock()
	log.Printf("[DEBUG] acquired kafka cluster mutex: %s", clusterID)
	return func() {
		mu.Unlock()
		log.Printf("[DEBUG] released kafka cluster mutex: %s", clusterID)
	}
}
