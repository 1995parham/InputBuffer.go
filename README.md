# InputBuffer
[![Travis branch](https://img.shields.io/travis/AUTProjects/InputBuffer.go/master.svg?style=flat-square)](https://travis-ci.org/AUTProjects/InputBuffer.go)
[![GoDoc](https://godoc.org/github.com/AUTProjects/InputBuffer?status.svg)](http://godoc.org/github.com/AUTProjects/InputBuffer)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/269347155f17432c85b21871da4b6fa2)](https://www.codacy.com/app/1995parham/InputBuffer.go?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=AUTProjects/InputBuffer.go&amp;utm_campaign=Badge_Grade)

## Introduction
Input buffered switches are desirable for high-speed switching, since the internal operation speed is only moderately
higher than the input line speed. But there are two problems:

1. Throughput limitation due to the head-of-line (HOL) blocking (throughput limited to 58.6 percent for FIFO buffering).
2. The need of arbitrating cells due to output port contention.

## Installation
```sh
go get
go build
./InputBuffer.go
```

## Algorithms
- [iSLIP](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.84.1219&rep=rep1&type=pdf)
- [DRRM]()
- [RRLQF](http://ieeexplore.ieee.org/abstract/document/7442149/)

## Simulations
### RR-LQF-1 (Uniform Traffic)

| Ports | Throughput | Delay | Load | Speedup | Simulation Cycles |
| :---- | :--------: | :---: | :--: | :-----: | ----------------: |
| 8     |  58.41666  | 0     | 0.10 | 1       | 300               |
| 8     |  58.41666  | 0     | 0.10 | 2       | 300               |
| 8     |  58.41666  | 0     | 0.10 | 4       | 300               |
| 8     |  58.41666  | 0     | 0.10 | 2 - 1/N | 300               |
| 8     |  89.03412  | 3     | 0.50 | 1       | 300               |
| 8     |  89.48055  | 1.3   | 0.50 | 2       | 300               |
| 8     |  88.57896  | 0.6   | 0.50 | 4       | 300               |
| 8     |  87.75952  | 0     | 0.50 | 2 - 1/N | 300               |
| 8     |  89.66666  | 38    | 1.00 | 1       | 300               |
| 8     |  92.00000  | 3.08  | 1.00 | 2       | 300               |
| 8     |  90.20833  | 1.7   | 1.00 | 4       | 300               |
| 8     |  90.04166  | 3.42  | 1.00 | 2 - 1/N | 300               |
| 16    |  77.14444  | 0     | 0.10 | 1       | 300               |
| 16    |  75.57777  | 0     | 0.10 | 2       | 300               |
| 16    |  75.57777  | 0     | 0.10 | 4       | 300               |
| 16    |  75.57777  | 0     | 0.10 | 2 - 1/N | 300               |
| 16    |  92.43680  | 1.4   | 0.50 | 1       | 300               |
| 16    |  91.63916  | 1.8   | 0.50 | 2       | 300               |
| 16    |  91.63916  | 1.8   | 0.50 | 4       | 300               |
| 16    |  91.25122  | 1.1   | 0.50 | 2 - 1/N | 300               |
| 16    |  85.375    | 47.30 | 1.00 | 1       | 300               |
| 16    |  92.6875   | 5.733 | 1.00 | 2       | 300               |
| 16    |  92.5      | 5.888 | 1.00 | 4       | 300               |
| 16    |  91.8125   | 6.923 | 1.00 | 2 - 1/N | 300               |

### RR-LQF-2 (Uniform Traffic)

| Ports | Throughput | Delay | Load | Speedup | Simulation Cycles |
| :---- | :--------: | :---: | :--: | :-----: | ----------------: |
| 8     |  52.83333  | 1.5   | 0.10 | 1       | 300               |
| 8     |  51.88888  | 0     | 0.10 | 2       | 300               |
| 8     |  51.88888  | 0     | 0.10 | 4       | 300               |
| 8     |  51.88888  | 0     | 0.10 | 2 - 1/N | 300               |
| 8     |  88.85952  | 0     | 0.50 | 1       | 300               |
| 8     |  88.98333  | 1.25  | 0.50 | 2       | 300               |
| 8     |  88.98333  | 1.33  | 0.50 | 4       | 300               |
| 8     |  88.94628  | 1.5   | 0.50 | 2 - 1/N | 300               |
| 8     |  89        | 26.14 | 1.00 | 1       | 300               |
| 8     |  91.79166  | 3.6   | 1.00 | 2       | 300               |
| 8     |  91.04166  | 2.25  | 1.00 | 4       | 300               |
| 8     |  98.83333  | 2.77  | 1.00 | 2 - 1/N | 300               |
| 16    |  77.53333  | 0     | 0.10 | 1       | 300               |
| 16    |  75.53333  | 0     | 0.10 | 2       | 300               |
| 16    |  75.53333  | 0     | 0.10 | 4       | 300               |
| 16    |  75.53333  | 0     | 0.10 | 2 - 1/N | 300               |
| 16    |  92.71579  | 2.22  | 0.50 | 1       | 300               |
| 16    |  92.79074  | 0.2   | 0.50 | 2       | 300               |
| 16    |  92.79074  | 0.2   | 0.50 | 4       | 300               |
| 16    |  92.85100  | 1     | 0.50 | 2 - 1/N | 300               |
| 16    |  85.8333   | 35.07 | 1.00 | 1       | 300               |
| 16    |  93.0416   | 4.928 | 1.00 | 2       | 300               |
| 16    |  92.8541   | 3.791 | 1.00 | 4       | 300               |
| 16    |  92.4375   | 4.052 | 1.00 | 2 - 1/N | 300               |


## References
[1] High Performance Switches and Routers, Wiley 2007
