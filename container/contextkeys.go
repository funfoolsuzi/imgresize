package container

import "context"

type contextKey string

const contextKeyOriginal contextKey = "original"

const contextKeyHeight contextKey = "height"

const contextKeyWidth contextKey = "width"

func getOriginalPath(ctx context.Context) string {
	opRaw := ctx.Value(contextKeyOriginal)
	op, ok := opRaw.(string)
	if !ok {
		return ""
	}

	return op
}

func getHeight(ctx context.Context) string {
	hRaw := ctx.Value(contextKeyHeight)
	h, ok := hRaw.(string)
	if !ok {
		return ""
	}

	return h
}

func getWidth(ctx context.Context) string {
	wRaw := ctx.Value(contextKeyWidth)
	w, ok := wRaw.(string)
	if !ok {
		return ""
	}

	return w
}
