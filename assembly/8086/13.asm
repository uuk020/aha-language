assume cs:code

a segment
  db 1,2,3,4,5,6,7,8
a ends

b segment
  db 1,2,3,4,5,6,7,8
b ends

c segment
  db 0,0,0,0,0,0,0,0
c ends

code segment
start: 
       mov ax, a
       mov es, ax

       mov ax, c
       mov ds, ax

       mov cx, 8
       mov bx, 0
    s: 
       mov al, es:[bx]
       add ds:[bx], al
       inc bx
       loop s

       mov ax, b
       mov es, ax

       mov cx, 8
       mov bx, 0

    s1:
       mov al, es:[bx]
       add ds:[bx], al
       inc bx
       loop s1

       mov ax, 4c00h
       int 21
code ends

end start
