// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IClickHouse interface {
		SetDBLink(link string) (err error)
		Flush(ctx context.Context) error
		OptimizeTable(ctx context.Context, table map[string]struct{}) (err error)
		DumpToDisk(ctx context.Context) (err error)
		RestoreFromDisk(ctx context.Context) (err error)
		SetCountFlush(count uint)
		SetCrontabFlush(ctx context.Context, crontabExpr string, isEnableOptimizeTable bool) (err error)
		SetCrontabOptimizeTable(ctx context.Context, crontabExpr string, table map[string]struct{}) (err error)
		Insert(ctx context.Context, table string, data []map[string]string) (err error)
	}
)

var (
	localClickHouse IClickHouse
)

func ClickHouse() IClickHouse {
	if localClickHouse == nil {
		panic("implement not found for interface IClickHouse, forgot register?")
	}
	return localClickHouse
}

func RegisterClickHouse(i IClickHouse) {
	localClickHouse = i
}
