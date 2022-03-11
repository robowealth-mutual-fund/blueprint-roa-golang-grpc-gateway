package httpServer

import "net/http"

func (s *Server) swager() {
	s.HttpMux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	s.HttpMux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))
}
func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger/product.swagger.json")
}
