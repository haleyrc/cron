// Package cron is an experimental cron-like job runner. The current
// implementation requires clients to do the actual scheduling, and only
// abstracts out the process of running a set of functions on a given time
// interval. This puts the onus of ensuring correct scheduling on the client,
// but gives maximum flexibility, since clients are now able to define, on an
// individual basis, what their definition of success is.
//
// See the examples for simple uses of the scheduling mechanism as well as more
// complicated examples which implement basic memory. Something like this could
// be used to ensure an event runs even if the ticker doesn't run at a precise
// time, as long as it's within a certain error bound.
package cron
