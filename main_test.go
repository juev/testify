package main

import (
	"fmt"
	"math"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func Test_f(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "10",
			args: args{num: 10},
			want: 0.2,
		},
		{
			name: "20",
			args: args{num: 20},
			want: 0.06666667,
		},
		{
			name: "5",
			args: args{num: 5},
			want: float32(math.Inf(1)),
		},
	}

	for _, tt := range tests {
		assert.Equal(f(tt.args.num),
			tt.want,
			"they should be equal")
	}
}

func Test_fuzz(t *testing.T) {
	fu := fuzz.New()
	for i := 0; i < 100000; i++ {
		var i int
		fu.Fuzz(&i)
		result := f(i)
		fmt.Printf("result(i): %v", result)
	}
}
