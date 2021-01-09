Christmas tree lighting control tool - ws2811 strip  
/sketch/

# Requirements
- Raspberry PI with Raspbian
- Go installed (apt install go)

# Get
```
go get github.com/mmalessa/christmasTree
```

# Settings (tree matrix)
- we wrap the Christmas tree with a led strip
```
    *
   *-*
  *-*-*
 *-*-*-*
*-*-*-*-*------------ [Raspberry PI / Power Supply]
```
- we map this in the tree_matrix variable in christmastree.go, e.g.:
```
{
    {14, 14, 14, 14, 14},
    {13, 13, 13, 12, 12},
    {11, 11, 10, 9,  9},
    {8,  7,  6,  6,  5},
    {4,  3,  2,  1,  0},
}
```
It is not perfect, but we try to get a reasonable representation.


# Build & run
```
cd ~/go/src/github.com/mmalessa/christmasTree
go build christmastree.go
sudo ./christmastree
```

# Create service
```
sudo ln -s ~/go/src/github.com/mmalessa/christmasTree/christmastree /usr/bin
sudo ln -s /go/src/github.com/mmalessa/christmasTree/system/christmastree.service /etc/systemd/system
sudo systemctl enable christmastree.service
sudo systemctl start christmastree.service
```

# Hardware
```
Power Supply (+)---------> Strip +

PI PIN 6 (GND)-------|
Power Supply (GND)---+---> Strip GND

PI PIN 12 (GPIO18)-------> Strip Data Input
```
