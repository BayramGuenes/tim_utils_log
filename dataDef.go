package tim_utils_log

type TimLoggerMicroservicesStruct struct {
	NameLogServer string
	PortLogServer string
}

type ExceptionStruct struct {
	Occured bool
	ErrTxt  string
}

type TimLogTransactPath struct {
	AppName   string
	TransName string
	SubDomain string
	utime     string
}
type TimLogTransactHeader struct {
	TransKey  string
	AppName   string
	TransName string
	SubDomain string
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

type InputParamStartTransact struct {
	TimLogTransactPath
	LogServerServiceAdr TimLoggerMicroservicesStruct
}
type OutputParamStartTransact struct {
	LogTrans  TimLogTransactHeader
	Exception ExceptionStruct
}

type InputParamLogStep struct {
	LogTransHeader TimLogTransactHeader
	StepName       string
	StepNum        int
	Context        string
}

type InputParamLogStepResult struct {
	LogTransHeader TimLogTransactHeader
	StepName       string
	StepNum        int
	StepResult     string
}

type InputParamFinishTransact struct {
	LogTransHeader TimLogTransactHeader
	StepNum        int
	Status         string
}

type TimExecLogging interface {
	StartLogTransaction(iInput InputParamStartTransact) (eOutput OutputParamStartTransact)
	LogTransStep(iInput InputParamLogStep) (eException ExceptionStruct)
	LogTransStepResult(iInput InputParamLogStepResult) (eException ExceptionStruct)
	FinishLogTransaction(iInput InputParamFinishTransact) (eException ExceptionStruct)
}

type LoggerClassProxy struct {
}
type LoggerServerClass struct{}
