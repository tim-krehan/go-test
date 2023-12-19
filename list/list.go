package list

import (
	"errors"
	"log"
	"os"
)

const list_file_path = "list.json"

func get_raw_list() string {
	data, error_read := os.ReadFile(list_file_path)
	if error_read != nil && !errors.Is(error_read, os.ErrNotExist) {
		log.Fatalf("cant read file: %v", error_read)
	}
	if errors.Is(error_read, os.ErrNotExist) {
		os.Create(list_file_path)
	}

	return string(data)
}

func parse_raw_list() any {
	raw_list := get_raw_list()

	if raw_list == "" {
		
	}
	return raw_list
}

func Add(item string) error {
	list := parse_raw_list()
	log.Printf("%v", list)
	return nil
}
func Get(index int) string {
	return ""
}
func Remove(index int) {}
func All() []string {
	return nil
}
func Search(pattern string) []string {
	return nil
}
