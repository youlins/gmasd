package coder

type ICoder interface {

	Parse(buffer []byte) error
	init()

};


func CreateCoder(name string) ICoder {
	c := &coder_mt{name:name}
	c.init()
	return c;
}
