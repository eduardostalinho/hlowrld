package server

import (
	"fmt"
	"net/http"

	"github.com/endocode/hlowrld/nicename"
)

// GetHandler handles all GET commands
func (s Server) GetHandler(w http.ResponseWriter, r *http.Request) {

	var res response

	res.Message = fmt.Sprintf("Hello %v", nicename.GeneratePair())

	res.Path = r.URL.RequestURI()
	res.Method = r.Method

	res.Status = http.StatusOK
	res.Error = nil

	s.writeResponse(w, res)
}
