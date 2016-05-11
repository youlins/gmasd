package access_point

import (
	"time"
)

const (

	//configuation for data center
	UDC_DEFAULT_PORT = 8888
	UDC_DEFAULT_HOST = "127.0.0.1"

	//configuation for web
	WEBSOCKET_DEFAULT_PORT = 8889
	WEB_SERVICE_DEFAULT_PORT = 8890

	//PROTOCOL
	MAX_NUMBER_OF_MACHINES = 1000000
	//
	
	//
	ACCEPT_MIN_SLEEP = 10 * time.Millisecond
	ACCEPT_MAX_SLEEP = 1 * time.Second

)