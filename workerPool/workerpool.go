package workerPool

import (
	"fmt"
	"github.com/devtron-labs/common-lib/constants"
	"github.com/devtron-labs/common-lib/pubsub-lib/metrics"
	"github.com/devtron-labs/common-lib/utils/reflectUtils"
	"github.com/devtron-labs/common-lib/utils/runTime"
	"github.com/gammazero/workerpool"
	"go.uber.org/zap"
	"reflect"
	"runtime/debug"
	"sync"
)

type WorkerPool[T any] struct {
	wp       *workerpool.WorkerPool
	mu       *sync.Mutex
	err      chan error
	service  constants.ServiceName
	logger   *zap.SugaredLogger
	response []T
}

func NewWorkerPool[T any](maxWorkers int, serviceName constants.ServiceName, logger *zap.SugaredLogger) *WorkerPool[T] {
	return &WorkerPool[T]{
		wp:       workerpool.New(maxWorkers),
		mu:       &sync.Mutex{},
		err:      make(chan error, 1),
		logger:   logger,
		service:  serviceName,
		response: []T{},
	}
}

func (impl *WorkerPool[T]) Submit(task func() (T, error)) {
	if task == nil {
		return
	}
	impl.wp.Submit(func() {
		defer func() {
			if r := recover(); r != nil {
				metrics.IncPanicRecoveryCount("go-routine", impl.service.ToString(), runTime.GetCallerFunctionName(), fmt.Sprintf("%s:%d", runTime.GetCallerFileName(), runTime.GetCallerLineNumber()))
				impl.logger.Errorw(fmt.Sprintf("%s %s", constants.GoRoutinePanicMsgLogPrefix, "go-routine recovered from panic"), "err", r, "stack", string(debug.Stack()))
			}
		}()
		if impl.Error() != nil {
			return
		}
		res, err := task()
		if err != nil {
			impl.logger.Errorw("error in worker pool task", "err", err)
			impl.SetError(err)
			return
		}
		val := reflect.ValueOf(res)
		impl.Lock()
		if reflectUtils.IsNullableValue(val) {
			if val.IsNil() {
				impl.Unlock()
				return
			}
		} else if val.IsZero() {
			impl.Unlock()
			return
		}
		impl.response = append(impl.response, res)
		impl.Unlock()
	})
}

func (impl *WorkerPool[_]) StopWait() {
	impl.wp.StopWait()
}

func (impl *WorkerPool[_]) Lock() {
	impl.mu.Lock()
}

func (impl *WorkerPool[_]) Unlock() {
	impl.mu.Unlock()
}

func (impl *WorkerPool[_]) Error() error {
	select {
	case err := <-impl.err:
		return err
	default:
		return nil
	}
}

func (impl *WorkerPool[_]) SetError(err error) {
	if err != nil && impl.Error() == nil {
		impl.err <- err
	}
}

func (impl *WorkerPool[T]) GetResponse() []T {
	return impl.response
}
