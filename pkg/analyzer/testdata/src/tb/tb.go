// Code generated by generator. DO NOT EDIT.

package tb

import (
	"context"
	"testing"
)

func nonTestHelper(tb int) {}

func helperWithoutHelper(tb testing.TB) {} // want "test helper function should start from tb.Helper()"

func helperWithHelper(tb testing.TB) {
	tb.Helper()
}

func helperWithEmptyStringBeforeHelper(tb testing.TB) {

	tb.Helper()
}

func helperWithHelperAfterAssignment(tb testing.TB) { // want "test helper function should start from tb.Helper()"
	_ = 0
	tb.Helper()
}

func helperWithHelperAfterOtherCall(tb testing.TB) { // want "test helper function should start from tb.Helper()"
	f()
	tb.Helper()
}

func helperWithHelperAfterOtherSelectionCall(tb testing.TB) { // want "test helper function should start from tb.Helper()"
	tb.Fail()
	tb.Helper()
}

func helperParamNotFirst(s string, i int, tb testing.TB) { // want "parameter testing.TB should be the first or after context.Context"
	tb.Helper()
}

func helperParamSecondWithoutContext(s string, tb testing.TB, i int) { // want "parameter testing.TB should be the first or after context.Context"
	tb.Helper()
}

func helperParamSecondWithContext(ctx context.Context, tb testing.TB) {
	tb.Helper()
}

func helperWithIncorrectName(o testing.TB) { // want "parameter testing.TB should have name tb"
	o.Helper()
}

func helperWithAnonymousHelper(tb testing.TB) {
	tb.Helper()
	func(tb testing.TB) {}(tb) // want "test helper function should start from tb.Helper()"
}

func helperWithNoName(_ testing.TB) {
}

func f() {}