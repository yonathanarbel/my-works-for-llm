package main

import (
    "fmt"
    "os"
    "time"
    "math/rand"
    "encoding/json"
    "net/http"
    "log"
    "strings"
)

const articleText = `
SLICING DEFAMATION BY CONTRACT

Commented [A1]: SLICES AND LUMPS:
DIVISION AND AGGREGATION IN LAW AND LIFE (2019)

Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

2 Chicago Law Review Online 2020
Commented [A2]:
Commented [A3]:
Commented [A4]:
Commented [A5]: Daniel Hemel and Ariel Porat, Free Speech
and Cheap Talk, J. LEG. ANALYSIS (Forthcoming, 2019); Yonathan
A. Arbel and Murat Mungan, The Case Against Expanding
Defamation Laws, ALABAMA L. REV. (Forthcoming, 2019);
Yonathan A. Arbel & Murat Mungan, Regulating Speech with
Bayesian Audiences, Alabama Working Paper 3452662 (2019).
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

Commented [A6]:
Commented [A7]:
Commented [A8]:
Commented [A9]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

4 Chicago Law Review Online 2020
Commented [A10]:
Commented [A11]:
Commented [A12]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

Commented [A13]:
Commented [A14]:
Commented [A15]:
Commented [A16]:
Commented [A17]:
Commented [A18]:
Commented [A19]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

6 Chicago Law Review Online 2020
Commented [A20]:
Commented [A21]:
Commented [A22]:
Commented [A23]:
Commented [A24]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

Commented [A25]:
Commented [A26]:
Commented [A27]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

8 Chicago Law Review Online 2020
Commented [A28]: Carlill v. Carbolic Smoke Ball Co. 1 Q.B.
256 (1893)
Commented [A29]:
Commented [A30]:
Commented [A31]:
Commented [A32]:
Commented [A33]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

Commented [A34]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

10 Chicago Law Review Online 2020
Commented [A35]:
Electronic copy available at: https://ssrn.com/abstract=3681083

<<PAGE_BREAK>>

Commented [A36]:
Electronic copy available at: https://ssrn.com/abstract=3681083`

func main() {
    fmt.Println(articleText)
}
