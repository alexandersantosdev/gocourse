package address_test

//package address_test // in tests files go will permit to name packages with _test

import (
	"fmt"
	"testing"

	. "github.com/alexandersantosdev/testsgo/address" //reference to principal package
)

// test archive must have _test.go extension
// go test -v ./...
// go test -v ./... --cover to show coverage
// go test -v ./... --coverprofile cover.txt to generate coverage report
// go tool cover -html=cover.txt to show coverage report in browser
// go tool cover -func=cover.txt to show coverage report in terminal
// The function name should start with 'Test'
// this test is a unity test

type TestScenario struct {
	input    string
	expected string
}

func TestAddressType(t *testing.T) {
	t.Parallel()

	var testsScenarios = []TestScenario{}
	for _, aType := range TypeAddress {
		testsScenarios = append(testsScenarios, TestScenario{
			input:    fmt.Sprintf("%s teste", aType),
			expected: aType,
		})
	}

	/* comment this 2 lines of scenarios above to test coverage report showing where is the problem */
	testsScenarios = append(testsScenarios, TestScenario{
		input:    "Casa 5",
		expected: "Tipo Inválido",
	})

	testsScenarios = append(testsScenarios, TestScenario{
		input:    "Casa",
		expected: "Tipo Inválido",
	})

	for _, scenario := range testsScenarios {
		typeReceived := AddressType(scenario.input)
		if typeReceived != scenario.expected {
			t.Errorf("Type received %s, expected %s", typeReceived, scenario.expected)
		}
	}

}

func AnotherTest(t *testing.T) {
	t.Parallel() // to run tests in parallel

	if 1 < 2 {
		t.Error("1 is less than 2")
	}
}
