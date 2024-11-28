package repository

// import (
// 	"context"
// 	"go-gin/app/user"
// 	"go-gin/app/user/repository"
// 	"go-gin/database/ent"
// 	"go-gin/database/ent/enttest" // Path untuk ent test utilities
// 	entUser "go-gin/database/ent/user"
// 	"testing"

// 	googleUUID "github.com/google/uuid"
// 	_ "github.com/mattn/go-sqlite3" // Import driver SQLite3
// 	"github.com/stretchr/testify/require"
// )

// func SetupTestDB(t *testing.T) (*ent.Client, context.Context) {
// 	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	t.Cleanup(func() {
// 		client.Close()
// 	})
// 	require.NoError(t, client.Schema.Create(context.Background()))
// 	ctx := context.Background()
// 	return client, ctx
// }

// func TestCreateUser(t *testing.T) {
// 	// Menggunakan builder SetupTestDB untuk setup client dan context
// 	client, ctx := SetupTestDB(t)

// 	repo := repository.NewUserRepository(client)

// 	t.Run("success", func(t *testing.T) {
// 		form := &user.UserForm{
// 			Name:     "Test User",
// 			Email:    "test@example.com",
// 			Username: "testuser",
// 			Password: "securepassword",
// 		}

// 		err := repo.CreateUser(ctx, form)
// 		require.NoError(t, err)

// 		createdUser, err := client.User.Query().Where(entUser.Username("testuser")).Only(ctx)
// 		require.NoError(t, err)
// 		require.Equal(t, "Test User", createdUser.Name)
// 		require.Equal(t, "test@example.com", *createdUser.Email)
// 		require.Equal(t, "testuser", createdUser.Username)
// 	})
// }

// func TestGetDetailUser(t *testing.T) {
// 	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	defer client.Close()

// 	require.NoError(t, client.Schema.Create(context.Background()))

// 	repo := repository.NewUserRepository(client)

// 	// Skenario Berhasil: Ketika user ditemukan berdasarkan ID yang valid
// 	t.Run("success", func(t *testing.T) {
// 		ctx := context.Background()

// 		// Menambahkan user ke dalam database
// 		newUser, err := client.User.Create().
// 			SetName("Test User").
// 			SetUsername("testuser").
// 			SetEmail("test@example.com").
// 			SetPassword("securepassword").
// 			Save(ctx)
// 		require.NoError(t, err)

// 		// Panggil fungsi GetDetailUser untuk mendapatkan user berdasarkan ID
// 		result, err := repo.GetDetailUser(ctx, newUser.ID)
// 		require.NoError(t, err)

// 		// Verifikasi hasil
// 		require.NotNil(t, result)
// 		require.Equal(t, newUser.ID, result.ID)
// 		require.Equal(t, "Test User", result.Name)
// 		require.Equal(t, "testuser", result.Username)
// 		require.Equal(t, "test@example.com", *result.Email)
// 	})

// 	// Skenario Gagal: Ketika user tidak ditemukan berdasarkan ID yang tidak ada
// 	t.Run("user_not_found", func(t *testing.T) {
// 		ctx := context.Background()

// 		// Gunakan ID yang tidak ada di database
// 		nonExistentID := googleUUID.New()

// 		// Panggil fungsi GetDetailUser untuk mendapatkan user berdasarkan ID yang tidak ada
// 		result, err := repo.GetDetailUser(ctx, nonExistentID)
// 		require.NoError(t, err)

// 		// Verifikasi hasil, harus nil karena user tidak ditemukan
// 		require.Nil(t, result)
// 	})
// }
