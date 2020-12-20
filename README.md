# christmasTree
```
go build christmastree.go
sudo ./christmastree
```

# create service
```
sudo ln -s ~/go/src/github.com/mmalessa/christmasTree/christmastree /usr/bin
sudo ln -s /go/src/github.com/mmalessa/christmasTree/system/christmastree.service /etc/systemd/system
sudo systemctl enable christmastree.service
sudo systemctl start christmastree.service
```
