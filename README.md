# txlog
in memory transaction log (for events and actions)

An in-memory structure to manage transactions of operations for a distributed environment.

# Basic Idea
A log structure to handle events from the past and for the future (actions) creating a save in-memory structure for handling this. TO be mapped to persistency independently and separately.

All operations are recorded into a in memory queue, which is indexed by a "transaction sequence number". There can be many queues, but a single queue is described by a unique seqeunce number - the "TSN". A TSN is identifying a position in the sequence uniquely, by that a TSN has to be guaranteed to be monotonically increasing over time. On the other side, there is no need to avoid gaps in the sequence numbers, they just need to be monotonically increasing.

# Clocks
A TSN is distributed by a logical "clock", which is associated to the queue. There can be many clocks, so that a general overall TSN has to have a per clock a unique clock id as spatial component, if seen more general. 



