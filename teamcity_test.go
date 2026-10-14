package main

import (
	"strings"
	"testing"
)

func TestFormatTestOutputUsesParallelTestNames(t *testing.T) {
	test := NewTeamCityTest("TestAccOne")
	test.ParseTestRunnerOutput("=== RUN   TestAccOne\n--- FAIL: TestAccOne (0.01s)\nFAIL\n", "")

	if out := test.FormatTestOutput(); strings.Contains(out, "TestAccOne.run") {
		t.Errorf("names should be unchanged without -parallel-tests:\n%s", out)
	}

	test.ParallelTestNames = true
	out := test.FormatTestOutput()
	for _, message := range []string{"testStarted", "testStdOut", "testStdErr", "testFailed", "testFinished"} {
		if !strings.Contains(out, "##teamcity["+message) || strings.Contains(out, "name='TestAccOne'") {
			t.Errorf("%s should report TestAccOne.run:\n%s", message, out)
		}
	}
	if strings.Count(out, "name='TestAccOne.run'") != 5 {
		t.Errorf("every message should use TestAccOne.run:\n%s", out)
	}
}
