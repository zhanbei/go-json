# Go Customizable Encoding/Json

<!-- > 2018-04-08T19:54:12+0800 -->

The customizable encoding/json package based on [`zhanbei/golang-encoding-json`][project-upstream] -- the [`encoding/json`][package-encoding-json] package separated from [`golang/go/master`][github-golang-go].

## Features

The goal is to add more customizable features while keeping compatible with the [`encoding/json`][package-encoding-json] package.

- [ ] Customizable/Different tags(`fromJson`, `toJson`) for decoding from and encoding to JSON string.
- [ ] Customizable JSON struct tag(`customizedJsonTag`), and customizable fall-back JSON struct tag(`json`).
- [ ] Customizable conversion policy to JSON keys from field names of a struct.

### Differences with the [`encoding/json`][package-encoding-json] Package

- [ ] By default, the tags `fromJson` and `toJson` will be checked first respectively, over the tag `json`.


[github-golang-go]: https://github.com/golang/go "Go Source Code"
[package-encoding-json]: https://github.com/golang/go/tree/master/src/encoding/json "Go Package `encoding/json`"
[project-upstream]: https://github.com/zhanbei/golang-encoding-json "Project Upstream: Separated `encoding/json` Package"
