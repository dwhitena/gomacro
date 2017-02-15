/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * string.go
 *
 *  Created on: Feb 13, 2015
 *      Author: Massimiliano Ghilardi
 */

package main

import (
	// "errors"
	// "fmt"
	"strconv"
)

const (
	eNormal = iota
	eBackslash
	eHex
	eOctal
	eUni4
	eUni8
)

func unescapeChar(str string) (rune, error) {
	// fmt.Printf("debug unescapeChar(): parsing CHAR %#v", str)
	rs := []rune(str)
	n := len(rs)
	if n >= 2 && rs[0] == '\'' && rs[n-1] == '\'' {
		rs = rs[1 : n-1]
	}
	/*
		rs = unescapeRunes(rs)
		if len(rs) != 1 {
			return 0, errors.New(fmt.Sprintf("invalid rune literal %#v, expecting exactly ONE rune", string(rs)))
		}
		return rs[0], nil
	*/
	ret, _, _, err := strconv.UnquoteChar(string(rs), '\'')
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func unescapeString(str string) (string, error) {
	/*
		rs := []rune(str)
		n := len(rs)
		if n >= 2 && rs[0] == '"' && rs[n-1] == '"' {
			rs = rs[1 : n-1]
		}
		return string(unescapeRunes(rs))
	*/
	return strconv.Unquote(str)
}

/*
func unescapeRunes(rs []rune) []rune {
	j := 0
	mode := eNormal
	var buf [8]rune
	bufi := 0

	for _, ch := range rs {
		switch mode {
		case eNormal:
			if ch == '\\' {
				mode = eBackslash
				continue
			}
			rs[j] = ch
			j++

		case eBackslash:
			switch ch {
			case '0', '1', '2', '3', '4', '5', '6', '7':
				mode = eOctal
				buf[0] = ch
				bufi = 1
				continue
			case 'U':
				mode = eUni8
				bufi = 0
				continue
			case 'a':
				ch = '\a'
			case 'b':
				ch = '\b'
			case 'f':
				ch = '\f'
			case 'n':
				ch = '\n'
			case 'r':
				ch = '\r'
			case 't':
				ch = '\t'
			case 'u':
				mode = eUni4
				bufi = 0
				continue
			case 'v':
				ch = '\v'
			case 'x':
				mode = eHex
				bufi = 0
				continue
			}
			rs[j] = ch
			j++
			mode = eNormal

		case eOctal:
			buf[bufi] = ch
			bufi++
			if bufi < 3 {
				continue
			}
			rs[j] = parseOctal(buf[:bufi])
			j++
			mode = eNormal

		case eHex, eUni4, eUni8:
			buf[bufi] = ch
			bufi++
			if mode == eHex && bufi < 2 {
				continue
			}
			if mode == eUni4 && bufi < 4 {
				continue
			}
			if mode == eUni8 && bufi < 8 {
				continue
			}
			rs[j] = parseHex(buf[:bufi])
			j++
			mode = eNormal
		}
	}
	return rs[0:j]
}

func parseOctal(rs []rune) rune {
	octal := ((rs[0] - '0') << 6) | ((rs[1] - '0') << 3) | (rs[2] - '0')
	// fmt.Printf("debug: parseOctal(%#v) -> %#v\n", string(rs), octal)
	return octal
}

func parseHex(rs []rune) rune {
	var hex rune = 0
	for _, ch := range rs {
		hex = hex<<4 | parseHexChar(ch)
	}
	// fmt.Printf("debug: parseHex(%#v) -> %#v\n", string(rs), hex)
	return hex
}

func parseHexChar(ch rune) rune {
	if ch >= '0' && ch <= '9' {
		ch -= '0'
	} else if ch >= 'A' && ch <= 'F' {
		ch -= 'A' - 10
	} else if ch >= 'a' && ch <= 'f' {
		ch -= 'a' - 10
	} else {
		ch = 0
	}
	return ch
}
*/
