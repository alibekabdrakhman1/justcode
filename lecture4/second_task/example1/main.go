package main

import (
	"fmt"
	"sync"
)

// eto ochen plohoi primer, t.k. ocherednost ne garantirovan
type UserMap struct {
	mutex sync.Mutex
	users map[int]string
}

func NewUserMap() *UserMap {
	return &UserMap{
		users: make(map[int]string),
	}
}

func (u *UserMap) create(name string) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.users[len(u.users)] = name
}
func (u *UserMap) delete(id int) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	delete(u.users, id)
}
func (u *UserMap) update(id int, newName string) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.users[id] = newName
}
func (u *UserMap) get(id int) (string, bool) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	val, ok := u.users[id]
	return val, ok
}

func main() {
	usersMap := NewUserMap()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		name := "Yernur"
		usersMap.create(name)
		fmt.Println(name, "is created")
	}()
	go func() {
		defer wg.Done()
		name := "Imangali"
		usersMap.create(name)
		fmt.Println(name, "is created")
	}()
	go func() {
		defer wg.Done()
		name := "Satzhan"
		usersMap.create(name)
		fmt.Println(name, "is created")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		id := 0
		pastName, ok := usersMap.get(id)
		if ok != false {
			//ya znauy chto takoe fmt.Printf, prosto toropilsya ochen silno
			fmt.Println("user by id not founded", id)
		} else {
			usersMap.update(0, "Yermurat")
			newName, _ := usersMap.get(id)
			//ya znauy chto takoe fmt.Printf, prosto toropilsya ochen silno

			fmt.Println(pastName, "is updated to", newName)
		}
	}()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			name, ok := usersMap.get(id)
			if ok {
				//ya znauy chto takoe fmt.Printf, prosto toropilsya ochen silno
				fmt.Println(name, "by id", id)
			} else {
				fmt.Println("not found name by ", id)
			}
		}(i)
	}
	wg.Wait()
}
