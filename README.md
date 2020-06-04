# Wu Manber Multi Pattern String Matching

## Original repository
https://github.com/loirsun/wumanber

## Example

```
patterns := []string{"brown", "jump"}

wm := wumanber.WuManber{}
wm.Init(patterns)

text := "the quick brown fox jumps over the lazy dog."

// matches are the keywords matched
// matchPositions are the start position (in bytes) of the matches in the text
// therefore if the the text has utf8 characters using more than 1 byte, then
// text[n] != []byte(text)[n]
matches, matchPositions := wm.Search(text)

for i, match := range matches {
    fmt.Println(match, matchPositions[i])
}
// output:
// brown 10
// jump 20
```
