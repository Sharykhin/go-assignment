package switcher

import (
	"errors"
	"fmt"
)

type (
	// Mode specifies one of three modes which logical switcher can work with
	Mode string
	// Formula describes general formula for calculating final result
	Formula func(d float64, e, f int) float64
	// Logical describes logical switcher with two phases, the first one calculates h result that affects
	// the second phase where final formula is set up
	Logical struct {
		m Formula
		p Formula
		t Formula
		h Formula
	}
)

const (
	Base Mode = "base"
	CustomOne Mode = "custom1"
	CustomTwo Mode = "custom2"
)

var (
	// ErrUnexpectedInput says that input has no matches to calculate h
	ErrUnexpectedInput = errors.New("unexpected input match")
	// ErrUnexpectedMode says that provided mode is invalid
	ErrUnexpectedMode = errors.New("mode is invalid")
)

// NewLogical creates a new logical struct instance
func NewLogical() *Logical {
	l := Logical{
		m : func(d float64, e, f int) float64 {
			return d + (d * float64(e) / 10)
		},
		p : func (d float64, e, f int) float64 {
			return d + d * ((float64(e) - float64(f)) / 25.5)
		},
		t : func (d float64, e ,f int) float64 {
			return d - (d * float64(f) / 30)
		},
		h: nil,
	}

	return &l
}

// Calculate calculates result taking into account a provided mode
func (l *Logical) Calculate(a, b, c bool, d float64, e, f int, mode Mode) (float64, error) {
	err := l.adjustMode(mode)
	if err != nil {
		return 0, err
	}
	err = l.calculateH(a, b, c, mode)
	if err != nil {
		return 0, nil
	}

	k := l.h(d, e, f)

	return k, nil
}

// adjustMode affects formulas that final result will be calculated with
func (l *Logical) adjustMode(mode Mode) error {
	switch mode {
	case Base:
		return nil
	case CustomOne:
		l.p = func (d float64, e, f int) float64 {
			return 2 * d + (d * float64(e) / 100)
		}
	case CustomTwo:
		l.m = func(d float64, e, f int) float64 {
			return float64(f) + d + (d * float64(e) / 100)
		}
	default:
		return fmt.Errorf("[switcher][Logical][adjustMode] failed to adjust mode: %w", ErrUnexpectedMode)
	}

	return nil
}

// calculateH calculates h result
func (l *Logical) calculateH(a, b, c bool, mode Mode) error {
	switch  {
	case a && b && !c:
		l.h = l.m
	case a && b && c:
		l.h = l.p
	case !a && b && c:
		l.h = l.t
	}

	switch mode {
	case CustomTwo:
		switch  {
		case a && b && !c:
			l.h = l.t
		case a && !b && c:
			l.h =l. m
		}
	}

	if l.h == nil {
		return fmt.Errorf("[switcher][Logical][calculateH] failed to calculate first result: %w", ErrUnexpectedInput)
	}

	return nil
}
