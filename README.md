# MD2TXT
Markdown to TXT

# How to use

```
go get -u github.com/shiyunjin/md2txt
```

# Demo

``` golang
package main

import (
	"log"

	"github.com/shiyunjin/md2txt"
)

func main() {
	txt := md2txt.ToTXT("##### MD2TXT\n- Optimize the display list\n- Pure text\n- Simple to use\n- Remove link\n##### Bug fix\n- hahahaha\n lalalala\n[aaaa](http://demo.com \"aaaa\")")
	log.Print(txt)
}
```
