package command

type Remote struct {
	command ICommand
}

func (r *Remote) SetCommand(command ICommand) {
	r.command = command
}

func (r *Remote) PressButton() {
	r.command.execute()
}
