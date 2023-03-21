package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

// A function to custom generate protobuf code
func main() {
	// Protoc passes pluginpb.CodeGeneratorRequest in via stdin
	// marshalled with Protobuf
	input, _ := ioutil.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	// Initialise our plugin with default options
	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		log.Println("File2: ", file)
		log.Println("Plugin2: ", plugin)
		/*
			if err := generateFile(plugin, file); err != nil {
				return err
			}
		*/
	}

	// create the response with all our files
	resp := plugin.Response()
	resp.File = append(resp.File, &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String("test.txt"),
		Content: proto.String("Hello World"),
	})
	out, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, string(out))
}
