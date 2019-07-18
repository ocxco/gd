# gd

[查看中文版本](http://www.591study.cn/archives/go/252/gd/)

GO TO DIRECTORY 

It's can set an alias to any directory, to help us enter the directory easily

Written in golang

```
◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇
◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇
◇◇◇◇◇◇◇◆◆◆◆◆◆◆◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◆◆◆◆◆◆◆◆◆◆◇◇◇◇◇◇
◇◇◇◇◇◇◆◆◆◇◇◇◇◆◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◆◇◇◆◆◆◆◇◇◇◇◇
◇◇◇◇◇◆◆◆◇◇◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◆◆◆◆◇◇◇◇
◇◇◇◇◇◆◆◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◆◆◆◆◇◇◇
◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◆◆◆◇◇◇
◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◆◆◆◇◇◇
◇◇◇◇◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◆◆◆◇◇◇
◇◇◇◇◆◆◆◇◇◇◇◇◆◆◆◆◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◆◆◆◇◇◇
◇◇◇◇◆◆◆◇◇◇◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◆◆◆◇◇◇
◇◇◇◇◆◆◆◇◇◇◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◇◆◆◆◇◇◇
◇◇◇◇◇◆◆◆◇◇◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◇◆◆◆◆◇◇◇
◇◇◇◇◇◆◆◆◇◇◇◇◇◇◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◇◇◇◇◆◆◆◆◇◇◇◇
◇◇◇◇◇◇◆◆◆◆◇◇◇◆◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◆◆◆◇◆◆◆◆◆◇◇◇◇◇
◇◇◇◇◇◇◇◆◆◆◆◆◆◆◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◆◆◆◆◆◆◆◆◆◆◇◇◇◇◇◇
◇◇◇◇◇◇◇◇◇◆◆◆◆◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇
◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇◇
```

### Installation

```$xslt
git clone git@github.com:ocxco/gd.git
cd gd && bash ./bin/install.sh
```

### usage
```
Usage: gd [command] [alias]
  add  <alias> add current directory as alias
  add! <alias> force add current directory as alias
  rm   <alias> remove an exist alias from list
  ls | list    list all of configs
  pwd          show current directory, it's useful on git bash on windows
  help         show this help content
  clear!       clear all config
```
![演示图片](https://github.com/ocxco/gd/blob/master/images/gd.gif)

### wrapper

wrapper is refer to [mfaerevaag/wd-c](https://github.com/mfaerevaag/wd-c)
