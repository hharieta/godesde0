package usuarios

import (
	"fmt"
	"time"

	"github.com/hharieta/godesde0/modelos"
)

func AltaUsuario() {
	usuario := new(modelos.User)
	usuario.AddUser(10, "Gatovsky", time.Now(), true)
	fmt.Println(usuario)
}
