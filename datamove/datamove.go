package datamove

import (
	"context"
	"errors"
	"fmt"
)

type DateMoveProgress struct {
	Error error
}

func DateMove(ctx context.Context, from DataSource, to DataSource, onProgress func(progress *DateMoveProgress)) (err error) {
	progress := &DateMoveProgress{}
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%x", x))
			progress.Error = err
		}

		onProgress(progress)
	}()
	var isStopped bool
	go func() {
		select {
		case <-ctx.Done():
			isStopped = true
		}
	}()

	for {
		if isStopped {
			break
		}
	}
	return
}
