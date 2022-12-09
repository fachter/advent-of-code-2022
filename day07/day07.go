package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	name          string
	parent        *Folder
	files         map[string]File
	folders       map[string]*Folder
	totalFileSize int
}

type File struct {
	name string
	size int
}

func newFolder(name string, parent *Folder) *Folder {
	f := Folder{}
	f.name = name
	f.parent = parent
	f.files = map[string]File{}
	f.folders = map[string]*Folder{}
	f.totalFileSize = 0
	return &f
}

func main() {
	//processDeviceText("day07-test2.txt")
	//processDeviceText("day07-test.txt")
	processDeviceText("day07.txt")

}

func processDeviceText(fileName string) {
	file, _ := os.Open(fileName)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	//rootFolder := Folder{name: "/"}
	rootFolder := newFolder("/", nil)
	workingFolder := rootFolder
	for fileScanner.Scan() {
		line := fileScanner.Text()
		commands := strings.Fields(line)
		if commands[0] == "$" {
			if commands[1] == "cd" {
				cdTo := commands[2]
				if cdTo == ".." {
					workingFolder = (*workingFolder).parent
				} else if cdTo == "/" {
					workingFolder = rootFolder
				} else {
					createFolderIfNotExist(workingFolder, cdTo)
					folder, _ := workingFolder.folders[cdTo]
					workingFolder = folder
				}
			}
		} else {
			if commands[0] == "dir" {
				createFolderIfNotExist(workingFolder, commands[1])
			} else {
				fileSize, _ := strconv.Atoi(commands[0])
				fileN := commands[1]
				(*workingFolder).files[fileN] = File{fileN, fileSize}
				(*workingFolder).totalFileSize += fileSize
				updateSizeOfParents(workingFolder.parent, fileSize)
			}
		}
	}

	totalDiskSpace := 70_000_000
	installDiskSpace := 30_000_000
	requiredDiskSpace := installDiskSpace - (totalDiskSpace - rootFolder.totalFileSize)
	fmt.Println(requiredDiskSpace)

	smallFolders := map[string]int{}
	checkAndAddSmallFolders(rootFolder, &smallFolders, requiredDiskSpace)
	fmt.Println(smallFolders)
	sum := 100000000000
	for _, v := range smallFolders {
		if v < sum {
			sum = v
		}
	}
	print(sum)
}

func updateSizeOfParents(parent *Folder, size int) {
	if parent != nil {
		(*parent).totalFileSize += size
		updateSizeOfParents((*parent).parent, size)
	}
}

func checkAndAddSmallFolders(folder *Folder, folders *map[string]int, limit int) {
	totalFolderSize := folder.totalFileSize
	//for _, f := range folder.folders {
	//	totalFolderSize += f.totalFileSize
	//}
	if totalFolderSize >= limit {
		_, exists := (*folders)[folder.name]
		if exists {
			(*folders)[folder.name] += totalFolderSize
			fmt.Println("Exists")
		} else {
			(*folders)[folder.name] = totalFolderSize
		}
	}
	for _, f := range folder.folders {
		checkAndAddSmallFolders(f, folders, limit)
	}
}

func createFolderIfNotExist(workingFolder *Folder, cdTo string) {
	_, exists := workingFolder.folders[cdTo]
	if !exists {
		workingFolder.folders[cdTo] = newFolder(cdTo, workingFolder)
	}
}

// 1814928
// > 1899879
// < 9......

// 11915402
