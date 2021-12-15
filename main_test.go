package main

import (
  "io/ioutil"
	"testing"
)

type AddResults struct {
  x int
  y int
  expected int
}

var addResults = []AddResults {
  {1, 1, 2},
}

func TestAdd(t *testing.T) {
  for _, test := range addResults {
    result := Add (test.x, test.y)
    if result != test.expected {
      t.Fatal ("Expected Result Not Given")
    }
  }
}

func TestReadFile(t *testing.T) {
  data, err := ioutil.ReadFile ("testdata/test.data")
  if err != nil {
    t.Fatal ("Could Not Open File")
  }
  if string(data) != "hello world" {
    t.Fatal ("String contents do not match expected")
  }
}
