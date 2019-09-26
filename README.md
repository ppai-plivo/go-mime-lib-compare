# go-mime-lib-compare

A simple project to compare different Go projects/packages that detect MIME types.

## Libraries

| Library                                                                      | Implementation | Stars                                                                              | Contributors                                                                                |
|------------------------------------------------------------------------------|----------------|------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| [http.DetectContentType](https://golang.org/pkg/net/http/#DetectContentType) | Pure Go        |                                                                                    |                                                                                             |
| [gabriel-vasile/mimetype](https://github.com/gabriel-vasile/mimetype)        | Pure Go        | ![stars](https://img.shields.io/github/stars/gabriel-vasile/mimetype?style=social) | ![contrib](https://img.shields.io/github/contributors/gabriel-vasile/mimetype?style=social) |
| [h2non/filetype](https://github.com/h2non/filetype)                          | Pure Go        | ![stars](https://img.shields.io/github/stars/h2non/filetype?style=social)          | ![contrib](https://img.shields.io/github/contributors/h2non/filetype?style=social)          |
| [vimeo/go-magic/magic](https://github.com/vimeo/go-magic)                    | libmagic       | ![stars](https://img.shields.io/github/stars/vimeo/go-magic?style=social)          | ![contrib](https://img.shields.io/github/contributors/vimeo/go-magic?style=social)          |
| [rakyll/magicmime](https://github.com/rakyll/magicmime)                      | libmagic       | ![stars](https://img.shields.io/github/stars/rakyll/magicmime?style=social)        | ![contrib](https://img.shields.io/github/contributors/rakyll/magicmime?style=social)        |

## Building

**Dependency:**

Some go packages for MIME type detection are [libmagic](https://github.com/file/file) wrappers.

```sh
$ brew install libmagic
```

**Build:**
```sh
$ go build
```

**Use:**
```sh
$ ./go-mime-lib-compare > out.md
```

See [comparison](out.md)
