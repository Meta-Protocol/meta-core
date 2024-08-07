package main

import (
	"os"
	"bytes"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func docsCmd(cmd *cobra.Command, args []string) error {
	var path string

	// Determine the output path
	if len(args) > 0 {
		path = args[0]
	} else {
		var err error
		path, err = cmd.Flags().GetString("path")
		if err != nil {
			return err
		}
	}

	// Create the output directory if it doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0750)
		if err != nil {
			return err
		}
	}

	// Set the output file
	outputFile := path + "/cli.md"
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Generate the documentation
	err = GenMarkdownToSingleFile(cmd.Root(), file)
	if err != nil {
		return err
	}
	return nil
}

// GenMarkdownToSingleFile generates markdown documentation for all commands into a single file
func GenMarkdownToSingleFile(cmd *cobra.Command, w *os.File) error {
	buf := new(bytes.Buffer)
	linkHandler := func(s string) string {
		anchor := strings.ReplaceAll(strings.ToLower(s), " ", "-")
		anchor = strings.ReplaceAll(anchor, "_", "-")
		return "#" + anchor
	}

	cmd.DisableAutoGenTag = true

	err := doc.GenMarkdownCustom(cmd, buf, linkHandler)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		err := GenMarkdownToSingleFile(c, w)
		if err != nil {
			return err
		}
	}

	return nil
}


func docsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docs [path]",
		Short: "Generate markdown documentation for zetacored",
		RunE:  docsCmd, 
		Args:  cobra.MaximumNArgs(1),
	}

	cmd.Flags().String("path", "docs/cli/zetacored", "Path where the docs will be generated")

	return cmd
}
