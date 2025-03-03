package clickhouse

import "sync/atomic"

func (s *sClickHouse) SetCountFlush(count uint32) {
	atomic.StoreUint32(&s.flushCount, count)
}
