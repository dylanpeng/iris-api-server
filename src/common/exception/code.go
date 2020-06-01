package exception

const (
	CodeInternalError = iota + 1
	CodeAccessDenied
	CodeQueryError
	CodeUnableConnect
	CodeUnauthConnection
	CodeCallServiceFailed
)

const (
	CodeInvalidParams = iota + 101
	CodeConvertError
	CodeDataNotExist
	CodeDataExist
	CodeDataCantCreate
	CodeCallGrpcFailed
	CodeAddJobFailed
	CodeNoPushRoute
	CodeAddKafkaFailed
	CodeServerInMaintain
	CodeOperateTooFast
	CodeRepeatedRequests
	CodeAntiFraudBlocked
	CodeFeatureDisable
)

const (
	CodeInvalidToken = iota + 201
	CodeAccountBlocked
	CodeVersionUnSupport
	CodePushFailed
	CodeUserOffline
	CodeInvalidPhoneNumber
	CodeNoServiceCity
	CodeNoServiceOrigin
	CodeNoServiceDest
)

var Desces = map[int64]string{
	CodeInternalError:     "server internal error",
	CodeAccessDenied:      "access denied",
	CodeQueryError:        "data query failed",
	CodeUnableConnect:     "unable to connect to server",
	CodeUnauthConnection:  "unauthenticated connection",
	CodeCallServiceFailed: "call service failed",

	CodeInvalidParams:    "invalid parameter",
	CodeConvertError:     "can't convert data",
	CodeDataNotExist:     "data not exist",
	CodeDataExist:        "data already exist",
	CodeDataCantCreate:   "can't create data",
	CodeCallGrpcFailed:   "call GRPC service failed",
	CodeAddJobFailed:     "add task failed",
	CodeNoPushRoute:      "user did't register a available route",
	CodeAddKafkaFailed:   "add kafka message failed",
	CodeServerInMaintain: "server is under maintenance",
	CodeOperateTooFast:   "operate too fast, please try again later",
	CodeRepeatedRequests: "repeated requests",
	CodeAntiFraudBlocked: "block by anti fraud",
	CodeFeatureDisable:   "current feature is disabled",

	CodeInvalidToken:       "invalid token",
	CodeAccountBlocked:     "this account has been disabled",
	CodeVersionUnSupport:   "Please download and install the upgraded version.",
	CodePushFailed:         "push message failed",
	CodeUserOffline:        "user is not online",
	CodeInvalidPhoneNumber: "the mobile phone number is invalid",
	CodeNoServiceCity:      "no service in current city",
	CodeNoServiceOrigin:    "no service in selected origin",
	CodeNoServiceDest:      "no service in selected destination",
}
