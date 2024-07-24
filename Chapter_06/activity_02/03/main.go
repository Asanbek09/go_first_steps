package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrBadData := errors.New("some bad data")
	fmt.Printf("ErrBadData type: %T", ErrBadData)

	var (
		ErrbodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
		ErrHijacked        = errors.New("http: connection has been hijacked")
		ErrContentLength   = errors.New("http: wrote more than the declared Content-Length")
		ErrWriteAfterFlush = errors.New("unused")
	)
	fmt.Println(ErrbodyNotAllowed)
	fmt.Println(ErrHijacked)
	fmt.Println(ErrContentLength)
	fmt.Println(ErrWriteAfterFlush)
}
