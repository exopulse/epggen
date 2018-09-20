package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/exopulse/files"

	"github.com/exopulse/epggen/generator"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "epggen",
	Short: "Exopulse generics generator",
	Long:  "Exopulse generics generator.",
}

var mapCmd = &cobra.Command{
	Use:   "map [key:value]",
	Short: "Create typed map",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mapping := args[0]
		colonAt := strings.Index(mapping, ":")

		if colonAt <= 0 {
			fmt.Printf("expected mapping in format key:value")
			os.Exit(1)
		}

		key := mapping[:colonAt]
		value := mapping[colonAt+1:]
		name := value

		if strings.HasPrefix(name, "*") {
			name = name[1:]
		}

		typeMap := fmt.Sprintf("%sMap", name)
		file := fmt.Sprintf("%s_map.go", strings.ToLower(name))

		params := map[string]string{"Key": key, "Type": value, "TypeMap": typeMap, "Name": name}

		if err := generateAndDeploy("map", params, file); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Example: "epggen map int:Person",
}

func init() {
	rootCmd.AddCommand(mapCmd)
}

var gen *generator.Generator

// Execute executes CLI commands.
func Execute(generator *generator.Generator) {
	gen = generator

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateAndDeploy(template string, params map[string]string, file string) error {
	params["Package"] = os.Getenv("GOPACKAGE")

	content, err := gen.Generate(template, params)

	if err != nil {
		return err
	}

	if files.FileExists(file) {
		bytes, err := files.ReadBytes(file)

		if err != nil {
			return err
		}

		if contentsAreEqual(string(bytes), content) {
			return nil
		}
	}

	if err := ioutil.WriteFile(file, []byte(content), 0666); err != nil {
		return err
	}

	fmt.Printf("generated %s\n", file)

	return nil
}

func contentsAreEqual(c1, c2 string) bool {
	return removeSpaces(c1) == removeSpaces(c2)
}

func removeSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		return r
	}, s)
}
