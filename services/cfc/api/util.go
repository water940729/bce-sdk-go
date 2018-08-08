/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package api

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	RegularFunctionName = `^[a-zA-Z0-9-_:]+|\$LATEST$`
)

const (
	functionNameInvalid = "the function name of %s Must Match " + RegularFunctionName
	strLenIllegal       = "the length of %s is illegal, must in %d~%d"
	intLenIllegal       = "the size of %d is illegal, must in %d~%d"
)

const (
	memoryError = "memory size of %d must be a multiple of 64 MB"
)

// The limit of memory, in MB, your cfc function is given. cfc uses this memory size to
// infer the amount of CPU and memory allocated to your function. Your function use-case
// determines your CPU and memory requirements. For example, a database operation might
// need less memory compared to an image processing function. The default value is 128 MB.
// The value must be a multiple of 64 MB.
// Valid Range: Minimum value of 128. Maximum value of 3008.
const (
	minMemoryLimit = 128
	maxMemoryLimit = 3008
)

func getInvocationsUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s/invocations", functionName)
}

func getFunctionsUri() string {
	return fmt.Sprintf("/v1/functions")
}

func getFunctionUri(functionName string) string {
	return fmt.Sprintf("/v1/functions/%s", functionName)
}

func matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

func checkCreateFunctionArgs(args *FunctionArgs) error {
	if err := checkFunctionName(args.FunctionName); err != nil {
		return err
	} else if err = checkPtrString(args.Description, 0, 255); err != nil {
		return err
	} else if err = checkPtrIntSize(args.Timeout, 1, 300); err != nil {
		return err
	} else if err = checkMemorySize(args.MemorySize); err != nil {
		return err
	}
	return nil
}

func checkFunctionName(functionName string) error {
	if !matches(functionName, RegularFunctionName) {
		return fmt.Errorf(functionNameInvalid, functionName)
	}
	return nil
}

func checkPtrString(s string, min, max int) error {
	l := len(s)
	if l < min || l > max {
		errStr := fmt.Sprintf(strLenIllegal, s, min, max)
		return errors.New(errStr)
	}
	return nil
}

func checkPtrIntSize(s int, min, max int) error {
	if s < min || s > max {
		errStr := fmt.Sprintf(intLenIllegal, s, min, max)
		return errors.New(errStr)
	}
	return nil
}

func checkMemorySize(size int) error {
	if err := checkPtrIntSize(size, minMemoryLimit, maxMemoryLimit); err != nil {
		return err
	}
	if size%64 != 0 {
		return fmt.Errorf(memoryError, size)
	}
	return nil
}

func ReadBase64FromFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
