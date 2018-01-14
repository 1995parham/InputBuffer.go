# InputBuffer
[![Travis branch](https://img.shields.io/travis/1995parham/InputBuffer/master.svg?style=flat-square)](https://travis-ci.org/1995parham/InputBuffer)
[![GoDoc](https://godoc.org/github.com/1995parham/InputBuffer?status.svg)](http://godoc.org/github.com/1995parham/InputBuffer)
[![Codacy grade](https://img.shields.io/codacy/grade/269347155f17432c85b21871da4b6fa2.svg?style=flat-square)](https://www.codacy.com/app/1995parham/InputBuffer?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=1995parham/InputBuffer&amp;utm_campaign=Badge_Grade)

## Introduction
Input buffered switches are desirable for high-speed switching, since the internal operation speed is only moderately
higher than the input line speed. But there are two problems:
1. Throughput limitation due to the head-of-line (HOL) blocking (throughput limited to 58.6 percent for FIFO buffering).
2. The need of arbitrating cells due to output port contention.

## Algorithms
- [iSLIP](algorithm/islip.go)
- [DRRM](algorithm/drrm.go)

## References
[1] High Performance Switches and Routers, Wiley 2007
