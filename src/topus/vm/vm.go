package vm

var ()

type Opcode struct {
	instruction byte
}

func NewOpcode(instruction byte) *Opcode {
	o := &Opcode{}
	o.instruction = instruction
	return o
}

func (o *Opcode) String() string {

}

func (o *Opcode) Value() byte {
	return (o.instruction)
}
