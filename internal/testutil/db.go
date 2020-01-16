package testutil

import "context"

// MockCounterDAO fulfills the db.CounterDAO interface to mock DB calls
type MockCounterDAO struct {
	Int   int
	Error error
}

// Increment returns the Int and Error in the mock DAO
func (mock MockCounterDAO) Increment(context.Context) (int, error) {
	return mock.Int, mock.Error
}

// Value returns the Int and Error in the mock DAO
func (mock MockCounterDAO) Value(context.Context) (int, error) {
	return mock.Int, mock.Error
}
