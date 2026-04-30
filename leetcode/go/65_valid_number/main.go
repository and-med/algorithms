package main

import "fmt"

func isNumber(s string) bool {
	parser := NewParser(s)

	return parser.Parse()
}

type Parser struct {
	currPos int
	buff    string
}

func NewParser(s string) *Parser {
	return &Parser{
		currPos: 0,
		buff:    s,
	}
}

func (p *Parser) Parse() bool {
	return p.parseDecimal() && p.outOfBounds()
}

func (p *Parser) parseDecimal() bool {
	hasIntegerPart := p.parseSignedNumber()
	hasDot := p.parseDot()

	if hasDot && !p.parseNumber() && !hasIntegerPart {
		// invalid: has dot but no decimal part or integer part
		return false
	}

	if !hasDot && !hasIntegerPart {
		// invalid: no dot and no integer part
		return false
	}

	if p.parseExponent() && !p.parseSignedNumber() {
		// invalid, exponent is present but the number is not present
		return false
	}

	return true
}

func (p *Parser) parseSignedNumber() bool {
	// Optional plus or minus, ignoring the result
	p.parsePlusOrMinus()

	return p.parseNumber()
}

func (p *Parser) parseNumber() bool {
	startPos := p.currPos
	for p.parseDigit() {
	}
	endPos := p.currPos

	return startPos != endPos
}

func (p *Parser) parseDigit() bool {
	if p.outOfBounds() {
		return false
	}

	curr := p.current()

	if curr >= '0' && curr <= '9' {
		return p.advance()
	}

	return false
}

func (p *Parser) parseDot() bool {
	if p.outOfBounds() {
		return false
	}

	if p.current() == '.' {
		return p.advance()
	}

	return false
}

func (p *Parser) parseExponent() bool {
	if p.outOfBounds() {
		return false
	}

	curr := p.current()

	if curr == 'e' || curr == 'E' {
		return p.advance()
	}

	return false
}

func (p *Parser) parsePlusOrMinus() bool {
	if p.outOfBounds() {
		return false
	}

	curr := p.current()

	if curr == '+' || curr == '-' {
		return p.advance()
	}

	// Positive by default
	return false
}

func (p *Parser) current() byte {
	if p.outOfBounds() {
		panic("attempting to get current at the end of array")
	}

	return p.buff[p.currPos]
}

func (p *Parser) outOfBounds() bool {
	return p.currPos >= len(p.buff)
}

func (p *Parser) advance() bool {
	p.currPos++
	return true
}

func main() {
	fmt.Println(verify(true, isNumber("1e1")))
	fmt.Println(verify(true, isNumber("2")))
	fmt.Println(verify(true, isNumber("0089")))
	fmt.Println(verify(true, isNumber("-0.1")))
	fmt.Println(verify(true, isNumber("+3.14")))
	fmt.Println(verify(true, isNumber("4.")))
	fmt.Println(verify(true, isNumber("-.9")))
	fmt.Println(verify(true, isNumber("2e10")))
	fmt.Println(verify(true, isNumber("-90E3")))
	fmt.Println(verify(true, isNumber("3e+7")))
	fmt.Println(verify(true, isNumber("+6e-1")))
	fmt.Println(verify(true, isNumber("53.5e93")))
	fmt.Println(verify(true, isNumber("-123.456e789")))
	fmt.Println(verify(true, isNumber("3.")))
	fmt.Println(verify(true, isNumber("3.e+1")))

	fmt.Println(verify(false, isNumber(".")))
	fmt.Println(verify(false, isNumber("abc")))
	fmt.Println(verify(false, isNumber("1a")))
	fmt.Println(verify(false, isNumber("1e")))
	fmt.Println(verify(false, isNumber("e3")))
	fmt.Println(verify(false, isNumber("99e2.5")))
	fmt.Println(verify(false, isNumber("--6")))
	fmt.Println(verify(false, isNumber("-+3")))
	fmt.Println(verify(false, isNumber("95a54e53")))
}

func verify(expected bool, actual bool) string {
	if expected == actual {
		return "PASS"
	} else {
		return "****************************FAIL****************************"
	}
}
