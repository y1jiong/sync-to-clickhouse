package clickhouse

import (
	"context"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcron"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync-to-clickhouse/internal/service"
)

type sClickHouse struct {
	db gdb.DB

	insertMu        sync.Mutex
	insertQueue     *gqueue.Queue
	insertQueuePath string

	flushCount uint32

	optimizeTableMu    sync.Mutex
	optimizeTableQueue *gqueue.Queue

	crontabMu          sync.Mutex
	crontab            *gcron.Cron
	flushEntry         *gcron.Entry
	optimizeTableEntry *gcron.Entry
}

func New() *sClickHouse {
	return &sClickHouse{
		insertQueuePath: "insert_queue.json",
	}
}

func init() {
	service.RegisterClickHouse(New())
}

func (s *sClickHouse) hasDB() error {
	if s.db == nil {
		return gerror.New("clickhouse db is nil")
	}
	return nil
}

func (s *sClickHouse) SetDBLink(link string) (err error) {
	s.db, err = gdb.New(gdb.ConfigNode{Link: link})
	return
}

func (s *sClickHouse) Flush(ctx context.Context) error {
	return s.flushInsertQueue(ctx)
}

func (s *sClickHouse) DumpToDisk(ctx context.Context) (err error) {
	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.dumpInsertQueueToDisk(egCtx)
	})

	if err = eg.Wait(); err != nil {
		return
	}

	return
}

func (s *sClickHouse) RestoreFromDisk(ctx context.Context) (err error) {
	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.restoreInsertQueueFromDisk(egCtx)
	})

	if err = eg.Wait(); err != nil {
		return
	}

	return
}

func (s *sClickHouse) Close(ctx context.Context) (errs []error) {
	if s.flushEntry != nil {
		s.flushEntry.Close()
		s.flushEntry = nil
	}

	if s.optimizeTableEntry != nil {
		s.optimizeTableEntry.Close()
		s.optimizeTableEntry = nil
	}

	if s.crontab != nil {
		s.crontab.Close()
		s.crontab = nil
	}

	if s.insertQueue != nil {
		s.insertQueue.Close()
		s.insertQueue = nil
	}

	if s.optimizeTableQueue != nil {
		s.optimizeTableQueue.Close()
		s.optimizeTableQueue = nil
	}

	if s.db != nil {
		if err := s.db.Close(ctx); err != nil {
			errs = append(errs, err)
		} else {
			s.db = nil
		}
	}

	return
}
