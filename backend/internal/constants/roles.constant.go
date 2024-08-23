package constants

var Role map[string]int

func LoadRoleConstants() {
	Role = map[string]int{
		"owner":  0,
		"editor": 1,
		"reader": 2,
	}
}
