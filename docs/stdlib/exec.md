
# Package: exec

## Functions

### ExecuteCommand

ExecuteCommand executes the given command and waits for it to complete, returning pid, stdout, stderr, exitCode, or any errors

``````
func ExecuteCommand(c string, args []interface{}) (int, string, string, int, error)
``````
#### Example
``````

``````

### ExecuteCommandAsync

ExecuteCommandAsync runs the command and does not wait for it to return.

``````
func ExecuteCommandAsync(c string, args []interface{}) (*executer.Cmd, error)
``````
#### Example
``````

``````
