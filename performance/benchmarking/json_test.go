package main

import (
	"bytes"
	"io"
	"testing"
)

// go test --run=XXX --bench=BenchmarkJSON

func BenchmarkJSONUnmarshal(b *testing.B) {

	r := bytes.NewReader([]byte(j))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		r.Seek(0, io.SeekStart)
		b.StartTimer()

		_, err := jsonUnmarshal(r)
		if err != nil {
			b.Fatalf("unexpected error [%d]: [%s]\n", i, err)
		}
	}

}

func BenchmarkJSONDecode(b *testing.B) {

	r := bytes.NewReader([]byte(j))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		r.Seek(0, io.SeekStart)
		b.StartTimer()

		_, err := jsonDecode(r)
		if err != nil {
			b.Fatalf("unexpected error [%d]: [%s]\n", i, err)
		}
	}

}

var j = `[{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234},{"A":"something", "B": 1043234}]`
