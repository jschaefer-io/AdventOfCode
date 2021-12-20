package day19

import (
    "errors"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "strconv"
    "strings"
)

type BeaconList []Vector

type Scanner struct {
    beacons    BeaconList
    variations [24]BeaconList
    index      map[int]map[int]map[int]struct{}
}

func (s *Scanner) BuildVariations() {
    for _, beacon := range s.beacons {
        for v, mBeacon := range beacon.Rotations() {
            s.variations[v] = append(s.variations[v], mBeacon)
        }
    }
}

func (s *Scanner) BuildIndex() {
    s.index = make(map[int]map[int]map[int]struct{})
    for _, beacon := range s.beacons {
        if _, ok := s.index[beacon.X]; !ok {
            s.index[beacon.X] = make(map[int]map[int]struct{})
        }
        if _, ok := s.index[beacon.X][beacon.Y]; !ok {
            s.index[beacon.X][beacon.Y] = make(map[int]struct{})
        }
        s.index[beacon.X][beacon.Y][beacon.Z] = struct{}{}
    }
}

func (s *Scanner) Overlay(target Scanner) (BeaconList, Vector, error) {
    for _, point := range s.beacons[11:] {
        for _, variation := range target.variations {
            for i, reference := range variation {
                matches := 1
                notMatched := make(BeaconList, 0)
                for u, check := range variation {
                    if i == u {
                        continue
                    }
                    mPoint := check.Sub(reference).Add(point)
                    if s.Has(mPoint) {
                        matches++
                    } else {
                        notMatched = append(notMatched, mPoint)
                    }
                }
                if matches >= 12 {
                    return notMatched, point.Sub(reference), nil
                }
            }
        }
    }
    return nil, Vector{}, errors.New("not enough matches found")
}

func (s *Scanner) Has(search Vector) bool {
    if _, ok := s.index[search.X]; !ok {
        return false
    }
    if _, ok := s.index[search.X][search.Y]; !ok {
        return false
    }
    _, ok := s.index[search.X][search.Y][search.Z]
    return ok
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
        scanners[i].BuildVariations()
    }

    startList := make([]Vector, 0)
    start := scanners[0]
    start.BuildIndex()
    list := scanners[1:]
    for len(list) > 0 {
        for i := range list {
            beaconList, scannerPos, err := start.Overlay(list[i])
            if err == nil {
                startList = append(startList, scannerPos)
                newStart := start.Extend(beaconList)
                newStart.BuildIndex()
                start = newStart
                list = append(list[:i], list[i+1:]...)
                break
            }
        }
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

func init() {
    orchestration.MainDispatcher.AddSolver("Day19", orchestration.NewSolver(Solve))
}
