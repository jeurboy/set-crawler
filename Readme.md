# set-crawler
Thai stock exchange crawler

```
package main

import crawler "github.com/jeurboy/set-crawler"

func main() {
	data, _ := crawler.GetSetPriceData("AAV", 0)

	//print data
	dump.DD(data)

	crawler.GetAllStock()

	fmt.Println("End")
}

```