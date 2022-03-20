package drain

import "io"

//go:generate go run github.com/yuzuki999/xray-core/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
