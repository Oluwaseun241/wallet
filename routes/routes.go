package routes

import (
	"github.com/Oluwaseun241/wallet/auths"
	"github.com/Oluwaseun241/wallet/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
  // users routes
  app.Get("", controllers.Demo)
  app.Get("/api/users", auth.AuthMiddleware, controllers.GetUser)
  app.Get("/api/users/:userId", controllers.GetUserId)
  app.Put("/api/users/:userId", controllers.UpdateUser)
  app.Post("/api/auth/register", controllers.NewUser)
  //auth routes
  app.Post("/api/auth/login", controllers.LoginUser)
  app.Post("/api/auth/refresh-token", auth.AuthMiddleware,controllers.ResfreshToken)

  //wallet routes
  app.Get("/api/users/wallet/:userId", auth.AuthMiddleware, controllers.GetWallet)
  app.Post("/api/users/wallet", auth.AuthMiddleware, controllers.NewWallet)
  app.Post("/api/users/wallet/:wallet_number", auth.AuthMiddleware, controllers.FundWallet)
  app.Post("/api/users/wallet/withdraw/:wallet_number", auth.AuthMiddleware, controllers.WihdrawFund)
  app.Post("/api/users/wallet/transfer/:wallet_number", auth.AuthMiddleware,controllers.TransferFund)
}
