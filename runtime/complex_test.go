// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grumpy

import (
	"math"
	"testing"
)

func TestComplexEq(t *testing.T) {
	cases := []invokeTestCase{
		{args: wrapArgs(complex(0, 0), 0), want: True.ToObject()},
		{args: wrapArgs(complex(1, 0), 0), want: False.ToObject()},
		{args: wrapArgs(complex(-12, 0), -12), want: True.ToObject()},
		{args: wrapArgs(complex(-12, 0), 1), want: False.ToObject()},
		{args: wrapArgs(complex(17.20, 0), 17.20), want: True.ToObject()},
		{args: wrapArgs(complex(1.2, 0), 17.20), want: False.ToObject()},
		{args: wrapArgs(complex(-4, 15), complex(-4, 15)), want: True.ToObject()},
		{args: wrapArgs(complex(-4, 15), complex(1, 2)), want: False.ToObject()},
		{args: wrapArgs(complex(math.Inf(1), 0), complex(math.Inf(1), 0)), want: True.ToObject()},
		{args: wrapArgs(complex(math.Inf(1), 0), complex(0, math.Inf(1))), want: False.ToObject()},
		{args: wrapArgs(complex(math.Inf(-1), 0), complex(math.Inf(-1), 0)), want: True.ToObject()},
		{args: wrapArgs(complex(math.Inf(-1), 0), complex(0, math.Inf(-1))), want: False.ToObject()},
		{args: wrapArgs(complex(math.Inf(1), math.Inf(1)), complex(math.Inf(1), math.Inf(1))), want: True.ToObject()},
	}
	for _, cas := range cases {
		if err := runInvokeTestCase(wrapFuncForTest(complexEq), &cas); err != "" {
			t.Error(err)
		}
	}
}

func TestComplexCompareNotSupported(t *testing.T) {
	cases := []invokeTestCase{
		{args: wrapArgs(complex(1, 2), 1), wantExc: mustCreateException(TypeErrorType, "no ordering relation is defined for complex numbers")},
		{args: wrapArgs(complex(1, 2), 1.2), wantExc: mustCreateException(TypeErrorType, "no ordering relation is defined for complex numbers")},
		{args: wrapArgs(complex(1, 2), math.NaN()), wantExc: mustCreateException(TypeErrorType, "no ordering relation is defined for complex numbers")},
		{args: wrapArgs(complex(1, 2), math.Inf(-1)), wantExc: mustCreateException(TypeErrorType, "no ordering relation is defined for complex numbers")},
		{args: wrapArgs(complex(1, 2), "abc"), want: NotImplemented},
	}
	for _, cas := range cases {
		if err := runInvokeTestCase(wrapFuncForTest(complexCompareNotSupported), &cas); err != "" {
			t.Error(err)
		}
	}
}

func TestComplexNE(t *testing.T) {
	cases := []invokeTestCase{
		{args: wrapArgs(complex(0, 0), 0), want: False.ToObject()},
		{args: wrapArgs(complex(1, 0), 0), want: True.ToObject()},
		{args: wrapArgs(complex(-12, 0), -12), want: False.ToObject()},
		{args: wrapArgs(complex(-12, 0), 1), want: True.ToObject()},
		{args: wrapArgs(complex(17.20, 0), 17.20), want: False.ToObject()},
		{args: wrapArgs(complex(1.2, 0), 17.20), want: True.ToObject()},
		{args: wrapArgs(complex(-4, 15), complex(-4, 15)), want: False.ToObject()},
		{args: wrapArgs(complex(-4, 15), complex(1, 2)), want: True.ToObject()},
		{args: wrapArgs(complex(math.Inf(1), 0), complex(math.Inf(1), 0)), want: False.ToObject()},
		{args: wrapArgs(complex(math.Inf(1), 0), complex(0, math.Inf(1))), want: True.ToObject()},
		{args: wrapArgs(complex(math.Inf(-1), 0), complex(math.Inf(-1), 0)), want: False.ToObject()},
		{args: wrapArgs(complex(math.Inf(-1), 0), complex(0, math.Inf(-1))), want: True.ToObject()},
		{args: wrapArgs(complex(math.Inf(1), math.Inf(1)), complex(math.Inf(1), math.Inf(1))), want: False.ToObject()},
	}
	for _, cas := range cases {
		if err := runInvokeTestCase(wrapFuncForTest(complexNE), &cas); err != "" {
			t.Error(err)
		}
	}
}

func TestComplexRepr(t *testing.T) {
	cases := []invokeTestCase{
		{args: wrapArgs(complex(0.0, 0.0)), want: NewStr("0j").ToObject()},
		{args: wrapArgs(complex(0.0, 1.0)), want: NewStr("1j").ToObject()},
		{args: wrapArgs(complex(1.0, 2.0)), want: NewStr("(1+2j)").ToObject()},
		{args: wrapArgs(complex(3.1, -4.2)), want: NewStr("(3.1-4.2j)").ToObject()},
	}
	for _, cas := range cases {
		if err := runInvokeTestCase(wrapFuncForTest(Repr), &cas); err != "" {
			t.Error(err)
		}
	}
}

func TestComplexHash(t *testing.T) {
	cases := []invokeTestCase{
		{args: wrapArgs(complex(0.0, 0.0)), want: NewInt(0).ToObject()},
		{args: wrapArgs(complex(0.0, 1.0)), want: NewInt(1000003).ToObject()},
		{args: wrapArgs(complex(1.0, 0.0)), want: NewInt(1).ToObject()},
		{args: wrapArgs(complex(3.1, -4.2)), want: NewInt(-1556830019620134).ToObject()},
		{args: wrapArgs(complex(3.1, 4.2)), want: NewInt(1557030815934348).ToObject()},
	}
	for _, cas := range cases {
		if err := runInvokeTestCase(wrapFuncForTest(complexHash), &cas); err != "" {
			t.Error(err)
		}
	}
}
