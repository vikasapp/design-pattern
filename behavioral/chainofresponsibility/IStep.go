package chainofresponsibility

type IStep interface {
	Execute()
	SetNextStep(IStep) IStep
}
