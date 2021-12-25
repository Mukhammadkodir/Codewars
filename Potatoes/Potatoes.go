/*
John bought potatoes: their weight is 100 kilograms.
Potatoes contain water and dry matter. The water content is 99 percent of the total weight.
He thinks they are too wet and puts them in an oven - at low temperature - for them to lose some water.
At the output the water content is only 98%.
What is the total weight in kilograms (water content plus dry matter) coming out of the oven?
He finds 50 kilograms and he thinks he made a mistake:
"So much weight lost for such a small change in water content!"
*/
package main

import "fmt"

func main() {
	fmt.Println(Potatoes(82, 127, 80))
}
func Potatoes(p0, w0, p1 int) int {

	return (100 - p0) * w0 / (100 - p1)
}
