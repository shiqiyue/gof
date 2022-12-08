package vipers

import (
	"context"
	"encoding/json"
	"github.com/shiqiyue/gof/asserts"
	"github.com/shiqiyue/gof/loggers"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitEtcdRemote() {
	viper.RemoteConfig = &EtcdRemoteConfig{}
}

// 使用etcd进行配置
func ConfigWithEtcd(valueInit, onValueChange func(v *viper.Viper), endpoint, paths, username, password, configType string) {
	v := viper.New()
	// 从配置中心取值
	loggers.Info(context.Background(), "获取配置文件", zap.String("endpoint", endpoint), zap.String("paths", paths), zap.String("username", username), zap.String("password", password))
	cert := &etcdCertificate{
		Username: username,
		Password: password,
	}
	certBs, err := json.Marshal(cert)
	asserts.Nil(err, err)

	err = v.AddSecureRemoteProvider("etcd", endpoint, paths, string(certBs))
	asserts.Nil(err, err)

	v.SetConfigType(configType)
	// 初始化
	go func() {
		err = v.ReadRemoteConfig()
		asserts.Nil(err, err)
		valueInit(v)
		asserts.Nil(err, err)
		for true {
			err := v.WatchRemoteConfig()
			if err != nil {
				loggers.Error(context.Background(), "监听etcd配置异常", zap.String("endpoint", endpoint), zap.String("paths", paths), zap.Error(err))
			} else {
				onValueChange(v)
			}
		}
	}()

}
