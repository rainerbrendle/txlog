// Package transaction log (txlog)

package txlog

// basic structure of a transaction log queue
//
// an entry into a log is uniquely identified by a kind of a "flake" ID.
//
// (Time is not unique, time is only unique at a given counter bound to a spatial location)

// Transactional Sequence Numbers - it is a kind of a "FLake"
//
// - have a spatial component - the Clock ID
// - a timely component - a counter, which is to be monotonically increasing
// - if the counter is going backword or reaches its limits, a new "clock" is to be created
//

type clockId struct {
	mac uint64 // mac address
	seq uint64 // sequence id (must be global per machine) TO BE DETERMINED
}

type seqId struct {
	time uint64 // a monotonically increasing, timely unique number
}

type operation struct {
	text []byte // to be improved
}

type queue struct {
	clockId       clockId // the ID of the clock, which is used to define the timely sequence in the queue
	name          string  // the name of the queue
	highWaterMark seqId   // the latest recorded seqID.time

	operations []operation // recorded operations (seqID is index to this entry)

        channel    chan<- operation  // channel to send operations to
}


func NewQueue ( name string) *queue {

    var q queue
    /* initialize */

    /* TODO: create operations buffer  */
    /* TODO: create channel */
    /* create receiver goroutine, receiver advances the seqID
    /* a new queue must be a Singleton per process and queue name: check */

    return &q

}

func (q* queue) Channel()  chan<- operation {
     return q.channel
}


func Test() {

    var q * queue
    var o operation

    q = NewQueue( "test" )
 
    // o.text create message here

    // send operation to channel
    q.Channel() <- o

}
