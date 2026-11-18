package vks

import "time"

var (
	CREATING  = []string{"CREATING", "DEGRADED"}
	ERROR     = []string{"ERROR"}
	DELETING  = []string{"DELETING"}
	ACTIVE    = []string{"ACTIVE"}
	UPDATING  = []string{"UPDATING", "DEGRADED"}
	UPGRADING = []string{"UPGRADING", "DEGRADED"}
	DELETED   = []string{"DELETED"}
	// Transitional cluster statuses during an async update/upgrade; used to wait for the
	// cluster to settle back to ACTIVE before issuing the next mutation. These are cluster
	// statuses (from the backend ClusterConst.Status enum) — note UPGRADING/DEGRADED are
	// node-group statuses a cluster never reports, so they are intentionally absent.
	PENDING_UPDATE = []string{
		"WAITING_UPDATE", "UPDATING",
		"WAITING_AUTO_UPGRADE", "AUTO_UPGRADING",
		"WAITING_FORCE_UPGRADE", "FORCE_UPGRADING", "SCHEDULING_UPGRADE",
	}
)

// Cluster add-on statuses (vks-api AddonConst.Status). DEGRADED means the Helm
// release is applied but its workloads are not ready yet (e.g. no nodes): the
// add-on is usable and turns ACTIVE on its own, so creates/updates accept it.
var (
	addonCreatePending = []string{"CREATING"}
	addonCreateTarget  = []string{"ACTIVE", "DEGRADED"}
	addonDeletePending = []string{"DELETING", "ACTIVE", "DEGRADED"}
	addonFailed        = []string{"CREATE_FAILED", "UPDATE_FAILED", "DELETE_FAILED"}

	// Add-on update (DescribeUpdate) statuses.
	addonUpdateInProgress = []string{"InProgress"}
	addonUpdateSuccessful = []string{"Successful"}
)

const (
	addonDelay      = 10 * time.Second
	addonMinTimeout = 10 * time.Second
)
