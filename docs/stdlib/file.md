
# Package: file

## Functions

### AppendFileBytes

AppendFileBytes takes a file and adds a byte array to the end of it

``````
func AppendFileBytes(targetPath string, addData []byte) error
``````
#### Example
``````

``````

### AppendFileString

AppendFileString takes a file and adds a string to the end of it

``````
func AppendFileString(targetPath, addString string) error
``````
#### Example
``````

``````

### CheckExists

CheckExists takes a file or directory and checks to see if it exists on the file system

``````
func CheckExists(targetPath string) bool
``````
#### Example
``````

``````

### CopyFile

CopyFile copies a file from the src to the dest with the original files permissions

``````
func CopyFile(srcPath, destPath string) (int, error)
``````
#### Example
``````

``````

### ReadFileAsBytes

ReadFileAsBytes takes a file path and reads that files contents and returns a byte array of the contents

``````
func ReadFileAsBytes(readPath string) ([]byte, error)
``````
#### Example
``````

``````

### ReadFileAsString

ReadFileAsString takes a file path and reads that files contents and returns a string representation of the contents

``````
func ReadFileAsString(readPath string) (string, error)
``````
#### Example
``````

``````

### ReplaceInFileWithRegex

ReplaceInFileWithRegex searches a file for a string and replaces each instance found of that string. Returns the amount of strings replaced

``````
func ReplaceInFileWithRegex(file string, regexString string, replaceWith string) (int, error)
``````
#### Example
``````

``````

### ReplaceInFileWithString

ReplaceInFileWithString searches a file for a string and replaces each instance found of that string. Returns the amount of strings replaced

``````
func ReplaceInFileWithString(file, match, replacement string) (int, error)
``````
#### Example
``````

``````

### SetPerms

SetPerms changes the file permissions of a givin file

``````
func SetPerms(targetPath string, perms int64) error
``````
#### Example
``````

``````

### WriteFileFromBytes

WriteFileFromBytes writes data from a byte array to a dest filepath with the dest parent dirs permissions.

``````
func WriteFileFromBytes(destPath string, fileData []byte) error
``````
#### Example
``````

``````

### WriteFileFromString

WriteFileFromString writes data from a string to a dest filepath with the dest parent dirs permissions.

``````
func WriteFileFromString(destPath string, fileData string) error
``````
#### Example
``````

``````
