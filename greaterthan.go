package gocheckext

import . "launchpad.net/gocheck"
import "fmt"

// (c) Artem Andreenko <mio@volmy.com>, 2014

// -----------------------------------------------------------------------
// Greater Than checker.

type greaterThanChecker struct {
	*CheckerInfo
}

// The Greater Than checker verifies that the obtained value is greater
// than the expected value, according to usual Go semantics for >.
//
// For example:
//
//     c.Assert(value, GreaterThan, 42)
//
var GreaterThan Checker = &greaterThanChecker{
	&CheckerInfo{Name: "GreaterThan", Params: []string{"obtained", "expected greater than"}},
}

func (checker *greaterThanChecker) Check(params []interface{}, names []string) (result bool, error string) {
	defer func() {
		if v := recover(); v != nil {
			result = false
			error = fmt.Sprint(v)
		}
	}()
	return params[0].(int) > params[1].(int), ""
}