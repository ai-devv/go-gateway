package oauth

import (
	"encoding/json"
	"fmt"
	"gateway/internal/repositories/state"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

func (c *Controller) Authorize(w http.ResponseWriter, r *http.Request) {
	var jsonBody map[string]any

	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	rawAuthorizeUrl, ok := jsonBody["authorizeUrl"].(string)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	authorizeUrl, err := url.Parse(rawAuthorizeUrl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	authorizeUrlQuery := authorizeUrl.Query()
	authorizeUrlRedirectUrl := authorizeUrlQuery.Get("redirect_uri")
	authorizeUrlState := authorizeUrlQuery.Get("state")
	// TODO ???
	redirectUrl := "http://127.0.0.1:3000/oauth/callback"
	// TODO to config?
	stateId := fmt.Sprintf("gateway-%s", uuid.New().String())

	authorizeUrlQuery.Set("redirect_uri", redirectUrl)
	authorizeUrlQuery.Set("state", stateId)

	authorizeUrl.RawQuery = authorizeUrlQuery.Encode()

	err = c.sr.Save(stateId, state.State{
		"redirectUrl": authorizeUrlRedirectUrl,
		"state":       authorizeUrlState,
	})

	if err != nil {
		panic(err)
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]any{
		"authorizeUrl": authorizeUrl.String(),
		"redirectUrl":  redirectUrl,
	})

	if err != nil {
		panic(err)
	}
}
