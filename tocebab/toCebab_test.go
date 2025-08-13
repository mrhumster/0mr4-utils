package tocebab

import "testing"

func TestConvert(t *testing.T) {
	data := []struct {
		name     string
		s        string
		expected string
	}{
		{
			"Example 1",
			"Average Salary Excluding the Minimum and Maximum Salary",
			"average-salary-excluding-the-minimum-and-maximum-salary",
		},
		{
			"Example 2",
			"Lemonade Change",
			"lemonade-change",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := convert(d.s)
			if resp != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` failed! Expected `%s`, but got `%s`!",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}
