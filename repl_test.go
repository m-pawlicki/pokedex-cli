package main
import ("testing")


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "Bulbasaur Charmander SQUIRTLE",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("Lengths do not match -> Expected: %d Actual: %d", len(expected), len(actual))
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match -> Expected: %s Actual: %s", expectedWord, word)
				t.Fail()
			}
		}
	}
}
