# Basic Go command-line tool
Go provides a basic library called flag. It refers to the command-line flags. Since it is
already packed with the Go distribution, there is no need to install anything externally. I
can see the absolute basics of writing the command-line tool. The flag package has
multiple functions, such as Int and String, to handle the input given as command-line
flags.

In the basic example, Suppose I need to take a name from the user and print it back to the console. I
use the flag.String method, as shown in the following code snippet:

```
import "flag"
var name = flag.String("name", "No Namer", "your wonderful name")

```
To run the basic CLI, you need to see `cd` into the current directory and build it first before you run it using the following commands:

```
go build flagMP.go

```

This create a binary in the same directory. Run it like a normal executable:

```
./flagMP

```

It gives out an output. In other to it with two parameters:

```
./flagMP -name Atanda0x -age 20

```

   