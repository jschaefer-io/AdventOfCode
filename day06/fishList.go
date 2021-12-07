package day06

type FishList struct {
    count int
    list  map[int8]int
}

func NewFishList(flatList []int8) FishList {
    list := make(map[int8]int)
    for _, f := range flatList {
        c := list[f]
        list[f] = c + 1
    }
    return FishList{
        count: len(flatList),
        list:  list,
    }
}

func (f *FishList) Count() int {
    return f.count
}

func (f *FishList) Tick() {
    var age int8
    for age = 0; age <= 9; age++ {
        f.list[age-1] = f.list[age]
    }
    birthCount := f.list[-1]
    f.list[6] += birthCount
    f.list[8] += birthCount
    f.count += birthCount
}
