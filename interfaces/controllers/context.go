package controllers

// Context is a wrapper interface of echo.Context.
type Context interface {
	JSON(int, interface{}) error
}
