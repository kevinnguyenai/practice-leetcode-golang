package isIsomorphic

import "fmt"

func isIsomorphic(s string, t string) bool {
	// tokenize str to char
	_sS := []rune(s)
	_sT := []rune(t)

	// map t , s
	//fmt.Println(_mS)
	_mTRaw := map[rune]bool{}
	_mT := map[rune]rune{}
	for _, _m := range _sT {
		if _mTRaw[_m] == true {
			continue
		}
		_mTRaw[_m] = false
	}
	for _i, _s := range _sS {
		if _mT[_s] != 0 || _mTRaw[_sT[_i]] {
			continue
		}
		_mTRaw[_sT[_i]] = true
		_mT[_s] = _sT[_i]

	}
	// map t to s
	fmt.Println(_mT)
	// compare
	for k, v := range _sS {
		_sS[k] = _mT[v]
	}
	if stringSlicesEqual(_sS, _sT) {
		return true
	}
	return false
}

func stringSlicesEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
