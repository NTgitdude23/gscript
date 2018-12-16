
# Package: net


### CheckForInUseTCP

CheckForInUseTCP is a function that checks all local IPv4 interfaces for to see if something is listening on the specified TCP port will timeout after 3 seconds

``````
func CheckForInUseTCP(port int) (bool, error)
``````
#### Example
``````

``````

### CheckForInUseUDP

CheckForInUseUDP will send a UDP packet to the local port and see it gets a response or will timeout

``````
func CheckForInUseUDP(port int) (bool, error)
``````
#### Example
``````

``````

