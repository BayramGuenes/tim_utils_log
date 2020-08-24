package tim_utils_log

var timLogger LoggerClassProxy

type UtilsLog struct {
	NameTimLogServer string
	PortTimLogServer string
	TransHeader      TimLogTransactHeader
	LoggingAppname   string
	LogItemTab       []BufferedLogItem
	ErrCase          bool
}

func NewLoggerTr(iAppName, iTransName, iNameTimLogServer, iPortTimLogServer, iUName string) (eLog UtilsLog) {

	lInput := InputParamStartTransact{}
	lInput.TransAppName = iAppName
	lInput.ClientAppName = iAppName
	lInput.LogServerServiceAdr.NameLogServer = iNameTimLogServer
	lInput.LogServerServiceAdr.PortLogServer = iPortTimLogServer
	lInput.UName = iUName
	lInput.ServiceName = iTransName

	lOutput := timLogger.StartLogTransaction(lInput)
	eLog.TransHeader.TransAppName = iAppName
	eLog.TransHeader.ClientAppName = iAppName
	eLog.TransHeader.TransAppName = iTransName
	eLog.LoggingAppname = iAppName
	eLog.NameTimLogServer = iNameTimLogServer
	eLog.PortTimLogServer = iPortTimLogServer
	eLog.TransHeader.UName = iUName
	eLog.TransHeader.TransKey = lOutput.LogTrans.TransKey
	return
}

func NewLoggerSvc(iTransHeader TimLogTransactHeader, iAppName, iServiceName, iNameTimLogServer, iPortTimLogServer, iUName string) (eLog UtilsLog) {
	lInput := InputParamStartTransact{}
	lInput.TransKey = iTransHeader.TransKey
	lInput.TransAppName = iAppName
	lInput.ClientAppName = iAppName
	lInput.LogServerServiceAdr.NameLogServer = iNameTimLogServer
	lInput.LogServerServiceAdr.PortLogServer = iPortTimLogServer
	lInput.ServiceName = iServiceName
	lInput.UName = iUName
	lInput.TransName = iTransHeader.TransName

	lOutput := timLogger.StartLogService(lInput)
	eLog.TransHeader.TransAppName = iAppName
	eLog.TransHeader.ClientAppName = iAppName
	eLog.LoggingAppname = iAppName
	eLog.NameTimLogServer = iNameTimLogServer
	eLog.PortTimLogServer = iPortTimLogServer
	eLog.TransHeader.UName = iUName
	eLog.TransHeader.TransKey = lOutput.LogTrans.TransKey
	return

}

func (ulog *UtilsLog) LogStep(iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInputLogStep := InputParamLogStep{}
	lInputLogStep.StepName = iStepName
	lInputLogStep.LogTransHeader = ulog.TransHeader
	lInputLogStep.Context = iContext
	eException = timLogger.LogTransStep(lInputLogStep)
	logItemCache := BufferedLogItem{
		ItemType:    "step",
		StepName:    lInputLogStep.StepName,
		StepContext: lInputLogStep.Context,
		StepResult:  "",
	}
	ulog.LogItemTab = append(ulog.LogItemTab, logItemCache)

	return
}
func (ulog *UtilsLog) LogStepExecOK(iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	lInputLogStepRes := InputParamLogStepResult{}
	lInputLogStepRes.LogTransHeader = ulog.TransHeader
	lInputLogStepRes.StepName = iStepName
	lInputLogStepRes.Context = iContext
	lInputLogStepRes.StepResult = CoResultTypeOk
	eException = timLogger.LogTransStepResult(lInputLogStepRes)
	logItemCache := BufferedLogItem{
		ItemType:    "result",
		StepName:    lInputLogStepRes.StepName,
		StepContext: lInputLogStepRes.Context,
		StepResult:  CoResultTypeOk,
	}
	ulog.LogItemTab = append(ulog.LogItemTab, logItemCache)
	return eException
}
func (ulog *UtilsLog) LogStepExecErr(iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInputLogStepRes := InputParamLogStepResult{}
	lInputLogStepRes.LogTransHeader = ulog.TransHeader
	lInputLogStepRes.StepName = iStepName
	lInputLogStepRes.Context = iContext
	lInputLogStepRes.StepResult = CoResultTypeErr
	//lInputLogStepRes.ErrCase = true
	eException = timLogger.LogTransStepResult(lInputLogStepRes)
	logItemCache := BufferedLogItem{
		ItemType:    "result",
		StepName:    lInputLogStepRes.StepName,
		StepContext: lInputLogStepRes.Context,
		StepResult:  CoResultTypeErr,
	}
	ulog.LogItemTab = append(ulog.LogItemTab, logItemCache)

	return eException
}

func (ulog *UtilsLog) CloseLoggerTrStatOK() (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	lInputFinishTr := InputParamFinishTransact{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	lInputFinishTr.Status = CoTransStatusFinishedOk
	eException = timLogger.FinishLogTransaction(lInputFinishTr)
	ulog = &UtilsLog{}
	return eException
}
func (ulog *UtilsLog) CloseLoggerSvcStatOK() (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	lInputFinishTr := InputParamFinishTransact{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	lInputFinishTr.Status = CoTransStatusFinishedOk
	eException = timLogger.FinishLogService(lInputFinishTr)
	ulog = &UtilsLog{}
	return eException
}
func (ulog *UtilsLog) CloseLoggerTrStatFailed() (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInputFinishTr := InputParamFinishTransact{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	lInputFinishTr.Status = CoTransStatusFinishedFailed
	lInputFinishTr.ErrCase = true
	eException = timLogger.FinishLogTransaction(lInputFinishTr)
	eException = ulog.LogEndFailedInFileSys()
	ulog = &UtilsLog{}
	return eException
}
func (ulog *UtilsLog) CloseLoggerSvcStatFailed() (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInputFinishTr := InputParamFinishTransact{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	lInputFinishTr.Status = CoTransStatusFinishedFailed
	lInputFinishTr.ErrCase = true
	eException = timLogger.FinishLogService(lInputFinishTr)
	eException = ulog.LogEndFailedInFileSys()
	ulog = &UtilsLog{}
	return eException

}

func (ulog *UtilsLog) LogEndFailedInFileSys() (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInpParamFailedToFilesys := InputParamFailedToFilesys{}
	lInpParamFailedToFilesys.LogTransHeader = ulog.TransHeader
	lInpParamFailedToFilesys.Items = ulog.LogItemTab
	eException = timLogger.LogEndFailedInFileSys(lInpParamFailedToFilesys)

	return eException

}
