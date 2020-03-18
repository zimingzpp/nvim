" Settings {{{
" Switch syntax highlighting on, when the terminal has colors
syntax on

"Enable mouse usage (all modes)
set mouse=a		

" Use vim, not vi api
set nocompatible

" No backup files
set nobackup

" No write backup
set nowritebackup

" No swap file
set noswapfile

" Command history
set history=100

" Always show cursor
set ruler

" Show incomplete commands
set showcmd

" Incremental searching (search as you type)
set incsearch

" Highlight search matches
"set hlsearch

" Ignore case in search
set smartcase

" Make sure any searches /searchPhrase doesn't need the \c escape character
set ignorecase

" A buffer is marked as ‘hidden’ if it has unsaved changes, and it is not currently loaded in a window
" if you try and quit Vim while there are hidden buffers, you will raise an error:
" E162: No write since last change for buffer “a.txt”
set hidden

" Turn word wrap off
set nowrap

" Allow backspace to delete end of line, indent and start of line characters
set backspace=indent,eol,start

" Convert tabs to spaces
set expandtab

" Set tab size in spaces (this is for manual indenting)
set tabstop=4

" The number of spaces inserted for a tab (used for auto indenting)
set shiftwidth=4

" Turn on line numbers
set number

" Highlight tailing whitespace
" See issue: https://github.com/Integralist/ProVim/issues/4
set list listchars=tab:\ \ ,trail:·

" Get rid of the delay when pressing O (for example)
" http://stackoverflow.com/questions/2158516/vim-delay-before-o-opens-a-new-line
set timeout timeoutlen=1000 ttimeoutlen=100

" Always show status bar
set laststatus=2

" Hide the toolbar
set guioptions-=T

" UTF encoding
set encoding=utf-8

" Autoload files that have changed outside of vim
set autoread

" Use system clipboard
" http://stackoverflow.com/questions/8134647/copy-and-paste-in-vim-via-keyboard-between-different-mac-terminals
set clipboard+=unnamed

" Don't show intro
set shortmess+=I

" Better splits (new windows appear below and to the right)
set splitbelow
set splitright

" Highlight the current line
set cursorline

" Ensure Vim doesn't beep at you every time you make a mistype
set visualbell

" Visual autocomplete for command menu (e.g. :e ~/path/to/file)
set wildmenu

" redraw only when we need to (i.e. don't redraw when executing a macro)
set lazyredraw

" highlight a matching [{()}] when cursor is placed on start/end character
set showmatch

" Set built-in file system explorer to use layout similar to the NERDTree plugin
let g:netrw_liststyle=3

" Always highlight column 80 so it's easier to see where
" cutoff appears on longer screens
autocmd BufWinEnter * highlight ColorColumn ctermbg=darkblue
set colorcolumn=120

" }}}

"Plugins{{{
call plug#begin('~/.vim/plugged')
Plug 'preservim/nerdtree'
Plug 'tomasr/molokai'
Plug 'liuchengxu/space-vim-dark'
Plug 'powerline/powerline'
Plug 'mattn/emmet-vim'
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
Plug 'suan/vim-instant-markdown', {'for': 'markdown'}
Plug 'mileszs/ack.vim'
Plug 'powerline/powerline'
Plug 'ryanoasis/vim-devicons'
Plug 'tiagofumo/vim-nerdtree-syntax-highlight'
Plug 'morhetz/gruvbox'
Plug 'vim-syntastic/syntastic'
Plug 'neoclide/coc.nvim', {'branch': 'release'}
"文本对齐
Plug 'https://github.com/godlygeek/tabular'
Plug 'https://github.com/junegunn/vim-easy-align'
Plug 'https://github.com/tommcdo/vim-lion'

"Cycle
Plug 'https://github.com/AndrewRadev/switch.vim'
Plug 'https://github.com/tpope/vim-speeddating'

"代码注释
Plug 'https://github.com/scrooloose/nerdcommenter'
Plug 'https://github.com/tomtom/tcomment_vim'
Plug 'https://github.com/tpope/vim-commentary'

"自动补全括号
Plug 'https://github.com/jiangmiao/auto-pairs'
Plug 'https://github.com/Raimondi/delimitMate'
Plug 'https://github.com/tpope/vim-endwise'

"
Plug 'https://github.com/majutsushi/tagbar'
Plug 'https://github.com/tpope/vim-surround'
call plug#end()

set t_Co=256
set background=dark
colorscheme gruvbox
let g:rehash256 = 1

let g:minBufExplForceSyntaxEnable = 1
set guifont=Source\ Code\ Pro\ for\ Powerline:h12
let g:instant_markdown_port = 8888

"}}}

"Commands{{{
autocmd BufRead * setlocal foldmethod=marker
autocmd BufRead * normal zM
"}}}
