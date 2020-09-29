# Cyclic dependencies test

### Plugin Example (Build Error cyclic imports are not allowed)
Plugins `one` directly depends on `two` and vice versa.
```
go build --buildmode=plugin -gcflags=all='-N -l' -o $PWD/out/plugins/one plugins/one/main.go  
go build --buildmode=plugin -gcflags=all='-N -l' -o $PWD/out/plugins/two plugins/two/main.go
go build --gcflags=all='-N -l' -o $PWD/out/plugin_test  cmd/plugin_example/main.go
```


### Static Example (Build Error cyclic imports are not allowed)
```
go build -gcflags=all='-N -l' -o $PWD/out/static_test cmd/static_example/main.go 
```

### Plugin Example Implicit imports
Plugins `three` indirectly depends on `four` and vice versa.
```
go build --buildmode=plugin -gcflags=all='-N -l' -o $PWD/out/plugins/three plugins/three/main.go  
go build --buildmode=plugin -gcflags=all='-N -l' -o $PWD/out/plugins/four plugins/four/main.go
go build --gcflags=all='-N -l' -o $PWD/out/plugin_test  cmd/plugin_example/main.go 
./out/plugin_test three
```
