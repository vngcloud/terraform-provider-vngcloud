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
)
