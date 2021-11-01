package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_F(t *testing.T) {
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
			// want: float32(math.Inf(1)),
			want: 0,
		},
	}

	for _, tt := range tests {
		result, _ := F(tt.args.num)
		assert.Equal(result,
			tt.want,
			"they should be equal")
	}
}
