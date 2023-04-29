package control

type Service interface {
	CheckPrivilege(username string) bool
}
