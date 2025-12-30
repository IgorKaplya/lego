package ctxreader

import (
	"context"
	"fmt"
	"io"
)

type CancellableReader struct {
	reader io.Reader
	ctx    context.Context
}

// Read implements [io.Reader].
func (c *CancellableReader) Read(p []byte) (n int, err error) {
	if errCtx := c.ctx.Err(); errCtx != nil {
		return 0, fmt.Errorf("read cancelled, %s", errCtx)
	}
	return c.reader.Read(p)
}

func NewCancellableReader(reader io.Reader, ctx context.Context) io.Reader {
	return &CancellableReader{reader: reader, ctx: ctx}
}
