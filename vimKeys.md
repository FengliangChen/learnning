***
* ****j*ump to line123: 123g
* delete until x: dtx
* delete to end: de
* to next "n": fn
* to previous "n": Fn
* visual block: ctrl+v
* delete to end: shift+c
* delete current line and insert mode: cc
* %s/foo/bar/g: globally replace "foo"with "bar"
* to column: 30| to 30 column
* to cursor: :call cursor(15,25) to line 15, column25
* to previous position ''
* add to end of line A
* to blank g _
* to middle line of screen: M, to end of screen:L
* d3d p ---> delete 3 line and paste to cursor
* dib ----> delete the line with the bracelte.
* da{ ----> delete the whole block of the {}
* { or } move between the blank lines.



## Vimscript

### Key Mapping

对应*normal*、*visual*、*insert*模式的 *nmap*、*vmap*和*imap*

按键：
```
<c-d>， <cr>, <esc>
```
非递归映射： noremap, nnoremap, vnoremap, inoremap

Leader键： let mapleader = ","

实例：
在一个分屏中打开~/.vimrc文件
```
nnoremap <leader>ev :vsplit $MYVIMRC<cr>
```
重读映射配置：
```
nnoremap <leader>sv :source $MYVIMRC<cr>
```
### abbreviations

```
iabbrev adn and
```