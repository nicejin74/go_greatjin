package go_greatjin

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		one int
		two int
	}
	test_data := []struct {
		title    string
		args     args
		expected int
	}{
		{
			title: "Testin Add Method",
			args: args{
				one: 1,
				two: 2,
			},
			expected: 3,
		},
	}
	for _, daum := range test_data {
		t.Run(daum.title, func(t *testing.T) {
			if answer := Add(daum.args.one, daum.args.two); answer != daum.expected {
				t.Errorf("Testing %v, expected %v", answer, daum.expected)
			}
		})
	}
}
