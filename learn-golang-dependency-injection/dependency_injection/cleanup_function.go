package dependency_injection

import "fmt"

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("File Close", f.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}

	return file, func() {
		file.Close()
	}
}

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Connection Close", c.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}

	return connection, func() {
		connection.Close()
	}
}
