package query

//QueryResult (Generic Result for query)
type QueryResult struct {
	Result interface{}
	Error  error
}

//BhinnekanerQuery (Generic query for Bhinnekaner)
type BhinnekanerQuery interface {
	FindAll() <-chan QueryResult
}
