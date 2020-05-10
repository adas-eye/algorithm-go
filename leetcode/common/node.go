package common

type Node struct {
	Val      int
	Children []*Node
}

var (
	NodeExample1 = &Node{
		Val: 0,
		Children: []*Node{
			{
				Val: 1,
				Children: []*Node{
					{
						Val:      4,
						Children: nil,
					},
					{
						Val:      5,
						Children: nil,
					},
					{
						Val:      6,
						Children: nil,
					},
				},
			},
			{
				Val:      2,
				Children: nil,
			},
			{
				Val:      3,
				Children: nil,
			},
		},
	}
)