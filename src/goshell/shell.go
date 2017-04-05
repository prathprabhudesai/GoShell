
package main

import(
	"bufio"
	"fmt"
	"shell/parser"
	"os"
	"github.com/fatih/color"
	"path/filepath"
)



var Pipes int = 0 

func makeString(command []byte) string {
	var row string
	row = ""
	for i:=0; i < len(command) -1; i++ {
		row = fmt.Sprintf("%v%v",row,string(command[i]))
	} 
	return row
}



func main(){
	for {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		c := color.New(color.FgRed)
		shell_line := fmt.Sprintf("$GoShell$ @ %v/ ",dir)
		c.Printf("%v",shell_line)
		buf := bufio.NewReader(os.Stdin) 
		command, err := buf.ReadBytes('\n') 
		
		if err != nil {
			fmt.Println("Runtime buffer error")
		} else {
			if(len(command) == 1){
				c.Printf("%v",shell_line) 
			} else {
				exec := parser.Parse(makeString(command))
				parser.PushToHistory(makeString(command))
				if (!exec) {
					break;
				}
			}
		}
	}
}



