/*
Copyright 2022 Zheng Dayu
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package windowsupdate

// OperationResultCode defines the possible results of a download, install, uninstall, or verification operation on an update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-operationresultcode
const (
	OperationResultCodeOrcNotStarted int32 = iota
	OperationResultCodeOrcInProgress
	OperationResultCodeOrcSucceeded
	OperationResultCodeOrcSucceededWithErrors
	OperationResultCodeOrcFailed
	OperationResultCodeOrcAborted
)

// UpdateOperation defines the types of operations that can be performed on updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-updateoperation
const (
	UpdateOperationUoInstallation int32 = iota + 1
	UpdateOperationUoUninstallation
)

// UpdateType defines the types of updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-updatetype
const (
	UpdateTypeUtSoftware int32 = iota + 1
	UpdateTypeUtDriver
)

// DeploymentAction defines the deployment actions for updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-deploymentaction
const (
	DeploymentActionDaNone int32 = iota
	DeploymentActionDaInstallation
	DeploymentActionDaUninstallation
	DeploymentActionDaDetection
)

// DownloadPriority defines the download priority levels.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-downloadpriority
const (
	DownloadPriorityDpLow int32 = iota
	DownloadPriorityDpNormal
	DownloadPriorityDpHigh
	DownloadPriorityDpExtraHigh
)

// ServerSelection defines the server selection options.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-serverselection
const (
	ServerSelectionSsDefault int32 = iota
	ServerSelectionSsManagedServer
	ServerSelectionSsWindowsUpdate
	ServerSelectionSsOthers
)

// UpdateLockdownOption defines the update lockdown options.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-updatelockdownoption
const (
	UpdateLockdownOptionUloForWebsiteAccess int32 = 1 << iota
)

// AutomaticUpdatesNotificationLevel defines the automatic updates notification levels.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-automaticupdatesnotificationlevel
const (
	AutomaticUpdatesNotificationLevelAunlNotConfigured int32 = iota
	AutomaticUpdatesNotificationLevelAunlDisabled
	AutomaticUpdatesNotificationLevelAunlNotifyBeforeDownload
	AutomaticUpdatesNotificationLevelAunlNotifyBeforeInstallation
	AutomaticUpdatesNotificationLevelAunlScheduledInstallation
)

// AutomaticUpdatesScheduledInstallationDay defines the days for scheduled installation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-automaticupdatesscheduledinstallationday
const (
	AutomaticUpdatesScheduledInstallationDayAusidEveryDay int32 = iota
	AutomaticUpdatesScheduledInstallationDayAusidEverySunday
	AutomaticUpdatesScheduledInstallationDayAusidEveryMonday
	AutomaticUpdatesScheduledInstallationDayAusidEveryTuesday
	AutomaticUpdatesScheduledInstallationDayAusidEveryWednesday
	AutomaticUpdatesScheduledInstallationDayAusidEveryThursday
	AutomaticUpdatesScheduledInstallationDayAusidEveryFriday
	AutomaticUpdatesScheduledInstallationDayAusidEverySaturday
)
