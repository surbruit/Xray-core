package udp

import (
	"github.com/yuzuki999/xray-core/common"
	"github.com/yuzuki999/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
