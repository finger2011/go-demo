package afterclass

import (
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	type fields struct {
		elements map[int]int
	}
	type args struct {
		ele int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{},
			},
			args: struct {
				ele int
			}{
				ele: 10,
			},
			wantErr: false,
		},
		{
			name: "test",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{
					10: 10,
				},
			},
			args: struct {
				ele int
			}{
				ele: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				elements: tt.fields.elements,
			}
			if err := s.Add(tt.args.ele); (err != nil) != tt.wantErr {
				t.Errorf("IntSet.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntSet_Delete(t *testing.T) {
	type fields struct {
		elements map[int]int
	}
	type args struct {
		ele int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{
					10: 10,
				},
			},
			args: struct {
				ele int
			}{
				ele: 10,
			},
			wantErr: false,
		},
		{
			name: "test",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{
					10: 10,
				},
			},
			args: struct {
				ele int
			}{
				ele: 9,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				elements: tt.fields.elements,
			}
			if err := s.Delete(tt.args.ele); (err != nil) != tt.wantErr {
				t.Errorf("IntSet.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntSet_Size(t *testing.T) {
	type fields struct {
		elements map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "test_iniset_size",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{
					1: 1,
					2: 2,
					3: 3,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := IntSet{
				elements: tt.fields.elements,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("IntSet.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Has(t *testing.T) {
	type fields struct {
		elements map[int]int
	}
	type args struct {
		ele int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{
					10: 10,
				},
			},
			args: struct {
				ele int
			}{
				ele: 10,
			},
			want: true,
		},
		{
			name: "test",
			fields: struct {
				elements map[int]int
			}{
				elements: map[int]int{
					10: 10,
				},
			},
			args: struct {
				ele int
			}{
				ele: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				elements: tt.fields.elements,
			}
			if got := s.Has(tt.args.ele); got != tt.want {
				t.Errorf("IntSet.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
