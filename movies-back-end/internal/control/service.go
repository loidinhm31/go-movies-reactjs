package control

type Service interface {
	CheckPrivilege(username string) bool
	CheckUser(username string) (bool, bool)
	CheckAdminPrivilege(username string) bool
}
