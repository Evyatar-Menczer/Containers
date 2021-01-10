package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "bufio"
    "os"
    "os/exec"
    "strings"
    "strconv"
    "bytes"
)

type conf struct {
    Name string `yaml:"Name"`
    Amount int `yaml:"Amount"`
    Image string `yaml:"Image"`
}

func (c *conf) getConf(path string) *conf {
    //Creats configuration struct with the values of: Name, Amount, Image.
    //Return:
    //  conf struct

    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func deleteAllContainers(){
    //Deletes all runnig containers


    // cmd := exec.Command("docker","rm", "-f", "$(docker ps -a -q)")
    cmd := exec.Command("docker", "ps", "-a", "-q")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()

    if err != nil {
        fmt.Println(err,"err in delete")
    }

    arr := strings.Split(out.String(), "\n") 
    for _,id := range arr{
        cmd := exec.Command("docker", "rm", "-f", id)
        err := cmd.Run()
        if err != nil {
            fmt.Println(err,"err in delete")
        }
    }
}

func createAgents(c conf){
    //Creates 2 agents processes, passing each "agent" its workload boundarys,
    //name of the image and name of the container.
    


    //splits workload
    var workload = c.Amount/2
    var rightBroderAgent1 = strconv.Itoa(workload)
    var leftBorderAgent2 = strconv.Itoa(c.Amount- workload + 1)

    //invoke agent1. Args are the range of the containers numbers for names(0 : workload),
    // container's name and image. 
    cmd1 := exec.Command("./agent","1", rightBroderAgent1, c.Name, c.Image)
    err1 := cmd1.Run()
    if err1 != nil {
        fmt.Println(err1)
    }

    //invoke agent2. Args are the range of the containers numbers for names(workload+1 : Amount),
    // container's name and image.
    cmd2 := exec.Command("./agent",leftBorderAgent2, strconv.Itoa(c.Amount),c.Name, c.Image)
    err2 := cmd2.Run()
    if err2 != nil {
        log.Fatal(err2)
    }
    
    return
}

func parseCommand(text string) ([]string){
    //Parses text to command and args
    //Args:
    //    text: string, the full text user entered.
    //Returns:
    //    array of the arguments in command.

    res := strings.Split(text, " ") 
    return res
}

func main() {
    i := 0
    var c conf
    for i < 1{
        var args []string
        reader := bufio.NewReader(os.Stdin)
        text, _ := reader.ReadString('\n')
        args = parseCommand(text)
        switch args[0] {
        case "create":
            c.getConf(args[1])
            createAgents(c)
        case "delete":
            deleteAllContainers()
        default:
            fmt.Println("Unknown command, try again")
        }
    }


}