package repo

import (
	"30New/entity"
	"strconv"
)

type Storage struct {
	Store map[int]*entity.User
}

func (r Storage) Create(name string, age int) string {
	newUser := entity.NewUser()
	newUser.Name = name
	newUser.Age = age
	r.Store[len(r.Store)] = newUser
	return "Пользователь id: " + strconv.Itoa(len(r.Store)) + " создан"
}

func (r Storage) Delete(id int) (string, bool) {
	if id > len(r.Store)-1 {
		return "Пользователя id: " + strconv.Itoa(id) + " не существует", false
	}
	name := r.Store[id].Name
	delete(r.Store, id)
	return "Пользователь id: " + name + " удален", true
}

func (r Storage) AddFriend(first, second int) (string, bool) {
	if first > len(r.Store)-1 {
		return "Пользователь " + strconv.Itoa(first) + " не существует", false
	}
	if second > len(r.Store)-1 {
		return "Пользователь " + strconv.Itoa(second) + " не существует", false
	}
	r.Store[first].Friends = append(r.Store[first].Friends, second)
	r.Store[second].Friends = append(r.Store[second].Friends, first)
	return r.Store[first].Name + " и " + r.Store[second].Name + " теперь друзья", true
}

func (r Storage) GetFriends(id int) (string, bool) {
	if id > len(r.Store)-1 {
		return "Пользователь " + strconv.Itoa(id) + " не существует", false
	}
	result := "Друзья пользователя" + strconv.Itoa(id) + ":\n"
	for _, iduser := range r.Store[id].Friends {
		result += "Name is: " + r.Store[iduser].Name + " and age: " + strconv.Itoa(r.Store[iduser].Age)
	}
	return result, true
}

func (r Storage) UpdateAge(id, newade int) (string, bool) {
	if id > len(r.Store)-1 {
		return "Пользователь " + strconv.Itoa(id) + " не существует", false
	}
	r.Store[id].Age = newade
	return "Возраст пользователя" + strconv.Itoa(id) + " успешно обновлен", true
}

func New() *Storage {
	return &Storage{Store: make(map[int]*entity.User)}
}
