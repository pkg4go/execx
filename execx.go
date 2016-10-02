package execx

import "os/exec"
import "strings"

func Run(name string, args ...string) (out string, err error) {
	data, err := exec.Command(name, args...).Output()
	if err != nil {
		return err.Error(), err
	}

	return string(data[:]), nil
}

// split command string safely
func Split(cmd string) []string {
	arr := strings.Split(strings.Trim(cmd, " "), " ")

	var result []string
	var inquote string
	var block string

	for _, v := range arr {
		if inquote == "" {
			if strings.HasPrefix(v, "'") || strings.HasPrefix(v, "\"") {
				inquote = string(v[0])
				block = strings.TrimPrefix(v, inquote) + " "
			} else if v != "" {
				result = append(result, v)
			}
		} else {
			if !strings.HasSuffix(v, inquote) {
				block += v + " "
			} else {
				block += strings.TrimSuffix(v, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}

	return result
}
