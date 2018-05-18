package paginator

import (
	"bytes"
	"errors"
	"sort"
	"strconv"
	"strings"
)

// Paginator struct for generating paginator
type Paginator struct {
	CurrentPage int
	TotalPages  int
	Boundaries  int
	Around      int

	SkipChar string
}

type boundaries struct {
	Begining []int
	Ending   []int
}

type around struct {
	Before []int
	After  []int
}

// Create Paginator struct with
func Create(currentPage int, totalPages int, boundaries int, around int, skipChar string) (Paginator, error) {
	if totalPages < 1 {
		err := errors.New("Total pages cannot be less then 1")
		return Paginator{}, err
	}
	if currentPage > totalPages {
		err := errors.New("Current page cannot be greater then total pages")
		return Paginator{}, err
	}
	if currentPage < 0 {
		err := errors.New("Current page cannot be less then zero")
		return Paginator{}, err
	}
	return Paginator{currentPage, totalPages, boundaries, around, skipChar}, nil
}

// Generate string paginator
func (p *Paginator) Generate() string {
	b := p.createBoundryPages()
	a := p.createAroundPages()
	r := p.createFinalResult(b, a)
	return r
}

func (p *Paginator) createBoundryPages() boundaries {
	var b, e []int
	if p.Boundaries > 0 {
		for i := 0; i < p.Boundaries; i++ {
			if i < p.TotalPages {
				b = append(b, i+1)
			}
			if p.TotalPages-i > 0 {
				e = append(e, p.TotalPages-i)
			}

		}
	}
	return boundaries{Begining: b, Ending: e}
}

func (p *Paginator) createAroundPages() around {
	var a, b []int
	for i := 1; i <= p.Around; i++ {
		if p.CurrentPage-i > 0 {
			b = append(b, p.CurrentPage-i)
		}
		if p.CurrentPage+i <= p.TotalPages {
			a = append(a, p.CurrentPage+i)
		}
	}
	return around{Before: b, After: a}
}

func (p *Paginator) createFinalResult(b boundaries, a around) string {
	var pages []int
	pages = append(pages, b.Begining...)
	pages = append(pages, a.Before...)
	pages = append(pages, p.CurrentPage)
	pages = append(pages, a.After...)
	pages = append(pages, b.Ending...)
	pages = p.removeDuplicates(pages)
	sort.Ints(pages)
	return p.buildString(pages)
}

func (p *Paginator) removeDuplicates(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (p *Paginator) buildString(pages []int) string {
	var buffer bytes.Buffer

	for i := range pages {
		buffer.WriteString(strconv.Itoa(pages[i]))
		buffer.WriteString(" ")

		if i+1 < len(pages) && pages[i]+1 != pages[i+1] {
			buffer.WriteString(p.SkipChar)
			buffer.WriteString(" ")
		}
	}

	// Handle edge case when boundaries are 0
	if p.Boundaries == 0 {
		if p.CurrentPage > 1 && (p.CurrentPage-p.Around) > 1 {
			s := buffer.String()
			buffer.Reset()
			buffer.WriteString(p.SkipChar)
			buffer.WriteString(" ")
			buffer.WriteString(s)
		}
		if p.CurrentPage < p.TotalPages && (p.CurrentPage+p.Around) < p.TotalPages {
			buffer.WriteString(p.SkipChar)
			buffer.WriteString(" ")
		}

	}

	return strings.TrimSpace(buffer.String())
}
