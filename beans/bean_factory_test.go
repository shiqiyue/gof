package beans

import (
	"fmt"
	"github.com/shiqiyue/inject"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type T1 struct {
	S string
}

type T2 struct {
	T1 *T1 `inject:""`
}

type T3 struct {
	T2 *T2 `inject:""`
	T1 *T1 `inject:""`
}

func TestGetObjectByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetObjectByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObjectByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetObjectByType(t *testing.T) {
	t1 := &T1{S: "str"}
	t2 := &T2{}
	t3 := &T3{}
	err := Provide(
		&inject.Object{Value: t1},
		&inject.Object{Value: t2},
		&inject.Object{Value: t3},
	)
	assert.Nil(t, err)
	err = Populate()
	assert.Nil(t, err)
	byType := GetObjectByType(t3)
	assert.NotNil(t, byType)
}

func TestPopulate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Populate(); (err != nil) != tt.wantErr {
				t.Errorf("Populate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvide(t *testing.T) {
	t1 := &T1{S: "str"}
	t2 := &T2{}
	err := Provide(
		&inject.Object{Value: t1},
		&inject.Object{Value: t2},
	)
	assert.Nil(t, err)
	err = Populate()
	assert.Nil(t, err)
	fmt.Println(t1.S)
	fmt.Println(t2.T1.S)
	t3 := &T3{}
	err = Provide(
		&inject.Object{Value: t3},
	)
	assert.Nil(t, err)
	err = Populate()
	assert.Nil(t, err)
	fmt.Println(t1.S)
	fmt.Println(t3.T2.T1.S)
	fmt.Println(t3.T1.S)
	fmt.Println(t2.T1.S)
}
