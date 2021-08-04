package vdb

import "time"

var (
	databaseCreatePending    = []string{"BUILDING", "BUILD", "BACKUP", "RESTART_REQUIRED", "REBOOT"}
	databaseCreateTarget     = []string{"ACTIVE"}
	databaseCreateTimeout    = 10 * time.Minute
	databaseCreateDelay      = 60 * time.Second
	databaseCreateMinTimeout = 10 * time.Second

	databaseDeletePending    = []string{"BUILDING", "BACKUP", "ACTIVE", "SHUTDOWN"}
	databaseDeleteTarget     = []string{"TRASH"}
	databaseDeleteTimeout    = 10 * time.Minute
	databaseDeleteDelay      = 10 * time.Second
	databaseDeleteMinTimeout = 10 * time.Second

	databaseDeleteInTrashTimeout = 10 * time.Minute

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
