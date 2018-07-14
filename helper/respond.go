package helper

type (
	Respond struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta,omitempty"`
	}

	ErrorRespond struct {
		Scope  string      `json:"scope"`
		Reason string      `json:"reason"`
		Detail interface{} `json:"detail,omitempty"`
	}
)

func Response(data interface{}, meta interface{}) Respond {
	return Respond{Data: data, Meta: meta}
}

func Abort(scope, message string, err interface{}) Respond {
	return Respond{
		Data: ErrorRespond{
			Scope:  scope,
			Reason: message,
			Detail: err,
		},
	}
}
