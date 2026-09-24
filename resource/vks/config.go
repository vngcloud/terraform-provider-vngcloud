package vks

var (
	CREATING = []string{"CREATING", "DEGRADED"}
	ERROR    = []string{"ERROR"}
	DELETING = []string{"DELETING"}
	ACTIVE   = []string{"ACTIVE"}
	UPDATING  = []string{"UPDATING", "DEGRADED"}
	UPGRADING = []string{"UPGRADING", "DEGRADED"}
	DELETED  = []string{"DELETED"}
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
