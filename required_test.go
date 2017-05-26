package required_test

import (
	"testing"
	"github.com/tiaguinho/required"
	"reflect"
)

var (
	i = required.Message{
		Field: "I",
		ErrMsg: "where is the value?",
	}

	s = required.Message{
		Field: "default",
		ErrMsg: "this field is required",
	}
)

type T struct {
	I    int     `required:"where is the number?"`
	S    string  `json:"default" required:"-"`
	A    []*A    `json:"array"`
	N    int     `json:"do_not_check"`
}

type A struct {
	I int    `required:"-"`
	S string `json:"s" required:"don't left this field blank!'"`
}

func TestValidate(t *testing.T) {
	v := T{}

	err := required.Validate(v)
	if  err == nil {
		t.Error("error expected! returned nil.")
	}

	sm := make([]required.Message, 2)
	sm[0] = i
	sm[1] = s

	e := required.New(sm...)
	if e == err {
		t.Errorf("\n expected: \n %s \n got: \n %s", e, err)
	}

	v.A = make([]*A, 1)
	v.A[0] = &A{
		I:50,
	}

	err = required.Validate(v)
	if err == nil {
		t.Errorf("error expected! returned nil.%s", err)
	}

	v.I = 100
	v.S = "ok"
	v.A[0].S = "sub message"

	err = required.Validate(v)
	if err != nil {
		t.Errorf("\n no error message expected \n got: \n %s", err)
	}
}

func TestValidateWithMessage(t *testing.T) {
	v := T{}

	err, msg := required.ValidateWithMessage(v)
	if err != nil {
		sm := make([]required.Message, 2)
		sm[0] = i
		sm[1] = s

		if reflect.DeepEqual(sm, msg) {
			t.Errorf("\n expected: \n %s \n got: \n %s", sm, msg)
		}
	}

	v.A = make([]*A, 1)
	v.A[0] = &A{
		I:50,
	}

	err = required.Validate(v)
	if err == nil {
		t.Error("error expected! returned nil.")
	}

	v.I = 100
	v.S = "ok"
	v.A[0].S = "sub message"

	err, msg = required.ValidateWithMessage(v)
	if err != nil || len(msg) > 0 {
		t.Errorf("\n no error message expected \n got: \n %s \n %+v", err, msg)
	}
}