package pkg

import "fmt"

type BadRequestException struct {
	Msj    string
	Detail string
}
type NotFoundException struct {
	Msj    string
	Detail string
}
type TimeoutException struct {
	Msj    string
	Detail string
}
type BadGatewayException struct {
	Msj    string
	Detail string
}
type InternalServerException struct {
	Msj    string
	Detail string
}

func (b BadRequestException) Error() string {
	return fmt.Sprintf("%s - %s", b.Msj, b.Detail)
}
func (b NotFoundException) Error() string {
	return fmt.Sprintf("%s - %s", b.Msj, b.Detail)
}
func (b TimeoutException) Error() string {
	return fmt.Sprintf("%s - %s", b.Msj, b.Detail)
}
func (b BadGatewayException) Error() string {
	return fmt.Sprintf("%s - %s", b.Msj, b.Detail)
}
func (b InternalServerException) Error() string {
	return fmt.Sprintf("%s - %s", b.Msj, b.Detail)
}
