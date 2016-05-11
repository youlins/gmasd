package access_point

import (
	"net"
	"fmt"
	"time"
)

type access_point struct {
	listener 		net.Listener
	machines  		map[uint64]*machine
	end       		chan bool
	running   		bool
	incrClientId  	uint64
	coder    		string
}

func New(coder string) *access_point {

	ap := &access_point{
		running : true,
		coder : coder,
	};

	ap.machines = make(map[uint64]*machine)

	return ap
}

func (ap *access_point) Listen(port uint16) bool {
	
	hp   := fmt.Sprintf("0.0.0.0:%d", port)
	l, e := net.Listen("tcp", hp)
	if e != nil {
		fmt.Printf("Error listening on %s, %q\n", hp, e)
		return false
	}
	
	ap.listener = l

	fmt.Printf("%s access_point is listening on %d\n", ap.coder, port)

	return true
}

func (ap *access_point) Start() {

	tmpDelay := ACCEPT_MIN_SLEEP
	l := ap.listener

	for ap.running {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				fmt.Printf("Temporary Client Accept Error(%v), sleeping %dms\n", ne, tmpDelay/time.Millisecond)
				time.Sleep(tmpDelay)
				tmpDelay *= 2				
				if tmpDelay > ACCEPT_MAX_SLEEP {
					tmpDelay = ACCEPT_MAX_SLEEP
				}
			} else if ap.running {
				fmt.Printf("Accept error: %v", err)
			}
			continue
		}
		tmpDelay = ACCEPT_MIN_SLEEP
		ap.createClient(conn)
	}

	ap.end <- true
}

func (ap *access_point) createClient(conn net.Conn) {
	machine := &machine{ap:ap, conn:conn}
	machine.init();	
}