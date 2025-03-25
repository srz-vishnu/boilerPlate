package api

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("insideeeeeeeeeee")

	fmt.Fprintf(w, "hello haiiii")
}
