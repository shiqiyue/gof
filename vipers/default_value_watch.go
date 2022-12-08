package vipers

import (
	"github.com/spf13/viper"
)

type DefaultValueWatch struct {
}

func (d DefaultValueWatch) OnChange(v *viper.Viper) {
	panic("implement me")
}

func (d DefaultValueWatch) OnInit(v *viper.Viper) {
	panic("implement me")
}
