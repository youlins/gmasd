package coder

import (
	"fmt"
	"strconv"
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
	ST_DATA_END
	ST_CHECKSUM_0
	ST_CHECKSUM_1
	ST_CTRL_0
	ST_CTRL_1
	ST_PACKAGE_END
)

const (
	MAX_PROTOBUF_SIZE = 1024
)

type coder_mt_context struct{
	state 		int32
	imei 		string
	cmd 		string
	checksum 	int16
	flag 		byte
	argBuf 		[]byte
	data    	[]byte
	protoBuf	[MAX_PROTOBUF_SIZE]byte
	protoPad	int
	dataSize   	int
}

type coder_mt struct {
	 name 		string
	 ctx  		*coder_mt_context
}

func (c *coder_mt) init() {
	c.ctx = &coder_mt_context{state:ST_FIRST, flag:0x0}
}

func (c *coder_mt) Parse(buffer []byte) error {

	var i int
	var b byte
	var s string
	ctx := c.ctx

	for i = 0; i < len(buffer); i++{
		b = buffer[i]

		if ctx.protoPad > MAX_PROTOBUF_SIZE {
			return fmt.Errorf("overflow protocol") 
		}

		switch ctx.state {
			case ST_FIRST:
				if b != '$' {
					return fmt.Errorf("error first byte")
				}
				ctx.protoPad = 1;
				ctx.protoBuf[0] = b;
				ctx.dataSize = MAX_PROTOBUF_SIZE;
				ctx.state = ST_SECOND
			case ST_SECOND:
				if b != '$' {
					return fmt.Errorf("error second byte")
				}
				ctx.protoBuf[1] = b;
				ctx.state = ST_PACKAGE_FLAG
			case ST_PACKAGE_FLAG:
				if b < 0x41 || b > 0x7A {
					return fmt.Errorf("error package flat :%c", b)
				}
				ctx.flag = b;
				ctx.protoBuf[2] = b;
				ctx.argBuf = ctx.protoBuf[ctx.protoPad:ctx.protoPad]
				ctx.state = ST_LENGTH
			case ST_LENGTH:
				ctx.argBuf = append(ctx.argBuf, b)
				if b == ',' {
					ctx.argBuf = ctx.argBuf[:len(ctx.argBuf) - 1]
					ctx.state = ST_LENGTH_END

					s = string(ctx.argBuf)
					len, err := strconv.Atoi(s)
					if err != nil {
						return err
					}
					ctx.dataSize = len - 1;
					ctx.argBuf = ctx.protoBuf[ctx.protoPad:ctx.protoPad]
					ctx.state = ST_IMEI
				} else if b < '0' || b > '9' {
					return fmt.Errorf("error length :%c", b) 
				}
			case ST_IMEI:
				ctx.argBuf = append(ctx.argBuf, b)
				if b == ',' {
					ctx.state = ST_IMEI_END

					ctx.imei = string(ctx.argBuf[0 : len(ctx.argBuf)-1])
					//fmt.Printf("imei:%s\n", ctx.imei)
					ctx.dataSize -= len(ctx.argBuf)
					ctx.argBuf = ctx.protoBuf[ctx.protoPad:ctx.protoPad]
					ctx.state = ST_CMD
				}
			case ST_CMD:
				ctx.argBuf = append(ctx.argBuf, b)							
				if b == ',' {
					ctx.state = ST_CMD_END

					ctx.cmd = string(ctx.argBuf[0 : len(ctx.argBuf) - 1])
					ctx.dataSize -= len(ctx.argBuf) + 5
					ctx.argBuf = ctx.protoBuf[ctx.protoPad:ctx.protoPad + ctx.dataSize]
					ctx.state = ST_DATA
				}
			case ST_DATA:
				toCopy := ctx.dataSize
				if toCopy > len(buffer) - i {
					toCopy = len(buffer) - i
				}

				if toCopy > 0 {
					start := len(ctx.argBuf) - ctx.dataSize
					copy(ctx.argBuf[start:], buffer[i:i+toCopy])
					ctx.dataSize -= toCopy
					ctx.protoPad += toCopy - 1
					i += toCopy - 1
				} else {
					ctx.argBuf = append(ctx.argBuf, b)
				}

				if ctx.dataSize <= 0 {
					ctx.state = ST_DATA_END
				}
			case ST_DATA_END:
				if b != '*' {
					return fmt.Errorf("error data end, %c\n", b) 
				}
				ctx.data = ctx.argBuf
				ctx.argBuf = append(ctx.argBuf, b)
				ctx.state = ST_CHECKSUM_0
			case ST_CHECKSUM_0:
				ctx.argBuf = ctx.protoBuf[ctx.protoPad:ctx.protoPad]
				ctx.argBuf = append(ctx.argBuf, b)
				ctx.state = ST_CHECKSUM_1
			case ST_CHECKSUM_1:
				ctx.argBuf = append(ctx.argBuf, b)
				cs := getChecksum(ctx.protoBuf[:ctx.protoPad-2])
				if cs != string(ctx.argBuf) {
					return fmt.Errorf("error checksum :%s\n", cs) 
				}
				ctx.state = ST_CTRL_0
			case ST_CTRL_0:
				ctx.state = ST_CTRL_1
			case ST_CTRL_1:
				ctx.state = ST_PACKAGE_END
				ctx.argBuf = ctx.protoBuf[:ctx.protoPad]
				c.processPacket()
			default :{
				return fmt.Errorf("unkown protocol state") 
			}
		}
		ctx.protoPad++
	}

	return nil
}

func getChecksum(buf []byte) string {
	var i int
	s := byte(0)
	for i = 0; i < len(buf); i++ {
		s += buf[i]
	}
	return fmt.Sprintf("%X", s)
}

func (c *coder_mt) processPacket() {
	fmt.Printf("processPacket:%s\n", string(c.ctx.protoBuf[0:c.ctx.protoPad]))
}
