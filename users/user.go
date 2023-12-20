package users

import (
	"fmt"
	"time"

	"github.com/agudelozc/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Pablo", time.Now(), true)
	fmt.Println(u)
}
