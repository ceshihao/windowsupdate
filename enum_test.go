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

import "testing"

func TestOperationResultCode(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"OrcNotStarted", OperationResultCodeOrcNotStarted, 0},
		{"OrcInProgress", OperationResultCodeOrcInProgress, 1},
		{"OrcSucceeded", OperationResultCodeOrcSucceeded, 2},
		{"OrcSucceededWithErrors", OperationResultCodeOrcSucceededWithErrors, 3},
		{"OrcFailed", OperationResultCodeOrcFailed, 4},
		{"OrcAborted", OperationResultCodeOrcAborted, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("OperationResultCode %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestDeploymentAction(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"DaNone", DeploymentActionDaNone, 0},
		{"DaDetection", DeploymentActionDaDetection, 1},
		{"DaInstallation", DeploymentActionDaInstallation, 2},
		{"DaUninstallation", DeploymentActionDaUninstallation, 3},
		{"DaOptionalInstallation", DeploymentActionDaOptionalInstallation, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("DeploymentAction %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestDownloadPriority(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"DpLow", DownloadPriorityDpLow, 1},
		{"DpNormal", DownloadPriorityDpNormal, 2},
		{"DpHigh", DownloadPriorityDpHigh, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("DownloadPriority %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestInstallationImpact(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"IiNormal", InstallationImpactIiNormal, 0},
		{"IiMinor", InstallationImpactIiMinor, 1},
		{"IiRequiresExclusiveHandling", InstallationImpactIiRequiresExclusiveHandling, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("InstallationImpact %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestInstallationRebootBehavior(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"IrbNeverReboots", InstallationRebootBehaviorIrbNeverReboots, 0},
		{"IrbAlwaysRequiresReboot", InstallationRebootBehaviorIrbAlwaysRequiresReboot, 1},
		{"IrbCanRequestReboot", InstallationRebootBehaviorIrbCanRequestReboot, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("InstallationRebootBehavior %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestUpdateOperation(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"UoInstallation", UpdateOperationUoInstallation, 1},
		{"UoUninstallation", UpdateOperationUoUninstallation, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("UpdateOperation %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestUpdateExceptionContext(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"UecGeneral", UpdateExceptionContextUecGeneral, 1},
		{"UecWindowsDriver", UpdateExceptionContextUecWindowsDriver, 2},
		{"UecWindowsInstaller", UpdateExceptionContextUecWindowsInstaller, 3},
		{"UecSearchIncomplete", UpdateExceptionContextUecSearchIncomplete, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("UpdateExceptionContext %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestServerSelection(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"SsDefault", ServerSelectionSsDefault, 0},
		{"SsManagedServer", ServerSelectionSsManagedServer, 1},
		{"SsWindowsUpdate", ServerSelectionSsWindowsUpdate, 2},
		{"SsOthers", ServerSelectionSsOthers, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("ServerSelection %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestAutomaticUpdatesNotificationLevel(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"AunlNotConfigured", AutomaticUpdatesNotificationLevelAunlNotConfigured, 0},
		{"AunlDisabled", AutomaticUpdatesNotificationLevelAunlDisabled, 1},
		{"AunlNotifyBeforeDownload", AutomaticUpdatesNotificationLevelAunlNotifyBeforeDownload, 2},
		{"AunlNotifyBeforeInstallation", AutomaticUpdatesNotificationLevelAunlNotifyBeforeInstallation, 3},
		{"AunlScheduledInstallation", AutomaticUpdatesNotificationLevelAunlScheduledInstallation, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("AutomaticUpdatesNotificationLevel %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestAutomaticUpdatesScheduledInstallationDay(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"AuisdEveryDay", AutomaticUpdatesScheduledInstallationDayAuisdEveryDay, 0},
		{"AuisdEverySunday", AutomaticUpdatesScheduledInstallationDayAuisdEverySunday, 1},
		{"AuisdEveryMonday", AutomaticUpdatesScheduledInstallationDayAuisdEveryMonday, 2},
		{"AuisdEveryTuesday", AutomaticUpdatesScheduledInstallationDayAuisdEveryTuesday, 3},
		{"AuisdEveryWednesday", AutomaticUpdatesScheduledInstallationDayAuisdEveryWednesday, 4},
		{"AuisdEveryThursday", AutomaticUpdatesScheduledInstallationDayAuisdEveryThursday, 5},
		{"AuisdEveryFriday", AutomaticUpdatesScheduledInstallationDayAuisdEveryFriday, 6},
		{"AuisdEverySaturday", AutomaticUpdatesScheduledInstallationDayAuisdEverySaturday, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("AutomaticUpdatesScheduledInstallationDay %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestDownloadPhase(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"Initializing", DownloadPhaseInitializing, 1},
		{"Downloading", DownloadPhaseDownloading, 2},
		{"Verifying", DownloadPhaseVerifying, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("DownloadPhase %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestAutoDownloadMode(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"ForbidAutoDownload", AutoDownloadModeForbidAutoDownload, 0},
		{"AllowAutoDownload", AutoDownloadModeAllowAutoDownload, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("AutoDownloadMode %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestAutoSelectionMode(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"LetWindowsUpdateDecide", AutoSelectionModeLetWindowsUpdateDecide, 0},
		{"AutoSelectIfDownloaded", AutoSelectionModeAutoSelectIfDownloaded, 1},
		{"NeverAutoSelect", AutoSelectionModeNeverAutoSelect, 2},
		{"AlwaysAutoSelect", AutoSelectionModeAlwaysAutoSelect, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("AutoSelectionMode %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestUpdateServiceRegistrationState(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"NotRegistered", UpdateServiceRegistrationStateNotRegistered, 1},
		{"RegistrationPending", UpdateServiceRegistrationStateRegistrationPending, 2},
		{"Registered", UpdateServiceRegistrationStateRegistered, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("UpdateServiceRegistrationState %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestAddServiceFlag(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"AsfAllowPendingRegistration", AddServiceFlagAsfAllowPendingRegistration, 1},
		{"AsfAllowOnlineRegistration", AddServiceFlagAsfAllowOnlineRegistration, 2},
		{"AsfRegisterServiceWithAU", AddServiceFlagAsfRegisterServiceWithAU, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("AddServiceFlag %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestUpdateType(t *testing.T) {
	tests := []struct {
		name     string
		constant int32
		expected int32
	}{
		{"Software", UpdateTypeSoftware, 1},
		{"Driver", UpdateTypeDriver, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("UpdateType %s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}
