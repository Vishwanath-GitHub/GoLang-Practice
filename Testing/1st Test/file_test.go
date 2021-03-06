// +build unit
//This tag is used to tell GoLang compiler that we are performing unit testing here.

//The test-file name must end with '_test.go'.
//To run the test, use command: "go test -v"
//"-v" is for verbose output
//Test Code is compiled ONLY at test time.
//Any test-file with name ending with '_test.go' will not seen by 'go build', 'go run' and 'go install'.

//If we have multiple tests in one test file and we want to execute only a certain test function, use command: "go test -run TestFunction"
package main

import (
	"errors"
	"testing" //this package provides all the methods for testing purposes
)

//Testing function must be an exported function and must start with "Test" keyword.
//It must not return something. But it can have a receiver (turning test function into a method).
//It must take one argument of data-type (pointer) '*testing.T' where 'T' is a struct.
//T is a type passed to Test functions to manage test state and support formatted test logs.
func Test1(t *testing.T) {
	t.Log("passing the 1st test") //works similar to 'Println()' and records the text in the error log
	//It will be only displayed when '-v' is used with 'go test' or if the test fails.
	t.Logf("%v, %T","1st test passed","1st test passed") //works similar to 'Printf()' with string formatting
}

func Test2(t *testing.T) {
	//This test-case will fail since we received an error and it is not nil.
	if f1()!=nil {
		t.Fatal("received error") //'Fatal()' is similar to 'Log()' but uses function 'FailNow()' that reports that test-case is failed.
		//test-case will compulsorily fail when 't.Fatal()' is used and program will stop execution.
	}
}

func f1() error { //function returning an error
	var1:=errors.New("this is an error")
	return var1
}

//'FailNow()' calls 'Fail()' and 'Fail()' marks the test-case as failed.