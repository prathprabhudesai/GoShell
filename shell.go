
package main

import(
	"bufio"
	"fmt"
	"os" 
	"io/ioutil"
	"strings"
	"path/filepath"
)


var IsPipe bool = false
var Pipes int = 0
var History = make([]string,0)  



func makeString(command []byte) string {
	var row string
	row = ""
	for i:=0; i < len(command) -1; i++ {
		row = fmt.Sprintf("%v%v",row,string(command[i]))
	} 
	return row
}


func Parse(command string) bool {
	result := strings.Split(command, " ") 
	switch result[0] {
	case "ls":
		if len(result) > 1 {
			if (result[1] == "-all") {
				executeLS(true)
			}
		} else {
			executeLS(false)
		}

	case "pwd":
		pwd()
		
	case "history":
		printHistory()

	case "exit":
		return false

	default:
		fmt.Println("This is default") 
	}
	
	return true
}











func main(){
	for {
		fmt.Printf("@GoShell# ")
		buf := bufio.NewReader(os.Stdin) 
		command, err := buf.ReadBytes('\n') 
		
		if err != nil {
			fmt.Println("Runtime buffer error")
		} else {
			if(len(command) == 1){
				fmt.Printf("@GoShell# ")
			} else {
				exec := Parse(makeString(command))
				pushToHistory(makeString(command))
				if (!exec) {
					break;
				}
			}
		}
	}
}



// ls and ls -all
func executeLS (isAll bool) {
	if (!isAll){
		files, _ := ioutil.ReadDir("./")
		for _, f := range files {
			fmt.Printf("%v\t",f.Name())
		}
		fmt.Println()
	}
	if(isAll){
		files, _ := ioutil.ReadDir("./")
		for _, f := range files {
			fmt.Printf("%v\t%v bytes\t%v\n",f.Name(),f.Size(),f.Mode())
		} 
	}
}


// pwd
func pwd(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Unable to fetch current directory!")
	}
	fmt.Printf("%v\n",dir)
}

// History

func pushToHistory(command string) {
	History = append(History,command)
}

func printHistory(){
	for i:=0; i<len(History); i++ {
		fmt.Printf("%v   %v\n",i,History[i])
	} 
}
