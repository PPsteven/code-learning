" excute vimscript :source :so <filename>
" :so %    % represent name of current open file
let g:animal = 'cat'
echo g:animal
" keyword: 'set' is used to set value to internal option
" keyword: 'let' is used to set value to custom value

let is_cat = 1 " bool value: 1=true 0=false

" scope 作用域
" g: global v: vim scopt l:local b: current buffer w: current window
" t: current tab a: args of function  s: scope when excuting by :source

let @a= 'cat' " set value to register a
let g:output = 'value of g:animal:' . g:animal " use dot to joint string
" echom will save to messages, use :messages to find, help file :help message-history
echom g:output

" condition
let animal_kind = 'cat'
let animal_name = 'Tommy'
if g:animal_kind == 'cat'
   echo animal_name . ' is a cat'
elseif animal_kind = 'dog'
   echo animal_name . ' is a dog'
else
   echo animal_name . ' is something else'
endif


let g:is_cat = 0
let g:is_dog = 0
if !(g:is_cat || g:is_dog)
   echo g:animal_name . ' is something else.'
endif

" ==? case sensitive
" ==# case not sensitive
" set ignorecase case sensitive

" list
let animals = ['cat', 'dog', 'parrot']
let cat = animals[0]
let dog = animals[1]
let parrot = animals[-1]
let last_two_animals = animals[1:]
let last_two_animals = animals[1:2] " be caution: right limit is closed
echo "last_two_animals: "  last_two_animals
" add elements
call add(animals, 'octopus') " insert at the end of queue
call insert(animals, 'bobcat') " insert at the head of queue
echo 'after add and insert,  animals: ' animals
" delete elements
unlet animals[0]
call remove(animals, -1)
echo 'after unlet and remove,  animals: ' animals
" joint list
echo ['cat', 'dog'] + ['parrot', 'octopus']
let g:animals = ['cat', 'dog']
call extend(g:animals, ['parrot', 'octopus'])
" sort list`
call sort(animals)
echo "after sort: " animals
" get index of elements
echo "dog index " index(animals, 'dog')
if !empty(animals)
   echo "animals len: " . len(animals)
else
   echo "animals is empty("
endif
" count the numbers of elements
let num = [1, 2, 3, 1]
echo "there are " . count(num, 1) . ' one in the list'

" dict, dont to forget \
let animal_names = {
  \ 'cat': 'Tommy',
  \ 'rice': 'Jerry'
  \ }
echo 'animal_names(dict): ' animal_names
echo 'cat name: ' . animal_names.cat . ' rice name: ' . animal_names['rice']
" keys
echo 'animal_names keys: ' keys(animal_names)
" remove dict
unlet animal_names['cat']
call remove(animal_names, 'rice')
" has key
echo 'animal_names has key rice: ' . has_key(animal_names, 'rice')
" empty?
if empty(animal_names)
  echo 'animal_names(dict) is empty'
else
  echo 'animal_names(dict) len ' . len(animal_names)
  echo animal_names
endif

" loop
for animal in animals
  echo 'animal '. animal
endfor
" loop dict
let animal_names = {
  \ 'cat': 'Tommy',
  \ 'rice': 'Jerry'
  \ }
for [animal, name] in items(animal_names)
  echo 'This animal ' . animal . '''s name is ' . name
endfor

" while loop
let animals = ['cat', 'dog']
while !empty(animals)
  let animal =  remove(animals, 0)
  if animal == 'cat'
    echo 'Encountered a dog, break loop'
    break
  endif
endwhile

" func
function! AnimalNames(animal, ...) "UpperCase function name can be called by other program, add ! after function allow this function can be load many times
  echo a:1 . ' and other ' . (a:0 - 1) . ' ' . a:animal . ' here.'
endfunction

call AnimalNames('cat', 'Tommy', 'Jack', 'Alice')

"class
let box = {
  \ 'x': 10,
  \ 'y': 10
  \ }

function box.Move(x, y)
  let self['x'] = self['x'] + a:x
  let self['y'] = self['y'] + a:y
  return [self['x'] + a:x, self['y'] + a:y]
endfunction

echo 'box new postion: ' box.Move(-1, -2)

" add new function to class
function Area() dict
  return self['x'] * self['y']
endfunction

let box['Area'] = function('Area')

echo 'box area: ' . box.Area()

" lambda
let Add = {x, y -> x + y}
echo '2 + 3 = ' . Add(2, 3)

" interactive with vim
execute 'echo "tomey is a cat"'

" check does vim support some feature
" get full feature list :help feature-list
" mac win64 wsl nvim python3
if has('python')
  echo 'there is Python 2.x'
endif
if has('python3')
  echo 'there is Python 3.x'
endif

" file-related commands
echom 'current full file path ' . expand('%:p')
echom 'current file head ' . expand('%:p:h')
echom 'current file tail ' . expand('%:p:t')
echom 'current file root ' . expand('%:p:r')
echom 'current file ext ' . expand('%:p:e')

if filereadable(expand('%'))
  echo 'current file is readable'
endif

" The input prompt
function InputPromptFunc()
  call inputsave()
  let answer = confirm('Is cat your favorite animal?', "&yes\n&no")
  call inputrestore()
  echo answer
  call inputsave()
  let animal = input("what is your favorite animal?\n")
  call inputrestore()
  echo 'your favorite animal is ' . animal
endfunction

" call InputPromptFunc()
