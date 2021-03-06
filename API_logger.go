package tim_utils_log

var timLogger LoggerClassProxy

type UtilsLog struct {
	NameTimLogServer string
	PortTimLogServer string
	TransHeader      TimLogTransactHeader
	LoggingAppname   string
	ServiceName      string
	LogItemTab       []BufferedLogItem
	Exception        ExceptionStruct
}

func NewLoggerTr(iAppName, iTransName, iNameTimLogServer, iPortTimLogServer, iUName string) (eLog UtilsLog) {

	lInput := InputParamStartTransact{}
	lInput.TransName = iTransName
	lInput.TransAppName = iAppName
	lInput.ClientAppName = iAppName
	lInput.NameLogServer = iNameTimLogServer
	lInput.PortLogServer = iPortTimLogServer
	lInput.UName = iUName
	lInput.ServiceName = iTransName

	lOutput := timLogger.StartLogTransaction(lInput)
	eLog.TransHeader.TransAppName = iAppName
	eLog.TransHeader.ClientAppName = iAppName
	eLog.TransHeader.TransName = iTransName
	eLog.LoggingAppname = iAppName
	eLog.NameTimLogServer = iNameTimLogServer
	eLog.PortTimLogServer = iPortTimLogServer
	eLog.ServiceName = iTransName
	eLog.TransHeader.UName = iUName
	eLog.TransHeader.TransKey = lOutput.LogTrans.TransKey
	return
}

func NewLoggerSvc(iTranskey string, iTransname string, iAppTransact, iAppClient, iAppLogging, iServiceName, iNameTimLogServer, iPortTimLogServer, iUName string) (eLog UtilsLog) {
	lInput := InputParamStartTransact{}
	lInput.TransKey = iTranskey
	lInput.TransAppName = iAppTransact
	lInput.TransName = iTransname
	lInput.ClientAppName = iAppClient
	lInput.LoggingAppName = iAppLogging
	lInput.NameLogServer = iNameTimLogServer
	lInput.PortLogServer = iPortTimLogServer
	lInput.ServiceName = iServiceName
	lInput.UName = iUName

	lOutput := timLogger.StartLogService(lInput)
	eLog.TransHeader.TransAppName = iAppTransact
	eLog.TransHeader.ClientAppName = iAppClient
	eLog.TransHeader.TransName = lInput.TransName
	eLog.LoggingAppname = iAppLogging
	eLog.NameTimLogServer = iNameTimLogServer
	eLog.PortTimLogServer = iPortTimLogServer
	eLog.ServiceName = iServiceName
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
	lInputLogStep.AppLogging = ulog.LoggingAppname
	lInputLogStep.AppSVName = ulog.ServiceName

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
	lInputLogStepRes.AppLogging = ulog.LoggingAppname
	lInputLogStepRes.AppSVName = ulog.ServiceName

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
	lInputLogStepRes.AppLogging = ulog.LoggingAppname
	lInputLogStepRes.AppSVName = ulog.ServiceName

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

	lInputFinishTr := InputParamFinishService{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	lInputFinishTr.AppLogging = ulog.LoggingAppname
	lInputFinishTr.Status = CoTransStatusFinishedOk
	lInputFinishTr.ServiceName = ulog.ServiceName
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
	lInputFinishTr := InputParamFinishService{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	lInputFinishTr.Status = CoTransStatusFinishedFailed
	lInputFinishTr.ErrCase = true
	lInputFinishTr.AppLogging = ulog.LoggingAppname
	lInputFinishTr.ServiceName = ulog.ServiceName
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
