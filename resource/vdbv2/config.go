package vdbv2

import "time"

var (
	databaseCreatePending    = []string{"BUILDING", "BUILD"}
	databaseCreateTarget     = []string{"ACTIVE"}
	databaseCreateTimeout    = 30 * time.Minute
	databaseCreateDelay      = 60 * time.Second
	databaseCreateMinTimeout = 10 * time.Second

	databaseDeletePending    = []string{"DELETING", "BACKUP", "deleting", "ACTIVE"}
	databaseDeleteTarget     = []string{"DELETED"}
	databaseDeleteTimeout    = 30 * time.Minute
	databaseDeleteDelay      = 30 * time.Second
	databaseDeleteMinTimeout = 10 * time.Second

	databaseResizePending    = []string{"BUILDING", "BUILD", "BACKUP", "RESIZING", "UPDATING", "resizing", "updating"}
	databaseResizeTarget     = []string{"ACTIVE", "RESTART_REQUIRED"}
	databaseResizeTimeout    = 30 * time.Minute
	databaseResizeDelay      = 30 * time.Second
	databaseResizeMinTimeout = 10 * time.Second

	databaseStartPending    = []string{"BUILDING", "BACKUP", "SHUTDOWN", "starting"}
	databaseStartTarget     = []string{"ACTIVE"}
	databaseStartTimeout    = 20 * time.Minute
	databaseStartDelay      = 30 * time.Second
	databaseStartMinTimeout = 10 * time.Second

	databaseStopPending    = []string{"BUILDING", "BACKUP", "ACTIVE", "stopping"}
	databaseStopTarget     = []string{"SHUTDOWN"}
	databaseStopTimeout    = 20 * time.Minute
	databaseStopDelay      = 20 * time.Second
	databaseStopMinTimeout = 10 * time.Second

	databaseRebootPending    = []string{"BUILDING", "BACKUP", "REBOOT", "rebooting"}
	databaseRebootTarget     = []string{"ACTIVE"}
	databaseRebootTimeout    = 10 * time.Minute
	databaseRebootDelay      = 20 * time.Second
	databaseRebootMinTimeout = 10 * time.Second

	databasePromoteTimeout = 10 * time.Minute

	kafkaClusterCreatePending = []string{
		"WAITING_WORKSPACE_CREATING",
		"WORKSPACE_CREATING",
		"INFRA_CREATING",
		"NODES_CREATING",
		"WAITING_CLUSTER_CREATING",
		"CLUSTER_CREATING",
		"CREATING_BILL",
	}
	kafkaClusterCreateTarget     = []string{"ACTIVE"}
	kafkaClusterCreateTimeout    = 60 * time.Minute
	kafkaClusterCreateDelay      = 60 * time.Second
	kafkaClusterCreateMinTimeout = 10 * time.Second

	// kafkaClusterCreateLockHold là khoảng grace giữ thêm sau khi CreateOrderCluster
	// trả về, nhằm phủ hết cửa sổ lock theo user ở phía server (server lock vài giây
	// khi nhận request accept). Trong khoảng này request create kafka kế tiếp phải chờ.
	kafkaClusterCreateLockHold = 5 * time.Second

	kafkaClusterDeletePending = []string{
		"ACTIVE",
		"WAITING_CLUSTER_DELETING",
		"CLUSTER_DELETING",
		"WAITING_WORKSPACE_DELETING",
		"WORKSPACE_DELETING",
		"INFRA_DELETING",
		"TOPICS_DELETING",
		"USERS_DELETING",
	}
	kafkaClusterDeleteTarget     = []string{"DELETED"}
	kafkaClusterDeleteTimeout    = 60 * time.Minute
	kafkaClusterDeleteDelay      = 30 * time.Second
	kafkaClusterDeleteMinTimeout = 10 * time.Second

	kafkaClusterUpdateBrokerCountPending = []string{
		"WAITING_UPDATING",
		"UPDATING",
		"WAITING_SCALING",
		"SCALING_OUT_GET_INFO",
		"INFRA_SCALING_OUT",
		"NODES_SCALING_OUT",
		"INFRA_SCALING_IN",
		"NODES_SCALING_IN",
		"WAITING_CLUSTER_SCALING_OUT",
		"CLUSTER_SCALING_OUT",
		"WAITING_CLUSTER_SCALING_IN",
		"CLUSTER_SCALING_IN",
		"WAITING_REBALANCING_OUT",
		"REBALANCING_OUT",
		"WAITING_REBALANCING_IN",
		"REBALANCING_IN",
	}
	kafkaClusterUpdateBrokerCountTarget = []string{"ACTIVE"}

	kafkaClusterUpdateStorageTypePending = []string{
		"WAITING_UPDATING",
		"UPDATING",
		"WAITING_UPDATING_STORAGE_TYPE",
		"UPDATING_STORAGE_TYPE",
	}
	kafkaClusterUpdateStorageTypeTarget = []string{"ACTIVE"}

	kafkaClusterUpdateStorageSizePending = []string{
		"WAITING_UPDATING",
		"UPDATING",
		"WAITING_UPDATING_STORAGE_SIZE",
		"UPDATING_STORAGE_SIZE",
	}
	kafkaClusterUpdateStorageSizeTarget = []string{"ACTIVE"}

	kafkaClusterUpdateConfigGroupPending = []string{
		"WAITING_UPDATING",
		"UPDATING",
		"WAITING_UPDATING_CONFIG_GROUP",
		"UPDATING_CONFIG_GROUP",
	}
	kafkaClusterUpdateConfigGroupTarget = []string{"ACTIVE", "WARNING_CONFIG_GROUP"}

	kafkaClusterUpdateAuthPending = []string{
		"WAITING_UPDATING",
		"UPDATING",
	}
	kafkaClusterUpdateAuthTarget = []string{"ACTIVE"}

	kafkaClusterUpdatePublicAccessPending = []string{
		"WAITING_UPDATING",
		"UPDATING",
		"WAITING_FLOATING_IPS_UPDATING",
		"FLOATING_IPS_UPDATING",
	}
	kafkaClusterUpdatePublicAccessTarget = []string{"ACTIVE"}

	kafkaClusterUpdateSecRulesPending = []string{
		"WAITING_UPDATING",
		"UPDATING",
		"SECURITY_GROUP_RULES_UPDATING",
	}
	kafkaClusterUpdateSecRulesTarget = []string{"ACTIVE"}

	kafkaClusterUpdateTimeout    = 60 * time.Minute
	kafkaClusterUpdateDelay      = 30 * time.Second
	kafkaClusterUpdateMinTimeout = 10 * time.Second

	kafkaClusterTopicUserPending = []string{
		"TOPIC_PROCESSING",
		"USER_PROCESSING",
	}
	kafkaClusterTopicUserTarget     = []string{"ACTIVE"}
	kafkaClusterTopicUserTimeout    = 10 * time.Minute
	kafkaClusterTopicUserDelay      = 10 * time.Second
	kafkaClusterTopicUserMinTimeout = 5 * time.Second

	kafkaTopicCreatePending    = []string{"WAITING_CREATING", "CREATING"}
	kafkaTopicCreateTarget     = []string{"ACTIVE"}
	kafkaTopicCreateTimeout    = 10 * time.Minute
	kafkaTopicCreateDelay      = 10 * time.Second
	kafkaTopicCreateMinTimeout = 5 * time.Second

	kafkaTopicUpdatePending    = []string{"WAITING_UPDATING", "UPDATING"}
	kafkaTopicUpdateTarget     = []string{"ACTIVE"}
	kafkaTopicUpdateTimeout    = 10 * time.Minute
	kafkaTopicUpdateDelay      = 10 * time.Second
	kafkaTopicUpdateMinTimeout = 5 * time.Second

	kafkaTopicDeletePending = []string{
		"ACTIVE",
		"WAITING_DELETING",
		"DELETING",
		"WAITING_DELETING_ALL",
		"DELETING_ALL",
	}
	kafkaTopicDeleteTarget     = []string{"DELETED"}
	kafkaTopicDeleteTimeout    = 10 * time.Minute
	kafkaTopicDeleteDelay      = 10 * time.Second
	kafkaTopicDeleteMinTimeout = 5 * time.Second

	kafkaUserCreatePending    = []string{"CREATING"}
	kafkaUserCreateTarget     = []string{"ACTIVE"}
	kafkaUserCreateTimeout    = 10 * time.Minute
	kafkaUserCreateDelay      = 10 * time.Second
	kafkaUserCreateMinTimeout = 5 * time.Second

	kafkaUserUpdatePending    = []string{"UPDATING"}
	kafkaUserUpdateTarget     = []string{"ACTIVE"}
	kafkaUserUpdateTimeout    = 10 * time.Minute
	kafkaUserUpdateDelay      = 10 * time.Second
	kafkaUserUpdateMinTimeout = 5 * time.Second

	kafkaUserDeletePending    = []string{"ACTIVE", "DELETING", "DELETING_ALL"}
	kafkaUserDeleteTarget     = []string{"DELETED"}
	kafkaUserDeleteTimeout    = 10 * time.Minute
	kafkaUserDeleteDelay      = 10 * time.Second
	kafkaUserDeleteMinTimeout = 5 * time.Second

	postgresClusterCreatePending    = []string{"BUILDING", "WAIT_BILLING"}
	postgresClusterCreateTarget     = []string{"ACTIVE"}
	postgresClusterCreateTimeout    = 60 * time.Minute
	postgresClusterCreateDelay      = 60 * time.Second
	postgresClusterCreateMinTimeout = 10 * time.Second

	postgresClusterUpdatePending    = []string{"BUILDING", "WAIT_BILLING"}
	postgresClusterUpdateTarget     = []string{"ACTIVE"}
	postgresClusterUpdateTimeout    = 60 * time.Minute
	postgresClusterUpdateDelay      = 30 * time.Second
	postgresClusterUpdateMinTimeout = 10 * time.Second

	postgresClusterRebootPending = []string{"BUILDING", "REBOOTING"}
	postgresClusterRebootTarget  = []string{"ACTIVE"}

	postgresClusterUpdateConfigGroupPending = []string{"BUILDING"}
	postgresClusterUpdateConfigGroupTarget  = []string{"ACTIVE", "RESTART_REQUIRED"}

	postgresClusterDeletePending    = []string{"DELETING"}
	postgresClusterDeleteTarget     = []string{"DELETED"}
	postgresClusterDeleteTimeout    = 60 * time.Minute
	postgresClusterDeleteDelay      = 30 * time.Second
	postgresClusterDeleteMinTimeout = 10 * time.Second
)
