package timer

// Clock with blocking ticks. Probably a spoonerism.
type Clock chan bool

// Tick waits for a signal from a clock source.
func (c Clock) Tick() {
	<-c
}

// Ticks waits for a signal from a clock source for a given number of cycles.
func (c Clock) Ticks(ticks int) {
	for ; ticks >= 0; ticks-- {
		<-c
	}
}
