package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	String string
	Int    int64
}

func parseEntry(s string) (*Entry, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return &Entry{}, err
	}
	return &Entry{
		String: s,
		Int:    i,
	}, nil
}

type EntryRangeFactory struct {
	Start string
	End   string
}

func (r *EntryRangeFactory) Create() ([]*Entry, error) {

	startEntry, err := parseEntry(r.Start)
	if err != nil {
		return []*Entry{}, err
	}
	result := []*Entry{startEntry}
	endEntry, err := parseEntry(r.End)
	if err != nil {
		return []*Entry{}, err
	}

	for i := startEntry.Int; i <= endEntry.Int; i++ {
		result = append(result, &Entry{
			String: strconv.FormatInt(i, 10),
			Int:    i,
		})
	}
	return result, nil
}

type EntryValidator struct {
}

func (v *EntryValidator) IsValid(entry *Entry) bool {
	if len(entry.String)%2 != 0 {
		return true
	}
	firstHalf := entry.String[:len(entry.String)/2]
	secondHalf := entry.String[len(entry.String)/2:]
	if firstHalf == secondHalf {
		return false
	}
	return true
}

func main() {
	file, err := os.ReadFile("/Users/mclarke/repos/adventofcode2025/two/part1/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n")[0]

	parts := strings.Split(input, ",")
	validator := EntryValidator{}
	invalidEntriesMap := map[string]int64{}
	for _, part := range parts {
		s := strings.Split(part, "-")
		start := s[0]
		end := s[1]
		rangeFactory := EntryRangeFactory{
			Start: start,
			End:   end,
		}
		entries, err := rangeFactory.Create()
		if err != nil {
			panic(err)
		}
		for _, entry := range entries {
			if !validator.IsValid(entry) {
				invalidEntriesMap[entry.String] = entry.Int
			}
		}
	}

	sum := int64(0)
	for _, entry := range invalidEntriesMap {
		sum += entry
	}
	fmt.Println(sum)
}
