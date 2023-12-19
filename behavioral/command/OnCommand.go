package command

type OnCommand struct {
	device IDevice
}

func (c *OnCommand) SetDevice(device IDevice) {
	c.device = device
}

func (c *OnCommand) execute() {
	c.device.On()
}
