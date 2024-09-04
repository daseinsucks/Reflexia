package prompt

import util "reflexia/lib/internal"

var DefaultCodePrompt = `
Describe each code symbol in short form as in the examples below.
Reduct your output as much as possible.
Don't lose original namings.
Always specify environment variables, cli arguments, flags, anything that can configure application behaviour.
Always omit empty sections! Write about section only if they present!
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.

Example input:
` + "```" + `
package foobar
import (
	"fmt"
	"os"
)

var zab = "zab"
const zoob = "zoob"

struct FooBar {
	Foo string
	Bar string
}

func SomeFunc() {
	// Unimplemented
	// TODO: implement SomeFunc
}

func main() {
	foobar := FooBar{
		Foo: "foo",
		Bar: "bar"
	}
	// Print current directory
	fmt.Println(os.Getenv("PWD"))
	fmt.Println(foobar.Foo + foobar.Bar + os.Getenv("BAZ") + zab + zoob)
	// Reverse of above output
	fmt.Println("boozbazZABraboff")
	//fmt.Println("TEST")
}
` + "```" + `

Example output:
package foobar
import: fmt, os

struct FooBar:
	- Fields: Foo, Bar
var:
	- zab: "zab"
const:
	- zoob: "zoob"

func SomeFunc(): unimplemented
func main():
	- assigns var foobar with FooBar struct with fields Foo: "foo", Bar: "bar"
	- prints current working directory
	- prints foobar.Foo + foobar.Bar + environment variable "BAZ" + var zab + const zoob
	- prints "boozbazZABraboff" which is reverse of above output

environment variables: "PWD", "BAZ"
` + util.LoadEnv("STOP_WORD") + `

Example input:
` + "```" + `
package main
import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
}
` + "```" + `

Example output:
package main
import: fmt

func main():
	- prints "Hello world!"
` + util.LoadEnv("STOP_WORD") + `

Provided code:
`

var DefaultPackageSummaryPrompt = `
Based on provided input from the summary of project package files create a markdown summary of what that package code does.
First write a short summary about provided project code summary.
Always specify which environment variables, flags, cmdline arguments can be used for configuration.
Always specify the edgecases of how application can be launched.
Try to guess package name from the file contents or file paths and add it as the markdown header of the summary.
Write out all file names as a project package structure.
Then write summary about every major code part, group it with the markdown subheaders.
Try to explain relations between code entities, try to find unclear places, possibly dead code.		
If unclear places or dead code are not present - don't write anything about their absense.
Try to be clear, concise, and brief.
The main goal is to summarize the logic of the whole package.
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.
`

var DefaultSummaryPrompt = `
Based on provided input from summary of the project files create a markdown summary of what that project code does.
First write a short summary about provided project code summary.
Always specify which environment variables, flags, cmdline arguments can be used for configuration.
Always specify the edgecases of how application can be launched.
Try to guess the project name from the filenames if possible and add it as the markdown header of the summary.
Write out all file names as a project structure.
Then write summary about every major code part, group it with the markdown subheaders.
Try to explain relations between code entities, try to find unclear places, possibly dead code.		
If unclear places or dead code are not present - don't write anything about their absense.
Try to be clear, concise, and brief.
The main goal is to summarize the logic of the whole project.
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.
`
var DefaultReadmePrompt = `
Based on provided input from technical summary of the project create a README.md contents for that project.
Try to guess the project name from filenames or other namings, if present.
Include short project description, possible configuration and/or run instructions.
Try to be clear, concise, and brief.
It is mandatory to prepend the ` + util.LoadEnv("STOP_WORD") + ` at the very end of your output.
`
