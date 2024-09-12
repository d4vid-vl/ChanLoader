package flags

import (
	"fmt"
	"strings"
)

type IExtension string

// These are all the current frameworks supported. If you want to add one, you
// can simply copy and past a line here. Do not forget to also add it into the
// AllowedProjectTypes slice too!
const (
	PNG  IExtension = "png"
	JPEG IExtension = "jpeg"
	WEBP IExtension = "webp"
)

var AllowedIExtensions = []string{string(PNG), string(JPEG), string(WEBP)}

func (f IExtension) String() string {
	return string(f)
}

func (f *IExtension) Type() string {
	return "IExtension"
}

func (f *IExtension) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedProjectTypes.Contains(value) {
	for _, project := range AllowedIExtensions {
		if project == value {
			*f = IExtension(value)
			return nil
		}
	}

	return fmt.Errorf("file image extension to use. Allowed values: %s", strings.Join(AllowedIExtensions, ", "))
}
