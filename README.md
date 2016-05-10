# txlog
In-memory transaction log (for events and actions)

An in-memory structure to manage transactions of operations for a distributed environment.

# Basic Idea
A log structure to handle events from the past (events) and for the future (actions) creating a save in-memory structure for handling this. The memory structure is be mapped to persistency independently and separately.

All operations are recorded into an in memory queue, which is indexed by a "transaction sequence number" - kind of a "flake" id. There can be many queues, but a single queue entry is described by a unique sequence number - a "TSN". A TSN is identifying a position in the sequence of a queue structure uniquely and to achieve this TSN has to be guaranteed to be monotonically increasing over time per queue. On the other side, there is no need to avoid gaps in the TSN sequence numbers, they just need to be monotonically increasing.

# Clocks
A TSN is distributed by a logical "clock", which is associated to the queue. There can be many clocks, so that a general overall TSN has to have a per clock a unique clock id as spatial component, if seen more general. There are many ways one can do this. Since we here have the combination of a log structure and the need for a distributable clock, we can use the log structure itself as the time producing entity and simply assume that the position of an entry in the queue actually gives the needed (non physical) "timestamp".

The base structure of a TSN is the nothing but a spatial component identifying the queue structure in memory in a distributed manner together with a queue index position. It is to be guranteed that the queue position is advanced by inserting requests into there without conflicts. 

The implementation is using a goroutine as a singleton per queue to advance the queue position. 




