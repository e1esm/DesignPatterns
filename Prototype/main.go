package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name:     "Folder1",
		Children: []INode{file1, file2, file3},
	}
	folder2 := &Folder{
		name:     "Folder2",
		Children: []INode{folder1},
	}
	folder2.print("  ")

	clonedFolder := folder2.clone()
	clonedFolder.print("  ")
}
