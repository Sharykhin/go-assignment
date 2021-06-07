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
			return d + (d * ((float64(e) - float64(f)) / 25.5))
		},
		t : func (d float64, e ,f int) float64 {
			return d - (d * float64(f) / 30)
		},
	}

	return &l
}

// Calculate calculates result taking into account a provided mode
func (l *Logical) Calculate(a, b, c bool, d float64, e, f int, mode Mode) (float64, error) {
	m, p, t, err := l.getFormulas(mode)
	if err != nil {
		return 0, fmt.Errorf("[switcher][Logical][Calculate] failed to get formulas: %w", err)
	}
	h, err := l.calculateH(a, b, c, m, p, t, mode)
	if err != nil {
		return 0, fmt.Errorf("[switcher][Logical][Calculate] failed to calculate H: %w", err)
	}

	k := h(d, e, f)

	return round(k, 0.01), nil
}

func (l *Logical) getFormulas(mode Mode) (Formula, Formula, Formula, error) {
	m, p, t := l.m, l.p, l.t

	switch mode {
	case Base:
		return m, p, t, nil
	case CustomOne:
		p = func (d float64, e, f int) float64 {
			return 2 * d + (d * float64(e) / 100)
		}
		return m, p, t, nil
	case CustomTwo:
		m = func(d float64, e, f int) float64 {
			return float64(f) + d + (d * float64(e) / 100)
		}
		return m, p, t, nil
	default:
		return nil, nil, nil, fmt.Errorf(
			"[switcher][Logical][getFormulas] failed to get formulas based on mode: %w", ErrUnexpectedMode,
		)
	}
}

// calculateH calculates h result
func (l *Logical) calculateH(a, b, c bool, m, p, t Formula, mode Mode) (Formula, error) {
	var h Formula
	switch  {
	case a && b && !c:
		h = m
	case a && b && c:
		h = p
	case !a && b && c:
		h = t
	}

	switch mode {
	case CustomTwo:
		switch  {
		case a && b && !c:
			h = t
		case a && !b && c:
			h = m
		}
	}

	if h == nil {
		return nil, fmt.Errorf("[switcher][Logical][calculateH] failed to calculate first result: %w", ErrUnexpectedInput)
	}

	return h, nil
}


