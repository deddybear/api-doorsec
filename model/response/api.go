package response

type Api struct {
	Code   int
	Status string
	Data   interface{}
}

type Login struct {
	Code          int
	Status, Token string
	Data          interface{}
}
