package memory

import (
	"fmt"

	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/a2-ito/todo-backend/domain/repository"

	"github.com/google/uuid"
)

type userRepository struct {
	users map[string]*model.User
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		users: map[string]*model.User{},
	}
}

func newUser(id uuid.UUID, name string) *model.User {
	return &model.User{
		ID:   id,
		Name: name,
	}
}

func (r *userRepository) Hello() {
	fmt.Println("UserRepository Hello")
}

func (u *userRepository) FindById(id uuid.UUID) (*model.User, error) {
	fmt.Println("UserRepository FindById")
	for _, ml := range u.users {
		if ml.ID == id {
			return ml, nil
		}
	}
	//return nil, fmt.Errorf("Error: %s", "no memo data")
	return nil, nil
}

func (u *userRepository) FindAll() ([]*model.User, error) {
	fmt.Println("UserRepository FindAll")
	users := make([]*model.User, len(u.users))
	fmt.Println(u.users)
	i := 0
	for _, user := range u.users {
		fmt.Println(user.ID)
		users[i] = newUser(user.ID, user.Name)
		i++
	}
	//return u.users, nil
	return users, nil
}

func (u *userRepository) Create(user *model.User) (*model.User, error) {
	fmt.Println("UserRepository Create " + user.Name)
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	user.ID = id
	u.users[id.String()] = user
	return user, nil
}

/*
func (u *userRepository) Update(user model.User) (*model.User, error) {
  return nil, nil
}

func (u *userRepository) Delete(id uuid.UUID) (error) {
  result := []*domain.User{}
	for _, ul := range u.userList {
		if ul.ID != id {
      result = append(result, ul)
    }
  }
  return nil
}
*/
