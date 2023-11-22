package size_test

import (
	"testing"

	"github.com/alessiodionisi/vmkit/size"
)

func TestToBytes(t *testing.T) {
	type args struct {
		size string
	}

	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "1GiB gibibytes",
			args: args{
				size: "1GiB",
			},
			want:    1073741824,
			wantErr: false,
		},
		{
			name: "1gib gibibytes lowercase",
			args: args{
				size: "1gib",
			},
			want:    1073741824,
			wantErr: false,
		},
		{
			name: "1 GiB gibibytes with space",
			args: args{
				size: "1 GiB",
			},
			want:    1073741824,
			wantErr: false,
		},
		{
			name: "1MiB mebibytes",
			args: args{
				size: "1MiB",
			},
			want:    1048576,
			wantErr: false,
		},
		{
			name: "1KiB kibibytes",
			args: args{
				size: "1KiB",
			},
			want:    1024,
			wantErr: false,
		},
		{
			name: "1GB gigabytes",
			args: args{
				size: "1GB",
			},
			want:    1000000000,
			wantErr: false,
		},
		{
			name: "1MB megabytes",
			args: args{
				size: "1MB",
			},
			want:    1000000,
			wantErr: false,
		},
		{
			name: "1KB kilobytes",
			args: args{
				size: "1KB",
			},
			want:    1000,
			wantErr: false,
		},
		{
			name: "1024 bytes",
			args: args{
				size: "1024",
			},
			want:    1024,
			wantErr: false,
		},
		{
			name: "KB missing number",
			args: args{
				size: "KB",
			},
			wantErr: true,
		},
		{
			name: "1UB unknown unit",
			args: args{
				size: "1UB",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := size.ToBytes(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUint(t *testing.T) {
	type args struct {
		size uint64
		unit string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1GiB gibibytes",
			args: args{
				size: 1073741824,
				unit: "gib",
			},
			want:    "1GiB",
			wantErr: false,
		},
		{
			name: "1MiB mebibytes",
			args: args{
				size: 1048576,
				unit: "mib",
			},
			want:    "1MiB",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := size.ParseUint(tt.args.size, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("ParseUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
