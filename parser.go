package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type valueRanges struct {
	name string
	min  int
	max  int
}

var (
	minute = valueRanges{"minute", 0, 59}
	hour   = valueRanges{"hour", 0, 23}
	dom    = valueRanges{"day of month", 1, 31}
	month  = valueRanges{"month", 1, 12}
	dow    = valueRanges{"day of week", 0, 6}
)

// this is where we store the parsed cron
type CronSchedule struct {

	// the raw cron string we get as input
	Raw string

	// parsed field values
	Minute []int
	Hour   []int
	DoM    []int
	Month  []int
	DoW    []int

	// command to run
	Command string
}

func NewCronSchedule(raw string) *CronSchedule {
	return &CronSchedule{
		Raw: raw,
	}
}

func (c *CronSchedule) Process() {

	cronSlc := strings.Split(c.Raw, " ")

	if len(cronSlc) != 6 {
		log.Fatalf("Invalid input: %s\n", c.Raw)
	}

	// parse minutes
	c.Minute = c.FieldProcessor(NewField(cronSlc[0], minute))
	c.Hour = c.FieldProcessor(NewField(cronSlc[1], hour))
	c.DoM = c.FieldProcessor(NewField(cronSlc[2], dom))
	c.Month = c.FieldProcessor(NewField(cronSlc[3], month))
	c.DoW = c.FieldProcessor(NewField(cronSlc[4], dow))

	c.Command = cronSlc[5]
}

func (c *CronSchedule) FieldProcessor(f *Field) []int {
	res, err := f.Parser()
	if err != nil {
		log.Fatalf("Invalid input %s: %s\n%s\n", f.kind.name, c.Raw, err)
	}

	valRes, err := f.Validator(res)
	if err != nil {
		log.Fatalf("Invalid input %s: %s\n%s\n", f.kind.name, c.Raw, err)
	}

	return valRes
}

// one cron field
type Field struct {
	raw  string
	kind valueRanges
}

func NewField(raw string, kind valueRanges) *Field {
	return &Field{
		raw:  raw,
		kind: kind,
	}
}

func (f *Field) Parser() ([]int, error) {

	// if the files is a number
	// e.g. "5" "99" etc...
	valInt, terr := strconv.Atoi(f.raw)
	if terr == nil {
		return []int{valInt}, nil
	}

	// check if it's a wildcard
	if f.raw == "*" {
		return rangeGenerator(f.kind.min, f.kind.max, 1), nil
	}

	// every X mimutes/hours "*/15"
	if strings.HasPrefix(f.raw, "*/") {
		valStr := strings.TrimLeft(f.raw, "*/")
		valInt, err := strconv.Atoi(valStr)
		if err != nil {
			return nil, err
		}

		return rangeGenerator(valInt, f.kind.max, valInt), nil
	}

	// series of numbers "5,6,7"
	if strings.Contains(f.raw, ",") {

		valStr := strings.Split(f.raw, ",")
		return strConverter(valStr)

	}

	// range "5-8"
	if strings.Contains(f.raw, "-") {
		valStr := strings.Split(f.raw, "-")

		if len(valStr) != 2 {
			return nil, fmt.Errorf("Invalid input %s: %s\n", f.kind.name, f.raw)
		}

		valInt, err := strConverter(valStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid input %s: %s\n", f.kind.name, f.raw)
		}

		return rangeGenerator(valInt[0], valInt[1], 1), nil
	}

	return nil, fmt.Errorf("Invalid input %s: %s\n", f.kind.name, f.raw)
}

func (f *Field) Validator(vals []int) ([]int, error) {

	sort.Ints(vals)

	if vals[0] < f.kind.min || vals[len(vals)-1] > f.kind.max {
		return nil, fmt.Errorf("Validator - Invalid input %s: %s\n", f.kind.name, f.raw)
	}

	return vals, nil

}

func isInt(v string) bool {
	_, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return true
}

func rangeGenerator(start, stop, step int) []int {

	var s []int
	for i := start; i <= stop; i = i + step {
		s = append(s, i)
	}

	return s
}

func strConverter(strSlice []string) ([]int, error) {
	var s []int

	for _, v := range strSlice {

		valInt, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		s = append(s, valInt)
	}
	return s, nil
}
