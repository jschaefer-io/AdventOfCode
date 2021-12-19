package day19

import (
    "fmt"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "strconv"
    "strings"
)

type Vector struct {
    X int
    Y int
    Z int
}

func (v Vector) Sub(s Vector) Vector {
    return Vector{
        X: v.X - s.X,
        Y: v.Y - s.Y,
        Z: v.Z - s.Z,
    }
}

func (v Vector) Add(s Vector) Vector {
    return Vector{
        X: v.X + s.X,
        Y: v.Y + s.Y,
        Z: v.Z + s.Z,
    }
}

func (v Vector) Mul(f int) Vector {
    return Vector{
        X: v.X * f,
        Y: v.Y * f,
        Z: v.Z * f,
    }
}

func (v Vector) Len() float64 {
    x := float64(v.X)
    y := float64(v.Y)
    z := float64(v.Z)
    return math.Sqrt(x*x + y*y + z*z)
}

func (v Vector) String() string {
    return fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z)
}

func (v Vector) Compare(s Vector) bool {
    return v.Equal(s) || v.Mul(-1).Equal(s)
}

func (v Vector) Equal(s Vector) bool {
    return s.X == v.X && s.Y == v.Y && s.Z == v.Z
}

func (v Vector) Spin(count int) Vector {
    if count <= 0 {
        return v
    }
    nV := Vector{v.X, v.Z, v.Y * -1}
    return nV.Spin(count - 1)
}

func (v Vector) Rotations() []Vector {
    x := v.X
    y := v.Y
    z := v.Z
    return []Vector{
        {x, y, z},
        {x, z, -y},
        {x, -y, -z},
        {x, -z, y},
        {y, -z, -x},
        {y, -x, z},
        {y, z, x},
        {y, x, -z},
        {z, x, y},
        {z, -y, x},
        {z, -x, -y},
        {z, y, -x},
        {-x, y, -z},
        {-x, z, y},
        {-x, -y, z},
        {-x, -z, -y},
        {-y, z, -x},
        {-y, x, z},
        {-y, -z, x},
        {-y, -x, -z},
        {-z, -x, y},
        {-z, -y, -x},
        {-z, x, -y},
        {-z, y, x},
    }
}

type BeaconList []Vector

type CheckVector struct {
    Vector Vector
    A      Vector
    B      Vector
}

func (bList BeaconList) CheckList() []CheckVector {
    count := len(bList)
    list := make([]CheckVector, 0)
    for a := 0; a < count; a++ {
        for b := a + 1; b < count; b++ {
            res := bList[a].Sub(bList[b])
            list = append(list, CheckVector{res, bList[a], bList[b]})
        }
    }
    return list
}

type Scanner struct {
    beacons    BeaconList
    variations [24]BeaconList
}

func (s *Scanner) Extend(list BeaconList) Scanner {
    scanner := Scanner{
        beacons: make(BeaconList, 0),
    }
    for _, source := range s.beacons {
        scanner.beacons = append(scanner.beacons, source)
    }
    for _, newSource := range list {
        scanner.beacons = append(scanner.beacons, newSource)
    }
    return scanner
}

func (s *Scanner) HasBeacon(check Vector) bool {
    for _, b := range s.beacons {
        if b.Equal(check) {
            return true
        }
    }
    return false
}

func (s *Scanner) BuildVariationList() {
    for _, beacon := range s.beacons {
        for v, mBeacon := range beacon.Rotations() {
            s.variations[v] = append(s.variations[v], mBeacon)
        }
    }
}

type Intersection struct {
    source CheckVector
    target CheckVector
}

func (i Intersection) ResolveTargetScanner() Vector {
    source := i.source.A.Sub(i.target.A)
    if source.Add(i.target.B).Equal(i.source.B) {
        return source
    }
    return i.source.A.Sub(i.target.B)
}

func (s *Scanner) Intersect(t Scanner) (int, Vector, BeaconList) {
    checkList := s.beacons.CheckList()

    max := -1
    var maxVariations BeaconList
    var sourcePoint Vector
    for _, variation := range t.variations {
        var intersection Intersection
        variations := make(map[Vector]struct{})
        count := 0
        for _, checkB := range variation.CheckList() {
            for _, checkA := range checkList {
                if checkA.Vector.Compare(checkB.Vector) {
                    count++
                    variations[checkA.A] = struct{}{}
                    variations[checkA.B] = struct{}{}
                    intersection = Intersection{checkA, checkB}
                }
            }
        }
        currCount := len(variations)
        if max == -1 || currCount > max {
            max = currCount
            maxVariations = variation
            sourcePoint = intersection.ResolveTargetScanner()
        }
    }
    mappedPoints := make(BeaconList, 0)
    for _, p := range maxVariations {
        tVec := sourcePoint.Add(p)
        if s.HasBeacon(tVec) {
            continue
        }
        mappedPoints = append(mappedPoints, tVec)
    }
    return max, sourcePoint, mappedPoints
}

func (s Scanner) String() string {
    var sb strings.Builder
    for _, beacon := range s.beacons {
        sb.WriteString(fmt.Sprintln(beacon))
    }
    return sb.String()
}

func Solve(data string, result *orchestration.Result) error {
    scanners := make([]Scanner, 0)
    for _, scanner := range strings.Split(data, "\n\n") {
        currScanner := Scanner{
            beacons: make(BeaconList, 0),
        }
        for _, beacon := range strings.Split(scanner, "\n")[1:] {
            if len(beacon) == 0 {
                continue
            }
            vals := strings.Split(beacon, ",")
            x, err := strconv.Atoi(vals[0])
            if err != nil {
                return err
            }
            y, err := strconv.Atoi(vals[1])
            if err != nil {
                return err
            }
            z, err := strconv.Atoi(vals[2])
            if err != nil {
                return err
            }
            currScanner.beacons = append(currScanner.beacons, Vector{x, y, z})
        }
        scanners = append(scanners, currScanner)
    }

    for i, _ := range scanners {
        scanners[i].BuildVariationList()
    }

    //startPoints[0] = Vector{0, 0, 0}

    //sCount := len(scanners)
    //for a := 0; a < sCount; a++ {
    //    for b := a + 1; b < sCount; b++ {
    //        i, _, _ := scanners[a].Intersect(scanners[b])
    //        //startPoints[b] = s
    //        //fmt.Println(b, iCount, s)
    //        if i == 12 {
    //            fmt.Println(a,b)
    //        }
    //    }
    //}
    //fmt.Println(startPoints)

    startList := make([]Vector, 0)

    start := scanners[0]
    list := scanners[1:]
    for len(list) > 0 {
        for i := range list {
            iCount, beaconPos, p := start.Intersect(list[i])
            if iCount >= 12 {
                startList = append(startList, beaconPos)
                newStart := start.Extend(p)
                newStart.BuildVariationList()
                start = newStart
                list = append(list[:i], list[i+1:]...)
                break
            }
        }
        //fmt.Println(len(list))
    }
    // a
    result.AddResult(strconv.Itoa(len(start.beacons)))

    // b
    maxD := -1
    for a := 0; a < len(startList); a++ {
        for b := a + 1; b < len(startList); b++ {
            v := startList[a].Sub(startList[b])
            d := int(math.Abs(float64(v.X)) + math.Abs(float64(v.Y)) + math.Abs(float64(v.Z)))
            if maxD == -1 || maxD < d {
                maxD = d
            }
        }
    }
    result.AddResult(strconv.Itoa(maxD))
    return nil
}

func
init() {
    orchestration.MainDispatcher.AddSolver("Day19", orchestration.NewSolver(Solve))
}
