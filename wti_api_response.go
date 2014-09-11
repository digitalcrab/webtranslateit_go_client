package webtranslateit_go_client

import "fmt"

type ErrorResponse struct {
	Message	string	`json:"error"`
	Request	string	`json:"request"`
}

func (self *ErrorResponse) Error() string {
	return fmt.Sprintf("wti: error on request %q: %s", self.Request, self.Message)
}

type ProjectResponse struct {
	Result	Project	`json:"project"`
}
