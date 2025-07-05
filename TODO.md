# TODOs

## v0.1
[X] Highlight each occurence of searched word in result list  
[X] Go up and down in result list  
[ ] Paginate through result list with TAB
[ ] Read DB schema through configuration  
[X] Don't delete if input list is empty (only char in prompt is '>')  
[ ] Disable special keys like CTRL + something, except CTRL-C and CTRL-Z  
    [ ] Allow putting program in background through CTRL-Z  
[ ] Add bi-modal system like viM(INSERTION and NORMAL mode),  
    [X] Enter NORMAL mode with ESC which shifts focus to the result list
    by selecting the very first list item  
    [ ] In NORMAL mode filter by any field set in the DB scheme for the table  
    [ ]* In NORMAL mode press 'i' to preview DB's table entry  
    [ ] See current mode in last line of terminal(get window's size)
[ ] See last results at program start  
[ ] Use byte slice instead of strings in prompt.go for performance improvements

