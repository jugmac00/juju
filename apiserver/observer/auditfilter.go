// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package observer

import (
	"fmt"
	"sync"

	"github.com/juju/errors"
	"github.com/juju/utils/set"

	"github.com/juju/juju/core/auditlog"
)

// bufferedLog defers writing records to its destination audit log
// until it sees an interesting request - then all buffered messages
// and subsequent ones get forwarded on.
type bufferedLog struct {
	mu          sync.Mutex
	buffer      []interface{}
	dest        auditlog.AuditLog
	interesting func(auditlog.Request) bool
}

// NewAuditLogFilter returns an auditlog.AuditLog that will only log
// conversations to the underlying log passed in if they include a
// request that satisfies the filter function passed in.
func NewAuditLogFilter(log auditlog.AuditLog, filter func(auditlog.Request) bool) auditlog.AuditLog {
	return &bufferedLog{
		dest:        log,
		interesting: filter,
	}
}

// AddConversation implements auditlog.AuditLog.
func (l *bufferedLog) AddConversation(c auditlog.Conversation) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	// We always buffer the conversation, since we don't know whether
	// it will have any interesting requests yet.
	l.deferMessage(c)
	return nil
}

// AddRequest implements auditlog.AuditLog.
func (l *bufferedLog) AddRequest(r auditlog.Request) error {
	l.mu.Lock()
	if len(l.buffer) > 0 {
		l.deferMessage(r)
		var err error
		if l.interesting(r) {
			err = l.flush()
		}
		l.mu.Unlock()
		return err
	}
	l.mu.Unlock()
	// We've already flushed messages, forward this on
	// immediately.
	return l.dest.AddRequest(r)
}

// AddResponse implements auditlog.AuditLog.
func (l *bufferedLog) AddResponse(r auditlog.ResponseErrors) error {
	l.mu.Lock()
	if len(l.buffer) > 0 {
		l.deferMessage(r)
		l.mu.Unlock()
		return nil
	}
	l.mu.Unlock()
	// We've already flushed messages, forward this on
	// immediately.
	return l.dest.AddResponse(r)
}

// Close implements auditlog.AuditLog.
func (l *bufferedLog) Close() error {
	return errors.Trace(l.dest.Close())
}

func (l *bufferedLog) deferMessage(m interface{}) {
	l.buffer = append(l.buffer, m)
}

func (l *bufferedLog) flush() error {
	for _, message := range l.buffer {
		var err error
		switch m := message.(type) {
		case auditlog.Conversation:
			err = l.dest.AddConversation(m)
		case auditlog.Request:
			err = l.dest.AddRequest(m)
		case auditlog.ResponseErrors:
			err = l.dest.AddResponse(m)
		default:
			err = errors.Errorf("unknown audit log message type %T %+v", m, m)
		}
		if err != nil {
			return errors.Trace(err)
		}
	}
	l.buffer = nil
	return nil
}

// InterestingRequest returns whether this API request is interesting enough
// to write the conversation to the audit log.
func InterestingRequest(req auditlog.Request) bool {
	return !readOnlyMethods.Contains(fmt.Sprintf("%s.%s", req.Facade, req.Method))
}

var readOnlyMethods = set.NewStrings(
	"Client.FullStatus",
)
