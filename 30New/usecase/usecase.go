package usecase

type Methods interface {
	Create(string, int) string
	Delete(int) (string, bool)
	GetFriends(int) (string, bool)
	AddFriend(int, int) (string, bool)
	UpdateAge(int, int) (string, bool)
}

type UseCase struct {
	meth Methods
}

func New(m Methods) *UseCase {
	return &UseCase{meth: m}
}

func (uc *UseCase) ContCreate(name string, age int) string {
	str := uc.meth.Create(name, age)
	return str
}

func (uc *UseCase) ContDelete(id int) (string, bool) {
	str, ok := uc.meth.Delete(id)
	return str, ok
}

func (uc *UseCase) ContGetFriends(id int) (string, bool) {
	str, ok := uc.meth.GetFriends(id)
	return str, ok
}

func (uc *UseCase) ContAddFriend(first, second int) (string, bool) {
	str, ok := uc.meth.AddFriend(first, second)
	return str, ok
}

func (uc *UseCase) ContUpdateAge(id, newage int) (string, bool) {
	str, ok := uc.meth.UpdateAge(id, newage)
	return str, ok
}
