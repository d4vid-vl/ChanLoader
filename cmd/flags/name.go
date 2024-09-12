package flags

import (
	"fmt"
	"strings"
)

type Name string

// These are all the current frameworks supported. If you want to add one, you
// can simply copy and past a line here. Do not forget to also add it into the
// AllowedProjectTypes slice too!
const (
	Numeric    Name = "numeric"
	Alphabetic Name = "alphabetic"
	ID         Name = "id"
	Random     Name = "random"
	Original   Name = "original"
)

var AllowedFileNames = []string{string(Numeric), string(Alphabetic), string(ID), string(Random), string(Original)}

func (f Name) String() string {
	return string(f)
}

func (f *Name) Type() string {
	return "Name"
}

func (f *Name) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedProjectTypes.Contains(value) {
	for _, project := range AllowedFileNames {
		if project == value {
			*f = Name(value)
			return nil
		}
	}

	return fmt.Errorf("categorized file names to use. Allowed values: %s", strings.Join(AllowedFileNames, ", "))
}
