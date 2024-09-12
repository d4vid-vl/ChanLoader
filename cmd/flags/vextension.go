package flags

import (
	"fmt"
	"strings"
)

type VExtension string

// These are all the current frameworks supported. If you want to add one, you
// can simply copy and past a line here. Do not forget to also add it into the
// AllowedProjectTypes slice too!
const (
	MP4  VExtension = "mp4"
	AVI  VExtension = "avi"
	WEBM VExtension = "webm"
)

var AllowedVExtensions = []string{string(MP4), string(AVI), string(WEBM)}

func (f VExtension) String() string {
	return string(f)
}

func (f *VExtension) Type() string {
	return "VExtension"
}

func (f *VExtension) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedProjectTypes.Contains(value) {
	for _, project := range AllowedVExtensions {
		if project == value {
			*f = VExtension(value)
			return nil
		}
	}

	return fmt.Errorf("file image extension to use. Allowed values: %s", strings.Join(AllowedVExtensions, ", "))
}
