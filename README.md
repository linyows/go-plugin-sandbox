go-plugin-sandbox
=================

use go v1.8 plugin package

```sh
$ docker-compose up
Starting gopluginsandbox_plugin_1
Attaching to gopluginsandbox_plugin_1
plugin_1  | go build -buildmode=plugin -o plugins/animal_sounds.so
plugins/animal.go
plugin_1  | go build -buildmode=plugin -o plugins/machine_sounds.so
plugins/machine.go
plugin_1  | go run main.go
plugin_1  |
plugin_1  | Plugins:
plugin_1  | ./plugins/animal_sounds.so
plugin_1  | animal plugin init!
plugin_1  | loaded: dog
plugin_1  | loaded: cat
plugin_1  | loaded: cow
plugin_1  | ./plugins/machine_sounds.so
plugin_1  | machine plugin init!
plugin_1  | loaded: windows
plugin_1  | loaded: ubuntu
plugin_1  | loaded: macos
plugin_1  |
plugin_1  | Noise1:
plugin_1  | cow < mu~u~~~~~~~!
plugin_1  | win < tete------n
plugin_1  | ubuntu < pecopon....
plugin_1  | mac < jyaaaaaaaan
plugin_1  | dog < bowwow!!!!!!
plugin_1  | cat < meow~ mewo~
plugin_1  |
plugin_1  | Noise2:
plugin_1  | tete------n (windows)
plugin_1  | pecopon.... (ubuntu)
plugin_1  | jyaaaaaaaan (macos)
plugin_1  | bowwow!!!!!! (dog)
plugin_1  | meow~ mewo~ (cat)
plugin_1  | mu~u~~~~~~~! (cow)
gopluginsandbox_plugin_1 exited with code 0
```
