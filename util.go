// Copyright 2015 Alex Goussiatiner. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Package godes  is the general-purpose simulation library
// which includes the  simulation engine  and building blocks
// for modeling a wide variety of systems at varying levels of details.

package godes

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
	"time"
)
const mAX_NUMBER_OF_SAMPLES = 100
const mAX_NUMBER_OF_PARAMETERS = 6

var curTime int64

func GetCurComputerTime() int64 {
	ct := time.Now().UnixNano()
	if ct > curTime {
		curTime = ct
		return ct
	} else if ct == curTime {
		curTime = ct + 1
		return curTime
	} else {
		curTime++
		return curTime
	}
}

//NewStatCollector creates a wrapper for samples data
func NewStatCollector(measures []string, samples [][]float64) *StatCollector {

	if measures == nil {
		panic("null measures array")
	}

	if samples == nil {
		panic("null samples array")
	}

	if len(measures) != len(samples[0]) {
		panic("invalid measures/samples arrays")
	}

	return &StatCollector{measures, samples}
}
//StatCollector is a wrapper which contains set of samples for statistical analyses
type StatCollector struct {
	measures []string
	samples  [][]float64
}

//Print calculates statistical parameters and output them to *bufio.Writer
// parameters flags are allowed
func (collector *StatCollector) Print(statWriter *bufio.Writer, measuresSwt bool, avgSwt bool, stdDevSwt bool, lBooundSwt bool, uBoundSwt bool, minSwt bool, maxSwt bool) (err error) {

	var results [mAX_NUMBER_OF_PARAMETERS][mAX_NUMBER_OF_SAMPLES]float64
	if statWriter == nil {
		panic("startWrite equal nil")
	}
	if measuresSwt {
		fmt.Fprintf(statWriter, " Replication\t")
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%v\t", collector.measures[i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	//Results
	fmt.Fprintf(statWriter, "\n")

	for i := 0; i < len(collector.measures); i++ {
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
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[0][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if stdDevSwt {
		fmt.Fprintf(statWriter, "StdDev\t")
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[1][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if lBooundSwt {
		fmt.Fprintf(statWriter, "L-Bound\t")
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[2][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if uBoundSwt {
		fmt.Fprintf(statWriter, "U-Bound\t")
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[3][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if minSwt {
		fmt.Fprintf(statWriter, "Minimum\t")
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[4][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}

	if maxSwt {
		fmt.Fprintf(statWriter, "Maximum\t")
		for i := 0; i < len(collector.measures); i++ {
			fmt.Fprintf(statWriter, "%6.3f \t", results[5][i])
		}
		fmt.Fprintf(statWriter, "\n")
	}
	return
}

// PrintStat calculates and prints statistical parameters
func (collector *StatCollector) PrintStat() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 10, 0, '\t', 0)
	fmt.Fprintln(w, "Variable\t#\tAverage\tStd Dev\tL-Bound\tU-Bound\tMinimum\tMaximum")
	for i := 0; i < len(collector.measures); i++ {
		obs, avg, std, lb, ub, min, max := collector.GetStat(i)
		fmt.Fprintf(w, "%s\t%d\t%6.3f\t%6.3f\t%6.3f\t%6.3f\t%6.3f\t%6.3f\n", collector.measures[i], obs, avg, std, lb, ub, min, max)
	}
	w.Flush()
	return
}

//GetStat returns size of the sample, average, standard deviation, low bound and uppe bounds of confidence inteval,  minimum and  maximum values
func (collector *StatCollector) GetStat(measureInd int) (int64, float64, float64, float64, float64, float64, float64) {
	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	avg := 0.
	std := 0.
	lb := 0.
	ub := 0.
	repl := int64(len(collector.samples))
	slice := []float64{}
	for i := 0; i < int(repl); i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	avg = Mean(slice)
	std = StandardDeviation(slice)
	lb, ub = NormalConfidenceInterval(slice)
	min, max := MinMax(slice)
	return repl, avg, std, lb, ub, min, max

}

//GetSize returns size of a sample
func (collector *StatCollector) GetSize(measureInd int) int {
	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	size := int(len(collector.samples))
	return size
}

//GetAverage returns average of a sample
func (collector *StatCollector) GetAverage(measureInd int) float64 {

	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	avg := 0.
	slice := []float64{}
	size := collector.GetSize(measureInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	avg = Mean(slice)
	return avg
}

//GetStandardDeviation returns standard deviation of a sample
func (collector *StatCollector) GetStandardDeviation(measureInd int) float64 {

	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	std := 0.
	slice := []float64{}
	size := collector.GetSize(measureInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	std = StandardDeviation(slice)
	return std
}

//GetLowBoundCI returns low bound of confidence interval for sample
func (collector *StatCollector) GetLowBoundCI(measureInd int) float64 {

	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	lb := 0.
	slice := []float64{}
	size := collector.GetSize(measureInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	lb, _ = NormalConfidenceInterval(slice)
	return lb
}

//GetUpperBoundCI returns upper bound of confidence interval for sample
func (collector *StatCollector) GetUpperBoundCI(measureInd int) float64 {

	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	ub := 0.
	slice := []float64{}
	size := collector.GetSize(measureInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	_, ub = NormalConfidenceInterval(slice)
	return ub
}

//GetMinimum returns minimum value for a sample
func (collector *StatCollector) GetMinimum(measureInd int) float64 {

	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	min := 0.
	slice := []float64{}
	size := collector.GetSize(measureInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	min, _ = MinMax(slice)
	return min
}

//GetMaximum returns maximum value for a sample
func (collector *StatCollector) GetMaximum(measureInd int) float64 {

	if measureInd < 0 || measureInd > len(collector.measures)-1 {
		panic("invalid index")
	}
	max := 0.
	slice := []float64{}
	size := collector.GetSize(measureInd)
	for i := 0; i < size; i++ {
		slice = append(slice, collector.samples[i][measureInd])
	}
	_, max = MinMax(slice)
	return max
}

// Mean returns float mean value
func Mean(nums []float64) (mean float64) {
	if len(nums) == 0 {
		return 0.0
	}
	for _, n := range nums {
		mean += n
	}
	return mean / float64(len(nums))
}

// StandardDeviation returns STD
func StandardDeviation(nums []float64) (dev float64) {
	if len(nums) == 0 {
		return 0.0
	}

	m := Mean(nums)
	for _, n := range nums {
		dev += (n - m) * (n - m)
	}
	dev = math.Pow(dev/float64(len(nums)-1.), 0.5)

	return dev
}

// NormalConfidenceInterval returns Confidence Interval
func NormalConfidenceInterval(nums []float64) (lower float64, upper float64) {
	conf := 1.95996 // 95% confidence for the mean, http://bit.ly/Mm05eZ
	mean := Mean(nums)
	dev := StandardDeviation(nums) / math.Sqrt(float64(len(nums)))
	return mean - dev*conf, mean + dev*conf
}

// MinMax returns minimum and maximum values amongst sample
func MinMax(nums []float64) (minimum float64, maximum float64) {

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
