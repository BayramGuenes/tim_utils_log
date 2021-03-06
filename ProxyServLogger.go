package tim_utils_log

import (
	"encoding/json"

	timHTTP "github.com/BayramGuenes/tim_utils_http"
)

var LogServer TimLoggerMicroservicesStruct

func (lcp LoggerClassProxy) StartLogTransaction(iInput InputParamStartTransact) (eOutput OutputParamStartTransact) {
	eOutput = OutputParamStartTransact{}
	eOutput.LogTrans.TransAppName = iInput.TransAppName
	eOutput.LogTrans.ClientAppName = iInput.ClientAppName
	eOutput.LogTrans.TransName = iInput.TransName
	eOutput.LogTrans.UName = iInput.UName
	LogServer.NameLogServer = iInput.NameLogServer
	LogServer.PortLogServer = iInput.PortLogServer
	StdOut("Start Transaction " + iInput.TransName + " {{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{ ")
	lData, err := json.Marshal(iInput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Marshall.Input:" + err.Error()
		return
	}
	//log.Println("LogServer.NameLogServer, LogServer.PortLogServer:"+LogServer.NameLogServer, LogServer.PortLogServer)
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/StartTransaction", lData)
	if lExcep.Occured {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "StartTransaction SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eOutput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}
func (lcp LoggerClassProxy) StartLogService(iInput InputParamStartTransact) (eOutput OutputParamStartTransact) {

	eOutput = OutputParamStartTransact{}
	eOutput.LogTrans.TransAppName = iInput.TransAppName
	eOutput.LogTrans.ClientAppName = iInput.ClientAppName
	eOutput.LogTrans.TransName = iInput.TransName
	eOutput.LogTrans.UName = iInput.UName
	LogServer.NameLogServer = iInput.NameLogServer
	LogServer.PortLogServer = iInput.PortLogServer
	StdOut("Start Service  {{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{ ")
	lData, err := json.Marshal(iInput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Marshall.Input:" + err.Error()
		return
	}
	//log.Println("LogServer.NameLogServer, LogServer.PortLogServer:"+LogServer.NameLogServer, LogServer.PortLogServer)
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/StartService", lData)
	if lExcep.Occured {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "StartService SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eOutput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}

func (lcp LoggerClassProxy) LogTransStep(iInput InputParamLogStep) (eException ExceptionStruct) {

	eException = ExceptionStruct{}
	StdOut(iInput.StepName + ":" + iInput.Context)
	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/LogServiceStep", lData)
	if lExcep.Occured {
		eException.Occured = true
		eException.ErrTxt = "LogTransStep SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eException)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}

func (lcp LoggerClassProxy) LogTransStepResult(iInput InputParamLogStepResult) (eException ExceptionStruct) {

	eException = ExceptionStruct{}
	StdOut(iInput.StepName + ":" + iInput.StepResult)
	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/LogServiceStepResult", lData)
	if lExcep.Occured {
		eException.Occured = true
		eException.ErrTxt = "LogTransStepResult SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eException)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}

func (lcp LoggerClassProxy) FinishLogTransaction(iInput InputParamFinishTransact) (eException ExceptionStruct) {

	eException = ExceptionStruct{}
	StdOut("End Transaction  }}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}} ")
	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/FinishLogTransaction", lData)
	if lExcep.Occured {
		eException.Occured = true
		eException.ErrTxt = "FinishLogTransaction SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eException)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}

func (lcp LoggerClassProxy) FinishLogService(iInput InputParamFinishService) (eException ExceptionStruct) {

	eException = ExceptionStruct{}
	StdOut("End Service  }}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}} ")
	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/FinishLogService", lData)
	if lExcep.Occured {
		eException.Occured = true
		eException.ErrTxt = "FinishLogTransaction SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eException)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}

func (lcp LoggerClassProxy) LogEndFailedInFileSys(iInput InputParamFailedToFilesys) (eException ExceptionStruct) {

	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/LogFailedIntoFilesys", lData)
	if lExcep.Occured {
		eException.Occured = true
		eException.ErrTxt = "LogEndFailedInFileSys SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eException)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}

/*func (lcp LoggerClassProxy) CheckDoTraceTransaction(iInput InputParamStartTransact) (eOutput OutputParamDoTrace) {
	eOutput = OutputParamDoTrace{}
	LogServer.NameLogServer = iInput.LogServerServiceAdr.NameLogServer
	LogServer.PortLogServer = iInput.LogServerServiceAdr.PortLogServer

	lData, err := json.Marshal(iInput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Marshall.Input:" + err.Error()
		return
	}
	//log.log.Println("LogServer.NameLogServer, LogServer.PortLogServer:"+LogServer.NameLogServer, LogServer.PortLogServer)
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/CheckDoTraceTransaction", lData)
	if lExcep.Occured {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "CheckDoTraceTransaction SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eOutput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}
func (lcp LoggerClassProxy) CheckDisableLogMedia(iInput InputParamCheckDisableLogMedia) (eOutput OutputParamDisableLogMedia) {
	eOutput = OutputParamDisableLogMedia{}
	LogServer.NameLogServer = iInput.LogServerServiceAdr.NameLogServer
	LogServer.PortLogServer = iInput.LogServerServiceAdr.PortLogServer

	lData, err := json.Marshal(iInput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Marshall.Input:" + err.Error()
		return
	}
	//log.Println("LogServer.NameLogServer, LogServer.PortLogServer:"+LogServer.NameLogServer, LogServer.PortLogServer)
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/CheckDisableLogMedia", lData)
	if lExcep.Occured {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "CheckDisableLogMedia SendPostMessage to tim_serv_logger:" + err.Error()
		return
	}
	err = json.Unmarshal(lResultArrByte, &eOutput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Unmarshall.lResultArrByte:" + err.Error()
		return
	}
	return
}*/
