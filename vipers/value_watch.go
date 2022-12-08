package vipers

import "github.com/spf13/viper"

type ValueWatch interface {
	OnChange(v *viper.Viper)

	OnInit(v *viper.Viper)
}
