package maintenance

import "errors"

func NeedsMaintenace(days int) (decision bool, err error) {
	if days < 0 {
		return false, errors.New("cannot accept less than zero days")
	}
	if days >= 30 {
		return true, nil
	} else {
		return false, nil
	}
}
