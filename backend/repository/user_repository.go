package repository

import (
	"errors"
	"event-management-backend/domain"
	"event-management-backend/infrastructure/hasura"
)

type UserRepository struct {
	hasura *hasura.Client
}

func NewUserRepository(hasura *hasura.Client) *UserRepository {
	return &UserRepository{hasura: hasura}
}

// Create - Used by signup action
func (r *UserRepository) Create(email, passwordHash, fullName string) (*domain.User, error) {
	query := `
		mutation($email: String!, $password_hash: String!, $full_name: String!) {
			insert_users_one(object: {
				email: $email
				password_hash: $password_hash
				full_name: $full_name
			}) {
				id email full_name phone avatar_url created_at updated_at
			}
		}
	`

	variables := map[string]interface{}{
		"email":         email,
		"password_hash": passwordHash,
		"full_name":     fullName,
	}

	var result struct {
		InsertUsersOne domain.User `json:"insert_users_one"`
	}

	err := r.hasura.Mutate(query, variables, &result)
	return &result.InsertUsersOne, err
}

// FindByEmail - Used by login action
func (r *UserRepository) FindByEmail(email string) (*domain.User, string, error) {
	query := `
		query($email: String!) {
			users(where: {email: {_eq: $email}}, limit: 1) {
				id email password_hash full_name phone avatar_url created_at updated_at
			}
		}
	`

	variables := map[string]interface{}{"email": email}

	var result struct {
		Users []struct {
			domain.User
			PasswordHash string `json:"password_hash"`
		} `json:"users"`
	}

	err := r.hasura.Mutate(query, variables, &result)
	if err != nil || len(result.Users) == 0 {
		return nil, "", err
	}

	user := result.Users[0]
	return &user.User, user.PasswordHash, nil
}

// FindByID - Used by payment usecase to get real user info
func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	query := `
		query($id: uuid!) {
			users_by_pk(id: $id) {
				id email full_name phone avatar_url created_at updated_at
			}
		}
	`

	variables := map[string]interface{}{"id": id}

	var result struct {
		UsersByPk *domain.User `json:"users_by_pk"`
	}

	err := r.hasura.Mutate(query, variables, &result)
	if err != nil {
		return nil, err
	}

	if result.UsersByPk == nil {
		return nil, errors.New("user not found")
	}

	return result.UsersByPk, nil
}
