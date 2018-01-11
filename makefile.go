// A Golang Generate file to act as a Makefile.
package main

//go:generate go fmt ./...
//go:generate go run create_empty_file.go
//go:generate go run truncate_a_file.go
//go:generate go run get_file_info.go
//go:generate go run rename_and_move_file.go
//go:generate go run delete_a_file.go
