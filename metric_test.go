package schema

import (
	"testing"
)

func BenchmarkSetId(b *testing.B) {
	metric := MetricData{
		OrgId:    1234,
		Name:     "key1=val1.key2=val2.my.test.metric.name",
		Metric:   "my.test.metric.name",
		Interval: 15,
		Value:    0.1234,
		Unit:     "ms",
		Time:     1234567890,
		Mtype:    "gauge",
		Tags:     []string{"key1:val1", "key2:val2"},
	}
	for i := 0; i < b.N; i++ {
		metric.SetId()
	}
}

func TestTagValidation(t *testing.T) {
	type testCase struct {
		tag       []string
		expecting bool
	}

	testCases := []testCase{
		{[]string{"abc=cba"}, true},
		{[]string{"a="}, false},
		{[]string{"a!="}, false},
		{[]string{"=abc"}, false},
		{[]string{"@#$%!=(*&"}, false},
		{[]string{"!@#$%=(*&"}, false},
		{[]string{"@#;$%=(*&"}, false},
		{[]string{"@#$%=(;*&"}, false},
		{[]string{"@#$%=(*&"}, true},
	}

	for _, tc := range testCases {
		if validateTags(tc.tag) != tc.expecting {
			t.Fatalf("Testcase %s returned %t, but expected %t", tc.tag, !tc.expecting, tc.expecting)
		}
	}
}
