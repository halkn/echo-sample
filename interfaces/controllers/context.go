package controllers

// Context is a wrapper interface of echo.Context.
type Context interface {
	Param(string) string
	JSON(int, interface{}) error
	Bind(interface{}) error
}
