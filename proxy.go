package goka

import (
	"fmt"

	"github.com/lovoo/goka/storage"
)

type storageProxy struct {
	storage.Storage
	topic     Stream
	partition int32
	stateless bool
	update    UpdateCallback

	openedOnce once
	closedOnce once
}

func (s *storageProxy) Open() error {
	if s == nil {
		return nil
	}
	return s.openedOnce.Do(s.Storage.Open)
}

func (s *storageProxy) Close() error {
	if s == nil {
		return nil
	}
	return s.closedOnce.Do(s.Storage.Close)
}

func (s *storageProxy) Update(ctx UpdateContext, k string, v []byte) error {
	fmt.Printf("storageProxy.Update: topic=%s, partition=%d, k=%s, v=%s\n", ctx.Topic(), ctx.Partition(), k, v)
	return s.update(ctx, s, k, v)
}

func (s *storageProxy) Stateless() bool {
	return s.stateless
}

func (s *storageProxy) MarkRecovered() error {
	return s.Storage.MarkRecovered()
}
