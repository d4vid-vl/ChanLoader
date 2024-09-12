// Package steps provides utility for creating
// each step of the CLI
package steps

import (
	"ChanLoader/cmd/flags"
)

// A StepSchema contains the data that is used
// for an individual step of the CLI
type StepSchema struct {
	StepName string // The name of a given step
	Options  []Item // The slice of each option for a given step
	Headers  string // The title displayed at the top of a given step
	Field    string
}

// Steps contains a slice of steps
type Steps struct {
	Steps map[string]StepSchema
}

// An Item contains the data for each option
// in a StepSchema.Options
type Item struct {
	Title, Desc string
}

// InitSteps initializes and returns the *Steps to be used in the CLI program
func InitSteps(imageType flags.IExtension, videoType flags.VExtension, nameType flags.Name) *Steps {
	steps := &Steps{
		map[string]StepSchema{
			"i_extension": {
				StepName: "File Extension for images",
				Options: []Item{
					{
						Title: "PNG",
						Desc:  "The standard in the internet, supports transparency",
					},
					{
						Title: "JPEG",
						Desc:  "Lossy Compression, lightweight, doesn't support transparency",
					},
					{
						Title: "WEBP",
						Desc:  "Extremely lightweight format without losing compression, might not work in your OS",
					},
				},
				Headers: "What file extension do you want to use for images?",
				Field:   imageType.String(),
			},
			"v_extension": {
				StepName: "File Extension for videos",
				Options: []Item{
					{
						Title: "MP4",
						Desc:  "Standard video format, kinda lightweight",
					},
					{
						Title: "AVI",
						Desc:  "Better quality, heavier weight, and native to Windows",
					},
					{
						Title: "WEBM",
						Desc:  "4Chan's favorite video format, extremely lightweight as webp, but limited support",
					},
				},
				Headers: "What file extension do you want to use for videos?",
				Field:   videoType.String(),
			},
			"name": {
				StepName: "File Names",
				Options: []Item{
					{
						Title: "Numeric",
						Desc:  "Number every file image",
					},
					{
						Title: "Alphabetic",
						Desc:  "Follows the alphabet like Excel's columns"},
					{
						Title: "Random",
						Desc:  "Random file names, like a Password Gen"},
					{
						Title: "ID",
						Desc:  "4Chan's id system"},
					{
						Title: "Original",
						Desc:  "Original file name submitted in 4Chan",
					},
				},
				Headers: "What file names should the images/videos follow?",
				Field:   nameType.String(),
			},
		},
	}

	return steps
}
