# GoTowerDefense
a Tower Defense with the inspiriation from Plants Vs Zombis

Using 2D libb ebitengine
https://ebitengine.org/en/documents/install.html

## install
´´´
$sudo dnf install mesa-libGL-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel libXxf86vm-devel alsa-lib-devel pkg-config
´´´
### test if all works
'''
go run github.com/hajimehoshi/ebiten/v2/examples/rotate@latest
'''
### for windows
'''
GOOS=windows go run github.com/hajimehoshi/ebiten/v2/examples/rotate@latest
'''