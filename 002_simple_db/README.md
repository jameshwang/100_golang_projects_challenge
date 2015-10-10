### Lessons learned

1. debugging is harder than Ruby.  I had to resort to `println` statements
everywhere which I hate.  Can't we just use a debugger like we do in Ruby?
Like `pry`?
1. Learning from my mistakes about type conversion
1. Why are the functions getting so long with all these `if err != nil` statements?
Am I not breaking these logical parts out enough?  Should I be checking each err
value or can I skip a few?
1. Research these other libraries out there like gorilla
