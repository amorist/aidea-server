package fromston

import "github.com/mylxsw/glacier/infra"

type Provider struct{}

func (Provider) Register(binder infra.Binder) {
	binder.MustSingleton(NewFromston)
}
