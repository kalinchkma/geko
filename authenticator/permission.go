package authenticator

type Permission struct {
	Owner  [3]rune // rwx base permission of a model for Owner
	Group  [3]rune // rwx base permission of a model for Group user
	Others [3]rune // rwx base permission of a model for others users
}
