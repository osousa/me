package main

import(
    "html/template"
    "log"
    "io"
    "os"
)


type Control struct{
    name string
    template *template.Template
}

type Controller interface{
    Execute(io.Writer)
}

func (c Control) Execute(w io.Writer){
    c.template.ExecuteTemplate(w, "layout", struct{Title string}{Title:"ok"})
}

func NewControl(name string)(Control){
    wd, err := os.Getwd()
    if err != nil {
       log.Fatal(err)
    }
    tpl, err := template.ParseFiles(wd + "/html/"+name+".tpl")
    if err != nil{
        log.Fatalln("Make sure your NewControl parameter equals a template name: %s", err.Error())
    }
    c := Control{name, tpl}
    return c
}