package resk

import (
	_ "github.com/vonnwang/account/apis/web"
	_ "github.com/vonnwang/account/core/accounts"
	"github.com/vonnwang/infra"
	"github.com/vonnwang/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
