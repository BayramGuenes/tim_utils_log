package tim_utils_log

type ExceptionStruct struct {
	Occured bool
	ErrTxt  string
}

type TimLogTransactHeader struct {
	TransKey  string
	AppName   string
	TransName string
}
type TimLogTransactItem struct {
	LogStep       string
	LogResultType string //Info, Error,Fatal
	LogTxt        string
	Context       string
}
type TimTransactionLog struct {
	TransHeader    TimLogTransactHeader
	TransExecItems []TimLogTransactItem
	TransStatus    string
}

const (
	CoResultTypeErr   = "Error"
	CoResultTypeOk    = "Ok"
	CoResultTypeFatal = "Fatal"
)
const (
	CoTransStatusRunning        = "running"
	CoTransStatusFinishedOk     = "succeeded"
	CoTransStatusFinishedFailed = "failed"
)

type TimExecLogging interface {
	StartLogTransaction(iApp, iTransName string) (eLogTrans TimLogTransactHeader, eException ExceptionStruct)
	LogTransStep(iLogTransHeader TimLogTransactHeader, iStepName string, iContext string) (eException ExceptionStruct)
	LogTransStepResult(iLogTransHeader TimLogTransactHeader, iStepResult string) (eException ExceptionStruct)
	FinishLogTransaction(iLogTransHeader, iStatus string) (eException ExceptionStruct)
}

type LoggerClassProxy struct {
}

var timLoggerViaServLogger LoggerClassProxy
