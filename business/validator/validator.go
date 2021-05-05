package validator

import (
	"context"
)

type Validator interface {
	Execute(ctx *context.Context, model interface{})
	NextInChain(Validator)
}
