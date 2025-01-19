package ownerauth

const (
	defaultOwnerUsername = "meet"
	defaultOwnerPassword = "meet"
)

func DoOwnerAuth(username, password string) bool {
	return defaultOwnerUsername == username && defaultOwnerPassword == password
}
