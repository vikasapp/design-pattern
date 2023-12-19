package command

func TestCommandPattern() {
	bulb := Bulb{}

	onCommand := OnCommand{}
	onCommand.SetDevice(&bulb)

	offCommand := OffCommand{}
	offCommand.SetDevice(&bulb)

	onRemote := Remote{}
	onRemote.SetCommand(&onCommand)
	onRemote.PressButton()

	offRemote := Remote{}
	offRemote.SetCommand(&offCommand)
	offRemote.PressButton()
}
