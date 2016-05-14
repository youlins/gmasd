package coder

import (
	"testing"
)

func dummyCoder() *coder_mt {
	c := &coder_mt{name:"coder_mt"}
	c.init();
	return c
}

func TestParser(t *testing.T) {
	
	msg := []byte("$$s136,861074020871516,AAA,35,36.653826,-4.641283,160509155353,A,9,10,0,21,0.9,224,8252377,4043610,214|1|CBA3|19D0,0000,0018|||02DE|010A,*B7\r\n")
	data := "35,36.653826,-4.641283,160509155353,A,9,10,0,21,0.9,224,8252377,4043610,214|1|CBA3|19D0,0000,0018|||02DE|010A,"

	c := dummyCoder()
	ctx := c.ctx
	err := c.Parse(msg)

	if  err != nil {
		t.Fatalf("Unexpected state %d : %v\n", c.ctx.state, err)
	}

	if ctx.imei != "861074020871516" {
		t.Fatalf("Unexpected imei %s, should be 861074020871516", ctx.imei)
	}

	if ctx.cmd != "AAA" {
		t.Fatalf("Unexpected cmd %s, should be AAA", ctx.cmd)
	}

	if string(ctx.data) != data {
		t.Fatalf("Unexpected data %s, should be %s", ctx.data, data)
	}
}

func TestChecksum(t *testing.T) {
	msg := []byte("$$s136,861074020871516,AAA,35,36.653826,-4.641283,160509155353,A,9,10,0,21,0.9,224,8252377,4043610,214|1|CBA3|19D0,0000,0018|||02DE|010A,*")

	if "B7" != getChecksum(msg) {
		t.Fatalf("Unexpected checksum %s, should be B7", getChecksum(msg))
	}
	
}
