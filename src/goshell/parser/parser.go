package parser

import(
	"fmt"
	"os" 
	"io/ioutil"
	"strings"
	"path/filepath"
)

var History = make([]string,0)


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

	case "cd": 
		if len(result) > 1 {
			cd(strings.Join(result[1:]," "))
		} else {
			cd(".")
		} 
		
	case "history":
		printHistory()

	case "exit":
		return false

	default:
		fmt.Println("This is default") 
	}
	
	return true
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

func PushToHistory(command string) {
	History = append(History,command)
}

func printHistory(){
	for i:=0; i<len(History); i++ {
		fmt.Printf("[%v]   %v\n",i,History[i])
	} 
}

// Change Directory

func cd (path string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	var abs string
	if err != nil {
		fmt.Println("Runtime buffer error")
	}

	current_dir := strings.Split(dir,"/")
	change_dir := strings.Split(path,"/")
	if(len(change_dir) > 1) {
		if (change_dir[1] == current_dir[1]){
			abs = path
		}
	} else {
		abs = fmt.Sprintf("%v/%v",dir,path)
	}
	err_cd := os.Chdir(abs)
	if(err_cd != nil){
		fmt.Printf("Directory not present!!\n")
	}
		
}
