package loggerv2

type Env int

const (
	//env's
	PROD Env = 0
	DEV  Env = 1

	// key constants
	timeKey      = "TimeStamp"
	requestIDKey = "RequestID"
	appIDKey     = "AppId"
	userIDKey    = "UserID"
	uriKey       = "URI"
	ipKey        = "IP"

	//colours
	greenColor     = "\x1b[32m"
	redColor       = "\x1b[91m"
	yellowColor    = "\x1b[33m"
	blueColor      = "\x1b[34m"
	pinkColor      = "\x1b[91m"
	purpleColor    = "\x1b[35m"
	lightBlueColor = "\x1b[36m"
	defaultStyle   = "\x1b[0m"
	lightGrayColor = "\x1b[30m"
)
