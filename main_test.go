package main

import (
	builder "go-design-patterns/Creational/Builder"
	"testing"
)

func TestBuilder(t *testing.T) {
	bldr := builder.NewNotificationBuilder()

	bldr.SetTitle("Test Notification")
	bldr.SetIcon("icon.png")
	bldr.SetPriority(4)
	bldr.SetTxt("blah blah blah")

	_, err := bldr.Build()
	if err != nil {
		t.Error(err)
	}
}
