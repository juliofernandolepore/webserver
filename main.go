package main

import (
	"log"
	"net/http"

	"github.com/juliofernandolepore/webserver/router"
	"github.com/juliofernandolepore/webserver/services"
	"github.com/juliofernandolepore/webserver/utils"
)

func main() {
	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
