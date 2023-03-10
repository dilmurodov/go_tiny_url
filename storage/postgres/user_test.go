package postgres

import (
	"context"
	pb "go_auth_api_gateway/genproto/auth_service"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func createUser(t *testing.T) (resp *pb.User) {

	repo := NewUserRepo(db)

	respCreateUser, err := repo.Create(context.Background(), &pb.CreateUserRequest{
		FirstName: fakeData.FirstName(),
		LastName:  fakeData.LastName(),
		Phone:     fakeData.PhoneNumber(),
		Username:  fakeData.UserName(),
		Password:  "123456",
	})

	resp, err = repo.GetByPK(context.Background(), &pb.UserPrimaryKey{
		Id: respCreateUser.GetId(),
	})
	assert.NoError(t, err)

	assert.NoError(t, err)
	return resp
}

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name    string
		give    *pb.CreateUserRequest
		wantErr error
	}{
		{
			name: "SUCCESS: Create user",
			give: &pb.CreateUserRequest{
				FirstName: "Said",
				LastName:  "Amir",
				Phone:     fakeData.PhoneNumber(),
				Username:  fakeData.UserName(),
			},
			wantErr: nil,
		},
		{
			name: "ERROR: Create user, Dublicate phone",
			give: &pb.CreateUserRequest{
				FirstName: "Said",
				LastName:  "Amir",
				Phone:     fakeData.PhoneNumber(),
				Username:  fakeData.UserName(),
			},
			wantErr: nil,
		},
	}

	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.Create(context.Background(), tt.give)
			assert.NoError(t, err)
		})
	}
}

func TestGetUserByPK(t *testing.T) {

	tests := []struct {
		name string
		give *pb.UserPrimaryKey
		want error
	}{
		{
			name: "SUCCESS: Get user",
			give: &pb.UserPrimaryKey{
				Id: createUser(t).GetId(),
			},
			want: nil,
		},
		{
			name: "ERROR: Get user, User not found",
			give: &pb.UserPrimaryKey{
				Id: createRandomId(t),
			},
			want: pgx.ErrNoRows,
		},
	}

	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.GetByPK(context.Background(), tt.give)
			assert.Equal(t, tt.want, err)
		})
	}

}

func TestUpdateUser(t *testing.T) {

	dublicateUsername := fakeData.UserName()

	tests := []struct {
		name string
		give *pb.UpdateUserRequest
		want error
	}{
		{
			name: "SUCCESS: Update user",
			give: &pb.UpdateUserRequest{
				Id:        createUser(t).GetId(),
				FirstName: fakeData.FirstName(),
				LastName:  fakeData.LastName(),
				Phone:     fakeData.PhoneNumber(),
				Username:  fakeData.UserName(),
			},
			want: nil,
		},
		{
			name: "SUCCESS: Update user,",
			give: &pb.UpdateUserRequest{
				Id:        createUser(t).GetId(),
				FirstName: fakeData.FirstName(),
				LastName:  fakeData.LastName(),
				Phone:     fakeData.PhoneNumber(),
				Username:  dublicateUsername,
			},
			want: nil,
		},
		{
			name: "ERROR: Update user, Dublicate user",
			give: &pb.UpdateUserRequest{
				Id:        createUser(t).GetId(),
				FirstName: fakeData.FirstName(),
				LastName:  fakeData.LastName(),
				Phone:     fakeData.PhoneNumber(),
				Username:  dublicateUsername,
			},
			want: nil,
		},
	}
	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.Update(context.Background(), tt.give)
			assert.NoError(t, err)
		})
	}
}

func TestDeleteUser(t *testing.T) {

	tests := []struct {
		name string
		give *pb.UserPrimaryKey
		want error
	}{
		{
			name: "SUCCESS: Delete user",
			give: &pb.UserPrimaryKey{
				Id: createUser(t).GetId(),
			},
			want: nil,
		},
		{
			name: "ERROR: Delete user, User not found",
			give: &pb.UserPrimaryKey{
				Id: createRandomId(t),
			},
			want: pgx.ErrNoRows,
		},
	}

	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.Delete(context.Background(), tt.give)
			assert.Equal(t, tt.want, err)
		})
	}

}

func TestGetByeUsername(t *testing.T) {

	tests := []struct {
		name string
		give string
		want error
	}{
		{
			name: "SUCCESS: Get user",
			give: createUser(t).GetUsername(),
			want: nil,
		},
		{
			name: "ERROR: Get user, User not found",
			give: fakeData.UserName(),
			want: pgx.ErrNoRows,
		},
	}

	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.GetByUsername(context.Background(), tt.give)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestResetPassword(t *testing.T) {

	tests := []struct {
		name string
		give *pb.ResetPasswordRequest
		want error
	}{
		{
			name: "SUCCESS: Update password",
			give: &pb.ResetPasswordRequest{
				UserId:   createUser(t).GetId(),
				Password: "12345678",
			},
			want: nil,
		},
		{
			name: "ERROR: Update user, User not found",
			give: &pb.ResetPasswordRequest{
				UserId:   createRandomId(t),
				Password: "12345678",
			},
			want: pgx.ErrNoRows,
		},
	}

	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.ResetPassword(context.Background(), tt.give)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestGetlist(t *testing.T) {

	tests := []struct {
		name string
		give *pb.GetUserListRequest
		want error
	}{
		{
			name: "SUCCESS: Get user list",
			give: &pb.GetUserListRequest{
				Offset: 0,
				Limit:  10,
			},
			want: nil,
		},
	}

	userRepo := NewUserRepo(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userRepo.GetList(context.Background(), tt.give)
			assert.NoError(t, err)
		})
	}
}
