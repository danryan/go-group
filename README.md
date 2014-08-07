# go-group

A supplemental package to `os/user` for working with groups. 

Extracted from [https://codereview.appspot.com/13454043](https://codereview.appspot.com/13454043). This package is a temporary solution until the aforementioned patch is merged into Go proper.

## Usage

```go
package main

import (
  "github.com/danryan/go-group/os/group"
  "fmt"
  "strconv"
)

type Group struct {
  g    *group.Group
  gid  int
  name string
}

func GroupByName(name string) (*Group, error) {
  g, err := group.LookupGroup(name)
  if err != nil {
    return nil, err
  }

  gid, err := strconv.Atoi(g.Gid)
  if err != nil {
    return nil, err
  }

  return &Group{
    name: g.Name,
    gid:  gid,
  }, nil
}

func main() {
  g := GroupByName("admin")
  fmt.Println(g.name, g.gid)
}
```
