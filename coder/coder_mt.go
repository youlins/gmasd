package coder

import (
	"fmt"
)

const (
	ST_FIRST = iota
	ST_SECOND
	ST_PACKAGE_FLAG
	ST_LENGTH
	ST_LENGTH_END
	ST_IMEI
	ST_IMEI_END
	ST_CMD
	ST_CMD_END
	ST_DATA
	ST_CHECKSUM
	ST_CTRL_0
	ST_CTRL_1
)

type coder_mt struct {
	 name string
	 state int32
}

func (c *coder_mt) init() {
	c.state = ST_FIRST
}

func (c *coder_mt) Parse(buffer []byte) error {

	var i int
	var b byte
	
	for i = 0; i < len(buffer); i++ {
		b = buffer[i]
		switch c.state {
			case ST_FIRST:
				if b != '@' {
					return fmt.Errorf("error first byte")
				}
				c.state = ST_SECOND
			case ST_SECOND:
				if b != '@' {
					return fmt.Errorf("error second byte")
				}
		}
		
	}

	fmt.Printf("data:%s\n", string(buffer));
	return nil
}

