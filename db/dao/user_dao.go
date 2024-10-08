package dao

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/o-ga09/note-app-backendapi/db/db"
	"github.com/o-ga09/note-app-backendapi/domain"
)

type userDao struct {
	query *db.Queries
}

func NewUserDao(d *sql.DB) *userDao {
	q := db.New(d)
	return &userDao{query: q}
}

func (dao *userDao) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	res, err := dao.query.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(res.UserID)
	if err != nil {
		return nil, err
	}
	user := domain.User{
		UserID:    userId,
		Username:  res.Email,
		Password:  res.Password,
		CreatedAt: res.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: res.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
	return &user, nil
}

func (dao *userDao) GetUsers(ctx context.Context) ([]domain.User, error) {
	res, err := dao.query.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := []domain.User{}
	for _, r := range res {
		userId, err := uuid.Parse(r.UserID)
		if err != nil {
			return nil, err
		}
		user := domain.User{
			UserID:    userId,
			Username:  r.Name,
			UserEmail: r.Email,
			Password:  r.Password,
			CreatedAt: r.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdatedAt: r.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		}
		users = append(users, user)
	}
	return users, nil
}

func (dao *userDao) CreateUser(ctx context.Context, user domain.User) error {
	_, err := dao.query.CreateUser(ctx, db.CreateUserParams{
		UserID:   user.UserID.String(),
		Name:     user.Username,
		Email:    user.Username,
		Password: user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (dao *userDao) UpdateUser(ctx context.Context, user domain.User) error {
	err := dao.query.UpdateUser(ctx, db.UpdateUserParams{
		UserID:   user.UserID.String(),
		Name:     user.Username,
		Email:    user.Username,
		Password: user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (dao *userDao) DeleteUserById(ctx context.Context, id string) error {
	err := dao.query.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
