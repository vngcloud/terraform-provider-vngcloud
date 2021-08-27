package vserver

var (
	networkCreating        = []string{"CREATING"}
	networkCreated         = []string{"ACTIVE"}
	secgroupRuleCreating   = []string{"CREATING"}
	secgroupRuleCreated    = []string{"ACTIVE"}
	secgroupCreating       = []string{"CREATING"}
	secgroupCreated        = []string{"ACTIVE"}
	serverCreating         = []string{"CREATING", "CREATING-BILLING"}
	serverCreated          = []string{"ACTIVE"}
	serverRebooting        = []string{"REBOOTING"}
	serverRebooted         = []string{"ACTIVE"}
	serverStopping         = []string{"TURNING-OFF"}
	serverStopped          = []string{"STOPPED"}
	serverStarting         = []string{"STARTING"}
	serverStarted          = []string{"ACTIVE"}
	subnetCreating         = []string{"CREATING"}
	subnetCreated          = []string{"ACTIVE"}
	volumeCreating         = []string{"CREATING", "CREATING-BILLING"}
	volumeCreated          = []string{"AVAILABLE"}
	volumeAttaching        = []string{"AVAILABLE", "ATTACHING"}
	volumeAttached         = []string{"IN-USE"}
	volumeDetaching        = []string{"IN-USE", "DETACHING"}
	volumeDetached         = []string{"AVAILABLE"}
	serverChangingSecGroup = []string{"CHANGING-SECURITY-GROUP"}
	serverChangedSecGroup  = []string{"ACTIVE", "STOPPED"}
)
