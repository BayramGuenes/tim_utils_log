package tim_utils_log

import (
	"encoding/json"

	timHTTP "github.com/BayramGuenes/tim_utils_http"
)

var LogServer TimLoggerMicroservicesStruct

func (lcp LoggerClassProxy) StartLogTransaction(iInput InputParamStartTransact) (eOutput OutputParamStartTransact) {
	eOutput = OutputParamStartTransact{}
	LogServer.NameLogServer = iInput.LogServerServiceAdr.NameLogServer
	LogServer.PortLogServer = iInput.LogServerServiceAdr.PortLogServer
	println("Start Transaction " + iInput.TransName + " {{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{ ")

	lData, err := json.Marshal(iInput)
	if err != nil {
		eOutput.Exception.Occured = true
		eOutput.Exception.ErrTxt = "json.Marshall.Input:" + err.Error()
		return
	}
	//println("LogServer.NameLogServer, LogServer.PortLogServer:"+LogServer.NameLogServer, LogServer.PortLogServer)
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

func (lcp LoggerClassProxy) LogTransStep(iInput InputParamLogStep) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	println(iInput.StepName + ":" + iInput.Context)
	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/LogTransactionStep", lData)
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
	println(iInput.StepName + ":" + iInput.StepResult)
	lData, err := json.Marshal(iInput)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	lResultArrByte, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/LogTransactionStepResult", lData)
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
	println("End Transaction  }}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}} ")
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

/*func (lc LoggerClass) NewTransactionKey(iName string) string {
	t := time.Now()

	//nanosec := t.UnixNano()
	//millisec := t.UnixNano() / int64(time.Millisecond)

	formatinfo := t.Format("20060102T150405")

	lUid := uid.New()
	//lKey := formatinfo + "-" + (strconv.FormatInt(nanosec, 10)) + "-" + iTransactionName
	//lKey := formatinfo + "-" + (strconv.FormatInt(millisec, 10)) + "-" + iTransactionName
	lKey := formatinfo + "-" + iName + "-" + lUid

	////println("lKey:"+lKey)
	return lKey
}*/
