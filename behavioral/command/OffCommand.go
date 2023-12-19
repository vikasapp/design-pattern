package command

type OffCommand struct {
	device IDevice
}

func (c *OffCommand) SetDevice(device IDevice) {
	c.device = device
}

func (c *OffCommand) execute() {
	c.device.Off()
}
