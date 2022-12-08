package validates

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required" vName:"姓"`
	Age            uint8      `validate:"gte=0,lte=130" vName:"年纪"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"isred" vName:"喜欢的颜色"`    // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func TestRegisterStructValidation(t *testing.T) {
	type args struct {
		fn    validator.StructLevelFunc
		types []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestValidateStruct(t *testing.T) {
	err2 := RegisterValidation("isred", func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		if s == "red" {
			return true
		}
		return false
	}, "{0}不是红色")
	assert.Nil(t, err2)

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "da",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "ww",
		Age:            130,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}
	err := ValidateStruct(address)
	assert.Nil(t, err)
	fmt.Println(err)

	err = ValidateStruct(user)
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestValidateVar(t *testing.T) {
	err2 := RegisterValidation("isred", func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		if s == "red" {
			return true
		}
		return false
	}, "{0}不是红色")
	assert.Nil(t, err2)
	err := ValidateVar("b", "isred")
	fmt.Println(err)
}
