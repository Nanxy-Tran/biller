package repositories

type RepositoryResult struct {
	Result interface{}
	Error  error
}

type ApiError struct {
	e string
}

func (apiError *ApiError) Error() string {
	return apiError.e
}
