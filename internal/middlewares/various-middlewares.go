package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// Basically middleware is a normal function which accepts an handler and returns an handler.
// It can be used as handler chaining in api calls.
func SampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// write some logic

		next.ServeHTTP(w,r)
	})
}


// if we want to provide arguments, then we need to write a wrapper function on top of normal middleware format 
func SampleMiddlewareWithArgs(value string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// write some logic
	
			next.ServeHTTP(w,r)
		})
	}
}

func SampleEchoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// write some logic
		fmt.Println("Sample middleware called.....")
		return next(c)
	}
}

func SampleEchoMiddlewareWithArgs(arg string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// write some logic
			fmt.Println("args", arg)
			return next(c)
		}
	}
}

func SampleGinMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context)  {
		// write some logic
		
		next(c)
	}
}

func SampleGinMiddlewareWithArgs(arg string) func(next gin.HandlerFunc) gin.HandlerFunc {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context)  {
			// write some logic
			
			next(c)
		}
	}
}