package errors

func RecoveryFromCrash(errResponse ErrorResponse, unknownErrorHandler func(data interface{})) {
	if recoveryData := recover(); recoveryData != nil {
		if err, ok := recoveryData.(error); ok {
			handled := false
			for err != nil {
				if e, ok := err.(ErrorResponseAggregator); ok {
					e.AppendData(errResponse)
					handled = true
				}
				if e, ok := err.(ErrorWithReason); ok {
					err = e.Reason()
					continue
				}
				break
			}
			if handled {
				return
			}
		}
		unknownErrorHandler(recoveryData)
	}
}
