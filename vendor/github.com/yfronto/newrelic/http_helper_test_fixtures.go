package newrelic

import (
	"errors"
	"net/http"
	"strings"
)

type testJSONInterface struct {
	data string `json:"data"id,omitempty"`
}

type testParamsInterface struct {
	data string
}

type doGetTestsInput struct {
	path   string
	params *testParamsInterface
	out    testJSONInterface
	status int
	data   string
}

type doGetTestsOutput struct {
	data testJSONInterface
	err  error
}

type doRequestTestsInput struct {
	req    *http.Request
	out    testJSONInterface
	status int
	data   string
}

type doRequestTestOutput struct {
	data testJSONInterface
	err  error
}

var (
	doGetTests = []struct {
		in  doGetTestsInput
		out doGetTestsOutput
	}{
		{
			doGetTestsInput{
				path:   "somePath",
				params: &testParamsInterface{"testData"},
				out:    testJSONInterface{"testData"},
				status: 200,
				data:   `{"data": "testStructData"}`,
			},
			doGetTestsOutput{
				data: testJSONInterface{"testData"},
				err:  nil,
			},
		},
		{
			doGetTestsInput{
				status: 404,
				data:   "Not Found",
			},
			doGetTestsOutput{
				err: errors.New("newrelic http error (404 Not Found): Not Found"),
			},
		},
	}
	testRequest, _ = http.NewRequest("GET", "http://testPath",
		strings.NewReader("testBody"))
	doRequestTests = []struct {
		in  doRequestTestsInput
		out doRequestTestOutput
	}{
		{
			doRequestTestsInput{
				req:    testRequest,
				out:    testJSONInterface{"testData"},
				status: 200,
				data:   `{"data": "testStructData"}`,
			},
			doRequestTestOutput{
				data: testJSONInterface{"testData"},
				err:  nil,
			},
		},
	}
	encodeGetParamsTests = []struct {
		in  map[string]interface{}
		out string
	}{
		{
			map[string]interface{}{
				"testInt":         5,
				"testString":      "test",
				"testStringSlice": []string{"test1", "test2"},
				"testArray":       Array{[]string{"test1", "test2"}},
				"testTime":        testTime,
			},
			"testArray=test1" +
				"&testArray=test2" +
				"&testInt=5" +
				"&testString=test" +
				"&testStringSlice=test1%2Ctest2" +
				"&testTime=" + testTimeStringEscaped,
		},
		{
			map[string]interface{}{
				"unexpectedType": map[string]string{"unexpected": "type"},
			},
			"unexpectedType=map%5Bunexpected%3Atype%5D",
		},
	}
)

func (m *testParamsInterface) String() string {
	return "data=testData"
}
