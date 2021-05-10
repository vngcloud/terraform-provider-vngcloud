package vserver

var (
	networkCreating      = []string{"CREATING"}
	networkCreated       = []string{"ACTIVE"}
	secgroupRuleCreating = []string{"CREATING"}
	secgroupRuleCreated  = []string{"ACTIVE"}
	secgroupCreating     = []string{"CREATING"}
	secgroupCreated      = []string{"ACTIVE"}
	serverCreating       = []string{"CREATING"}
	serverCreated        = []string{"ACTIVE"}
	serverRebooting      = []string{"REBOOTING"}
	serverRebooted       = []string{"ACTIVE"}
	serverStopping       = []string{"ACTIVE"}
	serverStopped        = []string{"STOPPED"}
	serverStarting       = []string{"STOPPED"}
	serverStarted        = []string{"ACTIVE"}
	subnetCreating       = []string{"CREATING"}
	subnetCreated        = []string{"ACTIVE"}
	volumeCreating       = []string{"CREATING"}
	volumeCreated        = []string{"AVAILABLE"}
	volumeAttaching      = []string{"AVAILABLE"}
	volumeAttached       = []string{"IN-USE"}
	volumeDetaching      = []string{"IN-USE"}
	volumeDetached       = []string{"AVAILABLE"}
)
