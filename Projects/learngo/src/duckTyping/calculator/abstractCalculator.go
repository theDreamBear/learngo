package calculator

import "duckTyping/operator"

type Calculator interface {
	operator.Iadd
	operator.Isub
	operator.Imultiply
	operator.Idivide
}
