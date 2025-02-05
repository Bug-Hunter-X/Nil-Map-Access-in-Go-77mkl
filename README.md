# Go Nil Map Access

This repository demonstrates a common, yet subtle, issue in Go: accessing keys in a nil map.  Unlike some other languages that might throw an exception, Go gracefully handles this by returning the zero value for the map's value type. While seemingly benign, this behavior can lead to unexpected results and difficult-to-debug issues if not handled carefully.

The `bug.go` file shows the problem.  The `bugSolution.go` file shows the solution for handling nil maps safely.
