package server

import(
   "net/http"
   "golang.org/x/oauth2"
)

func newServer(port int) *http.Server{
   port = strconv.Itoa(port)
   server := &http.Server{
      Addr: fmt.Sprintf(":%s", port)
   }
   print(server)
   return server
}

func (server *http.Server) startServer{
   server.ListenAndServe()
}
