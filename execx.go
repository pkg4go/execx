package execx

import "os/exec"

func Run(name string, arg ...string) (out string, err error) {
	data, err := exec.Command(name, arg...).Output()
	if err != nil {
		return err.Error(), err
	}

	return string(data[:]), nil
}
