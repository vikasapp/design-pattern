package chainofresponsibility

func TestChainOfResponsibility() {
	openFlap := OpenFlap{}
	openFlap.SetNextStep(&StartFlap{}).SetNextStep(&LoginFlap{})
	openFlap.Execute()
}
