# ABTest

```go
 import (
 	"fmt"
 	"github.com/KunBetter/ABTest/core"
 	"io/ioutil"
 )
 
 func main() {
 	ab := &core.ABTest{}
 	ab.Init()
 
 	buf, err := ioutil.ReadFile("ABExpConfig.json")
 	if err != nil {
 		fmt.Print(err)
 		return
 	}
 
 	ab.LoadConfig([]string{string(buf)})
 
 	reqMap := make(map[string]string)
 	tagMap := ab.Distribute(reqMap)
 
 	fmt.Println(tagMap)
 }
```

#### Next Plan
-[] Experimental variables matching supports multiple operators.