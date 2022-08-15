package main

import (
	"log"
	"os"
	"testing"
)

func Test_Contains(t *testing.T) {
	type Testdata struct {
		inp1     []string
		inp2     string
		expected bool
	}

	var tests = []Testdata{
		{[]string{"tom", "dick", "harry"}, "harry", true},
		{[]string{"tom", "dick", "harry"}, "harry potter", false},
	}

	for _, tst := range tests {
		output := contains(tst.inp1, tst.inp2)
		if output != tst.expected {
			t.Errorf("%t not equal to expected %t", output, tst.expected)
		}
	}
}

func Test_path_trawl(t *testing.T) {
	type Testdata struct {
		inp      string
		expected []string
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Default()
	}

	var tests = []Testdata{
		{
			(pwd + "/fixtures"),
			[]string{
				(pwd + "fixtures/__init__.py"),
				(pwd + "fixtures/foo.py"),
				(pwd + "fixtures/bar/__init__.py"),
				(pwd + "fixtures/bar/baz.py"),
			}},
	}

	for _, tst := range tests {
		output := path_trawl(tst.inp)

		for _, entry := range output {
			if contains(tst.expected, entry) {
				t.Errorf("%q not equal to expected %v", entry, tst.expected)
			}
		}
	}
}
