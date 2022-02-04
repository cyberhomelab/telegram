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
	"strings"
	"os"
	"fmt"
)

func StringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IntegerIsEmpty(i int64) bool {
	return i == 0
}

func ListIsEmpty(l []string) bool {
	return len(l) == 0
}

func IsStringInSlice(val string, slice []string) bool {
	for _, sliceVal := range slice {
		if strings.EqualFold(sliceVal, val) {
			return true
		}
	}
	return false
}

func GetEnvVariable(env string) (string, error) {
	envContent, ok := os.LookupEnv(env)
	if !ok {
		return "", fmt.Errorf("variable %s doesn't exists", env)
	}
	if StringIsEmpty(envContent) {
		return "", fmt.Errorf("variable %s is empty", env)
	}
	return envContent, nil
}
