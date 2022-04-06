package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	expectedResult := "Hello"
	receivedResult := Greeting()
	assert.Equal(t, expectedResult, receivedResult)
}

func TestGoodBye(t *testing.T) {
	expectedResult := "Goodbye, John..."
	receivedResult := GoodBye("John")
	assert.Equal(t, expectedResult, receivedResult)
}

func TestAdd(t *testing.T) {
	type testCase struct {
		a        int
		b        int
		expected int
	}

	testCases := []testCase{
		{a: 2, b: 2, expected: 4},
		{a: 1, b: 4, expected: 5},
	}

	for _, test := range testCases {
		expectedResult := test.expected
		receivedResult := Add(test.a, test.b)
		assert.Equal(t, expectedResult, receivedResult)

		fmt.Printf("case : a = %d, b = %d while expected result equals to %d, "+
			"received result equals to %d\n", test.a, test.b, test.expected, receivedResult)
	}
}

func TestAddWithTableDrivenTesting(t *testing.T) {
	type args struct {
		a int
		b int
	}

	type testCase struct {
		name string
		args args
		want int
	}

	tests := []testCase{
		{
			name: "should adding correctly a and b when they are equal",
			args: args{a: 2, b: 2},
			want: 4,
		},
		{
			name: "should adding correctly a and b when a is less than b",
			args: args{a: 1, b: 3},
			want: 4,
		},
		{
			name: "should adding correctly a and b when a is bigger than b",
			args: args{a: 5, b: 3},
			want: 8,
		},
	}

	for _, test := range tests {
		// Bir metodun birden fazla case'ini çalıştırmak için t.Run() kullanıyoruz.
		t.Run(test.name, func(t *testing.T) {
			assert.Equalf(t, test.want, Add(test.args.a, test.args.b), "Add(%v, %v)", test.args.a, test.args.b)
		})
	}
}

func TestGreetingAPI(t *testing.T) {
	// Initialize a new dummy http.Request.
	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	rec := httptest.NewRecorder()

	// Call the GreetingAPI handler function, passing in the
	// httptest.ResponseRecorder and http.Request.
	GreetingAPI(rec, req)

	expected := "Hello Gopher"
	received := rec.Body.String()

	// We can also use `assert.Equal(t, expected, actual)`
	if expected != received {
		t.Errorf("expected %s must equal to received %s", expected, received)
	}
}
