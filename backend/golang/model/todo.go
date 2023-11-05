package model

import (
	"context"
	"fmt"
	"time"

	"github.com/km1110/calendar-app/backend/golang/utils"
	"github.com/km1110/calendar-app/backend/golang/view/request"
	"github.com/km1110/calendar-app/backend/golang/view/response"
)

type TodoModel struct {
}

func NewTodoModel() *TodoModel {
	return &TodoModel{}
}

func (tm *TodoModel) GetTodos(user_id string) ([]*response.TodosResponse, error) {
	sql := `
				SELECT 
						t.id, 
						t.name, 
						t.date, 
						t.status, 
						p.id AS project_id, 
						p.title AS project_title, 
						tg.id AS tag_id, 
						tg.name AS tag_name
				FROM 
						todos t
				LEFT JOIN 
						projects p ON t.project_id = p.id
				LEFT JOIN 
						tags tg ON t.tag_id = tg.id
				WHERE 
						t.user_id = ?
				`

	rows, err := Db.Query(sql, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []*response.TodosResponse

	for rows.Next() {
		var (
			id, name, project_id, project_title, tag_id, tag_name string
			date                                                  time.Time
			status                                                bool
		)

		if err := rows.Scan(&id, &name, &date, &status, &project_id, &project_title, &tag_id, &tag_name); err != nil {
			return nil, err
		}

		todos = append(todos, &response.TodosResponse{
			Id:     id,
			Name:   name,
			Date:   date,
			Status: status,
			Project: response.ProjectsResponse{
				Id:    project_id,
				Title: project_title,
			},
			Tag: response.TagResponse{
				Id:   tag_id,
				Name: tag_name,
			},
		})
	}

	return todos, nil
}

func (tm *TodoModel) GetYearCount(user_id string, start_year string, end_year string) ([]*response.TodoCountResponse, error) {
	getQuery := `	SELECT 
										DATE(date) AS date, COUNT(*) AS count 
								FROM 
										todos 
								WHERE 
										user_id = ? AND status = true AND date >= ? AND date < ? 
								GROUP BY 
										DATE(date) 
								ORDER BY 
										DATE(date)
							`

	rows, err := Db.Query(getQuery, user_id, start_year, end_year)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*response.TodoCountResponse

	for rows.Next() {
		var (
			date  time.Time
			count int
		)

		if err := rows.Scan(&date, &count); err != nil {
			return nil, err
		}

		res = append(res, &response.TodoCountResponse{
			Date:  date,
			Count: count,
		})
	}

	fmt.Println(res)

	return res, nil
}

func (tm *TodoModel) AddTodos(ctx context.Context, user_id string, r request.CreateTodoRequest) (response.TodosResponse, error) {
	// todoのidを生成
	id := utils.GenerateId()

	// responseを作成
	res := response.TodosResponse{
		Id:     id,
		Name:   r.Name,
		Date:   r.Date,
		Status: r.Status,
		Project: response.ProjectsResponse{
			Id:    r.Project.Id,
			Title: r.Project.Title,
		},
		Tag: response.TagResponse{
			Id:   r.Tag.Id,
			Name: r.Tag.Name,
		},
	}

	// sqlの作成と実行
	insertQuery := `INSERT INTO todos(id, user_id, name, date, status, project_id, tag_id) VALUES(?, ?, ?, ?, ?, ?, ?)`

	_, err := Db.Exec(insertQuery, res.Id, user_id, res.Name, res.Date, res.Status, res.Project.Id, res.Tag.Id)

	if err != nil {
		return response.TodosResponse{}, err
	}

	return res, nil
}

func (tm *TodoModel) UpdateTodos(ctx context.Context, r response.TodosResponse) (response.TodosResponse, error) {
	// responseを作成
	res := response.TodosResponse{
		Id:     r.Id,
		Name:   r.Name,
		Date:   r.Date,
		Status: r.Status,
		Project: response.ProjectsResponse{
			Id:    r.Project.Id,
			Title: r.Project.Title,
		},
		Tag: response.TagResponse{
			Id:   r.Tag.Id,
			Name: r.Tag.Name,
		},
	}

	// sqlの作成と実行
	updateQuery := `UPDATE todos SET name = ?, date = ?, status = ?, project_id = ?, tag_id = ? WHERE id = ?`

	_, err := Db.Exec(updateQuery, res.Name, res.Date, res.Status, res.Project.Id, res.Tag.Id, res.Id)

	if err != nil {
		return response.TodosResponse{}, err
	}

	return res, nil
}

func (tm *TodoModel) UpdateTodoStatus(ctx context.Context, id string, r request.UpdateTodoStatusRequest) (response.TodosResponse, error) {
	// responseを作成
	res := response.TodosResponse{
		Id:     id,
		Status: r.Status,
	}

	// sqlの作成と実行
	updateQuery := `UPDATE todos SET status = ? WHERE id = ?`

	_, err := Db.Exec(updateQuery, res.Status, res.Id)

	if err != nil {
		return response.TodosResponse{}, err
	}

	return res, nil
}

func (tm *TodoModel) DeleteTodo(ctx context.Context, id string) error {
	// sqlの作成と実行
	deleteQuery := `DELETE FROM todos WHERE id = ?`

	_, err := Db.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	return nil
}
