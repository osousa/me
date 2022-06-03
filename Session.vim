let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Desktop/Go/osousa.me
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
set shortmess=aoO
argglobal
%argdel
$argadd me.go
set stal=2
tabnew
tabnew
tabnew
tabrewind
edit me.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe '2resize ' . ((&lines * 28 + 29) / 59)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
exe '3resize ' . ((&lines * 27 + 29) / 59)
exe 'vert 3resize ' . ((&columns * 207 + 119) / 238)
argglobal
enew
file NvimTree_7
balt router.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
wincmd w
argglobal
balt router.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 105 - ((0 * winheight(0) + 14) / 28)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 105
normal! 05|
wincmd w
argglobal
if bufexists("term://~/Desktop/Go/osousa.me//1331266:/bin/bash") | buffer term://~/Desktop/Go/osousa.me//1331266:/bin/bash | else | edit term://~/Desktop/Go/osousa.me//1331266:/bin/bash | endif
if &buftype ==# 'terminal'
  silent file term://~/Desktop/Go/osousa.me//1331266:/bin/bash
endif
balt me.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 16 - ((1 * winheight(0) + 13) / 27)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 16
normal! 012|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe '2resize ' . ((&lines * 28 + 29) / 59)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
exe '3resize ' . ((&lines * 27 + 29) / 59)
exe 'vert 3resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit router.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
enew
file NvimTree_6
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
wincmd w
argglobal
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 133 - ((9 * winheight(0) + 28) / 56)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 133
normal! 016|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit models.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
enew
file NvimTree_4
balt models.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
wincmd w
argglobal
balt me.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 74 - ((55 * winheight(0) + 28) / 56)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 74
normal! 05|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit database.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
enew
file NvimTree_5
balt database.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
wincmd w
argglobal
balt database.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 114 - ((35 * winheight(0) + 28) / 56)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 114
normal! 05|
wincmd w
2wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext 4
set stal=1
badd +98 me.go
badd +1 router.go
badd +89 models.go
badd +179 database.go
badd +1 middlewares.go
badd +26 term://~/Desktop/Go/osousa.me//742110:/bin/bash
badd +35 term://~/Desktop/Go/osousa.me//734383:/bin/bash
badd +1 db/database.go
badd +68 term://~/Desktop/Go/osousa.me//734399:/bin/bash
badd +1 go.mod
badd +1 go.sum
badd +1 ~/Desktop/Go/osousa.me/.env
badd +170 term://~/Desktop/Go/osousa.me//98525:/bin/bash
badd +0 term://~/Desktop/Go/osousa.me//1331266:/bin/bash
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20 shortmess=filnxtToOF
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
