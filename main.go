package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
    fmt.Println("Server is running...")
    // time module
    time_now := time.Now()
    fmt.Println(time_now)

    // Echo instance
    e := echo.New()

    // Middleware
    e.Use(MW)
    
    // Routes
    // e.GET("/", hello)
    e.GET("/status", Handler)

    // Start server
    err := e.Start(":8080")
    if err != nil {
        log.Fatal(err)
    }
}

// func hello(ctx echo.Context) error {
//     return ctx.String(http.StatusOK, "Hello, World!")
// }

func Handler(ctx echo.Context) error {
    d := time.Date(2025, time.January, 0, 0, 0, 0, 0, time.UTC)
    dur := int64(time.Until(d).Hours())/24
    s := fmt.Sprintf("%d days left.", dur)
    err := ctx.String(http.StatusOK, s)
    if err != nil {
        return err
    }
    return nil
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        val := ctx.Request().Header.Get("User-Role")
        
        if val == "admin" {
            log.Println("red button user detected")
        }
        err := next(ctx)
        if err != nil {
            return err
        }
        return nil
    }
}
