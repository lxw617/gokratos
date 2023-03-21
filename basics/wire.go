//go:build wireinject
// +build wireinject

// 上面这一行代码一定要在包最上面声明，表明这是一个准备被编译的 injector

package main

import (
	"context"

	"github.com/google/wire"
)

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

// 参考官方文档 https://github.com/google/wire/blob/main/docs/guide.md

func initializeBaz(ctx context.Context) (Baz, error) {
	wire.Build(SuperSet)
	return Baz{}, nil
}
