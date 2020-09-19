# MD2TXT
Markdown to TXT

# How to use

```
go get -u github.com/shiyunjin/md2txt
```

# Demo
## Code
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
## Output
```
MD2TXT
• Optimize the display list
• Pure text
• Simple to use
• Remove link
Bug fix
• hahahaha
lalalala
```
