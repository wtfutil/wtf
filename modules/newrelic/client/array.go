package newrelic

// An Array is a type expected by the NewRelic API that differs from a comma-
// separated list. When passing GET params that expect an 'Array' type with
// one to many values, the expected format is "key=val1&key=val2" but an
// argument with zero to many values is of the form "key=val1,val2", and
// neither can be used in the other's place, so we have to differentiate
// somehow.
type Array struct {
	arr []string
}
