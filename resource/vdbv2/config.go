package vdbv2

import "time"

var (
	databaseCreatePending    = []string{"BUILDING", "BUILD"}
	databaseCreateTarget     = []string{"ACTIVE"}
	databaseCreateTimeout    = 20 * time.Minute
	databaseCreateDelay      = 10 * time.Second
	databaseCreateMinTimeout = 10 * time.Second

	databaseDeletePending    = []string{"DELETING", "BACKUP", "deleting", "ACTIVE"}
	databaseDeleteTarget     = []string{"DELETED"}
	databaseDeleteTimeout    = 20 * time.Minute
	databaseDeleteDelay      = 10 * time.Second
	databaseDeleteMinTimeout = 10 * time.Second

	databaseResizePending    = []string{"BUILDING", "BUILD", "BACKUP", "resizing", "updating"}
	databaseResizeTarget     = []string{"ACTIVE", "RESTART_REQUIRED"}
	databaseResizeTimeout    = 20 * time.Minute
	databaseResizeDelay      = 10 * time.Second
	databaseResizeMinTimeout = 10 * time.Second

	databaseStartPending    = []string{"BUILDING", "BACKUP", "SHUTDOWN"}
	databaseStartTarget     = []string{"ACTIVE"}
	databaseStartTimeout    = 10 * time.Minute
	databaseStartDelay      = 10 * time.Second
	databaseStartMinTimeout = 10 * time.Second

	databaseStopPending    = []string{"BUILDING", "BACKUP", "ACTIVE"}
	databaseStopTarget     = []string{"SHUTDOWN"}
	databaseStopTimeout    = 10 * time.Minute
	databaseStopDelay      = 10 * time.Second
	databaseStopMinTimeout = 10 * time.Second

	databaseRebootPending    = []string{"BUILDING", "BACKUP", "REBOOT"}
	databaseRebootTarget     = []string{"ACTIVE"}
	databaseRebootTimeout    = 10 * time.Minute
	databaseRebootDelay      = 10 * time.Second
	databaseRebootMinTimeout = 10 * time.Second

	databasePromoteTimeout = 10 * time.Minute
)
