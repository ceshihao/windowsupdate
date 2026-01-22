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

// DeploymentAction defines the action for which an update is eligible.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-deploymentaction
const (
	DeploymentActionDaNone int32 = iota
	DeploymentActionDaDetection
	DeploymentActionDaInstallation
	DeploymentActionDaUninstallation
	DeploymentActionDaOptionalInstallation
)

// DownloadPriority defines the priority of a download.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-downloadpriority
const (
	DownloadPriorityDpLow    int32 = 1
	DownloadPriorityDpNormal int32 = 2
	DownloadPriorityDpHigh   int32 = 3
)

// InstallationImpact defines the impact of installing an update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-installationimpact
const (
	InstallationImpactIiNormal int32 = iota
	InstallationImpactIiMinor
	InstallationImpactIiRequiresExclusiveHandling
)

// InstallationRebootBehavior defines the restart behavior of an update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-installationrebootbehavior
const (
	InstallationRebootBehaviorIrbNeverReboots int32 = iota
	InstallationRebootBehaviorIrbAlwaysRequiresReboot
	InstallationRebootBehaviorIrbCanRequestReboot
)

// UpdateOperation defines the operation for which an update is being installed or uninstalled.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-updateoperation
const (
	UpdateOperationUoInstallation   int32 = 1
	UpdateOperationUoUninstallation int32 = 2
)

// UpdateExceptionContext defines the context in which an IUpdateException object can be provided.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-updateexceptioncontext
const (
	UpdateExceptionContextUecGeneral int32 = iota + 1
	UpdateExceptionContextUecWindowsDriver
	UpdateExceptionContextUecWindowsInstaller
	UpdateExceptionContextUecSearchIncomplete
)
