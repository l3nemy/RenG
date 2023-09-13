package file

import (
	"internal/compiler/core/code"
	"internal/compiler/core/object"
	"strconv"
)

func (f *File) WriteConstant(os []object.Object) error {
	err := f.WriteFileByte('C')
	if err != nil {
		return err
	}

	for _, o := range os {
		switch o.Type() {
		case object.INTEGER_OBJ:
			f.WriteFileBytes([]byte(strconv.Itoa(int(o.(*object.Integer).Value))))
		}
	}
	return f.WriteFileByte('E')
}

func (f *File) WriteInstruction(is code.Instructions, os []object.Object) error {
	err := f.WriteFileByte('B')
	if err != nil {
		return err
	}

	for ip := 0; ip < len(is); ip++ {
		op := code.Opcode(is[ip])

		switch op {
		case code.OpConstant:
			err = f.WriteFileByte(byte(op))
			if err != nil {
				return err
			}

			switch os[code.ReadUint32(is[ip+1:])].Type() {
			case object.INTEGER_OBJ:
				err = f.WriteFileByte(0x04)
				if err != nil {
					return err
				}

				err = f.WriteFileBytes(is[ip+1 : ip+5])
				if err != nil {
					return err
				}
			}
			ip += 4
		default:
			err = f.WriteFileByte(byte(op))
			if err != nil {
				return err
			}
		}
	}
	return f.WriteFileByte('E')
}

func (f *File) WriteFileByte(b byte) error {
	_, err := f.file.Write([]byte{b})
	return err
}

func (f *File) WriteFileBytes(bs []byte) error {
	_, err := f.file.Write(bs)
	return err
}
