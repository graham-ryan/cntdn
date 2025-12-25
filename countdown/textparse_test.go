package countdown

import (
	"testing"
	"time"
)

func Test_parseTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{name: "1", args: args{s: "30s"}, want: time.Second * 30, wantErr: false},
		{name: "2", args: args{s: "1m30s"}, want: time.Second * 90, wantErr: false},
		{name: "3", args: args{s: "1m"}, want: time.Second * 60, wantErr: false},
		{name: "4", args: args{s: "0m5s"}, want: time.Second * 5, wantErr: false},
		{name: "5", args: args{s: "-10m5s"}, want: 0, wantErr: true},
		{name: "6", args: args{s: "5a"}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTime(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("parseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
