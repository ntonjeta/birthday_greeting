package greeting

type FriendRepository interface {
	Get() ([]Friend, error)
}
