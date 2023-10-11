package vserver

var (
	loadBalancerCreating     = []string{"CREATING"}
	loadBalancerCreated      = []string{"CREATED"}
	networkCreating          = []string{"CREATING"}
	networkCreated           = []string{"ACTIVE"}
	networkDeleting          = []string{"DELETING"}
	networkDeleted           = []string{"DELETED"}
	secgroupCreating         = []string{"CREATING"}
	secgroupCreated          = []string{"ACTIVE"}
	serverCreating           = []string{"CREATING", "CREATING-BILLING"}
	serverCreated            = []string{"ACTIVE"}
	serverRebooting          = []string{"REBOOTING"}
	serverRebooted           = []string{"ACTIVE"}
	serverStopping           = []string{"TURNING-OFF"}
	serverStopped            = []string{"STOPPED"}
	serverStarting           = []string{"STARTING"}
	serverStarted            = []string{"ACTIVE"}
	serverResizing           = []string{"CHANGING-FLAVOR", "VERIFYING-FLAVOR"}
	serverResize             = []string{"ACTIVE", "STOPPED"}
	serverDeleting           = []string{"DELETING"}
	serverDeleted            = []string{"DELETED"}
	subnetCreating           = []string{"CREATING"}
	subnetCreated            = []string{"ACTIVE"}
	subnetDeleting           = []string{"DELETING"}
	subnetDeleted            = []string{"DELETED"}
	volumeCreating           = []string{"CREATING", "CREATING-BILLING"}
	volumeCreated            = []string{"AVAILABLE"}
	volumeAttaching          = []string{"AVAILABLE", "ATTACHING"}
	volumeAttached           = []string{"IN-USE"}
	volumeDetaching          = []string{"IN-USE", "DETACHING"}
	volumeResizing           = []string{"RESIZING", "CHANGING-IOPS"}
	volumeResized            = []string{"IN-USE", "AVAILABLE"}
	volumeDetached           = []string{"AVAILABLE"}
	volumeDeleting           = []string{"DELETING"}
	volumeDeleted            = []string{"DELETED"}
	serverChangingSecGroup   = []string{"CHANGING-SECURITY-GROUP"}
	serverChangedSecGroup    = []string{"ACTIVE", "STOPPED"}
	k8sClusterCreating       = []string{"CREATING", "CREATING-BILLING"}
	k8sClusterCreated        = []string{"ACTIVE"}
	k8sClusterDeleting       = []string{"DELETING"}
	k8sClusterDeleted        = []string{"DELETED"}
	k8sClusterAttachingLB    = []string{"PROCESSING"}
	k8sClusterAttachedLB     = []string{"FINISH"}
	clusterNodeGroupCreating = []string{"CREATING"}
	clusterNodeGroupCreated  = []string{"ACTIVE"}
	clusterNodeGroupDeleting = []string{"DELETING"}
	clusterNodeGroupDeleted  = []string{"DELETED"}
	clusterNodeGroupScaling  = []string{"PROCESSING"}
	clusterNodeGroupScaled   = []string{"FINISH"}
	routeTableCreating       = []string{"CREATING"}
	routeTableCreated        = []string{"ACTIVE"}
	routeTableDeleting       = []string{"DELETING"}
	routeTableDeleted        = []string{"DELETED"}
	routeTableUpdating       = []string{"UPDATING"}
	routeTableUpdated        = []string{"ACTIVE"}
)
