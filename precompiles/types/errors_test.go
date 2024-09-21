package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ErrInvalidAddr(t *testing.T) {
	e := ErrInvalidAddr{
		Got:    "foo",
		Reason: "bar",
	}
	got := e.Error()
	expect := "invalid address foo, reason: bar"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidAddr{"foo", "bar"}, e)
}

func Test_ErrInvalidNumberOfArgs(t *testing.T) {
	e := ErrInvalidNumberOfArgs{
		Got:    1,
		Expect: 2,
	}
	got := e.Error()
	expect := "invalid number of arguments; expected 2; got: 1"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidNumberOfArgs{1, 2}, e)
}

func Test_ErrInvalidArgument(t *testing.T) {
	e := ErrInvalidArgument{
		Got: "foo",
	}
	got := e.Error()
	expect := "invalid argument: foo"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidArgument{"foo"}, e)
}

func Test_ErrInvalidMethod(t *testing.T) {
	e := ErrInvalidMethod{
		Method: "foo",
	}
	got := e.Error()
	expect := "invalid method: foo"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidMethod{"foo"}, e)
}

func Test_ErrInvalidCoin(t *testing.T) {
	e := ErrInvalidCoin{
		Got:      "foo",
		Negative: true,
		Nil:      false,
		Empty:    false,
	}
	got := e.Error()
	expect := "invalid coin: denom: foo, is negative: true, is nil: false, is empty: false"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidCoin{"foo", true, false, false}, e)
}

func Test_ErrInvalidAmount(t *testing.T) {
	e := ErrInvalidAmount{
		Got: "foo",
	}
	got := e.Error()
	expect := "invalid token amount: foo"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidAmount{"foo"}, e)
}

func Test_ErrUnexpected(t *testing.T) {
	e := ErrUnexpected{
		When: "foo",
		Got:  "bar",
	}
	got := e.Error()
	expect := "unexpected error in foo: bar"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrUnexpected{"foo", "bar"}, e)
}

func Test_ErrInsufficientBalance(t *testing.T) {
	e := ErrInsufficientBalance{
		Requested: "foo",
		Got:       "bar",
	}
	got := e.Error()
	expect := "insufficient balance: requested foo, current bar"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInsufficientBalance{"foo", "bar"}, e)
}

func Test_ErrInvalidToken(t *testing.T) {
	e := ErrInvalidToken{
		Got:    "foo",
		Reason: "bar",
	}
	got := e.Error()
	expect := "invalid token foo: bar"
	if got != expect {
		t.Errorf("Expected %v, got %v", expect, got)
	}
	require.ErrorIs(t, ErrInvalidToken{"foo", "bar"}, e)
}
