package data

import (
	"errors"
	"fmt"
	"user/internal/conf"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewCasdoor)

// Data .
type Data struct {
	// TODO wrapped database client
	cs *casdoorsdk.Client
}

// NewData .
func NewData(
	c *conf.Data, 
	cs *casdoorsdk.Client, 
	logger log.Logger,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		cs: cs,
	}, cleanup, nil
}

func NewCasdoor(cc *conf.Auth) (*casdoorsdk.Client, error) {
	if cc == nil {
		return nil, errors.New("auth config is nil") // ❌ 防止传空值
    }
	fmt.Printf("Casdoor Config: %+v\n", cc.Casdoor) // ✅ 调试打印，看看配置是否完整
	client := casdoorsdk.NewClient(
		cc.Casdoor.Endpoint,
		cc.Casdoor.ClientId,
		cc.Casdoor.ClientSecret,
		cc.Casdoor.Certificate,
		cc.Casdoor.Organization,
		cc.Casdoor.Application,
	)
    if client == nil {  // 确保 casdoor.NewClient 真的返回了一个实例
        return nil, errors.New("failed to create casdoor client")
    }
	return client, nil
}