package business

type ValidatorService interface {
	SomeClientValidatorMethod(request DeviceRequest, countryCode string)
}

type validatorServiceImpl struct {
}

func (service *validatorServiceImpl) SomeClientValidatorMethod(request DeviceRequest, countryCode string) {
	// 1. Map the request to the related entity
	// 2. Use a business validator to check the entity
	// error := DeviceRequestValidator(request)

}
