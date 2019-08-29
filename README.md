# Jump

### Outline
The source for this is contained in `jump.go`. To run (after installing Go)
`go run jump.go`

To Test:
`go test -v`

There are two main functions, `addAction` which takes a string, and returns an error (or nil). This will
add an action to the running list of actions and times, and compute the new average

The next function is `getStats`, which returns a string containing all the actions and their averages

I have also included a `main` function for quick use/demonstration of the two functions above. I call a couple addActions in goroutines, and sleep for 5 seconds
to ensure they all complete. In retrospect, using waitgroups would have been a better route, but it would have added un-needed complexity for just the quick demonstration. 

### Future Considerations
One spot for future considerations is the use of the mutex around reading and writing to the data store. I wrap a pretty big chunk of code in the R/W mutuex,
and so it cuts down on the true concurrency of the function. Taking more time to see if there is a better design to allow for more concurrency would be a good area of improvement.
Also, using a persistent data store could be an area of future consideration.
