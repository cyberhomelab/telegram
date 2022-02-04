/*
   Copyright (c) 2022 Cyber Home Lab authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package telegram

import (
	"fmt"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestStringIsEmpty(t *testing.T) {
	if !StringIsEmpty("") {
		t.Fatal()
	}
	if StringIsEmpty("Hello") {
		t.Fail()
	}
}

func TestIntegerIsEmpty(t *testing.T) {
	if !IntegerIsEmpty(0) {
		t.Fatal()
	}
	if IntegerIsEmpty(1) {
		t.Fail()
	}
}

func TestListIsEmpty(t *testing.T) {
	if !ListIsEmpty([]string{}) {
		t.Fatal()
	}
	if ListIsEmpty([]string{"Hello"}) {
		t.Fail()
	}
}

func TestGetEnvVariableHappyPath(t *testing.T) {
	// Setting the env variable
	key := "CYBERHOMELAB_TEST_HAPPYPATH"
	value := "HAPPYPATH"
	os.Setenv(key, value)

	// Checks
	curentValue, err := GetEnvVariable(key)
	if err != nil {
		t.Fail()
	}
	if curentValue != value {
		t.Fail()
	}

	// Deleting the env variable
	os.Unsetenv(key)
}

func TestGetEnvVariableNegativePath(t *testing.T) {
	key := "CYBERHOMELAB_TEST_NEGATIVEPATH"
	_, err := GetEnvVariable(key)
	assert.ErrorContains(t, err, fmt.Sprintf("%s doesn't exists", key))
}

func TestIsStringInSlice(t *testing.T) {
	slice := []string{"Cyber", "Home", "Lab"}

	// Happy Path 1
	if !IsStringInSlice("Cyber", slice) {
		t.Fail()
	}

	// Happy Path 2
	if !IsStringInSlice("home", slice) {
		t.Fail()
	}

	// Negative Path
	if IsStringInSlice("cyberhomelab.com", slice) {
		t.Fail()
	}
}
