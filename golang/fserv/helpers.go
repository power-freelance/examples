package fserv

import "os"

func checkDir(dir string, perm os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, perm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
