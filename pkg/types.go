package pkg

import (
	"context"
	"github.com/SameeranB/go-heavy/pkg/providers"
)

type Workflow struct {
	Name     string
	SetUp    func(pctx PassedContext) (PassedContext, error)
	Init     func(pctx PassedContext) (PassedContext, error)
	Steps    []func(pctx PassedContext) (PassedContext, error)
	Teardown func(pctx PassedContext) (PassedContext, error)
}

type PassedContext struct {
	Ctx        context.Context
	CancelFunc context.CancelFunc
	Config     interface{}
	Params     interface{}
	Iter       int64
	UserId     string
	Providers  providers.Providers
}
