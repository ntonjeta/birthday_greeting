package main

type FriendRepository interface {
	get() ([]Friend, error)
}
