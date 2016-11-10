package goo

var _ Chan = ChanFloat32(nil)

var _ Pointer = (*ChanFloat32)(nil)

// ChanFloat32 is a channel of float32.
type ChanFloat32 chan float32

// Cap implements Chan.
func (c ChanFloat32) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanFloat32) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanFloat32) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanFloat32) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanFloat32) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanFloat32) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanFloat32) Make(cap int) Chan {
	return make(ChanFloat32, cap)
}

// Receive implements Chan.
func (c ChanFloat32) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanFloat32) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanFloat32) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanFloat32) Send(v interface{}) {
	c <- v.(float32)
}
