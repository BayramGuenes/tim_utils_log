package tim_utils_log

type TimLoggerMicroservicesStruct struct {
	NameServLogger string
	PortServLogger string
}

func (lcp LoggerClassProxy) StartLogTransaction(iApp, iTransName string) (eLogTrans TimLogTransactHeader, eException ExceptionStruct) {
	eLogTrans = TimLogTransactHeader{}
	eException = ExceptionStruct{}
	return
}

func (lcp LoggerClassProxy) LogTransStep(iLogTransHeader TimLogTransactHeader, iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	return
}

func (lcp LoggerClassProxy) LogTransStepResult(iLogTransHeader TimLogTransactHeader, iStepName string, iContext string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	return
}

func (lcp LoggerClassProxy) FinishLogTransaction(iLogTransHeader, iStatus string) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
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
