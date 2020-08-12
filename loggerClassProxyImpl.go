package tim_utils_log

import timHTTP "github.com/BayramGuenes/tim_utils_http"

var LogServer TimLoggerMicroservicesStruct

func (lcp LoggerClassProxy) StartLogTransaction(iInput InputParamStartTransact) (eOutput OutputParamStartTransact) {
	eOutput = OutputParamStartTransact{}

	println("Start Transaction " + iInput.TransName + " {{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{ ")

	lData := []byte{}
	_, _, _, lExcep := timHTTP.SendPostMsg(LogServer.NameLogServer, LogServer.PortLogServer, "/StartTransaction", lData)
	if lExcep.Occured {

	}
	return
}

func (lcp LoggerClassProxy) LogTransStep(iInput InputParamLogStep) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	println(iInput.StepName + ":" + iInput.Context)
	return
}

func (lcp LoggerClassProxy) LogTransStepResult(iInput InputParamLogStepResult) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	println(iInput.StepName + ":" + iInput.StepResult)
	return
}

func (lcp LoggerClassProxy) FinishLogTransaction(iInput InputParamFinishTransact) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	println("End Transaction  }}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}} ")

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
