package algorithms

import "errors"

// Todo make this a method on the dominoes struct like below
func DominoEffect(in []int) []int {
	var out []int
	vals := make(map[int]int, 0)
	for i, v := range in {

		vals[i] = v

		if v == 0 {
			out = append(out, i)
		}
		if in[i] > 0 {
			ran := i + in[i]
			if contains(in, ran) {
				if in[ran] > 0 {
					ran = ran + in[ran]
				}
			}
			out = append(out, ran)
		}
	}
	return out
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type dominoes struct {
	Dominoes []int
}

func (d *dominoes) Add(value int) []int {
	d.Dominoes = append(d.Dominoes, value)
	return d.Dominoes
}

func (d *dominoes) Update(position int, value int) ([]int, error) {
	if len(d.Dominoes) < position {
		return nil, errors.New("Position does not exist")
	}
	d.Dominoes[position] = value
	return d.Dominoes, nil
}

func (d *dominoes) Delete(position int) ([]int, error) {
	if d.Dominoes == nil || len(d.Dominoes) == 0 || len(d.Dominoes) < position {
		return nil, errors.New("There is no slice or the position does not exist")
	}

	d.Dominoes = append(d.Dominoes[:position], d.Dominoes[position+1:]...)
	return d.Dominoes, nil
}

func (d *dominoes) Query() ([]int, error) {
	if d.Dominoes == nil || len(d.Dominoes) == 0 {
		return nil, errors.New("There is no slice")
	}
	return d.Dominoes, nil
}
