package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/Jamescog/gosh/utils"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Description() string {
	return "list directory contents"
}

func (c *LsCommand) Usage() string {
	return "ls [OPTIONS] [PATH]"
}

func (c *LsCommand) Examples() []Example {
	return []Example{
		{"ls", "list files in current directory"},
		{"ls -a", "list all files including hidden"},
		{"ls -l", "list files in long format"},
		{"ls -lh", "list files in long format with human-readable sizes"},
		{"ls /home", "list files in /home directory"},
	}
}

func (c *LsCommand) Execute(paths []string, flags []string) error {
	return handleLs(paths, flags)
}

func init() {
	Register(&LsCommand{})
}

func handleLs(paths []string, flags []string) error {
	fs := flag.NewFlagSet("ls", flag.ContinueOnError)
	showAll := fs.Bool("a", false, "do not ignore entries starting with .")
	showAllLong := fs.Bool("all", false, "do not ignore entries starting with .")
	longFormat := fs.Bool("l", false, "use a long listing format")
	humanReadable := fs.Bool("h", false, "with -l, print sizes in human readable format")
	longHuman := fs.Bool("lh", false, "use a long listing format with human readable sizes")

	if err := fs.Parse(flags); err != nil {
		return err
	}

	// Consolidate flag combinations
	if *longHuman {
		*longFormat = true
		*humanReadable = true
	}
	if *showAllLong {
		*showAll = true
		*longFormat = true
	}

	targetDir := "."
	if len(paths) > 0 {
		targetDir = paths[0]
	}

	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return fmt.Errorf("cannot access '%s': No such file or directory", targetDir)
	}

	for idx, entry := range entries {
		if shouldSkipHiddenFile(entry.Name(), *showAll) {
			continue
		}

		if *longFormat {
			fileInfo, err := entry.Info()
			if err != nil {
				return fmt.Errorf("error getting file info: %v", err)
			}
			printLongFormat(entry, fileInfo, *humanReadable)
		} else {
			printShortFormat(entry, idx, len(entries))
		}
	}

	if !*longFormat && len(entries)%5 != 0 {
		fmt.Println()
	}
	return nil
}

func shouldSkipHiddenFile(filename string, showAll bool) bool {
	return filename[0] == '.' && !showAll
}

func printLongFormat(entry os.DirEntry, info os.FileInfo, humanReadable bool) {
	name := formatName(entry)
	sizeStr := formatSize(info.Size(), humanReadable)
	modifiedTime := info.ModTime().Format("Jan 02 15:04")
	mode := info.Mode().String()

	output := fmt.Sprintf("%s %10s %s %s", mode, sizeStr, modifiedTime, name)
	fmt.Println(output)
}

func printShortFormat(entry os.DirEntry, index, totalEntries int) {
	name := formatName(entry)
	fmt.Print(name)

	if (index+1)%5 == 0 {
		fmt.Println()
	} else if index < totalEntries-1 {
		fmt.Print("\t")
	}
}

func formatName(entry os.DirEntry) string {
	name := entry.Name()
	if entry.IsDir() {
		return utils.Blue(name)
	}
	return utils.White(name)
}

func formatSize(size int64, humanReadable bool) string {
	if humanReadable {
		return utils.HumanReadableSize(size)
	}
	return fmt.Sprintf("%d", size)
}
