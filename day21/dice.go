package day21

type Dice interface {
    Roll() map[int]int
    Count() int
}

type DeterministicDice struct {
    next  int
    count int
}

func (d *DeterministicDice) Count() int {
    return d.count
}

func (d *DeterministicDice) Roll() map[int]int {
    v := d.next + 1
    value := v*3 + 3
    value %= 100
    d.next += 3 % 100
    d.count += 3
    return map[int]int{value: 1}
}

type DiracDice struct{}

func (d *DiracDice) Count() int {
    return 0
}

func (d *DiracDice) Roll() map[int]int {
    list := make(map[int]int, 0)
    for a := 1; a <= 3; a++ {
        for b := 1; b <= 3; b++ {
            for c := 1; c <= 3; c++ {
                list[a+b+c]++
            }
        }
    }
    return list
}
