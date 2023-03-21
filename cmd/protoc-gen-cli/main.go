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

	enums := make(map[string]*protogen.Enum)
	services := make(map[string]*protogen.Service)
	messages := make(map[string]*protogen.Message)
	svcNameCounts := make(map[string]int)
	msgNameCounts := make(map[string]int)
	enumNameCounts := make(map[string]int)

	incBaseName := func(baseNameCounts map[string]int, name string) int {
		if _, ok := baseNameCounts[name]; !ok {
			baseNameCounts[name] = 0
		}
		baseNameCounts[name]++
		return baseNameCounts[name]
	}

	// Process each file and extracts messages etc so they can be handled
	// when we come to generating the CLI file
	processFile := func(file *protogen.File) error {
		// log.Println("File2: ", file)
		for _, svc := range file.Services {
			fqn := fmt.Sprintf("%s.%s", file.GoPackageName, svc.GoName)
			log.Println("Found Service: ", fqn)
			if _, ok := services[fqn]; ok {
				log.Fatal("Duplicate service name: ", fqn)
			}
			services[fqn] = svc
			incBaseName(svcNameCounts, svc.GoName)
		}

		for _, msg := range file.Messages {
			fqn := fmt.Sprintf("%s.%s", file.GoPackageName, msg.GoIdent.GoName)
			log.Println("Found Message: ", fqn)
			if _, ok := messages[fqn]; ok {
				log.Fatal("Duplicate message name: ", fqn)
			}
			messages[fqn] = msg
			incBaseName(msgNameCounts, msg.GoIdent.GoName)
		}

		for _, enum := range file.Enums {
			fqn := fmt.Sprintf("%s.%s", file.GoPackageName, enum.GoIdent.GoName)
			log.Println("Found Enum: ", fqn)
			if _, ok := enums[fqn]; ok {
				log.Fatal("Duplicate enum name: ", fqn)
			}
			enums[fqn] = enum
			incBaseName(enumNameCounts, enum.GoIdent.GoName)
		}

		/*
			if err := generateFile(plugin, file); err != nil {
				return err
			}
		*/
		return nil
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		processFile(file)
	}

	// now go through each service and each method in the service and generate a corresponding CLI subcommand for it
	log.Println("All Enums: ", enums)
	log.Println("All Messages: ", messages)
	for fqn, svc := range services {
		log.Println("Generating Service: ", fqn, svc.GoName)

		// what should the sub command be?
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

/**
 * A CommandGroup is a group of commands that are related to each other and allow
 * subcommands to be added to them along with containing things like help strings
 * and other metadata.
 */
type ComamndGroup struct {
	Name       string
	Title      string
	HelpString string
}
