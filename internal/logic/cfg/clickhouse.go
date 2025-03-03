package cfg

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
)

func (s *sCfg) ClickHouseCountFlush(ctx context.Context) uint32 {
	return gcfg.Instance().MustGet(ctx, "clickhouse.flush.count", 0).Uint32()
}

func (s *sCfg) ClickHouseCrontabFlush(ctx context.Context) string {
	return gcfg.Instance().MustGet(ctx, "clickhouse.flush.crontab", "").String()
}

func (s *sCfg) IsClickHouseOptimizeTableAfterInsert(ctx context.Context) bool {
	return gcfg.Instance().MustGet(ctx, "clickhouse.optimizeTable.enableAfterInsert", false).Bool()
}

func (s *sCfg) ClickHouseCrontabOptimizeTable(ctx context.Context) string {
	return gcfg.Instance().MustGet(ctx, "clickhouse.optimizeTable.crontab", "").String()
}
