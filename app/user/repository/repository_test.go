package repository_test

// import (
// 	"context"
// 	"go-gin/app/user"
// 	"go-gin/app/user/repository"
// 	"go-gin/database/ent"
// 	"go-gin/database/ent/enttest"
// 	entUser "go-gin/database/ent/user"
// 	"testing"

// 	googleUUID "github.com/google/uuid"
// 	_ "github.com/mattn/go-sqlite3"
// 	"github.com/stretchr/testify/require"
// )

// type TestEnv struct {
// 	Client *ent.Client
// 	Repo   *repository.UserRepository
// 	Ctx    context.Context
// }

// // SetupTestEnv membuat environment untuk semua test
// func SetupTestEnv(t *testing.T) *TestEnv {
// 	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	t.Cleanup(func() {
// 		client.Close()
// 	})
// 	require.NoError(t, client.Schema.Create(context.Background()))
// 	ctx := context.Background()
// 	return &TestEnv{
// 		Client: client,
// 		Repo:   repository.NewUserRepository(client),
// 		Ctx:    ctx,
// 	}
// }

// func TestUserRepository(t *testing.T) {
// 	env := SetupTestEnv(t)

// 	t.Run("CreateUser_Success", func(t *testing.T) {
// 		form := &user.UserForm{
// 			Name:     "Test User",
// 			Email:    "test@example.com",
// 			Username: "testuser",
// 			Password: "securepassword",
// 		}

// 		err := env.Repo.CreateUser(env.Ctx, form)
// 		require.NoError(t, err)

// 		createdUser, err := env.Client.User.Query().Where(entUser.Username("testuser")).Only(env.Ctx)
// 		require.NoError(t, err)
// 		require.Equal(t, "Test User", createdUser.Name)
// 		require.Equal(t, "test@example.com", *createdUser.Email)
// 		require.Equal(t, "testuser", createdUser.Username)
// 	})

// 	t.Run("GetDetailUser_Success", func(t *testing.T) {
// 		// Menambahkan user ke database
// 		newUser, err := env.Client.User.Create().
// 			SetName("Test User").
// 			SetUsername("testuser").
// 			SetEmail("test@example.com").
// 			SetPassword("securepassword").
// 			Save(env.Ctx)
// 		require.NoError(t, err)

// 		// Memanggil fungsi GetDetailUser
// 		result, err := env.Repo.GetDetailUser(env.Ctx, newUser.ID)
// 		require.NoError(t, err)

// 		// Verifikasi hasil
// 		require.NotNil(t, result)
// 		require.Equal(t, newUser.ID, result.ID)
// 		require.Equal(t, "Test User", result.Name)
// 		require.Equal(t, "testuser", result.Username)
// 		require.Equal(t, "test@example.com", *result.Email)
// 	})

// 	t.Run("GetDetailUser_NotFound", func(t *testing.T) {
// 		// Gunakan ID yang tidak ada di database
// 		nonExistentID := googleUUID.New()

// 		// Memanggil fungsi GetDetailUser
// 		result, err := env.Repo.GetDetailUser(env.Ctx, nonExistentID)
// 		require.NoError(t, err)

// 		// Verifikasi hasil
// 		require.Nil(t, result)
// 	})
// }
