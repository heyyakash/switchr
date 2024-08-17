package constants

import "fmt"

type Role int

const (
	Owner Role = iota
	Editor
	Viewer
)

func (r Role) String() string {
	switch r {
	case Owner:
		return "Owner"
	case Viewer:
		return "Viewer"
	case Editor:
		return "Editor"
	default:
		return fmt.Sprintf("Status(%d)", int(r))
	}
}
