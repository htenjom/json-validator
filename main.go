package main

import (
	"context"
	"github.com/htenjom/json-validator/business"
	"github.com/htenjom/json-validator/business/validator"
)

func main() {
	//dirPath, _ := os.Getwd()
	//web.Validate(dirPath)
	configureValidatorsChain()
}

func configureValidatorsChain() {
	ctx := context.Background()
	service := business.NewDeviceService("TEST")
	request := business.DeviceRequest{
		Id:      "1",
		Name:    "Device-1",
		Price:   "900",
		OwnerId: "UserX",
	}

	v1 := validator.NewDeviceOwnerByUser(service)
	v2 := validator.NewInstallmentsByCountry()
	v1.NextInChain(&v2)
	v1.Execute(&ctx, request)

}
