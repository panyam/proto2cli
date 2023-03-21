package main

import (
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

// A function to custom generate protobuf code
func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			log.Println("File: ", file)
			log.Println("Plugin: ", plugin)
			/*
				if err := generateFile(plugin, file); err != nil {
					return err
				}
			*/
		}

		return nil
	})
}
