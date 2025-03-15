package services

import (
	"context"
	"database/sql"
	"errors"
	"event-planner/internal/entities"
	"event-planner/internal/models"
	"event-planner/internal/models/mocks"
	authMock "event-planner/pkg/auth/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_service_RegisterUser(t *testing.T) {

	modelMock := mocks.Model{}
	hashMocks := authMock.Auth{}
	modelMock.On("CreateUser", context.Background(), &entities.User{
		ID:       0,
		Name:     "",
		Email:    "test@example.com",
		Password: "password123",
	}, mock.Anything).Return(nil).Once()

	modelMock.On("CreateUser", context.Background(), &entities.User{
		ID:       0,
		Name:     "",
		Email:    "test@example.com",
		Password: "password123",
	}, mock.Anything).Return(errors.New("error")).Once()

	hashMocks.On("GenerateHash", "password123").Return([]byte("some-hash"), nil).Once()
	hashMocks.On("GenerateHash", "password123").Return([]byte(""), errors.New("some-error")).Once()
	hashMocks.On("GenerateHash", "password123").Return([]byte("some-hash"), nil).Once()

	type fields struct {
		models models.Model
		auth   *authMock.Auth
	}
	type args struct {
		ctx  context.Context
		user *entities.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful registration",
			fields: fields{
				models: &modelMock,
				auth:   &hashMocks,
			},
			args: args{
				ctx: context.Background(),
				user: &entities.User{
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			wantErr: false,
		},
		{
			name: "error generating password hash",
			fields: fields{
				models: &modelMock,
				auth:   &hashMocks,
			},
			args: args{
				ctx: context.Background(),
				user: &entities.User{
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			wantErr: true,
		},
		{
			name: "error creating user",
			fields: fields{
				models: &modelMock,
				auth:   &hashMocks,
			},
			args: args{
				ctx: context.Background(),
				user: &entities.User{
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				models: tt.fields.models,
				auth:   tt.fields.auth,
			}
			if err := s.RegisterUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_AuthenticateUser(t *testing.T) {
	modelMock := mocks.Model{}
	authMocks := authMock.Auth{}

	// case 1
	modelMock.On("GetUserByEmail", context.Background(), "test@example.com").Return(&entities.User{
		Email: "test@example.com",
	}, "hashed-pass", nil).Once()
	authMocks.On("CompareHash", mock.Anything, []byte("pass")).
		Return(true, nil).
		Once()
	authMocks.On("GenerateJWTToken", mock.Anything).
		Return("some-jwt-token", nil).
		Once()

	// case 2
	modelMock.On("GetUserByEmail", context.Background(), "invalid@example.com").Return(nil, "", sql.ErrNoRows).Once()

	// case 3
	modelMock.On("GetUserByEmail", context.Background(), "error@example.com").Return(nil, "", errors.New("error")).Once()

	// case 4
	modelMock.On("GetUserByEmail", context.Background(), "test@example.com").Return(&entities.User{
		Email: "test@example.com",
	}, "wrongpassword", nil).Once()
	authMocks.On("CompareHash", mock.Anything, []byte("wrongpassword")).
		Return(false, nil).
		Once()

	// case 5
	modelMock.On("GetUserByEmail", context.Background(), "test@example.com").Return(&entities.User{
		Email: "test@example.com",
	}, "hashed-pass", nil).Once()
	authMocks.On("CompareHash", mock.Anything, []byte("pass")).
		Return(true, nil).
		Once()
	authMocks.On("GenerateJWTToken", mock.Anything).
		Return("", errors.New("failed to generate jwt")).
		Once()

	type fields struct {
		models models.Model
		auth   *authMock.Auth
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successful authentication",
			fields: fields{
				models: &modelMock,
				auth:   &authMocks,
			},
			args: args{
				ctx:      context.Background(),
				email:    "test@example.com",
				password: "pass",
			},
			want:    "some-jwt-token",
			wantErr: false,
		},
		{
			name: "user not found",
			fields: fields{
				models: &modelMock,
				auth:   &authMocks,
			},
			args: args{
				ctx:      context.Background(),
				email:    "invalid@example.com",
				password: "pass",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error getting user",
			fields: fields{
				models: &modelMock,
				auth:   &authMocks,
			},
			args: args{
				ctx:      context.Background(),
				email:    "error@example.com",
				password: "pass",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid credentials",
			fields: fields{
				models: &modelMock,
				auth:   &authMocks,
			},
			args: args{
				ctx:      context.Background(),
				email:    "test@example.com",
				password: "wrongpassword",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed to create jwt",
			fields: fields{
				models: &modelMock,
				auth:   &authMocks,
			},
			args: args{
				ctx:      context.Background(),
				email:    "test@example.com",
				password: "pass",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				models: tt.fields.models,
				auth:   tt.fields.auth,
			}
			got, err := s.AuthenticateUser(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.AuthenticateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
