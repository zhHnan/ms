package code_gen

import "testing"

func TestGenStruct(t *testing.T) {
	GenStruct("ms_project", "Project")
}

func TestGenProtoMessage(t *testing.T) {
	GenProtoMessage("ms_project", "Project")
}
