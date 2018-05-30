package rudp

import "sync/atomic"

type counter uint32

func (c *counter) GetAndIncrement() uint32 {
	var next uint32

	for {
		next = uint32(*c) + 1
		if atomic.CompareAndSwapUint32((*uint32)(c), uint32(*c), next) {
			return next - 1
		}
	}
}

func (c *counter) Get() uint32 {
	return atomic.LoadUint32((*uint32)(c))
}
