package list

import (
	"errors"

	"dinodb/pkg/repl"

	"strings" //self import

	"fmt" //self import
)

// Use these provided errors instead of defining your own!
var (
	ErrListPrintInvalidArgs = errors.New("invalid arguments, usage: list_print")

	ErrListPushHeadInvalidArgs = errors.New("invalid arguments, usage: list_push_head <elt>")

	ErrListPushTailInvalidArgs = errors.New("invalid arguments, usage: list_push_tail <elt>")

	ErrListRemoveValueNotFound = errors.New("link with given value was not found")
	ErrListRemoveInvalidArgs   = errors.New("invalid arguments, usage: list_remove <elt>")

	ErrListContainsInvalidArgs = errors.New("invalid arguments, usage: list_contains <elt>")
)

const (
	// Use these help strings for each command instead of defining your own!
	HelpListPrint    = "Input: List of anything. Prints out all of the elements in the list in order. usage: list_print"
	HelpListPushHead = "Inserts the given element to the head of the list as a string. usage: list_push_head <elt>"
	HelpListPushTail = "Inserts the given element to the end of the list as a string. usage: list_push_tail <elt>"
	HelpListRemove   = "Removes the given element from the list. usage: list_remove <elt>"
	HelpListContains = "Check whether the element is in the list or not. usage: list_contains <elt>"

	// Output strings for the list_contains command
	OutputListContainsFound    = "value was found"
	OutputListContainsNotFound = "value was not found"
)

/*
Create a NewRepl() and use repl.AddCommand() to create these following commands:
- list_print
- list_push_head <elt>
- list_push_tail <elt>
- list_remove <elt> TODO: removes the First instance of an object
- list_contains <elt>

[Notes]:
Remember that AddCommand() takes in a:

1. Trigger (name of the command),

2. ReplCommand (function containing code to run command),
  - A ReplCommand is a function that returns a string and error (string, error).
  - Return any output in the string, and any errors in error
  - We check error based on whether it is nil or not
  - If an error exists, the returned string is NOT used! Instead, return an error (using the above defined INVALID_ARGS_ERR)
  - There are also custom error checks for some functions. Hint: They should match the above defined error messages

3. Help string (string explaining how to use the command)
  - Help strings should be one line! Do NOT use '\n' in the Help Strings.
  - Use the above help str vars for your help strings!
*/
func ListRepl(list *List[Str]) *repl.REPL {
	rep := repl.NewRepl()

	rep.AddCommand("list_print", 
		func (input string, config *repl.REPLConfig) (string, error){  //function for replcommand, must follow constructor of replcommand then cast
			fields := strings.Fields(input)
			if len(fields) != 1 || fields[0] != "list_print" {
				return "", ErrListPrintInvalidArgs
			}

			temp := list.head
			if temp == nil {
				return "", nil
			}
			printedVals := ""

			for {
				printedVals += fmt.Sprintf("%s", temp.value)
				if temp.next != nil {
					printedVals += "\n"
					temp = temp.next
				} else {
					return printedVals, nil
				}
			}
			
		}, 
		HelpListPrint,
	)


   	rep.AddCommand("list_push_head", 
		func (input string, config *repl.REPLConfig) (string, error){  //function for replcommand, must follow constructor of replcommand then cast
			fields := strings.Fields(input)
			if len(fields) != 2 || fields[0] != "list_push_head" {
				return "", ErrListPushHeadInvalidArgs
			}
			list.PushHead(Str(fields[1]))
			return "", nil //input, nil
		}, 
		HelpListPushHead,
	)


	rep.AddCommand("list_push_tail", 
		func (input string, config *repl.REPLConfig) (string, error){  //function for replcommand, must follow constructor of replcommand then cast
			fields := strings.Fields(input)
			if len(fields) != 2 || fields[0] != "list_push_tail" {
				return "", ErrListPushTailInvalidArgs
			}
			
			//link = 
			list.PushTail(Str(input))
			return input, nil
		}, 
		HelpListPushTail,
	)

	rep.AddCommand("list_remove", 
		func (input string, config *repl.REPLConfig) (string, error){  //function for replcommand, must follow constructor of replcommand then cast
			fields := strings.Fields(input)
			if len(fields) != 2 || fields[0] != "list_remove" {
				return "", ErrListRemoveInvalidArgs
			}
				
			link := list.Find(func(link *Link[Str]) bool { //match the method signature of Find
								return string(link.value) == fields[1]
							})
			if link == nil { //link not found with value
				return "", ErrListRemoveValueNotFound
			} else {
				link.PopSelf()
				return input, nil
			}
		}, 
		HelpListRemove,
	)


	rep.AddCommand("list_contains", 
		func (input string, config *repl.REPLConfig) (string, error){  //function for replcommand, must follow constructor of replcommand then cast
			fields := strings.Fields(input)
			if len(fields) != 2 || fields[0] != "list_contains" {
				return "", ErrListContainsInvalidArgs
			}
				
			link := list.Find(func(link *Link[Str]) bool { //match the method signature of Find, find with matching value
								return string(link.value) == fields[1] 
							})
			if link == nil { //link not found with value
				return OutputListContainsNotFound, nil
			} else {
				return OutputListContainsFound, nil
			}
		}, 
		HelpListContains,
	)

   	return rep
}
