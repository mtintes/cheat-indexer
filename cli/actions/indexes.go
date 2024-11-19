package actions

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mtintes/cheat-indexer/types"
	"github.com/spf13/afero"
)

var AppFs = afero.NewOsFs()
var Comments []string

func AddIndex(indexPathToAdd, configFilePath string) {
	config, err := types.ReadConfig(configFilePath)
	if err != nil {
		fmt.Println("Error reading config file")
		return
	}

	config.Repositories = append(config.Repositories, types.Repository{Location: indexPathToAdd})
	RunIndexer(indexPathToAdd)
	//types.WriteConfig(config, configFilePath)
}

func RunIndexer(indexPath string) {
	err := filepath.WalkDir(indexPath, ProcessFile)
	if err != nil {
		fmt.Println("Error walking directory")
	}

	fmt.Println(Comments)
}

func ProcessFile(path string, d os.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		return nil
	}

	extType := filepath.Ext(path)

	if extType != ".go" {
		return nil
	}

	file, err := AppFs.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if isComment(extType, text) {
			CollectComments(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
	}

	fmt.Println(path)
	return nil
}

func CollectComments(comment string) {
	Comments = append(Comments, comment)
}

func isComment(ext, line string) bool {
	if ext == ".go" {
		return strings.Contains(line, "//")
	}
	return false
}

func isCheatComment(line string) bool {
	return strings.Contains(line, "cheat")
}
