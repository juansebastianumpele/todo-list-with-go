package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	userData := entity.User{}
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Find(&userData)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return userData, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	userData := entity.User{}
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Find(&userData)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return userData, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	// userData := entity.User{}
	err := r.db.WithContext(ctx).Create(&user)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	// userData := entity.User{}
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Where("id =?", id).Delete(&entity.User{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
