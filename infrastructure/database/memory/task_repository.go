package memory

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/a2-ito/todo-backend/domain/model"
	"github.com/a2-ito/todo-backend/domain/repository"
)

type taskRepository struct {
	tasks map[string]*model.Task
}

func NewTaskRepository() repository.TaskRepository {
	return &taskRepository{
		tasks: map[string]*model.Task{},
	}
}

/*
func newTask(
	id uuid.UUID,
	owner uuid.UUID,
	title string,
	status string,
	desc string,
	startDate time.Time,
	endDate time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *model.Task {
	return &model.Task{
		ID:          id,
		Owner:       owner,
		Title:       title,
		Status:      status,
		Description: desc,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
*/

func (u *taskRepository) FindById(id uuid.UUID) (*model.Task, error) {
	fmt.Println("TaskRepository FindById")
	for _, ml := range u.tasks {
		if ml.ID == id {
			return ml, nil
		}
	}
	//return nil, fmt.Errorf("Error: %s", "no memo data")
	return nil, nil
}

func (u *taskRepository) FindAllByUserId(id uuid.UUID) ([]*model.Task, error) {
	fmt.Println("TaskRepository FindAll")
	tasks := make([]*model.Task, len(u.tasks))
	fmt.Println(u.tasks)

	i := 0
	for _, ml := range u.tasks {
		if ml.Owner == id {
			tasks[i] = ml
		}
		i++
	}

	/*
		i := 0
		for _, task := range u.tasks {
			fmt.Println(task.ID)
			tasks[i] = newTask(
				task.ID,
				task.Owner,
				task.Title,
				task.Status,
				task.Description,
				task.StartDate,
				task.EndDate,
				task.CreatedAt,
				task.UpdatedAt,
			)
			i++
		}
	*/

	return tasks, nil
}

func (u *taskRepository) Create(task *model.Task) (*model.Task, error) {
	fmt.Println("TaskRepository Create " + task.Title)
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	task.ID = id
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	u.tasks[id.String()] = task
	return task, nil
}

func (u *taskRepository) Update(task *model.Task) (*model.Task, error) {
	task.UpdatedAt = time.Now()
	u.tasks[task.ID.String()] = task

	return task, nil
}

func (u *taskRepository) Delete(id uuid.UUID) error {
	// Todo
	/*
		result := []*model.Task{}
		for _, ul := range u.tasks {
			if ul.ID != id {
				result = append(result, ul)
			}
		}
	*/
	return nil
}
