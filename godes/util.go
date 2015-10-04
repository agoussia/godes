// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package godes

import (
	"bufio"
	"fmt"
	"os"
	//"godes"
	//"github.com/fatih/color"
	//"github.com/Kenshin/cprint"
	//"github.com/shiena/ansicolor"
	//"github.com/daviddengcn/go-colortext"
	
	"math"
	"text/tabwriter"
)

// local error wrapper so we can distinguish errors we want to return
// as errors from genuine panics (which we don't want to return as errors)
type osError struct {
	err error
}

func handlePanic(err *error, op string) {
	if e := recover(); e != nil {
		if nerr, ok := e.(osError); ok {
			*err = nerr.err
			return
		}
		panic("tabwriter: panic during " + op)
	}
}

func NewStatCollector(titles []string, measures [][]float64) *StatCollector {

	return &StatCollector{titles, measures}
}

type StatCollector struct {
	Titles   []string
	Measures [][]float64
	//StatWriter *bufio.Writer
}

func (collector *StatCollector) Clear() {
	collector.Titles = nil
	collector.Measures = nil
}

func (collector *StatCollector) Print(statWriter *bufio.Writer, titleSwt bool, avgSwt bool, stdDevSwt bool, lBooundSwt bool, uBoundSwt bool, minSwt bool, maxSwt bool) (err error) {

	var results [6][100]float64

	if statWriter == nil {
		panic("startWrite equal nil")
	}

	if collector.Measures == nil {
		panic("no data is available")
	}

	if titleSwt {

		fmt.Fprintf(statWriter, " Replication\t")

		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%v\t", collector.Titles[i])

		}
		fmt.Fprintf(statWriter, "\n")
	}

	//Results
	fmt.Fprintf(statWriter, "\n")

	for i := 0; i < len(collector.Titles); i++ {
		_, avg, std, lb, ub, min, max := collector.GetStat(i)
		results[0][i] = avg
		results[1][i] = std
		results[2][i] = lb
		results[3][i] = ub
		results[4][i] = min
		results[5][i] = max

	}
	if avgSwt {
		fmt.Fprintf(statWriter, "Avg. \t")
		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[0][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if stdDevSwt {
		fmt.Fprintf(statWriter, "StdDev\t")
		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[1][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if lBooundSwt {
		fmt.Fprintf(statWriter, "L-Bound\t")
		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[2][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if uBoundSwt {
		fmt.Fprintf(statWriter, "U-Bound\t")
		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[3][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if minSwt {
		fmt.Fprintf(statWriter, "Minimum\t")
		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[4][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if maxSwt {
		fmt.Fprintf(statWriter, "Maximum\t")
		for i := 0; i < len(collector.Titles); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[5][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	return
}

func (collector *StatCollector) PrintStat() (err error) {
	defer handlePanic(&err, "PrintStat")
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 10, 0, '\t', 0)

	
	fmt.Fprintln(w, "Variable\t#\tAverage\tStd Dev\tL-Bound\tU-Bound\tMinimum\tMaximum")




	for i := 0; i < len(collector.Titles); i++ {
		obs, avg, std, lb, ub, min, max := collector.GetStat(i)
		fmt.Fprintf(w, "%s\t%d\t%6.3f\t%6.3f\t%6.3f\t%6.3f\t%6.3f\t%6.3f\n", collector.Titles[i], obs, avg, std, lb, ub, min, max)

	}

	w.Flush()
	return
}
//GetStat returns size of the collection, average, standard deviation, low bound of confidence inteval, upper bould confidence inteval, minimum, maximum
func (collector *StatCollector) GetStat(titleInd int) (int64, float64, float64, float64, float64, float64, float64) {

	if collector.Measures == nil {
		panic("no data is available")
	}

	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}

	avg := 0.
	std := 0.
	lb := 0.
	ub := 0.

	repl := int64(len(collector.Measures))
	slice := []float64{}

	for i := 0; i < int(repl); i++ {
		slice = append(slice, collector.Measures[i][titleInd])

	}
	avg = MeanFloat(slice)
	std = StandardDeviationFloat(slice)
	lb, ub = NormalConfidenceIntervalFloat(slice)
	min, max := MinMaxFloat(slice)

	return repl, avg, std, lb, ub, min, max

}

//GetSize returns size of a collection
func (collector *StatCollector) GetSize(titleInd int) int{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	size := int(len(collector.Measures))
	return  size

}

//GetAverage returns average of a collection
func (collector *StatCollector) GetAverage(titleInd int) float64{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	avg := 0.
	slice := []float64{}
	size:=collector.GetSize(titleInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.Measures[i][titleInd])
	}
	avg = MeanFloat(slice)
	return  avg
}

//GetStandardDeviation returns standard deviation of a collection
func (collector *StatCollector) GetStandardDeviation(titleInd int) float64{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	std := 0.
	slice := []float64{}
	size:=collector.GetSize(titleInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.Measures[i][titleInd])
	}
	std = StandardDeviationFloat(slice)
	return  std
}

//GetLowBoundCI returns low bound of confidence interval for collection
func (collector *StatCollector) GetLowBoundCI(titleInd int) float64{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	lb := 0.
	slice := []float64{}
	size:=collector.GetSize(titleInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.Measures[i][titleInd])
	}
	lb, _ = NormalConfidenceIntervalFloat(slice)
	return  lb
}

//GetUpperBoundCI returns upper bound of confidence interval for collection
func (collector *StatCollector) GetUpperBoundCI(titleInd int) float64{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	ub := 0.
	slice := []float64{}
	size:=collector.GetSize(titleInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.Measures[i][titleInd])
	}
	_, ub = NormalConfidenceIntervalFloat(slice)
	return  ub
}



//GetMinimum returns minimum value for collection
func (collector *StatCollector) GetMinimum(titleInd int) float64{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	min := 0.
	slice := []float64{}
	size:=collector.GetSize(titleInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.Measures[i][titleInd])
	}
	min,_ = MinMaxFloat(slice)
	return  min
}

//GetMaximum returns maximum value for collection
func (collector *StatCollector) GetMaximum(titleInd int) float64{

	if collector.Measures == nil {
		panic("no data is available")
	}
	if titleInd < 0 || titleInd >= len(collector.Measures)-1 {
		panic("invalid index")
	}
	max := 0.
	slice := []float64{}
	size:=collector.GetSize(titleInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.Measures[i][titleInd])
	}
	_,max = MinMaxFloat(slice)
	return  max
}



// MeanFloat returns float mean value
func MeanFloat(nums []float64) (mean float64) {
	if len(nums) == 0 {
		return 0.0
	}
	for _, n := range nums {
		mean += n
	}
	return mean / float64(len(nums))
}

// StandardDeviationFloat returns STD
func StandardDeviationFloat(nums []float64) (dev float64) {
	if len(nums) == 0 {
		return 0.0
	}

	m := MeanFloat(nums)
	for _, n := range nums {
		dev += (n - m) * (n - m)
	}
	dev = math.Pow(dev/float64(len(nums)-1.), 0.5)

	return dev
}

// NormalConfidenceIntervalFloat returns Confidence Interval
func NormalConfidenceIntervalFloat(nums []float64) (lower float64, upper float64) {
	conf := 1.95996 // 95% confidence for the mean, http://bit.ly/Mm05eZ
	mean := MeanFloat(nums)
	dev := StandardDeviationFloat(nums) / math.Sqrt(float64(len(nums)))
	return mean - dev*conf, mean + dev*conf
}

// NormalConfidenceIntervalFloat returns Confidence Interval
func MinMaxFloat(nums []float64) (minimum float64, maximum float64) {

	min := nums[0]
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return min, max
}
