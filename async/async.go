package async

import "context"

// Future interface has the method signature for await
type Future[K any] interface {
	Await() K
}
type future[K any] struct {
	await func(ctx context.Context) K
}

func (f future[K]) Await() K {
	return f.await(context.Background())
}

// Exec executes the async function
func Exec[K any, F func() K](f F) Future[K] {
	var r interface{}
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		r = f()
	}()
	return future[K]{
		await: func(ctx context.Context) K {
			select {
			case <-ch:
				return r.(K)
			case <-ctx.Done():
				return r.(K)
			}
		},
	}
}
