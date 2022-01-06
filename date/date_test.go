package date

import (
	"reflect"
	"testing"
	"time"
)

func TestDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		d       Date
		want    []byte
		wantErr bool
	}{
		{
			name:    "2016-08-15 should success",
			d:       Date(time.Date(2016, time.August, 15, 0, 0, 0, 0, time.Local)),
			want:    []byte(`"2016-08-15"`),
			wantErr: false,
		},
		{
			name:    "2016-08-15 at 23:00 should success",
			d:       Date(time.Date(2016, time.August, 15, 23, 0, 0, 0, time.Local)),
			want:    []byte(`"2016-08-15"`),
			wantErr: false,
		},
		{
			name:    "1980-12-31 should success",
			d:       Date(time.Date(1980, time.December, 31, 0, 0, 0, 0, time.Local)),
			want:    []byte(`"1980-12-31"`),
			wantErr: false,
		},
		{
			name:    "1950-01-01 should success",
			d:       Date(time.Date(1950, time.January, 1, 0, 0, 0, 0, time.Local)),
			want:    []byte(`"1950-01-01"`),
			wantErr: false,
		},
		{
			name:    "nil should return null",
			d:       Date{},
			want:    []byte("null"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		d       Date
		args    args
		wantErr bool
		want    Date
	}{
		{
			name: "2016-08-15 should success",
			d:    Date{},
			args: args{
				[]byte(`"2016-08-15"`),
			},
			wantErr: false,
			want:    Date(time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "1980-12-31 should success",
			d:    Date{},
			args: args{
				[]byte(`"1980-12-31"`),
			},
			wantErr: false,
			want:    Date(time.Date(1980, time.December, 31, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "1950-01-01 should success",
			d:    Date{},
			args: args{
				[]byte(`"1950-01-01"`),
			},
			wantErr: false,
			want:    Date(time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "RFC3339 format should error",
			d:    Date{},
			args: args{
				[]byte(`"1950-01-01T15:04:05Z07:00"`),
			},
			wantErr: true,
			want:    Date{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.d, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", tt.d.Format(time.RFC3339), tt.want.Format(time.RFC3339))
			}
		})
	}
}

func TestDate_Format(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		d    Date
		args args
		want string
	}{
		{
			name: "'02 Jan 06' format should success",
			d:    Date(time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)),
			args: args{
				"02 Jan 06",
			},
			want: "15 Aug 16",
		},
		{
			name: "'Monday, 02-Jan-2006' format should success",
			d:    Date(time.Date(1980, time.December, 31, 0, 0, 0, 0, time.UTC)),
			args: args{
				"Monday, 02-Jan-2006",
			},
			want: "Wednesday, 31-Dec-1980",
		},
		{
			name: "RFC3339 format should success",
			d:    Date(time.Date(1950, time.January, 1, 0, 0, 0, 0, time.UTC)),
			args: args{
				time.RFC3339,
			},
			want: "1950-01-01T00:00:00Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Format(tt.args.s); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
