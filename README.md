# Identicon

Identicon is github's [Identicons](https://github.com/blog/1586-identicons) clone in Go

## Simple example

Type: Horizontal, Theme: White (Default)
```go
package main

import "github.com/1l0/identicon"

func main() {
	id := identicon.New()
	id.GeneratePNGToFile("identicons/default")
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/default.png)

Type: Vertical, Theme: Black
```go
	id.Type = identicon.Mirrorv
	id.Theme = identicon.Black
	id.GeneratePNGToFile("identicons/vertical_black")
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/vertical_black.png)

Divisions: 7, Theme: Gray
```go
	id.Type = identicon.Mirrorh
	id.Theme = identicon.Gray
	id.Q = 50
	id.Div = 7
	id.GeneratePNGToFile("identicons/div7_gray")
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/div7_gray.png)

Margin: 140, Theme: Free
```go
	id.Theme = identicon.Free
	id.Q = 70
	id.Div = 5
	id.Margin = 140
	id.GeneratePNGToFile("identicons/margin140_free")
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/margin140_free.png)

Type: Normal, Theme: White
```go
	id.Type = identicon.Normal
	id.Theme = identicon.White
	id.Margin = 35
	id.GeneratePNGToFile("identicons/normal_white")
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/normal_white.png)

Random batch
```go
	id = identicon.New()
	id.GenerateRandomThemes("identicons/rand", 4)
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/rand1_white.png)![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/rand2_free.png)![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/rand3_white.png)![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/rand4_free.png)

Sequential batch
```go
	id.GenerateSequentialThemes("identicons/seq", 1)
}
```
![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/seq1_white.png)![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/seq1_black.png)![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/seq1_gray.png)![](https://raw.githubusercontent.com/1l0/identicon/master/example/identicons/seq1_free.png)
