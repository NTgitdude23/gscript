
# Package: crypto

## Functions

### GenerateRSASSHKeyPair

GenerateRSASSHKeyPair Generates a new rsa key pair the size of the arg, and returns the pub and priv key as strings

``````
func GenerateRSASSHKeyPair(size int) (string, string, error)
``````
#### Example
``````

``````

### GetMD5FromBytes

GetMD5FromBytes takes a byte array and returns the md5 string of it

``````
func GetMD5FromBytes(data []byte) string
``````
#### Example
``````

``````

### GetMD5FromString

GetMD5FromString takes a string and returns the md5 string of it

``````
func GetMD5FromString(data string) string
``````
#### Example
``````

``````

### GetSHA1FromBytes

GetSHA1FromBytes takes a byte array and returns the md5 string of it

``````
func GetSHA1FromBytes(data []byte) string
``````
#### Example
``````

``````

### GetSHA1FromString

GetSHA1FromString takes a byte array and returns the md5 string of it

``````
func GetSHA1FromString(data string) string
``````
#### Example
``````

``````

### GetSHA256FromBytes

GetSHA256FromBytes takes a byte array and returns the md5 string of it

``````
func GetSHA256FromBytes(data []byte) string
``````
#### Example
``````

``````

### GetSHA256FromString

GetSHA256FromString takes a byte array and returns the md5 string of it

``````
func GetSHA256FromString(data string) string
``````
#### Example
``````

``````
