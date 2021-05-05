package business

type DeviceService interface {
	FindDeviceOwner(deviceId string) string
}

type deviceServiceImpl struct {
	namePrefix string
}

func NewDeviceService(devicePrefix string) DeviceService {
	return &deviceServiceImpl{
		namePrefix: devicePrefix,
	}
}

func (service *deviceServiceImpl) FindDeviceOwner(deviceId string) string {
	return service.namePrefix + "-User-" + deviceId
}
