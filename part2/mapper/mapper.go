package mapper

import (
	"fmt"
	"unicode"
)

//StringMapper interface
type StringMapper interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

//SkipString struct
//skips defines what char of string to transform, ex. 3 => every third char
//skipCount is a counter that increments everytime an alphanumeric char is found in string
//stringToMap is the string that will be transformed
type SkipString struct {
	skips       int
	skipCount   int
	stringToMap string
}

//NewSkipString takes as arguments a number of skips,
//skipCount is set to 0, counts the number of alphanumeric chars,
//and the string that will be transformed
func NewSkipString(skips int, stringToMap string) SkipString {
	return SkipString{skips, 0, stringToMap}
}

//GetValueAsRuneSlice return the string as a rune slice
func (s *SkipString) GetValueAsRuneSlice() []rune {
	return []rune(s.stringToMap)
}

//TransformRune takes in the position of the rune to transform,
//checks if alphanumeric, if alphanumeric increase skipCount,
//if skipcount == skips then transform rune to uppercase, else tranform to lowercase,
//save modified string to stringToMap
func (s *SkipString) TransformRune(pos int) {
	if s.skips < 1 {
		return
	}

	if unicode.IsLetter(rune(s.stringToMap[pos])) || unicode.IsNumber(rune(s.stringToMap[pos])) {
		s.skipCount++
		runeSlice := s.GetValueAsRuneSlice()
		if s.skipCount == s.skips {
			runeSlice[pos] = unicode.ToUpper(runeSlice[pos])
			s.stringToMap = string(runeSlice)
			s.skipCount = 0
			return
		}
		runeSlice[pos] = unicode.ToLower(runeSlice[pos])
		s.stringToMap = string(runeSlice)
	}
}

//String returns the string, satisfies stringer interface
func (s SkipString) String() string {
	return fmt.Sprintf(s.stringToMap)
}

//MapString accepts StringMapper interface,
//loop over rune slice and transforms rune in current pos
func MapString(i StringMapper) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}
