# xtemplate
从 go 模板文件中提取翻译文本，并保存为 pot 文件。  
extract msgid from go template file and save to a pot file.


```bash
go install github.com/youthlin/t/cmd/xtemplate
xtemplate -i <input-file-pattern> -k keywords
```

## translations of this project
本项目的 `lang` 目录包含一个 po 文件，可以将其翻译为需要的语言，
然后设置环境变量 `LOCALEDIR` 以加载翻译，默认的加载路径是 `./lang` 目录。

You can find a po file in `lang` dir.
set `LOCALEDIR=/path/to/po/dir` to load your translations.
the default dir is `./lang`.
