# xtemplate
从 go 模板文件中提取翻译文本，并保存为 pot 文件。  
extract msgid from go template file and save to a pot file.


```bash
go install github.com/youthlin/t/cmd/xtemplate@latest
xtemplate -i <input-file-pattern> -k keywords

# e.g.
# help
xtemplate -h
# extract and save (with keywords(-k), function names(-f))
xtemplate -i "path/to/**/*.tmpl" -k "T;N:1,2;N64:1,2;X:1c,2;XN:1c,2,3;XN64:1c,2,3" -f FunName,Fun2 -o path/to/save/name.pot
```


```bash
## usage:
xtemplate -i input-pattern -k keywords [-f functions] [-o output]
  -d    debug mode
  -f string
        function names of template
  -h    show help message
  -i string
        input file pattern
  -k string
        keywords e.g.: gettext;T:1;N1,2;X:1c,2;XN:1c,2,3
  -left string
        left delim (default "{{")
  -o string
        output file, - is stdout
  -right string
        right delim (default "}}")
  -v    show version

## 用法

xtemplate -i 输入文件 -k 关键字 [-f 模版中函数] [-o 输出文件]
  -d    debug 模式
  -f string
        模板中用到的函数名
  -h    显示帮助信息
  -i string
        输入文件
  -k string
        关键字，例： gettext;T:1;N1,2;X:1c,2;XN:1c,2,3
  -left string
        左分隔符 (default "{{")
  -o string
        输出文件，- 表示标准输出
  -right string
        右分隔符 (default "}}")
  -v    显示版本号
```

## translations of this project
本项目的 `lang` 目录包含一个 po 文件，可以将其翻译为需要的语言，
然后设置环境变量 `LANG_PATH` 以加载翻译，默认的加载路径是 `./lang` 目录。

You can find a po file in `lang` dir.
set `LANG_PATH=/path/to/po/dir` to load your translations.
the default dir is `./lang`.
