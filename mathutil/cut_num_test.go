package mathutil

import "testing"

func TestLastNumRange(t *testing.T) {
	type args struct {
		num  int64
		last int64
		len  int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"invalid: len 0", args{1234500890123, 3, 0}, -1, true},
		{"last 3, len 1", args{1234500890123, 3, 1}, 1, false},
		{"last 3, len 2", args{1234500890123, 3, 2}, 12, false},
		{"last 3, len 3", args{1234500890123, 3, 3}, 123, false},
		{"invalid: last 0", args{1234500890123, 0, 2}, -1, true},
		{"last 8, len 1, meet padding zero", args{1234500890123, 8, 1}, 0, false},
		{"last 8, len 2, meet padding zero", args{1234500890123, 8, 2}, 0, false},
		{"last 8, len 3, meet padding zero", args{1234500890123, 8, 3}, 8, false},
		{"last 8, len 4, meet padding zero", args{1234500890123, 8, 4}, 89, false},
		{"last long", args{1234500890123, 13, 4}, 1234, false},
		{"last longer", args{1234500890123, 14, 4}, 123, false},
		{"last too long", args{1234500890123, 50, 4}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastNumRange(tt.args.num, tt.args.last, tt.args.len)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastNumRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastNumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastNum(t *testing.T) {
	type args struct {
		num  int64
		last int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"Last 0: invalid", args{1234500890123, 0}, -1, true},
		{"Last -1: invalid", args{1234500890123, -1}, -1, true},
		{"Last 3", args{1234500890123, 3}, 123, false},
		{"Last 4", args{1234500890123, 4}, 123, false},
		{"Last 7", args{1234500890123, 7}, 890123, false},
		{"Last 8", args{1234500890123, 8}, 890123, false},
		{"Last 9", args{1234500890123, 9}, 500890123, false},
		{"Last 13: long", args{1234500890123, 50}, 1234500890123, false},
		{"Last 14: longer", args{1234500890123, 50}, 1234500890123, false},
		{"Last 50: too long", args{1234500890123, 50}, 1234500890123, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastNum(tt.args.num, tt.args.last)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow10(t *testing.T) {
	type args struct {
		p int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Pow -1", args{-1}, -1},
		{"Pow 0", args{0}, 1},
		{"Pow 1", args{1}, 10},
		{"Pow 1", args{2}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow10(tt.args.p); got != tt.want {
				t.Errorf("pow10() = %v, want %v", got, tt.want)
			}
		})
	}
}
