package validator

import (
	"context"
	"fmt"
	"github.com/htenjom/json-validator/business"
)

type DeviceOwnedByUser struct {
	nextValidator Validator
	service       business.DeviceService
}

func NewDeviceOwnerByUser(service business.DeviceService) DeviceOwnedByUser {
	return DeviceOwnedByUser{
		service: service,
	}
}

func (v *DeviceOwnedByUser) Execute(ctx *context.Context, request business.DeviceRequest) {
	if v.service.FindDeviceOwner(request.Id) == "TEST-User-"+request.Id {
		fmt.Printf("::: Validating DeviceOwnedByUser\n")
	}

	if v.nextValidator != nil {
		v.nextValidator.Execute(ctx, request)
	}
}

func (v *DeviceOwnedByUser) NextInChain(validator Validator) {
	v.nextValidator = validator
}
