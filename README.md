# DDDA Save Corrupter

PC `.sav` (version 21) editor. A CLI application converted to a simple web GUI (Vue or React). This is just a prototype.

![Tests](https://github.com/var512/ddda-save-corrupter/workflows/Tests/badge.svg)

```sh
# /etc/hosts => 127.0.0.1 dddasc.test
# vue.config.js
# constants/app.js
local_build_example() {
    local frontend="react"
    local webdir="web-react/build"
    local npmcmd="npm -C web-react run-script build"
    
    if [[ "${frontend}" = "vue" ]]; then
        webdir="web-vue/dist"
        npmcmd="npm -C web-vue run-script build"
    fi
    
    eval "${npmcmd}"
    
    go run \
        -tags=dev \
        -ldflags="-X 'github.com/var512/ddda-save-corrupter/internal/assets.WebDir=${webdir}'" \
        ./tools/assets/generate.go
    
    go build \
        -v \
        -race \
        -o "ddda-save-corrupter-${frontend}" \
        -tags=!dev \
        -ldflags="
          -X 'main.AppVersion=local-prod-build'
          -X 'github.com/var512/ddda-save-corrupter/internal/api.CorsAllowOrigin=*'
        " \
        ./cmd/ddda-save-corrupter
}
```
