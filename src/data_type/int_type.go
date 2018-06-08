package data_type

type number struct {
}

func (a number)Int8s(num int8) int {
	    c := int(num)+int(num)
	  return  c
}

func NewInt() *number {
	return &number{}
}
