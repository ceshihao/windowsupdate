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
