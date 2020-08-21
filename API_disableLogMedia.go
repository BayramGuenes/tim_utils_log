package tim_utils_log

func CheckDisableLogMedia(iNameTimLogServer, iPortTimLogServer string, iErrCase bool) (eOutput OutputParamDisableLogMedia) {
	eOutput = OutputParamDisableLogMedia{}
	lInput := InputParamCheckDisableLogMedia{}
	lInput.LogServerServiceAdr.NameLogServer = iNameTimLogServer
	lInput.LogServerServiceAdr.PortLogServer = iPortTimLogServer
	lInput.ErrCase = iErrCase
	eOutput = timLogger.CheckDisableLogMedia(lInput)
	return
}
