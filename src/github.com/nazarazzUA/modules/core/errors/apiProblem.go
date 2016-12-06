package errors

type ApiProblem struct {

	statusCode string
	statusMessage string
    mapMessages	map[string]string

}