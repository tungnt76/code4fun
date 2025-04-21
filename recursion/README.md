# Recursion

## Backtrack pattern

```golang
func backtrack(path []int, options []int, res *[][]int) {
 if isSolution(path) {
  // Make a copy of path and save it
  tmp := make([]int, len(path))
  copy(tmp, path)
  *res = append(*res, tmp)
  return
 }

 for i := 0; i < len(options); i++ {
  if isValid(options[i], path) {
   path = append(path, options[i])         // choose
   backtrack(path, options, res)           // explore
   path = path[:len(path)-1]               // un-choose (backtrack)
  }
 }
}
```
