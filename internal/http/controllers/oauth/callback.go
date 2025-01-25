package oauth

import (
	"fmt"
	"net/http"
)

func (c *Controller) Callback(w http.ResponseWriter, r *http.Request) {
	stateId := r.URL.Query().Get("state")
	state, err := c.sr.Pull(r.Context(), stateId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	query := r.URL.Query()

	if originalState := state["state"]; originalState != "" {
		query.Set("state", originalState.(string))
	} else {
		query.Del("state")
	}

	http.Redirect(w, r, fmt.Sprintf("%s?%s", state["redirectUrl"], query.Encode()), http.StatusFound)
}
