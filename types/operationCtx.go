package types

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"
)

type CtxType int

const (
	_ CtxType = iota
	CtxOperationTracker
)

type OperationTracker struct {
	traceID     string
	startTime   time.Time
	authAccount Account
}

func NewOperationCtx(
	baseContext context.Context,
	traceID string,
	startTime time.Time,
	authAccount Account,
) context.Context {
	if startTime.IsZero() {
		startTime = time.Now()
	}
	if traceID == "" {
		// Worst case scenario we get some colliding traceIDs.
		// Extremely unlikely, worth the risk to avoid bailing
		traceID = safeishUUIDV4()
	}
	operationTracker := OperationTracker{
		traceID:     traceID,
		startTime:   startTime,
		authAccount: authAccount,
	}

	return context.WithValue(baseContext, CtxOperationTracker, operationTracker)
}

func safeishUUIDV4() (uuid string) {
	bytes := make([]byte, 16)

	// safe-ish because rand.Read could error
	// but it is very, very unlikely to do so
	_, _ = rand.Read(bytes)

	// ensure 13th char is 4
	bytes[6] = (bytes[6] | 0x40) & 0x4F
	// ensure 17th char is 8, 9, a, or b
	bytes[8] = (bytes[8] | 0x80) & 0xBF

	return fmt.Sprintf(
		"%04X-%04X-%04X-%04X-%04X",
		bytes[0:4],
		bytes[4:6],
		bytes[6:8],
		bytes[8:10],
		bytes[10:],
	)
}
