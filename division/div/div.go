package div

import "errors"

func Division(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}

	return a / b, nil

}
