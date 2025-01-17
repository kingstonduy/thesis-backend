package cmdregister

import (
	"context"
	"fmt"
)

// Define an interface for handlers
type handler interface {
	handle(ctx context.Context, input any) (output any, err error)
}

// handlerWrapper adapts specific handler functions to the Handler interface
type handlerImpl[T any, K any] struct {
	f func(ctx context.Context, input T) (output K, err error)
}

func (hw *handlerImpl[T, K]) handle(ctx context.Context, input any) (output any, err error) {
	inputVal, ok := input.(T)
	if !ok {
		return nil, fmt.Errorf("expected input of type %T, got %T", inputVal, input)
	}
	return hw.f(ctx, inputVal)
}

// map with string keys and functions as values
var m = make(map[string]handler, 1)

// Register function to add handlers to the map
func Register[T any, K any](con string, f func(ctx context.Context, input T) (output K, err error)) {
	_, exist := m[con]
	if exist {
		// each request in request/response strategy should have just one handler
		s := fmt.Sprintf("registered handler already exists in the registry for message %s", con)
		panic(s)
	}

	// Use a wrapper function to handle type assertion
	m[con] = &handlerImpl[T, K]{
		f: f,
	}
}

// look up in the map, then send the inpụt to the handler
func Process[T any, K any](ctx context.Context, con string, input T) (output K, err error) {
	h, ok := m[con]
	if !ok { // exist handler func or not
		err = fmt.Errorf("the condition=%s is not registered", con)
		return output, err
	}

	res, err := h.handle(ctx, input)
	output, ok = res.(K)
	if err != nil || !ok {
		if !ok {
			err = fmt.Errorf("expected output of type %T, got %T", output, res)
		}
		return output, err
	}

	return output, nil
}
