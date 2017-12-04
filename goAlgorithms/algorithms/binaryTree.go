package algorithms

import "errors"

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value string, data string) error {
	if n == nil {
		return errors.New("There is no tree to add to")
	}
	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{
				Value: value,
				Data:  data,
			}
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{
				Value: value,
				Data:  data,
			}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s string) (string, bool) {
	if n == nil {
		return "", false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)

	}
}
