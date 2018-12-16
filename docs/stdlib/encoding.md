
# Package: encoding


### DecodeBase64

DecodeBase64 decodes a base64 string and returns a string

``````
func DecodeBase64(data string) (string, error)
``````
#### Example
``````

``````

### EncodeBase64

EncodeBase64 takes a string and turns it into a base64 string

``````
func EncodeBase64(data string) string
``````
#### Example
``````

``````

### EncodeBytesAsString

EncodeBytesAsString takes a byte array and returns a string representation

``````
func EncodeBytesAsString(data []byte) string
``````
#### Example
``````

``````

### EncodeStringAsBytes

EncodeStringAsBytes takes a string and returns the bytes that make it

``````
func EncodeStringAsBytes(data string) []byte
``````
#### Example
``````

``````

