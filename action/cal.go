package action

import (
	"fmt"
	"net/http"
)

func CalAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cal")
}
