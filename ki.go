
package main

import "fmt"
import "os"
import "os/exec"
import "regexp"


func runcmd(path, param string){
    cmd := exec.Command(path, param)
    _, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("StdoutPipe: " + err.Error())
        return
    }
    cmd.Start()
    cmd.Wait()
    os.Exit(1)
}

func main(){
    // myConfig := new(cf.Config)
    // myConfig.InitConfig(".kiconfig")
    // fmt.Println(myConfig.Read("browser", "path"))
    // fmt.Printf("%v", myConfig.Mymap)


    args := os.Args
    if len(args) < 2 {
        fmt.Println("please input path.")
        os.Exit(0)
    }

    param := args[1]

    // url
    isurl, _:= regexp.MatchString("^(http|https):", param)
    if isurl {
        runcmd("/opt/google/chrome/chrome", param)
    } else {
        // local file
        _, err := os.Stat(param)
        if err != nil {
            fmt.Println("No such file or directory: ", param)
            os.Exit(0)
        }

        // pdf, xlsx, doc, docx, ppt
        isoffice, _:= regexp.MatchString("(\\.pdf|\\.xlsx|\\.doc|\\.docx|\\.ppt)$", param)
        if isoffice {
            fmt.Println(isoffice)
            runcmd("/usr/bin/libreoffice", param)
        }
        // bmp, gif, jpeg, jpg, png
        isimg, _:= regexp.MatchString("(\\.bmp|\\.gif|\\.jpeg|\\.jpg|\\.png)", param)
        if isimg {
            runcmd("/usr/bin/eog", param)
        }
        // text file  .js, .php, .html
        istext, _:=regexp.MatchString("(\\.php|\\.js|\\.html)", param)
        if istext {
            runcmd("/usr/bin/atom", param)
        }

        runcmd("/usr/bin/nautilus", param)
    }
    os.Exit(0)
}
