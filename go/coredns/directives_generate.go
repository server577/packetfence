//+build ignore

package main

import (
	"bufio"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	mi := make(map[string]string, 0)
	md := []string{}

	file, err := os.Open(pluginFile)
	if err != nil {
		log.Fatalf("Failed to open %s: %q", pluginFile, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		items := strings.Split(line, ":")
		if len(items) != 2 {
			// ignore empty lines
			continue
		}
		name, repo := items[0], items[1]

		if _, ok := mi[name]; ok {
			log.Fatalf("Duplicate entry %q", name)
		}

		md = append(md, name)
		mi[name] = pluginPath + repo // Default, unless overridden by 3rd arg

		if _, err := os.Stat(pluginFSPath + repo); err != nil { // External package has been given
			mi[name] = repo
		}
	}

	genImports("core/zplugin.go", "core", mi)
	genDirectives("core/dnsserver/zdirectives.go", "dnsserver", md)
}

func genImports(file, pack string, mi map[string]string) {
	outs := header + "package " + pack + "\n\n" + "import ("

	if len(mi) > 0 {
		outs += "\n"
	}

	outs += "// Include all plugins.\n"
	for _, v := range mi {
		outs += `_ "` + v + `"` + "\n"
	}
	outs += ")\n"

	if err := formatAndWrite(file, outs); err != nil {
		log.Fatalf("Failed to format and write: %q", err)
	}
}

func genDirectives(file, pack string, md []string) {

	outs := header + "package " + pack + "\n\n"
	outs += `
// Directives are registered in the order they should be
// executed.
//
// Ordering is VERY important. Every plugin will
// feel the effects of all other plugin below
// (after) them during a request, but they must not
// care what plugin above them are doing.

var directives = []string{
`

	for i := range md {
		outs += `"` + md[i] + `",` + "\n"
	}

	outs += "}\n"

	if err := formatAndWrite(file, outs); err != nil {
		log.Fatalf("Failed to format and write: %q", err)
	}
}

func formatAndWrite(file string, data string) error {
	res, err := format.Source([]byte(data))
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(file, res, 0644); err != nil {
		return err
	}
	return nil
}

const (
	pluginPath   = "github.com/inverse-inc/packetfence/go/coredns/plugin/"
	pluginFile   = "plugin.cfg"
	pluginFSPath = "plugin/" // Where the plugins are located on the file system
	header       = "// generated by directives_generate.go; DO NOT EDIT\n\n"
)
