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
	TransActApp  string
	TransActName string
}
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
	TransKey       string
	TransName      string
	ServiceName    string
	TransAppName   string
	ClientAppName  string
	LoggingAppName string
	UName          string
	NameLogServer  string
	PortLogServer  string
}

type OutputParamStartTransact struct {
	LogTrans  TimLogTransactHeader
	Exception ExceptionStruct
}

type InputParamLogStep struct {
	LogTransHeader TimLogTransactHeader
	StepName       string
	Context        string
	AppLogging     string
	AppSVName      string
}

type InputParamLogStepResult struct {
	LogTransHeader TimLogTransactHeader
	StepName       string
	StepResult     string
	Context        string
	AppLogging     string
	AppSVName      string
}

type InputParamFinishTransact struct {
	LogTransHeader TimLogTransactHeader
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
	StartLogTransaction(iInput InputParamStartTransact) (eOutput OutputParamStartTransact)
	LogTransStep(iInput InputParamLogStep) (eException ExceptionStruct)
	LogTransStepResult(iInput InputParamLogStepResult) (eException ExceptionStruct)
	FinishLogTransaction(iInput InputParamFinishTransact) (eException ExceptionStruct)
}

type LoggerClassProxy struct {
}
type LoggerServerClass struct{}

/*type InputParamCheckDisableLogMedia struct {
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
*/
