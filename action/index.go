package action

import (
	"fmt"
	"net/http"
)

func IndexAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
}
