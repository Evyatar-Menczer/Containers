package main

import(
	"fmt"
	"os/exec"
	"os"
	"strconv"
)

func main(){
	// Args[] structure : {FILE, left border, right border, container's name, image}
	
	left,_ := strconv.Atoi(os.Args[1])
	right,_ := strconv.Atoi(os.Args[2])

	//this loop creates the containers by the workload partition.
	for i := left; i <= right; i++{

		var name string = os.Args[3] + strconv.Itoa(i) 

		//first cmd1 runs the dokcer, cmd2 copies the binary file to its dir.
		cmd1 := exec.Command("docker", "run", "--name", name, "-it", "-d", os.Args[4])
		cmd2 := exec.Command("docker", "cp", "showtime" ,name + ":/showtime")
		
		err1 := cmd1.Run()
		if err1 != nil {
			fmt.Println(err1)
			os.Exit(2)
		}		
		err2 := cmd2.Run()
		if err2 != nil {
			fmt.Println(err2)
			os.Exit(2)
		}

		//cmd3 execute within the container the binary.
		cmd3 := exec.Command("docker", "exec" ,name , "./showtime")
		err3 := cmd3.Run()
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(2)
			}
	}

}
