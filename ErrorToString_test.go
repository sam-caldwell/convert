package convert

import (
	"fmt"
	"testing"
)

func TestErrorToString(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "With error",
			err:  fmt.Errorf("test error"),
			want: "test error",
		},
		{
			name: "Without error",
			err:  nil,
			want: "",
		},
	}

	for _, tt := range tests {
		if got := ErrorToString(tt.err); tt.want != got {
			t.Fatalf("test failed: %s\n"+
				"\t got:%v\n"+
				"\twant:%v\n",
				tt.name, got, tt.want)
		}
	}
}
