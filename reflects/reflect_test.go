package reflects

import (
	"reflect"
	"testing"
)

type People struct {
	Name string
}

func TestNewByType(t *testing.T) {
	peopleSlice := make([]*People, 0)

	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "Ptr test",
			args: args{t: reflect.TypeOf(&People{})},
			want: reflect.TypeOf(&People{}),
		},
		{
			name: "Struct test",
			args: args{t: reflect.TypeOf(People{})},
			want: reflect.TypeOf(People{}),
		},
		{
			name: "Int test",
			args: args{t: reflect.TypeOf(int(1))},
			want: reflect.TypeOf(int(0)),
		},
		{
			name: "Bool test",
			args: args{t: reflect.TypeOf(bool(true))},
			want: reflect.TypeOf(bool(false)),
		},
		{
			name: "Slice test",
			args: args{t: reflect.TypeOf(make([]*People, 0))},
			want: reflect.TypeOf(make([]*People, 0)),
		},
		{
			name: "Slice Ptr test",
			args: args{t: reflect.TypeOf(&peopleSlice)},
			want: reflect.TypeOf(&peopleSlice),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewByType(tt.args.t); !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("NewByType() = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}
