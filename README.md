# Go Customizable Encoding/Json

<!-- > 2018-04-08T19:54:12+0800 -->

The customizable encoding/json package based on [zhanbei/golang-encoding-json](https://github.com/zhanbei/golang-encoding-json) -- the [encoding/json](https://github.com/golang/go/tree/master/src/encoding/json) package separated from [golang/go/master](https://github.com/golang/go).

## Goals

- [ ] Customizable/Different tags(`jsonFrom`, `jsonTo`) for decoding from and encoding to JSON string.
- [ ] Customizable JSON struct tag(`customizedJsonTag`), and customizable fall-back JSON struct tag(`json`).
- [ ] Customizable conversion policy to JSON keys from field names of a struct.
