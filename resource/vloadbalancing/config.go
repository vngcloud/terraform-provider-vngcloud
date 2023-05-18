package vloadbalancing

var (
	loadBalancerCreating = []string{"DELETING", "UPDATING", "CREATING", "CREATING-BILLING"}
	loadBalancerCreated  = []string{"CREATED"}
	loadBalancerDeleting = []string{"DELETING"}
	loadBalancerDeleted  = []string{"DELETED"}

	listenerCreating = []string{"DELETING", "UPDATING", "CREATING"}
	listenerCreated  = []string{"CREATED"}
	listenerUpdating = []string{"DELETING", "UPDATING", "CREATING"}
	listenerUpdated  = []string{"CREATED"}
	listenerDeleting = []string{"DELETING"}
	listenerDeleted  = []string{"DELETED"}

	poolCreating = []string{"DELETING", "UPDATING", "CREATING"}
	poolCreated  = []string{"CREATED"}
	poolUpdating = []string{"DELETING", "UPDATING", "CREATING"}
	poolUpdated  = []string{"CREATED"}
	poolDeleting = []string{"DELETING"}
	poolDeleted  = []string{"DELETED"}

	l7PolicyCreating = []string{"DELETING", "UPDATING", "CREATING"}
	l7PolicyCreated  = []string{"CREATED"}
	l7PolicyUpdating = []string{"DELETING", "UPDATING", "CREATING"}
	l7PolicyUpdated  = []string{"CREATED"}
	l7PolicyDeleting = []string{"DELETING"}
	l7PolicyDeleted  = []string{"DELETED"}
)
