package access_point

import (
	"net"
	"sync/atomic"
	"fmt"
	"../coder"
)

type machine struct {
	conn   		net.Conn
	info 		string
	ap 			*access_point
	id 			uint64
}

func (m *machine) init() {

	m.id = atomic.AddUint64(&m.ap.incrClientId, 1)

	m.info = ""
	if ip, ok := m.conn.(*net.TCPConn); ok {
		addr := ip.RemoteAddr().(*net.TCPAddr)
		m.info = fmt.Sprintf("%s:%d", addr.IP, addr.Port)
	}

	m.info = fmt.Sprintf("%s - access client id:%d\n", m.info, m.id)

	go m.readLoop();
}

func (m* machine)readLoop() {

	if m.conn == nil {
		return
	}

	buffer := make([]byte, 1501)

	c  := coder.CreateCoder(m.ap.coder);

	for {
		n, err := m.conn.Read(buffer[0:1500])
		buffer[n] = 0;
		if err != nil {
			m.close()
			return
		}

		if err := c.Parse(buffer[:n]); err != nil {
				m.close()
			return
		}
	}
	m.close()
}

//func (m *machine) parse(buffer []byte) error {
//	fmt.Printf("size：%d, %s\n", len(buffer), string(buffer));
//	return nil
//}

func (c *machine) close() {
	c.conn.Close();
}
