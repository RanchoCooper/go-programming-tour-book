package convert

import (
    "strconv"
)

/**
 * @author Rancho
 * @date 2021/11/27
 */

type StrTo string

func (s StrTo) String() string {
    return string(s)
}

func (s StrTo) Int() (int, error) {
    return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
    v, _ := s.Int()

    return v
}

func (s StrTo) UInt32() (uint32, error) {
    v, err := s.Int()

    return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
    v, _ := s.UInt32()

    return v
}