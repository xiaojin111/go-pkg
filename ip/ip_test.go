package ip

import "testing"

func TestIPV4ToInt(t *testing.T) {
	type args struct {
		ipAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name:    "case 1",
			args:    args{"127.0.0.1"},
			want:    2130706433,
			wantErr: false,
		},
		{
			name:    "case 2",
			args:    args{"localhost"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "case 3",
			args:    args{"1.1.0"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "case 4",
			args:    args{"0.0.0.0"},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IPV4ToInt(tt.args.ipAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPV4ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IPV4ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToIPv4(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{2130706433},
			want: "127.0.0.1",
		},
		{
			name: "case 2",
			args: args{0},
			want: "0.0.0.0",
		},
		{
			name: "case 3",
			args: args{0xFFFFFFFF},
			want: "255.255.255.255",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToIPv4(tt.args.i); got != tt.want {
				t.Errorf("IntToIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}
