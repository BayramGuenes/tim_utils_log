package tim_utils_log

type TimLoggerMicroservicesStruct struct {
	NameLogServer string
	PortLogServer string
}

type ExceptionStruct struct {
	Occured bool
	ErrTxt  string
}

/*type TimLogTransactPath struct {
	TransAppname string
	TransName    string
}*/
type TimLogTransactHeader struct {
	TransKey      string
	TransName     string
	TransAppName  string
	ClientAppName string
	UName         string
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
	CoTransStatusFinishedOk     = "OK"
	CoTransStatusFinishedFailed = "FAILED"
)

type InputParamStartTransact struct {
	TransKey            string
	TransName           string
	ServiceName         string
	TransAppName        string
	ClientAppName       string
	UName               string
	LogServerServiceAdr TimLoggerMicroservicesStruct
}
type InputParamCheckDisableLogMedia struct {
	LogServerServiceAdr TimLoggerMicroservicesStruct
	ErrCase             bool
}
type OutputParamDoTrace struct {
	DoTrace   bool
	Exception ExceptionStruct
}
type OutputParamDisableLogMedia struct {
	DisableFilesys bool
	DisableDB      bool
	Exception      ExceptionStruct
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
	ErrCase        bool
}

type InputParamLogStepResult struct {
	LogTransHeader TimLogTransactHeader
	StepName       string
	StepNum        int
	StepResult     string
	Context        string
	ErrCase        bool
}

type InputParamFinishTransact struct {
	LogTransHeader TimLogTransactHeader
	StepNum        int
	Status         string
	ErrCase        bool
}

type BufferedLogItem = struct {
	ItemType    string //"step","result"
	StepName    string
	StepContext string
	StepResult  string
}
type InputParamFailedToFilesys struct {
	LogTransHeader TimLogTransactHeader
	Items          []BufferedLogItem
}

type TimExecLogging interface {
	CheckDoTraceTransaction(iInput InputParamStartTransact) (eOutput OutputParamDoTrace)
	StartLogTransaction(iInput InputParamStartTransact) (eOutput OutputParamStartTransact)
	LogTransStep(iInput InputParamLogStep) (eException ExceptionStruct)
	LogTransStepResult(iInput InputParamLogStepResult) (eException ExceptionStruct)
	FinishLogTransaction(iInput InputParamFinishTransact) (eException ExceptionStruct)
}

type LoggerClassProxy struct {
}
type LoggerServerClass struct{}
