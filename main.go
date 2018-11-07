package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

//go:generate go-bindata -pkg main templates/
func main() {
	t := template.New("templates")
	fs, err := AssetDir("templates")
	if err != nil {
		log.Fatal(err)
	}

	d, _ := os.Getwd()

	data := tpl{
		DockerRoot: "docker.astuart.co/astuart/gotest",
		Image:      path.Base(d),
		Name:       path.Base(d),
		Namespace:  "default",
	}

	for _, file := range fs {
		t2, err := t.New(file).Parse(string(MustAsset("templates/" + file)))
		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(file); err == nil {
			fmt.Println(file + " already exists; skipping")
			continue
		}

		f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0640)
		if err != nil {
			fmt.Printf("error opening file %q: %s\n", file, err)
			continue
		}

		err = t2.Execute(f, data)
		if err != nil {
			fmt.Println("error templating: ", err)
		}
	}

}

type tpl struct {
	DockerRoot string
	Image      string
	Name       string
	Namespace  string
}
