package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open tags.txt
	file, err := os.Open("tags.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the Go module from the first line
	if !scanner.Scan() {
		fmt.Println("Error: tags.txt is empty.")
		return
	}
	goModule := scanner.Text()

	// Process each subsequent line in the file
	for scanner.Scan() {
		line := scanner.Text()
		tagNames := strings.Split(line, " ")
		if len(tagNames) != 2 {
			fmt.Printf("Invalid format in line: %s\n", line)
			continue
		}

		tagCamel := tagNames[0]
		tagSnake := tagNames[1]
		err := createDomainFile(goModule, tagCamel, tagSnake)
		if err != nil {
			fmt.Printf("Error creating domain file for %s: %v\n", tagCamel, err)
		}

		err = createUsecaseFile(goModule, tagCamel, tagSnake)
		if err != nil {
			fmt.Printf("Error creating usecase file for %s: %v\n", tagCamel, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

func createDomainFile(goModule, tagCamel, tagSnake string) error {
	dir := "./internal/domain/"
	filename := fmt.Sprintf("%sapi_%s.go", dir, tagSnake)

	// Create directory if not exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Generate content
	content := fmt.Sprintf(`package domain

import "context"

type %sRequest struct {
}

type %sResponse struct {
}

type %sHandler interface {
	Handle(ctx context.Context, req *%sRequest) (res *%sResponse, err error)
}
`, tagCamel, tagCamel, tagCamel, tagCamel, tagCamel)

	// Write to file
	return os.WriteFile(filename, []byte(content), 0644)
}

func createUsecaseFile(goModule, tagCamel, tagSnake string) error {
	dir := fmt.Sprintf("./internal/usecase/%s/", strings.ReplaceAll(tagSnake, "_", "-"))
	filename := dir + "usecase.go"

	// Create directory if not exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Generate content
	content := fmt.Sprintf(`package usecase

import (
	"context"

	"%s/internal/domain"
)

type handler struct {
}

func New%sHandler() domain.%sHandler {
	return &handler{}
}

// Handle implements domain.%sHandler.
func (h *handler) Handle(ctx context.Context, req *domain.%sRequest) (res *domain.%sResponse, err error) {
	panic("unimplemented")
}
`, goModule, tagCamel, tagCamel, tagCamel, tagCamel, tagCamel)

	// Write to file
	return os.WriteFile(filename, []byte(content), 0644)
}
