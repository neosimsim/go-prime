# go-prime
An academic tool generating prime numbers to all eternity.

The prime numbers are generated using a go routine oriented approach.

For each found prime number a new go routine is started filtering multiples
of the found prime out of a generating channel.
