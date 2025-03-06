package router

import (
	"context"
	"go-gin/app/middleware"
	"go-gin/app/tools"
	"go-gin/database/ent"
	"go-gin/worker"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	exampleHandler "go-gin/app/example/handler"
	exampleRepo "go-gin/app/example/repository"
	exampleUC "go-gin/app/example/usecase"

	userHandler "go-gin/app/user/handler"
	userRepo "go-gin/app/user/repository"
	userUC "go-gin/app/user/usecase"

	authHandler "go-gin/app/auth/handler"
	authRepo "go-gin/app/auth/repository"
	authUC "go-gin/app/auth/usecase"

	questionHandler "go-gin/app/question/handler"
	questionRepo "go-gin/app/question/repository"
	questionUC "go-gin/app/question/usecase"

	formResponseHandler "go-gin/app/form_response/handler"
	formResponseRepo "go-gin/app/form_response/repository"
	formResponseUC "go-gin/app/form_response/usecase"

	historyAnswerHandler "go-gin/app/history_answer/handler"
	historyAnswerRepo "go-gin/app/history_answer/repository"
	historyAnswerUC "go-gin/app/history_answer/usecase"

	mailerHandler "go-gin/app/mailer/handler"
	mailerRepo "go-gin/app/mailer/repository"
	mailerUC "go-gin/app/mailer/usecase"
)

type Handlers struct {
	Ctx    context.Context
	DB     *ent.Client
	R      *gin.Engine
	Log    *logrus.Entry
	Gomail *gomail.Dialer
	Redis  *redis.Client
	Worker *worker.Worker
}

func (h *Handlers) Routes() {
	middleware.Add(h.R, middleware.CORSMiddleware())
	v1 := h.R.Group(os.Getenv("PREFIX_API"))

	h.R.Use(func(c *gin.Context) {
		go routine()
		c.Next()
	})

	v1.GET("/check-connection", h.CheckConnection)

	// Swagger
	v1.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Repository
	ExampleRepo := exampleRepo.NewExampleRepository(h.DB)
	UserRepo := userRepo.NewUserRepository(h.DB)
	QuestionRepo := questionRepo.NewQuestionRepository(h.DB)
	FormResponseRepo := formResponseRepo.NewFormResponseRepository(h.DB)
	HistoryAnswerRepo := historyAnswerRepo.NewHistoryAnswerRepository(h.DB)
	MailAnswerRepo := mailerRepo.NewEmailRepository(h.Gomail, h.Redis)
	authRepo := authRepo.NewAuthRepository(h.DB)

	// Usecase
	ExampleUC := exampleUC.NewExampleUsecase(ExampleRepo, h.Ctx)
	UserUC := userUC.NewUserUsecase(UserRepo, h.Ctx)
	questionUC := questionUC.NewQuestionUsecase(QuestionRepo, h.Ctx)
	formResponseUC := formResponseUC.NewFormResponseUsecase(FormResponseRepo, h.Ctx)
	HistoryAnswerUC := historyAnswerUC.NewHistoryAnswerUsecase(HistoryAnswerRepo, h.Ctx)
	MailerUC := mailerUC.NewEmailUsecase(MailAnswerRepo, h.Redis, h.Ctx)
	AuthUC := authUC.NewAuthUsecase(authRepo, UserRepo, h.Ctx, MailerUC)

	// Handler
	exampleHandler.ExampleRoute(ExampleUC, v1, h.Log)
	userHandler.UserRoute(UserUC, v1, h.Log)
	authHandler.AuthRoute(AuthUC, v1, h.Log)
	questionHandler.QuestionRoute(questionUC, v1, h.Log)
	formResponseHandler.FormResponseRoute(formResponseUC, v1, h.Log)
	historyAnswerHandler.HistoryAnswerRoute(HistoryAnswerUC, v1, h.Log)
	mailerHandler.MailerRoute(MailerUC, v1, h.Log)
}

func routine() {
	time.Sleep(1 * time.Second)
}

// @BasePath /api/v1
// @Router /check-connection [get]
// @Accept json
// @Produce json
func (h *Handlers) CheckConnection(c *gin.Context) {
	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "Success Check Connect to API",
	})
}
