package validator

import (
	"context"
	"fmt"
	"github.com/htenjom/json-validator/business"
	"strconv"
)

type InstallmentsByCountry struct {
	nextValidator         Validator
	installmentsByCountry map[string]int
}

func NewInstallmentsByCountry() InstallmentsByCountry {
	return InstallmentsByCountry{
		installmentsByCountry: map[string]int{
			"MLB": 100,
			"MLA": 200,
		},
	}
}

func (v *InstallmentsByCountry) Execute(ctx *context.Context, request business.DeviceRequest) {
	currentInstallments, _ := strconv.Atoi(request.Price)
	currentInstallments = currentInstallments / 10

	if v.installmentsByCountry["MLB"] > currentInstallments {
		fmt.Printf("::: Validating InstallmentsByCountry\n")
	}

	if v.nextValidator != nil {
		v.nextValidator.Execute(ctx, request)
	}
}

func (v *InstallmentsByCountry) NextInChain(validator Validator) {
	v.nextValidator = validator
}
