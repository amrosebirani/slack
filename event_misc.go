package slack


// ApiEvent is the main wrapper. You will find all the other messages attached
type ApiEvent struct {
	Type string
	Data interface{}
}
