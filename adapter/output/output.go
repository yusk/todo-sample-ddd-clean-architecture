package output

type RenderOutputPort struct {
	Context map[string]interface{}
	Error   error
}

type RedirectOutputPort struct {
	Error error
}
