package operator

type Iadd interface {
	Add(a interface{}, b interface{}) interface{}
}

type Isub interface {
	Sub(a interface{}, b interface{}) interface{}
}

type Imultiply interface {
	Mul(a interface{}, b interface{}) interface{}
}

type Idivide interface {
	Div(a interface{}, b interface{}) interface{}
}
