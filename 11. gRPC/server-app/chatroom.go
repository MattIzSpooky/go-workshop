package main

import (
	"fmt"
	"slices"
	"sync"
	"workshop/grpc/generated/chat"
)

type chatRoom struct {
	sync.Mutex
	name  string
	users []user
}

func (c *chatRoom) GetName() string {
	return c.name
}

func (c *chatRoom) findUser(name string) *user {
	c.Lock()
	defer c.Unlock()
	usersIdx := slices.IndexFunc(c.users, func(u user) bool { return u.name == name })
	if usersIdx > -1 {
		user := &c.users[usersIdx]

		return user
	}

	return nil
}

func (c *chatRoom) addUser(name string) {
	c.Lock()
	defer c.Unlock()

	msgChannel := make(chan *chat.ChatMessage)

	c.users = append(c.users, user{
		name:       name,
		msgChannel: msgChannel,
	})
}

func (c *chatRoom) deleteUser(name string) error {
	c.Lock()
	defer c.Unlock()

	usersIdx := slices.IndexFunc(c.users, func(u user) bool { return u.name == name })

	if usersIdx == -1 {
		return fmt.Errorf("User with name [%s] cannot be deleted as it does not exist.", name)
	}

	user := &c.users[usersIdx]
	close(user.msgChannel)

	c.users = slices.Delete(c.users, usersIdx, usersIdx+1)

	return nil
}

func (c *chatRoom) isEmpty() bool {
	return len(c.users) == 0
}
