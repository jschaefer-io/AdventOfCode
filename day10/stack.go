package day10

type Stack struct {
    length int
    list   []rune
}

func NewStack() Stack {
    return Stack{
        length: 0,
        list:   make([]rune, 0),
    }
}

func (s *Stack) Peek() rune {
    return s.list[s.length-1]
}

func (s *Stack) Pop() rune {
    v := s.Peek()
    s.list = s.list[:s.length-1]
    s.length--
    return v
}

func (s *Stack) Push(c rune) {
    s.list = append(s.list, c)
    s.length++
}

func (s *Stack) Length() int {
    return s.length
}
