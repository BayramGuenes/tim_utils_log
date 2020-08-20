package tim_utils_log

var timLogger LoggerClassProxy

type BufferedLogItem = struct {
	ItemType    string //"step","result"
	StepName    string
	StepContext string
	StepResult  string
}
type UtilsLog struct {
	NameTimLogServer string
	PortTimLogServer string
	TransHeader      TimLogTransactHeader
	CurrentStepnum   int
	LogItemTab       []BufferedLogItem
	LogOnlyErr       bool
}

func NewLogger(iAppName, iSubdomain, iNameTimLogServer, iPortTimLogServer, iUName string, iOnlyErr bool) (eLog UtilsLog) {
	eLog.TransHeader.AppName = iAppName
	eLog.TransHeader.SubDomain = iSubdomain
	eLog.NameTimLogServer = iNameTimLogServer
	eLog.PortTimLogServer = iPortTimLogServer
	eLog.LogOnlyErr = iOnlyErr
	eLog.TransHeader.UName = iUName
	return
}

func (ulog *UtilsLog) LogStart(iTransname string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	ulog.TransHeader.TransName = iTransname

	lInputStartTr := InputParamStartTransact{}
	lInputStartTr.TimLogTransactPath.AppName = ulog.TransHeader.AppName
	lInputStartTr.TimLogTransactPath.SubDomain = ulog.TransHeader.SubDomain
	lInputStartTr.TimLogTransactPath.TransName = ulog.TransHeader.TransName
	lInputStartTr.LogServerServiceAdr.NameLogServer = ulog.NameTimLogServer
	lInputStartTr.LogServerServiceAdr.PortLogServer = ulog.PortTimLogServer
	lInputStartTr.UName = ulog.TransHeader.UName

	ulog.CurrentStepnum = 1
	if !ulog.LogOnlyErr {
		lOutput := timLogger.StartLogTransaction(lInputStartTr)
		ulog.TransHeader = lOutput.LogTrans
		eException = lOutput.Exception
	}
	return
}

func (ulog *UtilsLog) LogStep(iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInputLogStep := InputParamLogStep{}
	lInputLogStep.StepName = iStepName
	lInputLogStep.LogTransHeader = ulog.TransHeader
	lInputLogStep.Context = iContext
	ulog.CurrentStepnum++
	lInputLogStep.StepNum = ulog.CurrentStepnum
	if !ulog.LogOnlyErr {
		eException = timLogger.LogTransStep(lInputLogStep)
	} else {
		logItemCache := BufferedLogItem{
			ItemType:    "step",
			StepName:    lInputLogStep.StepName,
			StepContext: lInputLogStep.Context,
			StepResult:  "",
		}
		ulog.LogItemTab = append(ulog.LogItemTab, logItemCache)
	}
	return
}
func (ulog *UtilsLog) LogStepExecOK(iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	lInputLogStepRes := InputParamLogStepResult{}
	lInputLogStepRes.LogTransHeader = ulog.TransHeader
	lInputLogStepRes.StepName = iStepName
	lInputLogStepRes.Context = iContext
	ulog.CurrentStepnum++
	lInputLogStepRes.StepNum = ulog.CurrentStepnum
	lInputLogStepRes.StepResult = CoResultTypeOk
	if !ulog.LogOnlyErr {
		eException = timLogger.LogTransStepResult(lInputLogStepRes)
	} else {
		logItemCache := BufferedLogItem{
			ItemType:    "result",
			StepName:    lInputLogStepRes.StepName,
			StepContext: lInputLogStepRes.Context,
			StepResult:  CoResultTypeOk,
		}
		ulog.LogItemTab = append(ulog.LogItemTab, logItemCache)
	}
	return eException
}
func (ulog *UtilsLog) LogStepExecErr(iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	lInputLogStepRes := InputParamLogStepResult{}
	lInputLogStepRes.LogTransHeader = ulog.TransHeader
	lInputLogStepRes.StepName = iStepName
	lInputLogStepRes.Context = iContext
	ulog.CurrentStepnum++
	lInputLogStepRes.StepNum = ulog.CurrentStepnum
	lInputLogStepRes.StepResult = CoResultTypeErr
	if !ulog.LogOnlyErr {
		eException = timLogger.LogTransStepResult(lInputLogStepRes)
	} else {
		logItemCache := BufferedLogItem{
			ItemType:    "result",
			StepName:    lInputLogStepRes.StepName,
			StepContext: lInputLogStepRes.Context,
			StepResult:  CoResultTypeErr,
		}
		ulog.LogItemTab = append(ulog.LogItemTab, logItemCache)
	}
	return eException
}

func (ulog *UtilsLog) LogEndOK() (eException ExceptionStruct) {
	eException = ExceptionStruct{}

	lInputFinishTr := InputParamFinishTransact{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	ulog.CurrentStepnum++
	lInputFinishTr.StepNum = ulog.CurrentStepnum
	lInputFinishTr.Status = CoTransStatusFinishedOk
	if !ulog.LogOnlyErr {
		eException = timLogger.FinishLogTransaction(lInputFinishTr)
	} else {
		bufferlogger := NewLogger(ulog.TransHeader.AppName, ulog.TransHeader.SubDomain, ulog.NameTimLogServer, ulog.PortTimLogServer, ulog.TransHeader.UName, false)
		eException := bufferlogger.LogStart(ulog.TransHeader.TransName)
		for i := 0; i < len(ulog.LogItemTab); i++ {
			if !eException.Occured {
				logItem := ulog.LogItemTab[i]
				switch logItem.ItemType {
				case "step":
					eException = bufferlogger.LogStep(logItem.StepName, logItem.StepContext)
				case "result":
					switch logItem.StepResult {
					case CoResultTypeOk:
						eException = bufferlogger.LogStepExecOK(logItem.StepName, logItem.StepContext)
					case CoResultTypeErr:
						eException = bufferlogger.LogStepExecErr(logItem.StepName, logItem.StepContext)
					}
				}
			}
		}
		if !eException.Occured {
			eException = bufferlogger.LogEndOK()
		}
	}
	return eException
}
func (ulog *UtilsLog) LogEndFailed() (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	lInputFinishTr := InputParamFinishTransact{}
	lInputFinishTr.LogTransHeader = ulog.TransHeader
	ulog.CurrentStepnum++
	lInputFinishTr.StepNum = ulog.CurrentStepnum
	lInputFinishTr.Status = CoTransStatusFinishedFailed

	if !ulog.LogOnlyErr {
		eException = timLogger.FinishLogTransaction(lInputFinishTr)
	} else {
		bufferlogger := NewLogger(ulog.TransHeader.AppName, ulog.TransHeader.SubDomain, ulog.NameTimLogServer, ulog.PortTimLogServer, ulog.TransHeader.UName, false)
		eException := bufferlogger.LogStart(ulog.TransHeader.TransName)
		for i := 0; i < len(ulog.LogItemTab); i++ {
			if !eException.Occured {
				logItem := ulog.LogItemTab[i]
				switch logItem.ItemType {
				case "step":
					eException = bufferlogger.LogStep(logItem.StepName, logItem.StepContext)
				case "result":
					switch logItem.StepResult {
					case CoResultTypeOk:
						eException = bufferlogger.LogStepExecOK(logItem.StepName, logItem.StepContext)
					case CoResultTypeErr:
						eException = bufferlogger.LogStepExecErr(logItem.StepName, logItem.StepContext)
					}
				}
			}
		}
		if !eException.Occured {
			eException = bufferlogger.LogEndFailed()
		}
	}
	return eException

}
